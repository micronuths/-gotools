package ocr

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

var QcrEngineMap = make(map[string]OcrEngine)

// Register registers an Extractor.
func Register(hostTag string, ocr OcrEngine) {
	QcrEngineMap[hostTag] = ocr
}

// Extract is the main function to extract the data.
func Recognize(imgbase64 string, ocrEngineTag string) string {
	extractor := QcrEngineMap[ocrEngineTag]
	if extractor == nil {
		return "没有匹配到extractor"
		// extractor = extractorMap[""]
	}
	vcode := extractor.Recognize(imgbase64)
	pngPath := fmt.Sprintf("./png/%s", vcode)
	Base64StrToFile(imgbase64, pngPath)
	return vcode
}

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
