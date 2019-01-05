// Front-end in Go server
// @jeffotoni
// 04/01/2019

package session

import (
    "net/http"
    "os"

    "github.com/gorilla/sessions"
)

// session key being set in makefile
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

/*
&sessions.Options{
        Path:     "/",
        MaxAge:   86400 * 7,
        HttpOnly: true,
    }
*/

// saving session
func Set(session_name, session_key, session_value string, w http.ResponseWriter, r *http.Request) {

    // If you do not pass, do not test or create a session.
    if session_name == "" || session_key == "" || session_value == "" {
        return
    }

    // Get a session. Get() always returns a session, even if empty.
    session, err := store.Get(r, session_name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set some session values.
    session.Values[session_key] = session_value

    // Save it before we write to the response/return from the handler.
    session.Save(r, w)
}

//getting the session
func Get(session_name, session_key string, w http.ResponseWriter, r *http.Request) string {

    // If you do not pass, do not test or create a session.
    if session_name == "" || session_key == "" {
        return ""
    }

    // Get a session. Get() always returns a session, even if empty.
    session, err := store.Get(r, session_name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return ""
    }

    if session.Values[session_key] == nil {
        return ""
    }

    return session.Values[session_key].(string)
}

//clearing session
func Destroy(session_name, session_key string, w http.ResponseWriter, r *http.Request) {

    // If you do not pass, do not test or create a session.
    if session_name == "" || session_key == "" {
        return
    }

    // Get a session. Get() always returns a session, even if empty.
    session, err := store.Get(r, session_name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set some session values.
    session.Values[session_key] = ""

    session.Options.MaxAge = -1
    // Save it before we write to the response/return from the handler.
    session.Save(r, w)
}
