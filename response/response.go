package response

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"Error"`
}
