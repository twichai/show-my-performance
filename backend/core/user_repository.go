package core

type UserRepository interface {
	Save(user User) error
	GetByEmail(userEmail string) (*User, error)
}
