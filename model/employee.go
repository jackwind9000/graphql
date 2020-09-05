package model

type Employee struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Position string  `json:"position"`
	Address  Address `json:"address"`
}

type Address struct {
	Ward     string `json:"ward"`
	District string `json:"district"`
	Province string `json:"province"`
	Country  string `json:"country"`
}
