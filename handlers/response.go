package handlers

type ApiError struct {
	Status int16  `json:"status"`
	Title  string `json:"title"`
}

type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type JsonErrorResponse struct {
	Error *ApiError `json:"error"`
}
