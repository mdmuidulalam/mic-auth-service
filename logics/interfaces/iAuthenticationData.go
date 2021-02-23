package logicinterface

type IAuthenticationData interface {
	/*
	* Action: Fetch password hash and active flag for an user
	* Return ([]byte, bool) => (password hash, active flag)
	 */
	FetchUserHashPasswordAndActive() ([]byte, bool)

	/*
	* Action: Set username class variable
	* Parameter: (string) => (username)
	 */
	SetUserName(string)
}
