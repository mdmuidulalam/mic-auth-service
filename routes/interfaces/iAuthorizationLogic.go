package routesinterface

type IAuthorizationLogic interface {
	/*
	* Action: Authorize a request
	* Return (int, string) => (Authorization Status)
	* Return => (1) => Authorization successfull
	* Return => (2) => Authorization unsuccessfull
	 */
	Authorize() int

	/*
	* Action: Set authToken class variable
	* Parameter: (string) => (authToken)
	 */
	SetAuthToken(string)
}
