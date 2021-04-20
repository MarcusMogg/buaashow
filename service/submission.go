package service

import (
	"buaashow/entity"
	"buaashow/global"
	"buaashow/utils"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
)

type workerType int

const (
	toZip workerType = iota + 1
	toExe
	toSrc
)

type info struct {
	gid     string
	tmpPath string
	tp      workerType
}

type expFile struct {
	eid      uint
	fileToDo chan info
	exit     chan struct{}
	//endtime  time.Time
}

// FIXME : how many goroutines we need?
const routineNums = 5

var (
	m     map[uint]*expFile
	mutex sync.RWMutex
)

// func clear() {
// 	now := time.Now()
// 	// 计算下一个执行时间 4:00
// 	next := now.Add(time.Hour * 24)
// 	next = time.Date(next.Year(), next.Month(), next.Day(),
// 		0, 4, 0, 0, next.Location())

// 	t := time.NewTimer(next.Sub(now))
// 	// will this channel will be destroyed?
// 	<-t.C
// 	for range time.Tick(24 * time.Hour) {
// 		mutex.Lock()
// 		for i, j := range m {
// 			if j.endtime.Before(now) {
// 				delete(m, i)
// 			}
// 		}
// 		mutex.Unlock()
// 	}
// }

// 创建文件夹 如果文件夹存在则清空内容
func createDir(dir string) error {
	s, err := os.Open(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			zap.S().Errorf("神必异常: %s", err)
			return err
		}
		if err := os.MkdirAll(dir, 0755); err != nil {
			zap.S().Errorf("创建目录失败: %s", err)
			return err
		}
	} else {
		defer s.Close()
		// 删除上一次
		names, err := s.Readdirnames(-1)
		if err != nil {
			zap.S().Errorf("神必异常: %s", err)
			return err
		}
		for _, name := range names {
			err = os.RemoveAll(filepath.Join(dir, name))
			if err != nil {
				zap.S().Errorf("神必异常: %s", err)
				return err
			}
		}
	}
	return nil
}

// 将压缩包解压至 global.GCoursePath/{eid}/{gid}/show/
// TODO: 失败时将失败信息写到 global.GCoursePath/{eid}/{gid}/show/index.html
func worker(dirPath string, file *info) {
	var err error
	switch file.tp {
	case toZip:
		err = moveUnzip(dirPath, file.gid, file.tmpPath)
	case toExe:
		err = moveExe(dirPath, file.gid, file.tmpPath)
	case toSrc:
		err = moveSrc(dirPath, file.gid, file.tmpPath)
	default:
		err = errors.New("unknown type")
	}
	zap.S().Debug(err)
}

// func updateEndtime(e *entity.MExperiment) {
// 	mutex.Lock()
// 	_, ok := m[e.ID]
// 	if ok {
// 		m[e.ID].endtime = e.EndTime
// 		mutex.Unlock()
// 	} else {
// 		mutex.Unlock()
// 		initWorker(e)
// 	}
// }

func initWorker(e *entity.MExperiment) error {
	exp := expFile{
		eid:      e.ID,
		fileToDo: make(chan info, routineNums*2),
		exit:     make(chan struct{}, 1),
		//endtime:  e.EndTime,
	}
	dirPath := fmt.Sprintf("%s%d/", global.GCoursePath, exp.eid)
	if _, err := os.Stat(dirPath); err != nil {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			zap.S().Errorf("创建目录失败: %s", err)
			return err
		}
	}
	f := func() {
		for {
			select {
			case <-exp.exit:
				return
			case file := <-exp.fileToDo:
				worker(dirPath, &file)
			}
		}
	}
	for i := 0; i < routineNums; i++ {
		go f()
	}
	mutex.Lock()
	m[e.ID] = &exp
	mutex.Unlock()
	return nil
}

// InitSubmitThread 初始化
func InitSubmitThread() {
	m = make(map[uint]*expFile)
	//go clear()
	//now := time.Now()
	var exps []entity.MExperiment
	global.GDB.Model(&entity.MExperiment{}).
		//Where("end_time >= ?", now.Format(global.TimeTemplateSec)).
		Find(&exps)
	for _, j := range exps {
		err := initWorker(&j)
		if err != nil {
			zap.S().Fatal(err)
		}
	}
}

func toWorker(eid uint, gid string, file string, tp workerType) error {
	var e *expFile
	var ok bool
	mutex.RLock()
	e, ok = m[eid]
	mutex.RUnlock()

	if !ok {
		return errors.New("ToWorker 初始化异常")
	}
	// no wait
	go func() {
		e.fileToDo <- info{
			gid:     gid,
			tmpPath: filepath.Join(global.GTmpPath, file),
			tp:      tp,
		}
	}()
	return nil
}

func copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func moveExe(dirPath, gid, tmpPath string) error {
	filetype := path.Ext(tmpPath)
	if filetype != ".zip" {
		return errors.New("不支持的类型")
	}
	dir := fmt.Sprintf("%s%s/show/", dirPath, gid)
	if err := createDir(dir); err != nil {
		return err
	}
	return copy(tmpPath, filepath.Join(dir, "release.zip"))
}

func moveSrc(dirPath, gid, tmpPath string) error {
	filetype := path.Ext(tmpPath)
	if filetype != ".zip" {
		return errors.New("不支持的类型")
	}
	dir := fmt.Sprintf("%s%s/", dirPath, gid)
	if err := os.MkdirAll(dir, 0755); err != nil {
		zap.S().Errorf("创建目录失败: %s", err)
		return err
	}
	return copy(tmpPath, filepath.Join(dir, "src.zip"))
}

func moveUnzip(dirPath string, gid, tmpPath string) error {
	// 创建作业根目录
	dir := fmt.Sprintf("%s%s/show", dirPath, gid)
	logfile := fmt.Sprintf("%s%s/log", dirPath, gid)
	err := createDir(dir)
	if err != nil {
		return err
	}
	var msg string
	err = utils.UnZip(tmpPath, dir)
	if err != nil {
		zap.S().Errorf("解压错误 %s\n", err.Error())
		msg = fmt.Sprintf("[%s] 解压错误 %s\n", time.Now().Local().Format(global.TimeTemplateSec), err.Error())
	} else {
		msg = fmt.Sprintf("[%s] 提交成功\n", time.Now().Local().Format(global.TimeTemplateSec))
	}
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		zap.S().Errorf("打开日志文件错误 %s\n", err.Error())
		return err
	}
	f.Write([]byte(msg))
	f.Close()
	copy(logfile, filepath.Join(dir, "log"))
	copy(tmpPath, fmt.Sprintf("%s%s/dist.zip", dirPath, gid))
	return nil
}

func DownloadSubmission(eid uint, uid string) (string, error) {
	var mid entity.MExperimentSubmit
	if err := global.GDB.Where("e_id = ? AND uid = ?", eid, uid).
		First(&mid).Error; err != nil {
		return "", err
	}
	if !mid.Status {
		return "", errors.New("未提交")
	}
	dirPath := fmt.Sprintf("%s%d/%s", global.GCoursePath, eid, mid.GID)
	dirPath = filepath.Clean(dirPath)
	filename := fmt.Sprintf("%d-%s.zip", eid, uid)
	outName := path.Join(global.GTmpPath, filename)

	outfile, err := os.Stat(outName)
	if err == nil {
		if outfile.ModTime().After(mid.UpdatedAt) {
			return outName, nil
		}
	}
	err = utils.ZipFiles(outName,
		[]string{dirPath},
		[]string{dirPath},
		[]string{uid})
	return outName, err
}

func DownloadAllSubmission(eid uint) (string, error) {
	dirPath := fmt.Sprintf("%s%d/", global.GCoursePath, eid)
	dirPath = filepath.Clean(dirPath)
	filename := fmt.Sprintf("%d.zip", eid)
	outName := path.Join(global.GTmpPath, filename)

	outfile, err := os.Stat(outName)
	if err == nil {
		var up time.Time
		err = global.GDB.Model(&entity.MExperimentSubmit{}).
			Select("updated_at").
			Order("updated_at").
			Limit(1).Find(&up).Error
		if err == nil && outfile.ModTime().After(up) {
			return outName, nil
		}
	}
	err = utils.ZipFiles(outName,
		[]string{dirPath},
		[]string{dirPath},
		[]string{fmt.Sprintf("%d", eid)})
	return outName, err
}

func CopyDir(src, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				zap.S().Debug(err)
			}
		} else {
			if err = copy(srcfp, dstfp); err != nil {
				zap.S().Debug(err)
			}
		}
	}
	return nil
}
