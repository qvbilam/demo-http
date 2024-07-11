package main

import (
	"encoding/json"
	"fmt"
	"github.com/namsral/flag"
	"net/http"
	"strconv"
)

var ServerPort int64

func main() {
	version := "1.1.1"
	type response struct {
		Version      string `json:"version"`
		ServerName   string `json:"server_name"`
		ServerSecret string `json:"server_secret"`
	}
	// 获取系统变量: Dockerfile中定义的env(需要注意在系统变量中为字符串)
	port := flag.String("port", "9001", "demo server port")
	serverName := flag.String("server_name", "default name", "demo server name")
	serverSecret := flag.String("server_secret", "default secret", "demo server secret")

	flag.Parse() // 配置映射到变量中

	pS := *port
	pI, _ := strconv.Atoi(pS)
	ServerPort = int64(pI)

	// 设置路由
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		res := response{
			Version:      version,
			ServerName:   *serverName,
			ServerSecret: *serverSecret,
		}
		resJson, _ := json.Marshal(&res)

		_, _ = writer.Write(resJson)
	})

	// 启动服务
	_ = http.ListenAndServe(fmt.Sprintf(":%d", ServerPort), nil)
}
