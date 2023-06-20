package utils

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

// 图片文件转base64 字符串
func Base64FileToStr(imgPath string) string {
	f, _ := os.Open(imgPath)
	all, _ := ioutil.ReadAll(f)
	imgbase64 := base64.StdEncoding.EncodeToString(all)
	return imgbase64
}

// base64转图片
func Base64StrToFile(imgBase64 string, vcodePath string) {

	tmp, _ := base64.StdEncoding.DecodeString(imgBase64)
	//buffer输出到jpg文件中（不做处理，直接写到文件）
	f, _ := os.OpenFile(vcodePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(tmp)
}

// base64.StdEncoding.DecodeString("./test.jpg")
// base64.StdEncoding.Encode(bufstore, ff)

// ioutil.WriteFile("./output.jpg", ddd, 0666)
// ioutil.ReadFile("output2.jpg")

//

// buffer
