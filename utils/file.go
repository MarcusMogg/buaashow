package utils

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
)

// SaveFile 将form中的文件保存到dst位置
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// DeleteFile 删除本地文件
func DeleteFile(file string) error {
	if err := os.Remove(file); err != nil {
		return errors.New("本地文件删除失败, err:" + err.Error())
	}
	return nil
}
