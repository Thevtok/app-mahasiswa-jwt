package model

type Students struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Major      string `json:"major"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	C_Username string `json:"c_username"`
}
