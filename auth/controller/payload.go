package authcontroller

type RegisterPayload struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

var RegisterValidationMessages = map[string]string{
	"Name.required":     "name is required",
	"Name.min":          "name is too short at leat 3 alphabate",
	"Name.max":          "name is too long it must less then 50 alphabate",
	"Email.required":    "email is required",
	"Email.email":       "Invalid Email",
	"Password.required": "password is required",
	"Password.min":      "password must be at least 8 character long",
}

type LoginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

var LoginValidationMessages = map[string]string{
	"Email.required":    "email is required",
	"Email.email":       "Invalid email input",
	"Password.required": "passowrd is required",
}

type VerifyOTPPayload struct {
	Code   string `json:"code" binding:"required"`
	UserId uint   `json:"user_id" binding:"required"`
}

var VerifyOTPValidationMessages = map[string]string{
	"Code.required":  "code is required",
	"UserId.user_id": "user_id is required",
}

type ResendOTPPayload struct {
	UserId uint `json:"user_id" binding:"required"`
}

var ResendOTPValidationMessages = map[string]string{
	"UserId.required": "user_id is required",
}
