package entity

type Account struct {
	AccountID    int     `json:"accountID"`
	CurrencyCode string  `json:"CurrencyCode"`
	StatusCode   string  `json:"StatusCode"`
	Balance      float64 `json:"Balance"`
	ClientID     int     `json:"ClientID"`
}
