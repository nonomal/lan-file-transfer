package router

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"lan-file-transfer/apps"
	"lan-file-transfer/asset"
	"lan-file-transfer/common"
	"lan-file-transfer/config"
	"net/http"
)

const (
	indexHtml = "index.html"
	dist      = "dist"
	static    = "static"
	css       = "css"
	js        = "js"
	fonts     = "fonts"
	img       = "img"
)

func Router(r *gin.Engine) {
	// 执行：go-bindata -o asset.go dist/...
	fsCss := assetfs.AssetFS{Asset: asset.Asset, AssetDir: asset.AssetDir, AssetInfo: nil, Prefix: common.CombinePath(false, dist, static, css), Fallback: indexHtml}
	fsJs := assetfs.AssetFS{Asset: asset.Asset, AssetDir: asset.AssetDir, AssetInfo: nil, Prefix: common.CombinePath(false, dist, static, js), Fallback: indexHtml}
	fsFonts := assetfs.AssetFS{Asset: asset.Asset, AssetDir: asset.AssetDir, AssetInfo: nil, Prefix: common.CombinePath(false, dist, static, fonts), Fallback: indexHtml}
	fsImg := assetfs.AssetFS{Asset: asset.Asset, AssetDir: asset.AssetDir, AssetInfo: nil, Prefix: common.CombinePath(false, dist, static, img), Fallback: indexHtml}

	r.StaticFS(common.CombinePath(true, static, css), &fsCss)
	r.StaticFS(common.CombinePath(true, static, fonts), &fsFonts)
	r.StaticFS(common.CombinePath(true, static, img), &fsImg)
	r.StaticFS(common.CombinePath(true, static, js), &fsJs)

	r.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		indexPath, _ := asset.Asset(common.CombinePath(false, dist, indexHtml))
		_, _ = c.Writer.Write(indexPath)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Flush()
	})
	//上传文件
	r.POST("/api/uploadFile", apps.UploadFile)
	//获取文件列表
	r.GET("/api/getPageListFile", apps.GetPageListFile)
	//删除文件
	r.DELETE("/api/deleteFile", apps.DeleteFile)
	//数据
	r.StaticFS(common.CombinePath(true, config.Get().DataDir), http.Dir(common.CombinePath(false, apps.GetCurrentDirectory(), config.Get().DataDir)))
	//获取url 地址
	r.GET("/api/getLocalUrls", apps.GetLocalUrls)
}
