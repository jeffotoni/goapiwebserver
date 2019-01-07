// Front-end in Go server
// @jeffotoni
// 2019-01-04
//
// Always searching for the best optimization in Go, I used the lib bufio
// to write directly to stdou and I eliminated the use of lib fmt.
// The use of the templates made a parse so that the htmls were declared inside
// go without having to make a Parsefiles whenever you need to read the html for web client.

package main

// native libs
import (
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/jeffotoni/goapiwebserver/gofrontend/config"
	"github.com/jeffotoni/goapiwebserver/gofrontend/handler"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/util" // external libs
)

// start templates
func init() {

	// set config
	config.Setenv()
}

func main() {

	// convert string to int
	HOST_MAXBYTE, _ := strconv.Atoi(config.HOST_MAXBYTE)

	serverCfg := config.Config{
		Host:           config.HOST_SERVER,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: HOST_MAXBYTE,
	}

	htmlServer := handler.Start(serverCfg)
	defer htmlServer.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	util.Print("\nmain : Shutting down frontend\n")

}
