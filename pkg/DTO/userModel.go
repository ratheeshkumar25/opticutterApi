package dto

// User represent the user data
type User struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

// OTP represent the otp verification
type OTP struct {
	Email string `json:"email" validate:"required"`
	OTP   string `json:"otp" validate:"required"`
}

// Login represent the login credentials
type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Password represent the updating password
type Password struct {
	Old     string `json:"old_password" validate:"required"`
	New     string `json:"new_password" validate:"required"`
	Confirm string `json:"confirm_password" validate:"required"`
}
