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

	// Creating limiter for all handlers
	// or one for each handler. Your choice.
	// This limiter basically says: allow at most NewLimiter request per 1 second.
	//limiter := tollbooth.NewLimiter(1, nil)

	// Limit only GET and POST requests.
	//limiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}).SetMethods([]string{"GET", "POST"})

	// DefaultServeMux
	mux := http.NewServeMux()

	// POST handler /v1/api/ping
	handlerApiPing := http.HandlerFunc(Ping)

	// http.Handle("/", Adapt(indexHandler, AddHeader("Server", "Mine"),
	//                                     CheckAuth(providers),
	//                                     CopyMgoSession(db),
	//                                     Notify(logger),
	//                                   )

	mux.Handle(SetEndPoint().Ping, middle.Adapt(handlerApiPing,
		middle.MaxClients(config.MaxClients)))

	//mux.Handle(confserv.Ping, tollbooth.LimitFuncHandler(limiter, HandlerFuncAuth(HandlerValidate, MaxClients(handlerApiPing, config.MaxClients))))

	// template index html
	// mux.HandleFunc("/", tpl.ShowHtml)
	mux.HandleFunc("/", homeHandler)

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

	util.Print("\n\033[37mkeys for the client to communicate\033[0m\n")
	util.Print("     ---------------------------------------------------\n")
	util.Print("     - export X_KEY:                " + config.X_KEY + "\n")
	util.Print("     - export AUTHORIZATION: " + config.AUTHORIZATION + "\n")

	util.Print("\n\033[37mEnd Points\033[0m\n")
	util.Print("\033[0;33m[POST]\033[0m....................\033[0;36m" + SetEndPoint().Ping)
	util.Print("\n")
	util.Print("\033[0;33m[POST/GET/PUT]\033[0m............\033[0;36m" + SetEndPoint().PostToken)
	util.Print("\n")
	util.Print("\033[0;33m[POST]\033[0m....................\033[0;36m" + SetEndPoint().PostLogin)
	util.Print("\n")
	util.Print("\033[0m\n")
}
