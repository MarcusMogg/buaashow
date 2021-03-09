package service

import (
	"buaashow/entity"
	"buaashow/global"
	"buaashow/utils"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
)

type info struct {
	gid     string
	tmpPath string
}

type expFile struct {
	eid      uint
	fileToDo chan info
	exit     chan struct{}
	endtime  time.Time
}

// FIXME : how many goroutines we need?
const routineNums = 5

var (
	m     map[uint]*expFile
	mutex sync.RWMutex
)

func clear() {
	now := time.Now()
	// 计算下一个执行时间 4:00
	next := now.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(),
		0, 4, 0, 0, next.Location())

	t := time.NewTimer(next.Sub(now))
	// will this channel will be destroyed?
	<-t.C
	for range time.Tick(24 * time.Hour) {
		mutex.Lock()
		for i, j := range m {
			if j.endtime.Before(now) {
				delete(m, i)
			}
		}
		mutex.Unlock()
	}
}

// 将压缩包解压至 global.GCoursePath/{eid}/{gid}/show/
// TODO: 失败时将失败信息写到 global.GCoursePath/{eid}/{gid}/show/index.html
func worker(dirPath string, file *info) {
	// 创建作业根目录
	dir := fmt.Sprintf("%s%s/", dirPath, file.gid)
	s, err := os.Open(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			zap.S().Errorf("神必异常: %s", err)
			return
		}
		if err := os.MkdirAll(dir, 0755); err != nil {
			zap.S().Errorf("创建目录失败: %s", err)
			return
		}
	} else {
		defer s.Close()
		// 删除上一次
		names, err := s.Readdirnames(-1)
		if err != nil {
			zap.S().Errorf("神必异常: %s", err)
			return
		}
		for _, name := range names {
			err = os.RemoveAll(filepath.Join(dir, name))
			if err != nil {
				zap.S().Errorf("神必异常: %s", err)
				return
			}
		}
	}
	err = utils.UnZip(file.tmpPath, filepath.Join(dir, "show"))
	if err != nil {
		zap.S().Errorf("解压错误 %s\n", err.Error())
		return
	}
}

func initWorker(e *entity.MExperiment) {
	exp := expFile{
		eid:      e.ID,
		fileToDo: make(chan info, routineNums*2),
		exit:     make(chan struct{}, 1),
		endtime:  e.EndTime,
	}
	dirPath := fmt.Sprintf("%s%d/", global.GCoursePath, exp.eid)
	if _, err := os.Stat(dirPath); err != nil {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			zap.S().Errorf("创建目录失败: %s", err)
			return
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
}

// InitSubmitThread 初始化
func InitSubmitThread() {
	m = make(map[uint]*expFile)
	go clear()
	now := time.Now()
	var exps []entity.MExperiment
	global.GDB.Model(&entity.MExperiment{}).
		Where("end_time >= ?", now.Format(global.TimeTemplateSec)).
		Find(&exps)
	for _, j := range exps {
		initWorker(&j)
	}
}

// ToUnzip 发送解压
func ToUnzip(eid uint, gid string, file string) error {
	var e *expFile
	var ok bool
	mutex.RLock()
	e, ok = m[eid]
	mutex.RUnlock()

	if !ok {
		return errors.New("ToUnzip 初始化异常")
	}
	e.fileToDo <- info{
		gid:     gid,
		tmpPath: filepath.Join(global.GTmpPath, file),
	}
	return nil
}
