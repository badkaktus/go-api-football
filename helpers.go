package gaf

type successResponse struct {
	Get     string `json:"get"`
	Results int    `json:"results"`
	Errors  any    `json:"errors"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response interface{} `json:"response"`
}
