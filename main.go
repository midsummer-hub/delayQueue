package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/delayQueue/api"
	"github.com/delayQueue/config"
	"github.com/delayQueue/daemon"
	"github.com/delayQueue/queueRpc"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", config.Address)
	if err != nil {
		panic(errors.New("listen:" + config.Address + "failed"))
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go daemon.Detect(ctx, wg)
	wg.Add(1)
	go daemon.Publish(ctx, wg)
	wg.Add(1)
	go func() {
		s := api.Server{}
		gRPCServer := grpc.NewServer()
		queueRpc.RegisterDelayQueueServer(gRPCServer, s)
		log.Println("grpc Server listening ", config.Address)
		go func() {
			time.Sleep(time.Second)
			<-ctx.Done()
			log.Println("grpc server try to stop ...")
			gRPCServer.GracefulStop()
			log.Println("grpc server stopped")
			wg.Done()
		}()
		if err = gRPCServer.Serve(l); err != nil {
			panic(err)
		}
		si := <-sig
		log.Println("server rect stop signal", si)
		cancel()
		wg.Wait()
		log.Println("all darmon stopped , now server stop")
	}()
	si := <-sig
	log.Println("server recv stop signal", si)
	cancel()
	log.Println("server  waiting all deamon stopping...")
	wg.Wait()
	log.Println("daemon all stopped, now server stopped. bye~")
}
