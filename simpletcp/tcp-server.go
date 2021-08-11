package main

import (
	"context"
	"github.com/golang/glog"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

const (
	PORT = ":8081"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	go processSignals(cancelFunc)

	if err := listenAndHandle(ctx); err != nil {
		log.Fatal(err)
	}
}

func processSignals(cancelFunc context.CancelFunc) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)

	for {
		sig := <-signalChan
		switch sig {

		case os.Interrupt:
			log.Println("Signal SIGINT is received, probably due to `Ctrl-C`, exiting ...")
			cancelFunc()
			return
		}
	}
}

func listenAndHandle(ctx context.Context) error {
	localAddr, err := net.ResolveTCPAddr("tcp", PORT)
	if err != nil {
		return err
	}

	l, err := net.ListenTCP("tcp", localAddr) // return (*TCPListener, error)
	if err != nil {
		return err
	}

	defer l.Close()
	log.Println("Start listening on the TCP socket", PORT, ".")

	for {
		select {
		case <-ctx.Done():
			log.Println("Stop listening on the TCP socket", PORT, ".")
			return nil

		default:
			if err := l.SetDeadline(time.Now().Add(time.Second * 10)); err != nil {
				return err
			}

			conn, err := l.Accept()
			if err != nil {
				if os.IsTimeout(err) {
					continue
				}
				return err
			}

			log.Println("New connection to the listening TCP socket", PORT, ".")

			go handleConnection(conn)
		}
	}
}

func handleConnection(conn net.Conn) error {
	for {
		time.Sleep(time.Second * 5)
		if _, err := io.Copy(conn, conn); err != nil {
			glog.Error(err.Error())
			return err
		}
		return nil
	}
}