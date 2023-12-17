package main

import (
	"cc_server/common"
	"cc_server/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	// 初始化配置文件
	common.InitConfig()
	// 初始化数据库
	common.InitDB()
}

func main() {
	fmt.Println(fmt.Printf("%+v", common.Conf.Mysql))
	r := routes.InitRoute()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Fatal("Server shutdown gracefully")
}
