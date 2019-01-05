// Front-end in Go server
// @jeffotoni
// 04/01/2019

package login

import (
	"net/http"
)

import (
	"github.com/jeffotoni/goapiwebserver/gofrontend/pkg/util"
)

func PostForm(w http.ResponseWriter, r *http.Request) {

	tplMap := map[string]interface{}{
		"title":   "GOCRUD - Login",
		"message": "",
	}

	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "form error", http.StatusInternalServerError)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		util.Print("username: " + username)
		util.Print("password: " + password)

		// req, err := http.NewRequest(http.MethodPost, "http://localhost:2015/gatekeeper", nil)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// req.SetBasicAuth(user, password)
		// client := http.Client{}
		// resp, err := client.Do(req)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// if resp.StatusCode == http.StatusForbidden {
		// 	tplMap["message"] = "User not found or incrorrect password"
		// 	err = templateHelper(w, tplMap, "login.html", "./assets/login.html")
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	}
		// 	return
		// }

		// if resp.StatusCode != http.StatusOK {
		// 	tplMap["message"] = "Some error message that can be seen by the user"
		// 	err = templateHelper(w, tplMap, "login.html", "./assets/login.html")
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	}
		// 	return
		// }

		// gatekeeperRet := struct {
		// 	Token string `json:"token"`
		// }{}
		// err = json.NewDecoder(resp.Body).Decode(&gatekeeperRet)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// // util.Print("token", gatekeeperRet.Token)

		// authCookie := http.Cookie{
		// 	Name:    "token",
		// 	Value:   gatekeeperRet.Token,
		// 	Expires: time.Now().Add(time.Duration(1) * time.Hour),
		// }

		// userDataCookie := http.Cookie{
		// 	Name:    "name",
		// 	Value:   user,
		// 	Expires: time.Now().Add(time.Duration(1) * time.Hour),
		// }

		// http.SetCookie(w, &authCookie)
		// http.SetCookie(w, &userDataCookie)

		// http.Redirect(w, r, "/", http.StatusSeeOther)

		// return
	}

	// err := templateHelper(w, tplMap, "web/templates/login-theme2.html", "./assets/login.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}
