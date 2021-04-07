package Models

type Signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
type SigninResponse struct {
	Role  string `json:"role"`
	Token string `json:"token"`
}
