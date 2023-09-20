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
	City        CodeName    `json:"city"`
	Emirate     CodeName    `json:"emirate"`
	Country     CodeName    `json:"country"`
	PostalCode  string      `json:"postalCode"`
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

type Attachment struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

type Company struct {
	Name               Name    `json:"name"`
	TrafficFileNumber  string  `json:"trafficFileNumber"`
	TradeLicenseNumber string  `json:"tradeLicenseNumber"`
	Contact            Contact `json:"contact"`
}

type Invoice struct {
	Number      string    `json:"number"`
	IssuedDate  time.Time `json:"issuedDate"`
	LastDate    time.Time `json:"lastDate"`
	Amount      Amount    `json:"amount"`
	Description string    `json:"description"`
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
	Name    Name    `json:"name"`
	UserID  string  `json:"userid"`
	Contact Contact `json:"contact"`
}

type Source struct {
	Name Name   `json:"name"`
	DID  string `json:"did"`
}

type PaymentRequest struct {
	Invoice Invoice `json:"invoice"`
	Service Service `json:"service"`
	Company Company `json:"company"`
	Ewallet Ewallet `json:"ewallet"`
	Person  Person  `json:"person"`
	Source  Source  `json:"source"`
}

type InvoiceRequest struct {
	Invoice Invoice `json:"invoice"`
	Company Company `json:"company"`
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

type Companies struct {
	ID                 int64     `json:"id"`
	Name               Name      `json:"name"`
	CustGroup          string    `json:"custgroup"`
	Email              string    `json:"email"`
	TradeLicenseNumber string    `json:"tradeLicenseNumber" validate:"required"`
	TrafficFileNumber  string    `json:"trafficFileNumber" validate:"required"`
	Account            string    `json:"account"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type ContactTable struct {
	ID              int64     `json:"id"`
	ContactableType string    `json:"contactable_type"`
	ContactableID   int64     `json:"contactable_id"`
	ContactableKey  string    `json:"contactable_key"`
	Addresses       string    `json:"addresses"`
	Phones          string    `json:"phones"`
	Emails          string    `json:"emails"`
	Attachments     string    `json:"attachments"`
	CreatedBy       string    `json:"created_by"`
	UpdatedBy       string    `json:"updated_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type InvoiceTable struct {
	ID                     int64      `json:"id"`
	Did                    string     `json:"did" validate:"required"`
	CompanyID              int64      `json:"company_id"`
	Number                 string     `json:"number" validate:"required"`
	IssuedDate             time.Time  `json:"issued_date" validate:"required"`
	LastDate               time.Time  `json:"last_date" validate:"required"`
	Description            string     `json:"description"`
	AmountValue            float64    `json:"amount_value" validate:"required"`
	AmountCurrency         string     `json:"amount_currency" validate:"required"`
	PaymentStatus          string     `json:"payment_status" validate:"required"`
	PaymentDate            *time.Time `json:"payment_date"`
	PaymentReferenceNumber *string    `json:"paymentReferenceNumber"`
	InvoiceReferenceNumber string     `json:"invoiceReferenceNumber"`
	PurchaseOrderNo        string     `json:"purchase_order_form_no"`
	DocumentDate           *time.Time `json:"document_date"`
	RMSInvoiceDate         *time.Time `json:"rms_invoice_date"`
	Requester              string     `json:"requester"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}
