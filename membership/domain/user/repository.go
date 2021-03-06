package user

type Repository interface {
	FindAuthInfo(username string) (*AuthInfo, error)

	FindUserByID(id string) (*User, error)

	HasAccount(username string) (bool, error)

	CreateAccount(r *Registration) (bool, error)
}
