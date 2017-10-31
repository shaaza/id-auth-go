package serializer
type SignupRequest struct {
	Username    string `json:username`
	Password    string `json:password`
	FirstName   string `json:first_name`
	LastName    string `json:last_name`
	PhoneNumber string `json:phone_number`
}

type SignupResponse struct {
	Status      string `json:status`
}
