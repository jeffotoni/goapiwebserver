// Front-end in Go server
// @jeffotoni
// 04/01/2019

package handler

import (
	"net/http"
	"sync"
)

// HTMLServer represents the web service that serves up HTML
type FrontEndServer struct {
	server *http.Server
	wg     sync.WaitGroup
}
