package authenticator

import "time"

type Authenticator struct {
	JWTAuth JWTAuthenticator
}

type AuthConfig struct {
	Token TokenConfig
}

type TokenConfig struct {
	Secret string
	Exp    time.Duration
	Iss    string
}
