package model

type Request struct {
	Start int `json:"start" validate:"required,min=1"`
	End   int `json:"end" validate:"required,gtefield=Start"`
}

type Response struct {
	PerfectNumbers []int `json:"perfect_numbers"`
}
