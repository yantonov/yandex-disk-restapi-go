// oauth_example.go provides a simple example implementing yandex disk oauth
//
// usage:
//   > go get github.com/yantonov/yandex-disk-restapi-go
//   > cd $GOPATH/github.com/yantonov/yandex-disk-restapi-go/examples
//   > go run oauth_example.go -id=youappid -secret=yourapppassword
//
//   Visit http://localhost:8080 in your webbrowser
//
//   Application id and secret can be found at https://oauth.yandex.ru/

package main

import (
	"flag"
	"fmt"
	"github.com/yantonov/yandex-disk-restapi-go/src/oauth"
	"net/http"
	"os"
)

const local_demo_server_port = 8080

var authenticator *oauth.OAuthAuthenticator

func main() {
	var credentials = oauth.ClientCredentials{}
	// setup the credentials for your app
	// can be defined here https://oauth.yandex.ru
	flag.StringVar(&credentials.ClientId, "id", "", "yandex disk client id")
	flag.StringVar(&credentials.ClientSecret, "secret", "", "yandex disk client secret")

	flag.Parse()

	if credentials.ClientId == "" || credentials.ClientSecret == "" {
		fmt.Println("\nPlease provide your application's client id and secret.")
		fmt.Println("For example: go run oauth_example.go -id=123 -secret=myappsecret")
		fmt.Println(" ")

		flag.PrintDefaults()
		os.Exit(1)
	}

	authenticator = &oauth.OAuthAuthenticator{
		CallbackURL:            fmt.Sprintf("http://localhost:%d/exchange_token", local_demo_server_port),
		RequestClientGenerator: nil,
	}

	http.HandleFunc("/",
		func(w http.ResponseWriter,
			r *http.Request) {
			indexHandler(credentials, w, r)
		})

	path, err := authenticator.CallbackPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	http.HandleFunc(path, authenticator.HandlerFunc(credentials,
		oAuthSuccess,
		oAuthFailure))

	// start the server
	fmt.Printf("Visit http://localhost:%d/ to view the demo\n", local_demo_server_port)
	fmt.Printf("ctrl-c to exit")
	http.ListenAndServe(fmt.Sprintf(":%d", local_demo_server_port), nil)
}

func indexHandler(credentials oauth.ClientCredentials,
	w http.ResponseWriter,
	r *http.Request) {
	// you should make this a template in your real application
	fmt.Fprintf(w, `<a href="%s">`, authenticator.AuthorizationURL(credentials, ""))
	fmt.Fprint(w, `<span>connect to yandex disk</span>`)
	fmt.Fprint(w, `</a>`)
}

func oAuthSuccess(auth *oauth.OAuthAuthorizationResponse,
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "SUCCESS:\nAt this point you can use this information to create a new user or link the account to one of your existing users\n")
	fmt.Fprintf(w, "Access Token: %s\n\n", auth.AccessToken)
	fmt.Fprintf(w, "Token type: %s\n\n", auth.TokenType)
	fmt.Fprintf(w, "Expires in: %d seconds\n\n", auth.ExpiresIn)
}

func oAuthFailure(err error,
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "Authorization Failure:\n")

	// some standard error checking
	if err == oauth.OAuthAuthorizationDeniedErr {
		fmt.Fprint(w, "The user clicked the 'Do not Authorize' button on the previous page.\n")
		fmt.Fprint(w, "This is the main error your application should handle.")
	} else if err == oauth.OAuthInvalidCredentialsErr {
		fmt.Fprint(w, "You provided an incorrect client_id or client_secret.\nDid you remember to set them at the begininng of this file?")
	} else if err == oauth.OAuthInvalidCodeErr {
		fmt.Fprint(w, "The temporary token was not recognized, this shouldn't happen normally")
	} else if err == oauth.OAuthServerErr {
		fmt.Fprint(w, "There was some sort of server error, try again to see if the problem continues")
	} else {
		fmt.Fprint(w, err)
	}
}
