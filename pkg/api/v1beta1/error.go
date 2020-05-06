package v1beta1

type Error struct {
	Message string `json:"msg"`
}

type ErrorWrapper struct {
	Error Error `json:"error"`
}
