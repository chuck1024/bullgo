package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/chuck1024/gd"
	"time"
)

type TestHttpClientReq struct {
	Data string
}

func main() {
	req := &TestHttpClientReq{Data: "chuck"}
	_, body, err := gd.NewHttpClient().Timeout(3 * time.Second).Post("http://127.0.0.1:10240/test").Send(req).End()
	if err != nil {
		fmt.Printf("occur error:%s\n", err)
		return
	}
	ret, _:= simplejson.NewJson([]byte(body))
	fmt.Println(ret.Get("result").Get("Ret"))
}
