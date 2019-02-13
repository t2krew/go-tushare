package tushare

import (
	"github.com/ddliu/go-httpclient"
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"strings"
)

const API = "http://api.tushare.pro"

type Params map[string]string

func (t *TuShare) Request(name string, param Params, fields ...[]string) (body map[string]interface{}, res *httpclient.Response, err error) {
	data := map[string]interface{}{
		"api_name": name,
		"token":    t.token,
		"params":   param,
	}
	if len(fields) > 0 {
		f := strings.Join(fields[0], ",")
		data["fields"] = f
	}
	res, err = httpclient.PostJson(API, data)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	bodyEnc, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = jsoniter.Unmarshal(bodyEnc, &body)
	return
}
