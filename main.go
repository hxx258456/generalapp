package main

import (
	"generalapp/pkg/sdkpool"
	"generalapp/router"
)

func main() {
	// 添加连接池释放
	defer func() {
		for _, v := range sdkpool.SdkPoll {
			v.GW.Close()
		}
	}()

	router.InitRouter()
}
