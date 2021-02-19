package logics

type AuthenticationLogic struct {
	username string
	password string
}

func (authLogic AuthenticationLogic) SetUserName(username string) {
	authLogic.username = username
}

func (authLogic AuthenticationLogic) SetPassword(password string) {
	authLogic.password = password
}

func (authLogic AuthenticationLogic) Authenticate() (int, string) {
	return 1, "dajsdhkddasdj"
}
