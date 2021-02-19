package routesinterface

type IAuthenticationLogic interface {
	/*
	* Action: Authenticate an user
	* Return (int, string) => (Authentication Status, Auth Token)
	* Return => (1, "token") => Authentication successfull
	* Return => (2, "") => Wrong authentication information
	 */
	Authenticate() (int, string)

	/*
	* Action: Set username class variable
	* Parameter: (string) => (username)
	 */
	SetUserName(string)

	/*
	* Action: Set password class variable
	* Parameter: (string) => (password)
	 */
	SetPassword(string)
}
