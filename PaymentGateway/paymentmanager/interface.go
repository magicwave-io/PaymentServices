package paymentmanager

import (
	"context"

	"paidpiper.com/payment-gateway/models"
)

type ClientHandler interface {
	ProcessCommand(nodeId models.PeerID, msg *PaymentCommand) error
	ProcessPayment(nodeId models.PeerID, msg *InitiatePayment)
	ProcessResponse(nodeId models.PeerID, msg *PaymentResponse)

	ValidatePayment(req *models.ValidatePaymentRequest) (uint32, error)
	CreatePaymentInfo(amount uint32) (*models.PaymentRequest, error)
	GetTransaction(sessionId string) (*models.PaymentTransaction, error)
}

type CallbackServer interface {
	Start()
	SetPort(port int)
	SetCallbackHandler(cb PPCallbackHandler)
	SetMetricsSource(source MetricsSource)
	Shutdown(ctx context.Context)
}
type PPConnection interface {
	ClientHandler
	CallbackServer
}
