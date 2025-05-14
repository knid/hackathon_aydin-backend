package requests

type UserCreateRequest struct {
	Name string
	Surname string
	Email string
	Password string
}

type UserLoginRequests struct {
	Email string
	Password string
}
