package model

type Store interface {
	Init() error
	AddUser(User) error
	FindUserById(uint) (error, User)
	Authenticate(string, string) (error, User)
	DeleteUser(User)
}
