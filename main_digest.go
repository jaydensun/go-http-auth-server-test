package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	auth "github.com/abbot/go-http-auth"
)

func secret(user, realm string) string {
	if user == "john" {
		// password is "hello"
		return "b98e16cbc3d01734b264adba7baa3bf9"
	}
	return ""
}

func handleDigest(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	fmt.Println("current time : " + time.Now().String())
	fmt.Fprintf(w, "{\"user\":\"%s\"}", r.Username)
}

func main() {
	authenticator := auth.NewDigestAuthenticator("example.com", secret)
	http.HandleFunc("/hellodigest", authenticator.Wrap(handleDigest))
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	fmt.Println("port is : " + port)
	http.ListenAndServe(":"+port, nil)
}
