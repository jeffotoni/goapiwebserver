// Go Api server
// @jeffotoni
// 2019-01-04
//
// Api server was implemented to use jwt for auth
// some middleware was created

package main

import (
	"os"
	"os/signal"
	"strconv"

	"github.com/jeffotoni/goapiwebserver/goapiserver/config"
	handler "github.com/jeffotoni/goapiwebserver/goapiserver/handler"
	"github.com/jeffotoni/goapiwebserver/goapiserver/pkg/util"
)

// set conf
var confserv = handler.SetEndPoint()

// init
func init() {
	// set config
	config.Setenv()
}

func main() {

	// convert string to int
	HOST_MAXBYTE, _ := strconv.Atoi(config.HOST_MAXBYTE)

	// conf server
	serverCfg := config.Config{
		Host:           config.HOST_SERVER,
		ReadTimeout:    config.READTIMEOUT,
		WriteTimeout:   config.WRITETIMEOUT,
		MaxHeaderBytes: HOST_MAXBYTE,
	}

	// star server
	hServer := handler.StartServer(serverCfg)

	// stop server
	defer hServer.StopServer()

	// channel with notify
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	util.Print("\nmain : Shutting down goapiserver\n")
}
