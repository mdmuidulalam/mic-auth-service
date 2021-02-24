package logicinterface

type IRegisterData interface {
	/*
	* Action: Create a new user with username and password
	* Return (bool) => (status flag success/failed)
	 */
	InsertAuthInformation()

	/*
	 * Action: Set username class variable
	 * Parameter: (string) => (username)
	 */
	SetUserName(string)

	/*
	 * Action: Set PasswordHash class variable
	 * Parameter: ([]byte) => (PasswordHash)
	 */
	SetPasswordHash([]byte)
}
