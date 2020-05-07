package v1beta1

//Error wraps errors to be returned as JSON
//in a HTTP response
type Error struct {
	Message string `json:"msg"`
}

type ErrorWrapper struct {
	Error Error `json:"error"`
}
