package routesinterface

type IRegisterLogic interface {
	/*
	* Action: Register an user
	* Return (int) => (Registration Status)
	* Return => (1) => Registration successfull
	* Return => (2) => Wrong authentication information
	 */
	Register() int

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
