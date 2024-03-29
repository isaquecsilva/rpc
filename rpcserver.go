package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
)

var listener net.Listener

func gracefulShutdown() {
	sig := make(chan os.Signal)
	defer close(sig)
	signal.Notify(sig, os.Interrupt)
	signalType := <- sig
	println("Signal received", signalType.String())

	if listener != nil {
		listener.Close()
	}
	println("goodbye.")
	os.Exit(0)
}

type Worker struct {}

func (w *Worker) ListDirFiles(nop *int, reply *[]string) error {
	log.Printf("RPC-CALL: { args: %d }\n", *nop)

	if direntry, err := os.ReadDir("."); err != nil {
		return err
	} else {
		dirs := make([]string, len(direntry))

		for _, file := range direntry {
			dirs = append(dirs, file.Name())
		}

		*reply = dirs		
	}

	return nil
}

func rpcConfiguration() {
	var worker *Worker = new(Worker)
	rpc.Register(worker)
	rpc.HandleHTTP()
}

func main() {
	// Setting Graceful Shutdown listener
	go gracefulShutdown()

	// Configuring RPC Server
	rpcConfiguration()

	var err error
	listener, err = net.Listen("tcp", "localhost:1234")
	checkError(err)

	println("RPC-Server running at port 1234")
	err = http.Serve(listener, nil)
	if err != nil {
		println(err.Error())
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}