// Front-end in Go server
// @jeffotoni
// 2019-01-04

package handler

import (
	"context"
	"net/http"
	"time"
)

import (
	"github.com/gorilla/mux"
	ctoken "github.com/jeffotoni/goapiwebserver/gofrontend/api/token"
	"github.com/jeffotoni/goapiwebserver/gofrontend/config"
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/util"
)

const (
	rlogin         = "/login"           //POST/GET
	rregister      = "/register"        // POST
	forgotpassword = "/forgot/password" // GET/POST
	admin          = "/admin"           //POST
	adminprofile   = "/admin/profile"   //POST/GET/PUT
	logout         = "/logout"          //POST
	// test destroy session
	destroy = "/destroy"
)

// Routes
// Start launches the HTML Server
func Start(cfg config.Config) *FrontEndServer {

	// Setup Context
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup Handlers
	router := mux.NewRouter()
	router.HandleFunc("/", LoginHandler)
	router.HandleFunc(rlogin, LoginHandler)
	router.HandleFunc(rregister, LoginHandlerRegister)
	router.HandleFunc(forgotpassword, ForgotPassHandlerRegister)
	router.HandleFunc(admin, AdminHandler)
	router.HandleFunc(adminprofile, AdminProfileHandler)
	router.HandleFunc(logout, LogoutHandler)

	// test destroy
	router.HandleFunc(destroy, DestroyHandler)

	router.PathPrefix("/web/static/").Handler(http.StripPrefix("/web/static/", http.FileServer(http.Dir("./web/static"))))

	// Create the HTML Server
	FrontEndServer := FrontEndServer{
		server: &http.Server{
			Addr:           cfg.Host,
			Handler:        router,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		},
	}

	ctoken.SetEnvKeys()

	// Add to the WaitGroup for the listener goroutine
	FrontEndServer.wg.Add(1)

	// Start the listener
	go func() {

		// optimized version
		util.Print("\n\033[0;33mServer FrontEnd:\033[0m \033[0;32mstarted:\033[0m \033[37mHost=" + cfg.Host + "\033[0m\n")
		util.Print("\033[0;31mINFO\033[0m\n")
		util.Print("\033[0;33m")
		util.Print("     ---------------------- API BACKEND ----------------------\n")
		util.Print("     - export API_HOST_SERVER:  " + ctoken.API_HOST_SERVER + "\n")
		util.Print("     - export API_X_KEY:        " + ctoken.API_X_KEY + "\n")
		util.Print("     - export AUTHORIZATION:  	" + ctoken.AUTHORIZATION_DEFAULT + "\n\n")
		util.Print("\033[0m\n")

		util.Print("\033[0;32m")
		util.Print("     ---------------------- API FRONTEND --------------------\n")
		util.Print("     - export HOST_SERVER:  " + config.HOST_SERVER + "\n")
		util.Print("     - export PORT_SERVER:  " + config.PORT_SERVER + "\n")
		util.Print("     - export HOST_MAXBYTE:  " + config.HOST_MAXBYTE + "\n")

		util.Print("     ---------------------------------------------------\n")
		util.Print("     - export X_KEY:                " + config.X_KEY + "\n")
		util.Print("     - export AUTHORIZATION: Bearer " + config.AUTHORIZATION + "\n")

		util.Print("     ---------------------------------------------------\n")
		util.Print("\033[37mStart Server\033[0m\n")
		util.Print("      [POST] - \033[0;33m" + rlogin + "\033[0m\n")
		util.Print("      [POST] - \033[0;33m" + rregister + "\033[0m\n")
		util.Print("      [POST] - \033[0;33m" + forgotpassword + "\033[0m\n")
		util.Print("      [POST] - \033[0;33m" + admin + "\033[0m\n")
		util.Print("      [POST] - \033[0;33m" + logout + "\033[0m")
		util.Print("\033[0m\n")

		util.Print("     ---------------------------------------------------\n")
		util.Print("\033[37mPublic directories\033[0m\n")
		util.Print("      - web/assets/img\n")
		util.Print("      - web/static/css\n")
		util.Print("      - web/static/js\n")
		util.Print("      - templates\n")

		FrontEndServer.server.ListenAndServe()
		FrontEndServer.wg.Done()
	}()

	return &FrontEndServer
}

// Stop turns off the HTML Server
func (FrontEndServer *FrontEndServer) Stop() error {

	// Create a context to attempt a graceful 6 second shutdown.
	const timeout = 6 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// optimized version
	util.Print("\n\033[0;31mServer FrontEnd:\033[0m \033[0;32mService stopped\033[0m\n")

	// Attempt the graceful shutdown by closing the listener
	// and completing all inflight requests
	if err := FrontEndServer.server.Shutdown(ctx); err != nil {
		// Looks like we timed out on the graceful shutdown. Force close.
		if err := FrontEndServer.server.Close(); err != nil {
			util.Print("\n\033[0;31mServer FrontEnd:\033[0m \033[0;32mService Error Close()" + err.Error() + "\033[0m\n")
			return err
		}
	}

	// Wait for the listener to report that it is closed.
	FrontEndServer.wg.Wait()
	util.Print("\n\033[0;31mServer FrontEnd:\033[0m \033[0;32mService Stopped")
	return nil
}
