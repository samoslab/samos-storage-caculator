package main

import (
	"net/http"
	"time"
	"os"
	"os/signal"
	"syscall"
	"context"
	"github.com/samoslab/samos-storage-caculator/router"
	"flag"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":8999", "监听地址和端口号，默认监听本机所有的IP的8999端口")
	flag.Parse()
}

func main() {
	router := &router.Router{}
	// 绑定server
	server := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 异步运行server
	go func() {
		server.ListenAndServe()
	}()

	// 监听系统信号
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for _ = range ch {
		server.Shutdown(context.TODO())
		break
	}
}
