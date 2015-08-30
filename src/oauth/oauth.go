package oauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const BaseOAuthPath = "https://oauth.yandex.ru"

type OAuthAuthenticator struct {
	CallbackURL            string
	RequestClientGenerator func(r *http.Request) *http.Client
}

type OAuthAuthorizationResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint64 `json:"expires_in"`
}

type OAuthErrorResponse struct {
	ErrorCode        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (err OAuthErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", err.ErrorCode, err.ErrorDescription)
}

func (auth OAuthAuthenticator) CallbackPath() (string, error) {
	if auth.CallbackURL == "" {
		return "", errors.New("callbackURL is empty")
	}
	url, err := url.Parse(auth.CallbackURL)
	if err != nil {
		return "", err
	}
	return url.Path, nil
}

func (auth OAuthAuthenticator) Authorize(credentials ClientCredentials,
	code string,
	client *http.Client) (*OAuthAuthorizationResponse, error) {
	if code == "" {
		return nil, OAuthInvalidCodeErr
	}

	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.PostForm(BaseOAuthPath+"/token",
		url.Values{
			"grant_type": {
				"authorization_code",
			},
			"code": {
				code,
			},
			"client_id": {
				credentials.ClientId,
			},
			"client_secret": {
				credentials.ClientSecret,
			},
		})

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 == 5 {
		return nil, OAuthServerErr
	}

	if resp.StatusCode/100 != 2 {
		var response OAuthErrorResponse
		contents, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(contents, &response)

		// TODO Create instances for documented error codes

		return nil, &response
	}

	var response OAuthAuthorizationResponse
	contents, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(contents, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (auth OAuthAuthenticator) HandlerFunc(
	credentials ClientCredentials,
	success func(auth *OAuthAuthorizationResponse, w http.ResponseWriter, r *http.Request),
	failure func(err error, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("error") == "access_denied" {
			failure(OAuthAuthorizationDeniedErr, w, r)
			return
		}

		if r.FormValue("error") == "unauthorized_client" {
			failure(OAuthUnauthorizedClientErr, w, r)
			return
		}

		client := http.DefaultClient
		if auth.RequestClientGenerator != nil {
			client = auth.RequestClientGenerator(r)
		}

		resp, err := auth.Authorize(credentials,
			r.FormValue("code"),
			client)

		if err != nil {
			failure(err, w, r)
			return
		}

		success(resp, w, r)
	}
}

func (auth OAuthAuthenticator) AuthorizationURL(
	credentials ClientCredentials,
	state string) string {
	path := fmt.Sprintf(
		"%s/authorize?response_type=code&client_id=%s&redirect_uri=%s",
		BaseOAuthPath,
		credentials.ClientId,
		auth.CallbackURL)

	if state != "" {
		path += "&state=" + state
	}

	return path
}
