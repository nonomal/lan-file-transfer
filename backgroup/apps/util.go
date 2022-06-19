package apps

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// commands 执行程序
var commands = map[string]string{
	"windows": "cmd /c start ",
	"darwin":  "open ",
	"linux":   "c-open ", //eog -w
}

func getURL(serverPort int) string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	url := ""
	for _, address := range addrList {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil &&
			strings.Index(ipNet.IP.String(), "192.168.1") >= 0 {
			url = "http://" + ipNet.IP.String() + ":" + strconv.Itoa(serverPort)
			break
		}
	}
	return url

}

// OpenUrl 打开浏览器
func OpenUrl(serverPort int) error {
	uri := getURL(serverPort)
	//runtime.GOOS
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	//exec.Command
	run = run + uri
	cmds := strings.Split(run, " ")
	cmd := exec.Command(cmds[0], cmds[1:]...)
	//cmd.Start
	fmt.Println("[CommandAs]", cmds)
	return cmd.Start()
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 创建文件夹
func CreateDir(path string) {
	exist, _ := PathExists(path)
	if !exist {
		os.Mkdir(path, os.ModePerm)
	}
}
