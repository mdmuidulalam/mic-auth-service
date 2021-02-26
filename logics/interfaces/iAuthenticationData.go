package logicinterface

type IAuthenticationData interface {
	/*
	* Action: Find auth information for a user with username
	* Return (*interface{}) => (authInfomation)
	 */
	FindOneAuthInformation() *IAuthInformation

	/*
	* Action: Set username class variable
	* Parameter: (string) => (username)
	 */
	SetUserName(string)
}
