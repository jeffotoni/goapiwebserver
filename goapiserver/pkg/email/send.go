package email

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jeffotoni/goapiwebserver/goapiserver/pkg/email"
)

func SendEmail(To, From, FromMsg, Nome, Projeto, EmailBounce, Bounce string) {

	var Demail = email.DataEmail{To, From, FromMsg, "Error accessing api", email.TemplateEmail(Nome, Projeto, EmailBounce, Bounce)}

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

	req, err := http.NewRequest("POST", email.END_POINT_SEND, bytes.NewBuffer(byteEmail))
	req = req.WithContext(ctx)

	req.Header.Set("X-Key", email.X_KEY)
	req.Header.Set("Authorization", email.AUTORIZATION)
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
