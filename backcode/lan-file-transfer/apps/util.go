package apps

import (
	"fmt"
	"lan-file-transfer/config"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	windows       = "windows"
	darwin        = "darwin"
	linux         = "linux"
	httpLocalHost = "http://localhost"
)

// commands 打开浏览器 不同环境命令
var openURLCommands = map[string]string{
	windows: "cmd /c start ",
	darwin:  "open ",
	linux:   "xdg-open ",
}

//GetLocalIps 获取本机ip集合
func GetLocalIps() []string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ips := make([]string, 0)
	//先添加192.168.1. 开头的ip(保持顺序)
	for _, address := range addrList {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil &&
			strings.Index(ipNet.IP.String(), "192.168.1.") >= 0 {
			ips = append(ips, ipNet.IP.String())
			break
		}
	}
	//再添加除 192.168. 开头以外的ip
	for _, address := range addrList {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil &&
			strings.Index(ipNet.IP.String(), "192.168.") >= 0 && strings.Index(ipNet.IP.String(), "192.168.1.") < 0 {
			ips = append(ips, ipNet.IP.String())
		}
	}
	//再添其它的ip
	for _, address := range addrList {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil &&
			strings.Index(ipNet.IP.String(), "192.168.") < 0 {
			ips = append(ips, ipNet.IP.String())
		}
	}
	return ips

}

// OpenUrl  打开 本地ip+端口 浏览器
func OpenUrl() error {
	url := fmt.Sprintf("%s:%d", httpLocalHost, config.Get().ServerPort)
	//runtime.GOOS
	run, ok := openURLCommands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	//exec.Command
	run = run + url
	cmds := strings.Split(run, " ")
	cmd := exec.Command(cmds[0], cmds[1:]...)
	//cmd.Start
	fmt.Printf("exec commad :[%s]", run)
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

// GetCurrentDirectory 获取当前应用程序的路径
func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}

// PortInOpen
// 传入查询的端口号
// 返回端口号对应的进程PID，若没有找到相关进程，返回-1
func PortInOpen(port int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	defer ln.Close()
	return true
}

// FindFreePort
// 寻找附近的空闲端口
func FindFreePort(portNumber int) int {
	if PortInOpen(portNumber) {
		return portNumber
	}
	temp := 1
	for {
		if PortInOpen(portNumber + temp) {
			return portNumber + temp
		}
		if PortInOpen(portNumber - temp) {
			return portNumber - temp
		}
		temp++
	}
}
