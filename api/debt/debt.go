package api

import "time"

type debt struct {
	name       string    `json:"name"`
	typeof     string    `json:"typeof"`
	amount     float32   `json:"amount"`
	periodicly bool      `json:"periodicly"`
	start_date time.Time `json:"start_date"`
	end_date   time.Time `json:"end_date"`
	completed  bool      `json:"completed"`
	active     bool      `json:"active"`
}

func DebtHandler() {

}
