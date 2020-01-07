package oauth

type OAuthError struct {
	message string
}

func (e *OAuthError) Error() string {
	return e.message
}

var (
	OAuthAuthorizationDeniedErr = &OAuthError{"authorization denied by user"}
	OAuthInvalidCredentialsErr  = &OAuthError{"invalid client_id or client_secret"}
	OAuthInvalidCodeErr         = &OAuthError{"unrecognized code"}
	OAuthServerErr              = &OAuthError{"server error"}
	OAuthUnauthorizedClientErr  = &OAuthError{"unauthorized client or this client is disabled/blocked"}
)
