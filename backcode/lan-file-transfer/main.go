package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"lan-file-transfer/apps"
	"lan-file-transfer/common"
	"lan-file-transfer/config"
	"lan-file-transfer/router"
)

const (
	defaultPort = 9999
	dataDir     = "data"
)

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o lan-file-transfer_linux
//CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o lan-file-transfer_windows
//CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o lan-file-transfer_mac
func init() {
	port := 0
	flag.IntVar(&port, "port", defaultPort, fmt.Sprintf("设置服务器端口是(默认:%d)", defaultPort))
	// 寻找指定端口附近空闲的端口
	port = apps.FindFreePort(port)
	_config := &config.Config{
		ServerPort: port,
		DataDir:    dataDir,
	}
	config.Init(_config)
	//创建文件夹
	apps.CreateDir(common.CombinePath(false, apps.GetCurrentDirectory(), config.Get().DataDir))
}

func main() {
	flag.Parse()
	r := gin.Default()
	router.Router(r)
	// 启动一个协程，打开浏览器
	go func() {
		apps.OpenUrl()
	}()
	r.Run(fmt.Sprintf(":%d", config.Get().ServerPort))
}
