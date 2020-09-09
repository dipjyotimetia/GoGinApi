package entity

type Account struct {
	ClientID       int            `json:"ClientID"`
	CurrencyCode   string         `json:"CurrencyCode"`
	StatusCode     string         `json:"StatusCode"`
	Balance        float64        `json:"Balance"`
}
