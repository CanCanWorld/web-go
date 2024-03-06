package serializer

type Response struct {
	Status int    `json:"status"`
	Data   any    `json:"data"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

type TokenData struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
