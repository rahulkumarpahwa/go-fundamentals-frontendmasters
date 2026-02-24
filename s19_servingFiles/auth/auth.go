package auth

import "net/http"

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the Auth and we will do the authentication here."))
}
