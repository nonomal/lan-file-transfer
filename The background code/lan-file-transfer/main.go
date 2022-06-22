package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"lan-file-transfer/apps"
	"lan-file-transfer/config"
	"lan-file-transfer/router"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	config.Set(&config.Config{
		ServerPort: 9999,
	})
	flag.IntVar(&config.Get().ServerPort, "port", 9999, "设置服务器端口是（默认是9999）")
}

func main() {

	flag.Parse()
	//寻找通过指定端口寻找附近空闲端口
	config.Get().ServerPort = apps.FindFreePort(config.Get().ServerPort)
	apps.CreateDir("./data")
	r := gin.Default()
	r.Use(Cors())
	router.Router(r)
	go func() {
		time.Sleep(1000 * time.Millisecond)
		apps.OpenUrl(config.Get().ServerPort)
	}()
	r.Run(":" + strconv.Itoa(config.Get().ServerPort))

}

// Cors 跨域设置
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		context.Next()
	}
}
