// Go Api server
// @jeffotoni
// 2019-01-04

package gjwt

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jeffotoni/goapiwebserver/goapiserver/pkg/util"
)

// Returns json without typing in http
func WriteJson(Status string, Msg string) {
	msgJsonStruct := &JsonMsg{Status, Msg}
	msgJson, errj := json.Marshal(msgJsonStruct)
	if errj != nil {
		msg := `{"status":"error","msg":"We could not generate the json error!"}`
		util.Print(msg)
		return
	}

	// byte
	util.Print(string(msgJson))
}

// Returns json by typing on http
func HttpWriteJson(w http.ResponseWriter, Status string, Msg string, httpStatus int) {
	msgJsonStruct := &JsonMsg{Status, Msg}
	msgJson, errj := json.Marshal(msgJsonStruct)
	if errj != nil {
		msg := `{"status":"error","msg":"We could not generate the json error!"}`
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, msg)
		return
	}
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	w.Write(msgJson)
}
