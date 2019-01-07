// Go Api server
// @jeffotoni
// 2019-01-04

package logg

import (
	"log"
	"time"
)

func Start() (s1 time.Time) {

	s1 = time.Now()
	return
}

func Show(endpoint, metodo, status string, s1 time.Time) {

	s2 := time.Now()
	sub := s2.Sub(s1)
	log.Println("\033[0;31m["+metodo+"]\033[0m", "\033[0;32m"+endpoint, "\033[0m\033[0;33m", "\033[0;31m{"+status+"}\033[0m", sub, "\033[0m")
}
