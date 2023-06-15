package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func PPrint(v ...interface{}) {
	for _, v := range v {
		b, err := json.Marshal(v)
		if err != nil {
			fmt.Println(v)
			return
		}
		var out bytes.Buffer
		err = json.Indent(&out, b, "", "  ")
		if err != nil {
			fmt.Println(v)
			return
		}

		fmt.Println(out.String())
	}

}
func PFile(v interface{}, fileName string) {

	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		return
	}

	ioutil.WriteFile(fmt.Sprintf("./%s.txt", fileName), out.Bytes(), 0777)

}
