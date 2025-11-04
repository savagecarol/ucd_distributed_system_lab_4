package model


type SecondRequest struct {
	Name string `json:"name" required:"true"`
	Age int `json:"age" required:"true"`
}