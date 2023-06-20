package aliyun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/micronuths/gotools/utils/ocr"
)

type ocrEngine struct {
	Info string
}

func init() {
	ocr.Register("AliyunOCR", New())
}

// New returns a new acfun bangumi extractor
func New() ocr.OcrEngine {
	return &ocrEngine{
		Info: "AliyunOCR Engine",
	}
}
func (o *ocrEngine) Recognize(imgbase64 string) string {
	// # OCR控制台：https://market.console.aliyun.com/#/autoRepurchase?_k=uuam3w
	api := "https://gjbsb.market.alicloudapi.com/ocrservice/advanced"

	// aliyunOCRApi := "http://127.0.0.1:8080/post"
	//api【168/500】
	apiConfig := map[string]string{
		"AppKey":    "204177193",
		"AppSecret": "BRGaHUr7IWbLs9v94KIwD0HuV5l8Pcn0",
		"AppCode":   "d9b3eb6ec5fd4e6280fdb3f2719d30ca",
	}

	client := http.DefaultClient
	proxy, _ := url.Parse("http://127.0.0.1:8888")
	client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	if imgbase64 == "" {
		panic("QcrEngine.ImgBase64 is None, please set ImgBase64 value.")
	}
	rawBody := strings.NewReader(fmt.Sprintf(`{"img":"%s"}`, imgbase64))
	req, err := http.NewRequest("POST", api, rawBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "APPCODE "+apiConfig["AppCode"])
	// resp, err := req.PostJson(api, data)

	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()

	// utils.PPrint(resp.Text())
	var filter struct {
		Content string `json:"content"`
	}
	txt, err := ioutil.ReadAll(resp.Body)
	// fmt.Println("print in aliyun ocr txt=", string(txt))
	json.Unmarshal(txt, &filter)
	return o.replace(filter.Content)
}

func (o *ocrEngine) replace(vcode string) string {
	for old, new := range map[string]string{
		" ": "",
		"0": "O",
		"1": "I",
	} {
		vcode = strings.ReplaceAll(vcode, old, new)
	}
	// 变成大写
	vcode = strings.ToUpper(vcode)
	return vcode
}
