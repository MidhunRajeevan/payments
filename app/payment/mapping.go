package payment

import "MidhunRajeevan/payments/app/gateways"

func TransactionRequestResponseMapping(lstGateways []gateways.PaymentGateway, record TransactionRequest) TransactionRequestResponse {
	var response TransactionRequestResponse

	response.ID = record.ID
	response.DID = record.DID
	response.Token = record.Token
	response.TransactionReference = record.TransactionReference
	response.Status = record.Status
	response.Attempts = record.Attempts
	response.Message = record.Message
	response.IsExpired = record.IsExpired
	response.PaymentInitiatedAt = record.PaymentInitiatedAt
	response.CreatedAt = record.CreatedAt
	response.UpdatedAt = record.UpdatedAt

	for _, channel := range lstGateways {
		var gatewayHeader GatewayHeader
		gatewayHeader.ID = channel.ID
		gatewayHeader.StandardName = channel.StandardName
		gatewayHeader.NameEn = channel.NameEn
		gatewayHeader.NameAr = channel.NameAr
		response.Channels = append(response.Channels, gatewayHeader)
	}

	return response
}
