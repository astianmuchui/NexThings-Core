package schemas

// User Register Request
// Will be unmarshaled when signing up the user via api

// type UserRegisterRequest struct {
// 	FirstName string `json:"firstname" validate:"required,min=2,max=50"`
// 	LastName  string `json:"lastname" validate:"required,min=2,max=50"`
// 	Username  string `json:"username" validate:"required,min=3,max=30"`
// 	Email     string `json:"email" validate:"required,email,max=255"`

// 	PhoneNumber string `json:"phonenumber" validate:"required,min=10,max=15"`
// 	City        string `json:"city" validate:"required,min=2,max=100"`
// 	Country     string `json:"country" validate:"required,min=2,max=100"`
// }

type UserRegisterRequest struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	City        string `json:"city"`
	Country     string `json:"country"`
}
