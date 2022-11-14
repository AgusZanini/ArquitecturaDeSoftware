package dtos

type ResponseDto struct {
	NumFound      int      `json:"numFound"`
	Start         int      `json:"start"`
	NumFoundexact bool     `json:"numFoundExact"`
	Docs          ItemsDto `json:"docs"`
}
