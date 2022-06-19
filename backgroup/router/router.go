package router

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"lan-file-transfer/apps"
	"lan-file-transfer/dist"
	"net/http"
)

func Router(r *gin.Engine) {
	//r.LoadHTMLGlob("dist/*")
	//r.LoadHTMLFiles("dist/static/*/*")
	//r.Static("/static", "./dist/static")
	//r.StaticFile("/", "./dist/index.html")

	fsCss := assetfs.AssetFS{Asset: dist.Asset, AssetDir: dist.AssetDir, AssetInfo: nil, Prefix: "dist1/static/css", Fallback: "index.html"}

	fsJs := assetfs.AssetFS{Asset: dist.Asset, AssetDir: dist.AssetDir, AssetInfo: nil, Prefix: "dist1/static/js", Fallback: "index.html"}

	fsFonts := assetfs.AssetFS{Asset: dist.Asset, AssetDir: dist.AssetDir, AssetInfo: nil, Prefix: "dist1/static/fonts", Fallback: "index.html"}

	fsImg := assetfs.AssetFS{Asset: dist.Asset, AssetDir: dist.AssetDir, AssetInfo: nil, Prefix: "dist1/static/img", Fallback: "index.html"}

	r.StaticFS("/static/css", &fsCss)
	r.StaticFS("/static/fonts", &fsFonts)
	r.StaticFS("/static/img", &fsImg)
	r.StaticFS("/static/js", &fsJs)

	r.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		indexHtml, _ := dist.Asset("dist1/index.html")
		_, _ = c.Writer.Write(indexHtml)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Flush()
	})
	//上传文件
	r.POST("/api/uploadFile", apps.UploadFile)
	//获取文件列表
	r.GET("/api/files", apps.GetPageListFile)
	//删除文件
	r.DELETE("/api/deleteFile", apps.DeleteFile)
	//静态文件
	r.StaticFS("/data", http.Dir("./data"))
}
