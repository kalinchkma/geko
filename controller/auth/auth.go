package authController

type AuthController interface {
	Register()
	VerifyOtp()
	ResendOtp()
	Login()
	ForgotPassword()
	ResetPassword()
	RefreshToken()
}
