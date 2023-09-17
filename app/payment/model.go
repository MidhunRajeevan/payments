package payment

import (
	"time"
)

type Amount struct {
	Value    float64 `json:"value" validate:"required"`
	Currency string  `json:"currency" validate:"required"`
}

type ServiceName struct {
	En string `json:"en"`
	Ar string `json:"ar"`
}

type Service struct {
	ServiceName ServiceName `json:"serviceName"`
	ID          string      `json:"id"`
	Description string      `json:"description"`
}

type Address struct {
	Label       string      `json:"label"`
	Address     string      `json:"address"`
	Makani      string      `json:"makani"`
	PoBox       string      `json:"poBox"`
	Building    CodeName    `json:"building"`
	Street      CodeName    `json:"street"`
	City        CodeName    `json:"city"`
	Region      CodeName    `json:"region"`
	Country     CodeName    `json:"country"`
	PostalCode  string      `json:"postalCode"`
	Emirate     CodeName    `json:"emirate"`
	GeoLocation GeoLocation `json:"geoLocation"`
	Preferred   bool        `json:"preferred"`
}

type GeoLocation struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type CodeName struct {
	Code string `json:"code"`
	Name struct {
		En string `json:"en"`
		Ar string `json:"ar"`
	} `json:"name"`
}

type Contact struct {
	Addresses []Address `json:"addresses"`
	Phones    []Phone   `json:"phones"`
	Emails    []Email   `json:"emails"`
}

type Phone struct {
	Label       string `json:"label"`
	CountryCode string `json:"countryCode"`
	Phone       string `json:"phone"`
	Extension   string `json:"extension"`
	Preferred   bool   `json:"preferred"`
}

type Email struct {
	Label     string `json:"label"`
	Email     string `json:"email"`
	Preferred bool   `json:"preferred"`
}

type Company struct {
	Name               CodeName `json:"name"`
	TrafficFileNumber  string   `json:"trafficFileNumber"`
	TradeLicenseNumber string   `json:"tradeLicenseNumber"`
	Contact            Contact  `json:"contact"`
}

type Invoice struct {
	Number     string `json:"number"`
	IssuedDate string `json:"issuedDate"`
	LastDate   string `json:"lastDate"`
	Amount     Amount `json:"amount"`
}

type Ewallet struct {
	AccountNumber string `json:"accountNumber"`
	PIN           string `json:"pin"`
}

type Name struct {
	En string `json:"en"`
	Ar string `json:"ar"`
}

type Person struct {
	Name       Name    `json:"name"`
	UserID     string  `json:"userid"`
	EmiratesID string  `json:"emiratesId"`
	Contact    Contact `json:"contact"`
}

type Source struct {
	Name Name   `json:"name"`
	DID  string `json:"did"`
}

type PaymentRequest struct {
	Kind    string  `json:"kind"`
	Invoice Invoice `json:"invoice"`
	Service Service `json:"service"`
	Company Company `json:"company"`
	Ewallet Ewallet `json:"ewallet"`
	Person  Person  `json:"person"`
	Source  Source  `json:"source"`
}

type TransactionRequest struct {
	ID                   int64     `json:"id"`
	DID                  string    `json:"did"`
	Payload              string    `json:"payload"`
	Token                string    `json:"token"`
	TransactionReference string    `json:"transactionReference"`
	Status               string    `json:"status"`
	Attempts             int       `json:"attempts"`
	Message              string    `json:"message"`
	IsExpired            bool      `json:"isExpired"`
	PaymentInitiatedAt   time.Time `json:"PaymentInitiatedAt"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

type GatewayHeader struct {
	ID           int64  `json:"-"`
	StandardName string `json:"standardName"`
	NameEn       string `json:"nameEn"`
	NameAr       string `json:"NameAr"`
	Description  string `json:"description"`
}

type TransactionRequestResponse struct {
	ID                   int64           `json:"id"`
	DID                  string          `json:"did"`
	Token                string          `json:"token"`
	TransactionReference string          `json:"transactionReference"`
	Status               string          `json:"status"`
	Attempts             int             `json:"attempts"`
	Message              string          `json:"message"`
	IsExpired            bool            `json:"isExpired"`
	PaymentInitiatedAt   time.Time       `json:"PaymentInitiatedAt"`
	CreatedAt            time.Time       `json:"createdAt"`
	UpdatedAt            time.Time       `json:"updatedAt"`
	Channels             []GatewayHeader `json:"channels"`
}
