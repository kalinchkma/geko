package mailers

type AuthMailer struct{}

func (a *AuthMailer) Welcome() string {
	return "Welcome to the mkr69"
}
