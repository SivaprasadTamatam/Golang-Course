package common

type User struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Relation  string  `json:"relation"`
	Income    float64 `json:"income"`
	Expenses  float64 `json:"expenses"`
	Balance   float64 `json:"balance"`
	Budget    float64 `json:"budget"`
	Bill      float64 `json:"bill"`
}

type Incoming struct {
	Date     string  `json:"date"`
	Account  string  `json:"account"` // actually user
	Vendor   string  `json:"vendor"`
	Category string  `json:"category"`
	Deposit  float64 `json:"deposit"`
	Balance  float64 `json:"balance"`
	Notes    string  `json:"notes"`
}

type Expenses struct {
	Date     string  `json:"date"`
	Account  string  `json:"account"` // actually user
	Category string  `json:"category"`
	Vendor   string  `json:"vendor"`
	Payment  float64 `json:"payment"`
	Balance  float64 `json:"balance"`
	Notes    string  `json:"notes"`
}
