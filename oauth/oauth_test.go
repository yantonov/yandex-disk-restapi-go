package yandexdiskapi

import (
	"./oauth"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_HttpClientIsGiven_OAuthServerDoesNotProvideCode(t *testing.T) {
	var credentials = oauth.ClientCredentials{}
	auth := OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return &http.Client{
				Transport: &storeRequestTransport{},
			}
		},
	}
	f := auth.HandlerFunc(credentials,
		func(auth *OAuthAuthorizationResponse,
			w http.ResponseWriter,
			r *http.Request) {
			t.Error("should handle request failure")
		},
		func(err error,
			w http.ResponseWriter,
			r *http.Request) {
			if err == nil {
				t.Error("error should not be if")
			}
			if err != OAuthInvalidCodeErr {
				t.Errorf("invalid error, got %v", err)
			}
		})

	req, _ := http.NewRequest("GET", "", nil)
	f(httptest.NewRecorder(), req)
}

func Test_UseDefaultHttpClient_OAuthServerDoesNotProvideCode(t *testing.T) {
	var credentials = ClientCredentials{}
	var auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return nil
		},
	}

	f := auth.HandlerFunc(credentials,
		func(auth *OAuthAuthorizationResponse,
			w http.ResponseWriter,
			r *http.Request) {
			t.Error("should handle request failure")
		},
		func(err error,
			w http.ResponseWriter,
			r *http.Request) {
			if err == nil {
				t.Error("error should not be nil")
			}
			if err != OAuthInvalidCodeErr {
				t.Errorf("invalid error, got %v", err)
			}
		})

	req, _ := http.NewRequest("GET", "", nil)
	f(httptest.NewRecorder(), req)
}

func Test_OAuthServerRespondWithAccessDeniedError(t *testing.T) {
	var credentials = ClientCredentials{}
	var auth = OAuthAuthenticator{}

	f := auth.HandlerFunc(credentials,
		func(auth *OAuthAuthorizationResponse,
			w http.ResponseWriter,
			r *http.Request) {
			t.Error("access denied should be failure")
		},
		func(err error,
			w http.ResponseWriter,
			r *http.Request) {
			if err != OAuthAuthorizationDeniedErr {
				t.Errorf("returned incorrect error, got %v", err)
			}
		})

	req, _ := http.NewRequest("GET", "?error=access_denied", nil)
	f(httptest.NewRecorder(), req)
}

func Test_OAuthServerRespondWithUnathorizedClientError(t *testing.T) {
	var credentials = ClientCredentials{}
	var auth = OAuthAuthenticator{}

	f := auth.HandlerFunc(credentials,
		func(auth *OAuthAuthorizationResponse,
			w http.ResponseWriter,
			r *http.Request) {
			t.Error("access denied should be failure")
		},
		func(err error,
			w http.ResponseWriter,
			r *http.Request) {
			if err != OAuthUnauthorizedClientErr {
				t.Errorf("returned incorrect error, got %v", err)
			}
		})

	req, _ := http.NewRequest("GET", "?error=unauthorized_client", nil)
	f(httptest.NewRecorder(), req)
}

func Test_OAuthServerReturnsInternalServerError(t *testing.T) {
	var credentials = ClientCredentials{}
	var auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient("{}", http.StatusInternalServerError).HttpClient
		},
	}

	f := auth.HandlerFunc(
		credentials,
		func(auth *OAuthAuthorizationResponse,
			w http.ResponseWriter,
			r *http.Request) {
			t.Error("should return error in case of invalid http code")
		},
		func(err error,
			w http.ResponseWriter,
			r *http.Request) {
			if err != OAuthServerErr {
				t.Errorf("returned incorrect error, got %v", err)
			}
		})

	req, _ := http.NewRequest("GET", "?code=123-456-789", nil)
	f(httptest.NewRecorder(), req)
}

func Test_OAuthServerRespondsWithBadRequest(t *testing.T) {
	var credentials = ClientCredentials{}
	var auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient(`{"error":"custom_code","error_description":"custom_description"}`,
				http.StatusBadRequest).HttpClient
		},
	}

	f := auth.HandlerFunc(credentials,
		func(auth *OAuthAuthorizationResponse,
			w http.ResponseWriter,
			r *http.Request) {
			t.Error("should return error")
		},
		func(err error,
			w http.ResponseWriter,
			r *http.Request) {
			if err == nil {
				t.Error("error expected")
			}
			var errMessage = err.Error()
			// TODO think about types and check code and description separately
			if errMessage != "custom_code: custom_description" {
				t.Errorf("returned incorrect error, got %v", errMessage)
			}
		})

	req, _ := http.NewRequest("GET", "?code=123-456-789", nil)
	f(httptest.NewRecorder(), req)
}

func Test_OAuthServerRespondsWithBadJson(t *testing.T) {
	var credentials = ClientCredentials{}
	var auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient(`bad json`, http.StatusOK).HttpClient
		},
	}

	f := auth.HandlerFunc(credentials,
		func(auth *OAuthAuthorizationResponse,
			w http.ResponseWriter,
			r *http.Request) {
			t.Error("should return error when server returned error")
		},
		func(err error,
			w http.ResponseWriter,
			r *http.Request) {
			if err == nil {
				t.Error("error should not be nil")
			}
			if _, ok := err.(*OAuthErrorResponse); ok {
				t.Error("invalid error")
			}
		})

	req, _ := http.NewRequest("GET", "", nil)
	f(httptest.NewRecorder(), req)
}

func Test_SuccessAuthorization(t *testing.T) {
	var credentials = ClientCredentials{}
	var auth = OAuthAuthenticator{
		RequestClientGenerator: func(r *http.Request) *http.Client {
			return NewStubResponseClient(`{}`, http.StatusOK).HttpClient
		},
	}

	f := auth.HandlerFunc(credentials,
		func(auth *OAuthAuthorizationResponse,
			w http.ResponseWriter,
			r *http.Request) {
		},
		func(err error,
			w http.ResponseWriter,
			r *http.Request) {
			t.Error("should be success")
		})

	req, _ := http.NewRequest("GET", "?code=123-456-789", nil)
	f(httptest.NewRecorder(), req)
}

func Test_CallbackUrlIsNotSet(t *testing.T) {
	auth := OAuthAuthenticator{}

	_, err := auth.CallbackPath()
	if err == nil {
		t.Error("should return error since callback url is not set")
	}
}

func Test_InvalidCallbackUrl(t *testing.T) {
	auth := OAuthAuthenticator{
		CallbackURL: "http://www.example.c%om/",
	}
	_, err := auth.CallbackPath()
	if err == nil {
		t.Error("should return error since not a callback url")
	}
}

func Test_ValidCallbackUrl_WaitExtractPath(t *testing.T) {
	auth := OAuthAuthenticator{
		CallbackURL: "http://abc.com/something/oauth",
	}
	s, _ := auth.CallbackPath()
	if s != "/something/oauth" {
		t.Error("incorrect path")
	}
}

func Test_AuthorizationUrlCustomState(t *testing.T) {
	var credentials = ClientCredentials{
		ClientId: "some_client_id",
	}
	auth := OAuthAuthenticator{
		CallbackURL: "http://abc.com/something/oauth",
	}

	url := auth.AuthorizationURL(credentials, "custom_state")
	if url != BaseOAuthPath+"/authorize?response_type=code&client_id=some_client_id&redirect_uri=http://abc.com/something/oauth&state=custom_state" {
		t.Errorf("incorrect oauth url, got %v", url)
	}
}

func Test_AuthorizationUrlStateIsUndefined(t *testing.T) {
	var credentials = ClientCredentials{
		ClientId: "some_client_id",
	}
	auth := OAuthAuthenticator{
		CallbackURL: "http://abc.com/something/oauth",
	}
	url := auth.AuthorizationURL(credentials, "")
	if url != BaseOAuthPath+"/authorize?response_type=code&client_id=some_client_id&redirect_uri=http://abc.com/something/oauth" {
		t.Errorf("incorrect oauth url, got %v", url)
	}
}

func Test_OAuthErrorError(t *testing.T) {
	var err = OAuthErrorResponse{
		ErrorCode:        "123",
		ErrorDescription: "descr",
	}
	if !strings.Contains(err.Error(), "123") {
		t.Error("error message should contain code")
	}
	if !strings.Contains(err.Error(), "descr") {
		t.Error("error message should contain description")
	}
}
