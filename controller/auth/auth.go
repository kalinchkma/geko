package authController

type AuthController interface {
	Register()
	VerifyOtp()
	Login()
	ForgotPassword()
	ResetPassword()
	RefreshToken()
}
