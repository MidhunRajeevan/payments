package sources

import (
	"MidhunRajeevan/payments/app/gateways"
	"time"
)

type SourceSystem struct {
	ID         int64     `json:"-"`
	Did        string    `json:"did"`
	NameEn     string    `json:"nameEn"`
	NameAr     string    `json:"NameAr"`
	Status     string    `json:"status"`
	IsActive   bool      `json:"isActive"`
	IsArchived bool      `json:"isArchived"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type SourceSystemConfig struct {
	ID                int64     `json:"-"`
	SourceSystemId    string    `json:"sourceSystemId"`
	SPCode            string    `json:"SPCode"`
	SPValue           string    `json:"SPValue"`
	RMSSourceId       string    `json:"RMSSourceId"`
	RMSSourceName     bool      `json:"RMSSourceName"`
	RMSSourceSubfield string    `json:"RMSSourceSubfield"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type SourceSystemRecord struct {
	ID                int64                     `json:"-"`
	Did               string                    `json:"did"`
	NameEn            string                    `json:"nameEn"`
	NameAr            string                    `json:"NameAr"`
	Status            string                    `json:"status"`
	IsActive          bool                      `json:"isActive"`
	IsArchived        string                    `json:"isArchived"`
	SPCode            string                    `json:"SPCode"`
	SPValue           string                    `json:"SPValue"`
	RMSSourceId       string                    `json:"RMSSourceId"`
	RMSSourceName     bool                      `json:"RMSSourceName"`
	RMSSourceSubfield string                    `json:"RMSSourceSubfield"`
	Gateways          []gateways.PaymentGateway `json:"gateway"`
	CreatedAt         time.Time                 `json:"createdAt"`
	UpdatedAt         time.Time                 `json:"updatedAt"`
}

type SourceSystemGateway struct {
	ID          int64  `json:"id"`
	DID         string `json:"did"`
	GatewayCode string `json:"gateway_code"`
	Status      string `json:"status"`
	IsActive    bool   `json:"is_active"`
	IsArchived  bool   `json:"is_archived"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
