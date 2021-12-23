package image

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"ginTemplate/api"
	"ginTemplate/config"
	"ginTemplate/model/image"
	"ginTemplate/utils"
)

// bmp,jpg,png,tif,gif,pcx,tga,exif,fpx,svg,psd,cdr,pcd,dxf,ufo,eps,ai,raw,WMF,webp,avif,apng
var imageType = []string{"bmp",
	"jpeg",
	"jpg",
	"png",
	"tif",
	"gif",
	"pcx",
	"tga",
	"exif",
	"fpx",
	"svg",
	"psd",
	"cdr",
	"pcd",
	"dxf",
	"ufo",
	"eps",
	"ai",
	"raw",
	"WMF",
	"webp",
	"avif",
	"apng",
	"svg+xml",
}

// 上传图片
func DoUploadImage(c *api.Context) (int, string, interface{}) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), nil
	}
	var contentType = "unKnow"
	if len(header.Header["Content-Type"]) > 0 {
		contentType = header.Header["Content-Type"][0]
		typeList := strings.Split(contentType, "/")
		if len(typeList) >= 2 {
			if !utils.IsInStringArr(imageType, typeList[1]) {
				return api.IllegalArgument, c.Error("不支持的格式：" + typeList[1]).Error(), nil
			}
		}
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return api.TypeTranslateFail, c.Error(err).Error(), nil
	}

	var prefix string

	switch config.String("mode") {
	case "test":
		prefix = "http://wengyifan.com:8080/"
	case "release":
		prefix = "http://wengyifan.com:8080/"
	default:
		prefix = "http://wengyifan.com:8080/"
	}

	// 这里做一层hash
	hashc := md5.New()
	hashc.Write(content)
	bytes := hashc.Sum(nil)
	hashcode := hex.EncodeToString(bytes)
	

	// 查询图片是否存在
	imageInfo, err := image.GetImage(" hashCode = ?", hashcode)
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}
	if imageInfo != nil {
		res := map[string]string{
			"imageUrl": fmt.Sprintf(prefix+"image/getImage?id=%d", imageInfo.ID),
		}
		return 0, "success", res
	}


	id, err := image.AddImage(&image.Image{
		ImageName:   header.Filename,
		Size:        int(header.Size),
		ContentType: contentType,
		Image:       content,
		HashCode:    hashcode,
		Creator:     "yifanweng",
		CreateTime:  time.Now(),
	})
	if err != nil {
		return api.DatabaseError, c.Error(err).Error(), nil
	}


	res := map[string]string{
		"imageUrl": fmt.Sprintf(prefix+"image/getImage?id=%d", id),
	}

	return 0, "success", res
}

// 获取图片
func DoGetImage(c *api.Context) (int, string, interface{}) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return api.IllegalArgument, c.Error(err).Error(), ""
	}

	imageInfo, err := image.FindImage(" id = ?", id)
	if err != nil {
		return 0, c.Error(err).Error(), ""
	}
	if len(imageInfo) == 0 {
		return api.DatabaseError, c.Error("图片不存在").Error(), nil
	}

	c.Data(200, imageInfo[0].ContentType, imageInfo[0].Image)

	return 0, "success", nil
}
