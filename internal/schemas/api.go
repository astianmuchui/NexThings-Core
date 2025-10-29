package schemas

// User Register Request
// Will be unmarshaled when signing up the user via api

type UserRegisterRequest struct {
	FirstName   string `json:"firstname" validate:"required,min=1,max=50"`
	LastName    string `json:"lastname" validate:"required,min=1,max=50"`
	Username    string `json:"username" validate:"required,alphanum,min=3,max=30"`
	Password    string `json:"password" validate:"required,min=8,max=72"`
	Email       string `json:"email" validate:"required,email,max=254"`
	PhoneNumber string `json:"phonenumber" validate:"required,e164"`
	City        string `json:"city" validate:"required,min=2,max=85"`
	Country     string `json:"country" validate:"required,min=2,max=56"`
}
