package main

import (
	"fmt"

	"github.com/micronuths/gotools/utils"
	"github.com/micronuths/gotools/utils/ocr"
	_ "github.com/micronuths/gotools/utils/ocr/aliyun"
	_ "github.com/micronuths/gotools/utils/ocr/yunma"
)

func main() {
	ocrEngineTags := []string{"AliyunOCR", "YunmaOCR"}
	imgbase64 := utils.Base64FileToStr(`D:\GOPATH\src\github.com\micronuths\gotools\utils\ocr\example\img\GGAY.png`)
	vcode := ocr.Recognize(imgbase64, ocrEngineTags[0])
	fmt.Println(vcode)

}
