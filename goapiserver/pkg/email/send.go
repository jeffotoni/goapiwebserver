// Go Api server
// @jeffotoni
// 2019-01-04

package email

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func SendEmail(To, From, FromMsg, Nome, Projeto string) {

	var Demail = DataEmail{To, From, FromMsg, "Error accessing api", TemplateEmail(Nome, Projeto)}

	//
	byteEmail, err := json.Marshal(Demail)
	if err != nil {
		log.Println("error while doing marshal in string mail:", err)
		return
	}

	ctx, cancel := context.WithCancel(context.TODO())
	afterFuncTimer := time.AfterFunc(2*time.Second, func() { cancel() })
	defer afterFuncTimer.Stop()
	// defer cancel() // cancel when we are finished consuming integers

	req, err := http.NewRequest("POST", END_POINT_SEND, bytes.NewBuffer(byteEmail))
	req = req.WithContext(ctx)

	req.Header.Set("X-Key", X_KEY)
	req.Header.Set("Authorization", AUTORIZATION)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		log.Println("############################## ERROR SENDING EMAIL ###########################")
		log.Println("http.client error: ", err)
		//////// parar aqui...
		return
	}

	defer resp.Body.Close()

	log.Println("############################## email error ###########################")
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	log.Println("#########################################################################")
}
