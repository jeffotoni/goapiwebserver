// Go Api server
// @jeffotoni
// 2019-01-04

package handler

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/jeffotoni/goapiwebserver/goapiserver/config"
	middle "github.com/jeffotoni/goapiwebserver/goapiserver/middleware"
	"github.com/jeffotoni/goapiwebserver/goapiserver/pkg/util"
)

// HTMLServer represents the web service that serves up HTML
type GoServerHttp struct {
	server *http.Server
	wg     sync.WaitGroup
}

func StartServer(cfg config.Config) *GoServerHttp {

	// DefaultServeMux
	mux := http.NewServeMux()

	// POST handler /api/v1/ping
	handlerApiPing := http.HandlerFunc(Ping)

	// handler ping
	mux.Handle(SetEndPoint().Ping, middle.Adapt(handlerApiPing,
		middle.Limit(),
		middle.MaxClients(config.MaxClients)))

	// POST handler /api/v1/token
	// Authorization: Basic key:keypass
	handlerApiToken := http.HandlerFunc(Token)

	// generate token jwt
	// handler token
	mux.Handle(SetEndPoint().PostToken, middle.Adapt(handlerApiToken,
		middle.Limit(),
		middle.MaxClients(config.MaxClients),
		middle.GtokenJwt()))

	// POST handler /api/v1/login
	// user and pass
	handlerApiTLogin := http.HandlerFunc(Login)

	// validate with method user jwt
	mux.Handle(SetEndPoint().GetLogin, middle.Adapt(handlerApiTLogin,
		middle.Limit(),
		middle.MaxClients(config.MaxClients),
		middle.AuthJwt()))

	// templates/index html
	// if you want to activate this handler, the directory templates
	// where the html file is located must
	// be sent next to the binary to work, as it needs to parse the html
	// mux.HandleFunc("/", tpl.ShowHtml)

	// this handler implements the version
	// that does not need the html file
	mux.Handle("/", http.HandlerFunc(homeHandler))

	// Create the HTML Server
	ApiServer := GoServerHttp{
		server: &http.Server{
			Addr:           cfg.Host,
			Handler:        mux,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		},
	}

	// Add to the WaitGroup for the listener goroutine
	ApiServer.wg.Add(1)

	// Start the listener
	go func() {

		Setenv(cfg)
		ApiServer.server.ListenAndServe()
		ApiServer.wg.Done()

	}()

	return &ApiServer
}

// Stop turns off the HTML Server
func (GoServerHttp *GoServerHttp) StopServer() error {

	// Create a context to attempt a graceful 6 second shutdown.
	const timeout = 6 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// optimized version
	util.Print("\n\033[0;31mServer Goapiserver:\033[0m \033[0;32mService stopped\033[0m\n")

	// Attempt the graceful shutdown by closing the listener
	// and completing all inflight requests
	if err := GoServerHttp.server.Shutdown(ctx); err != nil {
		// Looks like we timed out on the graceful shutdown. Force close.
		if err := GoServerHttp.server.Close(); err != nil {
			util.Print("\n\033[0;31mServer Goapiserver:\033[0m \033[0;32mService Error Close()" + err.Error() + "\033[0m\n")
			return err
		}
	}

	// Wait for the listener to report that it is closed.
	GoServerHttp.wg.Wait()
	util.Print("\n\033[0;31mServer Goapiserver:\033[0m \033[0;32mService Stopped")
	return nil
}

func Setenv(cfg config.Config) {

	// optimized version
	util.Print("\n\033[0;33mApi Server Run \033[0m \033[0;32mstarted:\033[0m \033[37mHost=" + config.HOST_SERVER + "\033[0m\n")
	util.Print("\033[0;31mInfo:\033[0m\n")
	util.Print("\033[0;32m")

	util.Print("     - export HOST_SERVER:   " + config.HOST_SERVER + "\n")
	util.Print("     - export PORT_SERVER:   " + config.PORT_SERVER + "\n")
	util.Print("     - export HOST_MAXBYTE:  " + config.HOST_MAXBYTE + "\n")
	util.Print("     - export WD_LEVEL:      " + config.WD_LEVEL + "\n")
	util.Print("     - export REQUEST_SEC:   " + config.REQUEST_SEC_STR2 + "\n")

	util.Print("\n\033[37mkeys for the client to communicate\033[0m\n")
	util.Print("     ---------------------------------------------------\n")
	util.Print("     - export X_KEY:                " + config.X_KEY + "\n")
	util.Print("     - export AUTHORIZATION: " + config.AUTHORIZATION + "\n")

	util.Print("\n\033[37mEnd Points\033[0m\n")
	util.Print("\033[0;33m[POST]\033[0m...........{public}................ -> \033[0;36m" + SetEndPoint().Ping)
	util.Print("\n")
	util.Print("\033[0;33m[GET]\033[0m............{private key}........... -> \033[0;36m" + SetEndPoint().PostToken)
	util.Print("\n")
	util.Print("\033[0;33m[POST]\033[0m...........{private token}......... -> \033[0;36m" + SetEndPoint().GetLogin)
	util.Print("\n")
	util.Print("\033[0;33m[POST]\033[0m...........{private token}......... -> \033[0;36m" + SetEndPoint().PostSignup)
	util.Print("\n")
	util.Print("\033[0;33m[POST]\033[0m...........{private token}......... -> \033[0;36m" + SetEndPoint().PostUpProfile)
	util.Print("\n")
	util.Print("\033[0;33m[POST]\033[0m...........{private token}......... -> \033[0;36m" + SetEndPoint().PostForgotPssword)
	util.Print("\n")
	util.Print("\033[0m\n")
}
