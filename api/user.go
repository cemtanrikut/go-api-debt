package api

type user struct {
	name       string  `json:"name"`
	phone      string  `json:"phone"`
	mail       string  `json:"mail"`
	avg_income float32 `json:"avg_income"`
}

func UserHandler() {

}
