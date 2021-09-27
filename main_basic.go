package main

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"net/http"
	"os"
	"time"
)

func Secret(user, realm string) string {
	if user == "john" {
		// password is "hello"
		return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
	}
	return ""
}

func handle(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	fmt.Println("current time : " + time.Now().String())
	fmt.Fprintf(w, "{\"user\":\"%s\"}", r.Username)
}

func main() {
	authenticator := auth.NewBasicAuthenticator("example.com", Secret)
	http.HandleFunc("/hello", authenticator.Wrap(handle))
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	fmt.Println("port is : " + port)
	http.ListenAndServe(":"+port, nil)
}
