package auth

import "net/http"

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is the Auth Method!"))
}
