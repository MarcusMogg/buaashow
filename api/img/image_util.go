package img

import (
	"buaashow/global"
	"buaashow/utils"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"path"

	"github.com/nfnt/resize"
)

// 支持的图像类型
// 图像类型
const (
	PNG  = "png"
	JPG  = "jpg"
	JPEG = "jpeg"
	GIF  = "gif"
)

var supportType map[string]int = map[string]int{
	PNG:  1,
	JPG:  1,
	JPEG: 1,
	GIF:  1,
}

type size struct {
	width  uint
	height uint
	suffix string
}

// var targetSize []size = []size{
// 	{width: 60, height: 45, suffix: "s"},
// 	{width: 130, height: 97, suffix: "m"},
// 	{width: 320, height: 240, suffix: "l"},
// }

func decode(file *multipart.FileHeader) (image.Image, string, error) {
	f, err := file.Open()
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	img, imgtype, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}
	_, ok := supportType[imgtype]
	if !ok {
		return nil, "", errors.New("not support type")
	}
	return img, imgtype, nil
}

func rename(fileType string) (string, error) {
	filePath := utils.NextID()
	_, err := os.Stat(path.Join(global.GImgPath, filePath+"."+fileType))
	if err == nil {
		return "", errors.New("file already exists")
	}
	return filePath, nil
}

func saveImage(file *multipart.FileHeader, imageF image.Image, fileType, dst string) error {
	out, err := os.Create(path.Join(global.GImgPath, dst+"."+fileType))
	if err != nil {
		return err
	}
	defer out.Close()

	switch fileType {
	case JPG, JPEG:
		if err := jpeg.Encode(out, imageF, nil); err != nil {
			return err
		}
	case PNG:
		if err := png.Encode(out, imageF); err != nil {
			return err
		}
	case GIF:
		f, err := file.Open()
		if err != nil {
			return err
		}
		defer f.Close()
		gifimg, err := gif.DecodeAll(f)
		if err != nil {
			return err
		}
		if err := gif.EncodeAll(out, gifimg); err != nil {
			return err
		}
	default:
		return errors.New("unknow err in saveImage")
	}
	return nil
}

// func resizeTFile(file *multipart.FileHeader, imageF image.Image, fileType, baseName string) ([]string, error) {
// 	var res []string = make([]string, 0)
// 	if fileType != GIF {
// 		start := time.Now()
// 		for _, i := range targetSize {
// 			if imageF.Bounds().Max.X <= int(i.width) {
// 				break
// 			}
// 			resFile := resize.Thumbnail(i.width, i.height, imageF, resize.Lanczos3)
// 			if err := saveImage(file, resFile, fileType, baseName+i.suffix); err != nil {
// 				return res, err
// 			}
// 			res = append(res, baseName+i.suffix+"."+fileType)
// 			zap.S().Debugf("次resize时间为%s", time.Now().Sub(start).String())
// 		}
// 		if err := saveImage(file, imageF, fileType, baseName); err != nil {
// 			return res, err
// 		}
// 		res = append(res, baseName+"."+fileType)
// 	} else {
// 		if err := saveImage(file, nil, fileType, baseName); err != nil {
// 			return res, err
// 		}
// 		res = append(res, baseName+"."+fileType)
// 	}
// 	return res, nil
// }

func resizeFile(file *multipart.FileHeader, imageF image.Image, fileType, baseName string, s size) (string, error) {
	var res string

	if fileType != GIF {
		if s.width != 0 || s.height != 0 {
			resFile := resize.Resize(s.width, s.height, imageF, resize.Lanczos3)
			if err := saveImage(file, resFile, fileType, baseName+s.suffix); err != nil {
				return res, err
			}
			res = baseName + s.suffix + "." + fileType
		} else {
			if err := saveImage(file, imageF, fileType, baseName); err != nil {
				return res, err
			}
			res = baseName + "." + fileType
		}
	} else {
		if err := saveImage(file, nil, fileType, baseName); err != nil {
			return res, err
		}
		res = baseName + "." + fileType
	}
	return res, nil
}
