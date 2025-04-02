package core

type User struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ProfileImageURL string `json:"profile_image_url"`
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
}
