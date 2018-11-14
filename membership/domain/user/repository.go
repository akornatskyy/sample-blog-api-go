package user

type Repository interface {
	FindAuthInfo(username string) (*AuthInfo, error)
}
