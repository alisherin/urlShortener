package http_handlers

type UrlPost struct {
	Url string `json:"url"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
