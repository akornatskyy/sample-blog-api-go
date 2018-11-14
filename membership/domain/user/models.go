package user

type (
	AuthInfo struct {
		ID       int
		IsLocked bool
		Password string
	}
)

func (a *AuthInfo) IsSamePassword(password string) bool {
	return a.Password == password
}
