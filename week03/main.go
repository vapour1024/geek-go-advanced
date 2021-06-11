package main

// 1. 基于 errgroup 实现一个 http server 的启动和关闭 ，
// 以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func server1() error {
	err := http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
	log.Print(err)
	return err
}

func server2() error {
	err := http.ListenAndServe("127.0.0.1:8002", http.DefaultServeMux)
	log.Print(err)
	return err
}
func main() {
	//define signals
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	//create an errgroup
	g, ctx := errgroup.WithContext(context.Background())
	//start servers
	g.Go(server1)
	g.Go(server2)
	//receive signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				close(c)
			}
		}
	})
	//wait for server goroutines finished
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
