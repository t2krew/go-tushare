package tushare

import (
	"log"
	"os"
	"testing"
)

var ts *TuShare

var token = ""

func init() {
	envToken := os.Getenv("TOKEN")
	if envToken != "" {
		token = envToken
	}
}

func TestMain(m *testing.M) {
	ts = New(token)
	os.Exit(m.Run())
}

func TestTuShare_Request(t *testing.T) {
	var param = Params{
		"list_status": "L",
	}
	var fields = []string{"ts_code", "name"}
	body, res, err := ts.Request("stock_basic", param, fields)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("request error, statusCode %d", res.StatusCode)
	}

	log.Println(body)
}
