package apps

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

func UploadFile(g *gin.Context) {
	file, err := g.FormFile("file")
	if err != nil {
		g.String(500, "上传文件出错")
	}
	g.SaveUploadedFile(file, "./data/"+file.Filename)

	g.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "上传成功",
	})
}

func DeleteFile(g *gin.Context) {
	fileName, exist := g.GetQuery("fileName") //获取查询关键字
	if !exist {
		g.JSON(http.StatusOK, map[string]interface{}{
			"code":    400,
			"message": "请选择文件名",
		})
	}
	err := os.Remove("./data/" + fileName)
	if err != nil {
		g.JSON(http.StatusNotFound, map[string]interface{}{
			"code": 404,
			"msg":  "删除失败",
		})

	} else {
		g.JSON(http.StatusOK, map[string]interface{}{
			"code": 200,
			"msg":  "删除成功",
		})
	}
}

func GetPageListFile(g *gin.Context) {
	key, exist := g.GetQuery("key") //获取查询关键字
	if !exist {
		key = ""
	}
	pageIndexStr, exist := g.GetQuery("pageIndex")
	if !exist {
		pageIndexStr = "1"
	}
	pageSizeStr, exist := g.GetQuery("pageSize")
	if !exist {
		pageSizeStr = "10"
	}
	path := "./data"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	//按 key 过滤
	newFiles := make([]os.FileInfo, len(files))
	copy(newFiles, files)
	if key != "" {
		linq.From(files).Where(func(i interface{}) bool {
			file := i.(os.FileInfo)
			return strings.Index(file.Name(), key) >= 0
		}).ToSlice(&newFiles)
	}
	//排序
	sort.Sort(ByModTime(newFiles))

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {

	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {

	}

	if pageIndex < 1 {
		g.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  "pageIndex不能小于1",
		})
	}
	total := len(newFiles)
	data := make([]string, 0)
	length := len(newFiles)
	for i := 0; i < pageSize; i++ {
		if length > pageSize*(pageIndex-1)+i {
			name := newFiles[pageSize*(pageIndex-1)+i].Name()
			data = append(data, name)
		}
	}
	g.JSON(http.StatusOK, map[string]interface{}{
		"data":  data,
		"total": total,
	})
}

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().After(fis[j].ModTime())
}
