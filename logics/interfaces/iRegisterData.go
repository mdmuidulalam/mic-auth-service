package logicinterface

type IRegisterData interface {
	/*
	* Action: Create a new user with username and password
	 */
	InsertAuthInformation()

	/*
	* Action: Find auth information for a user with username
	* Return (*interface{}) => (authInfomation)
	 */
	FindOneAuthInformation() *IRegisterData

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
