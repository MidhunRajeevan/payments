package gateways

import "time"

type PaymentGateway struct {
	ID           int64     `json:"-"`
	Code         string    `json:"code" validate:"required"`
	StandardName string    `json:"standardName" validate:"required"`
	NameEn       string    `json:"nameEn" validate:"required"`
	NameAr       string    `json:"NameAr" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Status       string    `json:"status"`
	IsActive     bool      `json:"isActive"`
	IsArchived   bool      `json:"isArchived"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type GatewayHeader struct {
	ID           int64  `json:"-"`
	StandardName string `json:"standardName"`
	NameEn       string `json:"nameEn"`
	NameAr       string `json:"NameAr"`
	Description  string `json:"description"`
}
