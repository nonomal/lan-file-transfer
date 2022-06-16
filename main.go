package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
func main() {
	fmt.Println("请访问下面的链接:")
	showip()
	http.HandleFunc("/", uploadFileHandler)
	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("./data"))))
	http.ListenAndServe(":10000", nil)
}
func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	/**/
	fmt.Fprintln(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>多平台局域网内文件互传</title>
</head>
<body style="text-align: center;"> s
    <h1>多平台局域网内文件互传</h1>
    <br>
    <br>
    <form action="UploadFile.ashx" method="post" enctype="multipart/form-data">
    <input type="file" name="fileUpload" />
    <input type="submit" name="上传文件">
    </form>
        <br>
    <br>
        <br>
    <br>
    <a href="/file">文件下载</a>
</body>
</html>
        `)
	if r.Method == "POST" {
		file, handler, err := r.FormFile("fileUpload") //name的字段
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		check(err)
		newFile, err := os.Create("./data/" + handler.Filename)
		check(err)
		defer newFile.Close()
		if _, err := newFile.Write(fileBytes); err != nil {
			check(err)
			return
		}
		fmt.Println(" upload successfully:" + "./data/" + handler.Filename)
		w.Write([]byte("SUCCESS"))
	}
}
func showip() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String() + ":8080")
			}
		}
	}
}
