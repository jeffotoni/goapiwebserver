// Front-end in Go server
// @jeffotoni
// 2019-01-04

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
