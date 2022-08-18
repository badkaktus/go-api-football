package gaf

type errorResponse struct {
	Message string `json:"message"`
}

type successResponse struct {
	Get     string `json:"get"`
	Results int    `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response interface{} `json:"response"`
}
