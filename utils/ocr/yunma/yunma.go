package yunma

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
	ocr.Register("YunmaOCR", New())
}

// New returns a new acfun bangumi extractor
func New() ocr.OcrEngine {
	return &ocrEngine{
		Info: "YunmaOCR Engine",
	}
}

// 云码控制台：https://www.jfbym.com/test/100.html
// http://api.jfbym.com/api/YmServer/customApi
// api token: cc05WEWO6pguZRUrVIehfrvLB7toX585cAGCmOAHeO4
// 账号：18181971727 密码：sj18181971727
// {"code":0,"data":"BWPM","time":0.04173636436462402,"unique_code":"76cd4ab4e30f2a69a62d8fdf2ed53556"}
func (o *ocrEngine) Recognize(imgbase64 string) string {
	const api = "http://api.jfbym.com/api/YmServer/customApi"
	// const api = "https://www.jfbym.com/api/YmServer/testCustomApi"
	const token = "cc05WEWO6pguZRUrVIehfrvLB7toX585cAGCmOAHeO4"
	client := http.DefaultClient

	if imgbase64 == "" {
		panic("QcrEngine.ImgBase64 is None, please set ImgBase64 value.")
	}
	resp, err := client.PostForm(api, url.Values{
		"image": {imgbase64},
		"token": {token},
		"type":  {"10110"},
	})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	var filter struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
		Data struct {
			Code       int     `json:"code"`
			Data       string  `json:"data"`
			Time       float64 `json:"time"`
			UniqueCode string  `json:"unique_code"`
		} `json:"data"`
	}
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	json.Unmarshal(rawBody, &filter)
	// utils.PPrint(resp.Text())
	r := o.querryMayunOCRSpareCount()
	fmt.Println("码云识别剩余积分", r)
	return o.replace(filter.Data.Data)
}
func (o *ocrEngine) querryMayunOCRSpareCount() string {
	const api = "http://api.jfbym.com/api/YmServer/getUserInfoApi"
	// const api = "https://www.jfbym.com/api/YmServer/testCustomApi"
	const token = "cc05WEWO6pguZRUrVIehfrvLB7toX585cAGCmOAHeO4"

	client := http.DefaultClient

	resp, err := client.PostForm(api, url.Values{
		"token": {token},
		"type":  {"score"},
	})
	if err != nil {
		panic(err)
	}
	var filter struct {
		Data struct {
			Score string `json:"score"`
		} `json:"data"`
	}
	txt, err := ioutil.ReadAll(resp.Body)
	// fmt.Println("yunma,txt=", string(txt))
	json.Unmarshal(txt, &filter)
	// utils.PPrint(resp.Text())
	return filter.Data.Score
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
