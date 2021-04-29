package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func addToZip(zw *zip.Writer, path, old, new string, fi os.FileInfo) error {
	// 通过文件信息，创建 zip 的文件信息
	fh, err := zip.FileInfoHeader(fi)
	if err != nil {
		return err
	}
	// 替换文件信息中的文件名
	// fh.Name = strings.TrimPrefix(path, string(filepath.Separator))
	if len(old) != 0 {
		fh.Name = strings.Replace(path, old, new, 1)
	}
	fh.Method = zip.Deflate

	if fi.IsDir() {
		fh.Name += "/"
	}

	// 写入文件信息，并返回一个 Write 结构
	w, err := zw.CreateHeader(fh)
	if err != nil {
		return err
	}

	// 检测，如果不是标准文件就只写入头信息，不写入文件数据到 w
	// 如目录，也没有数据需要写
	if !fh.Mode().IsRegular() {
		return nil
	}

	fr, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fr.Close()

	_, err = io.Copy(w, fr)
	return err
}

// ZipFiles 压缩文件
func ZipFiles(outName string, fileNames, olds, news []string) error {
	outFile, err := os.Create(outName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := zip.NewWriter(outFile)
	defer w.Close()

	for i := range fileNames {
		err := filepath.Walk(fileNames[i], func(path string, fi os.FileInfo, errBack error) error {
			if errBack != nil {
				return errBack
			}
			return addToZip(w, path, olds[i], news[i], fi)
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(dst string, file *zip.File) error {
	path := filepath.Join(dst, file.Name)

	// 如果是目录，就创建目录
	if file.FileInfo().IsDir() {
		err := os.MkdirAll(path, file.Mode())
		return err
	}

	// 获取到 Reader
	fr, err := file.Open()
	if err != nil {
		return err
	}
	defer fr.Close()
	// 创建要写出的文件对应的 Write
	fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer fw.Close()

	_, err = io.Copy(fw, fr)
	return err
}

// UnZip 将src文件解压到 dst文件夹下
func UnZip(src, dst string) (err error) {
	zr, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer zr.Close()

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	// 遍历 zr ，将文件写入到磁盘
	for _, file := range zr.File {
		err := unzipFile(dst, file)
		if err != nil {
			return err
		}
	}
	return nil
}
