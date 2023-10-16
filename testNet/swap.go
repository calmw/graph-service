package main

import (
	"graph-service/services"
	"graph-service/tasks"
)

func main() {
	services.InitTestNet()
	// Approve
	//blockchain.ApproveTox("0xC0f98D355dc08f8187cc742A2fCb00e62De60E2F")
	// 部署，初始化
	//services.InitSwap()

	// event task
	tasks.Start()
}
