package apps

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// commands 打开浏览器 不同环境命令
var openURLCommands = map[string]string{
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

// OpenUrl  打开 本地ip+端口 浏览器
func OpenUrl(serverPort int) error {
	uri := getURL(serverPort)
	//runtime.GOOS
	run, ok := openURLCommands[runtime.GOOS]
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

// PortInUse
// 传入查询的端口号
// 返回端口号对应的进程PID，若没有找到相关进程，返回-1
func PortInUse(portNumber int) int {
	res := -1
	var outBytes bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmdStr := fmt.Sprintf("netstat -ano -p tcp | findstr %d", portNumber)
		cmd = exec.Command("cmd", "/c", cmdStr)
	} else if runtime.GOOS == "linux" {
		cmdStr := fmt.Sprintf("netstat -anp |grep %d", portNumber)
		cmd = exec.Command(cmdStr)
	}

	cmd.Stdout = &outBytes
	cmd.Run()
	resStr := outBytes.String()
	r := regexp.MustCompile(`\s\d+\s`).FindAllString(resStr, -1)
	if len(r) > 0 {
		pid, err := strconv.Atoi(strings.TrimSpace(r[0]))
		if err != nil {
			res = -1
		} else {
			res = pid
		}
	}
	return res
}

// FindFreePort
// 寻找附近的空闲端口
func FindFreePort(portNumber int) int {
	if PortInUse(portNumber) == -1 {
		return portNumber
	}
	temp := 1
	for {
		if PortInUse(portNumber+temp) == -1 {
			return portNumber + temp
		}
		if PortInUse(portNumber-temp) == -1 {
			return portNumber - temp
		}
		temp++
	}
	return portNumber

}
