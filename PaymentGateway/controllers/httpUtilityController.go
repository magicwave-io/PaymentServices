package controllers

import (
	"encoding/json"
	"net/http"
	"paidpiper.com/payment-gateway/node/local/paymentregestry/database/entity"
	"strconv"

	"paidpiper.com/payment-gateway/log"

	"github.com/gorilla/mux"
	"paidpiper.com/payment-gateway/common"

	"paidpiper.com/payment-gateway/models"
	"paidpiper.com/payment-gateway/node/local"
)

type HttpUtilityController struct {
	local.LocalPPNode
}

func NewHttpUtilityController(n local.LocalPPNode) *HttpUtilityController {
	return &HttpUtilityController{
		n,
	}
}

func (u *HttpUtilityController) ListTransactions(w http.ResponseWriter, r *http.Request) {
	_, span := spanFromRequest(r, "requesthandler:ListTransactions")
	defer span.End()

	trx := u.GetTransactions()

	Respond(w, trx)
}

func (u *HttpUtilityController) HttpGetTransactionInfo(w http.ResponseWriter, r *http.Request) {
	_, span := spanFromRequest(r, "requesthandler:GetTransactionInfo")
	defer span.End()

	trx := u.GetTransactionInfo()

	Respond(w, trx)
}
func (u *HttpUtilityController) HttpGetTransaction(w http.ResponseWriter, r *http.Request) {
	_, span := spanFromRequest(r, "requesthandler:GetTransaction")
	defer span.End()

	vars := mux.Vars(r)
	sessionId := vars["sessionId"]

	trx := u.GetTransaction(sessionId)

	Respond(w, trx)
}

func (u *HttpUtilityController) HttpFlushTransactions(w http.ResponseWriter, r *http.Request) {

	ctx, span := spanFromRequest(r, "requesthandler:FlushTransactions")
	defer span.End()
	err := u.FlushTransactions(ctx)
	if err != nil {
		log.Errorf("Error flushing transactions: %s", err.Error())
		Respond(w, MessageWithStatus(http.StatusBadRequest, "Error in FlushTransactions: "+err.Error()))
	}

	Respond(w, MessageWithStatus(http.StatusOK, "Transactions committed"))
}

func (u *HttpUtilityController) HttpGetStellarAddress(w http.ResponseWriter, r *http.Request) {
	response := u.GetStellarAddress()
	Respond(w, response)
}

func (u *HttpUtilityController) HttpNewPaymentRequest(w http.ResponseWriter, r *http.Request) {
	ctx, span := spanFromRequest(r, "requesthandler:HttpNewPaymentRequest")

	defer span.End()

	request := &models.CreatePaymentInfo{}
	err := json.NewDecoder(r.Body).Decode(request)
	log.Infof("CreatePaymentInfo: %v %v %v", request.CommodityType, request.ServiceType, request.Amount)
	if err != nil {
		log.Errorf("Error decoding new payment request: %s", err.Error())
		Respond(w, MessageWithStatus(http.StatusBadRequest, "Invalid request"))
		return
	}

	pr, err := u.NewPaymentRequest(ctx, request)

	if err != nil {
		log.Errorf("Error creating new payment request (invalid commodity): %s", err.Error())
		Respond(w, MessageWithStatus(http.StatusBadRequest, "Invalid commodity"))
		return
	}
	Respond(w, pr)
}

func (u *HttpUtilityController) HttpValidatePayment(w http.ResponseWriter, r *http.Request) {

	ctx, span := spanFromRequest(r, "ValidatePayment")
	defer span.End()

	request := &models.ValidatePaymentRequest{}

	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		log.Error("Error decoding payment validation request : %s", err.Error())
		Respond(w, MessageWithStatus(http.StatusBadRequest, "Bad request"))
		return
	}

	response, err := u.ValidatePayment(ctx, request)

	if err != nil {
		log.Errorf("Error validating payment: %s", err.Error())
		Respond(w, MessageWithStatus(http.StatusBadRequest, err.Error()))
		return
	}

	Respond(w, response)
}

func (u *HttpUtilityController) HttpProcessCommand(w http.ResponseWriter, r *http.Request) {
	ctx, span := spanFromRequest(r, "requesthandler:ProcessCommand")
	defer span.End()

	command := &models.UtilityCommand{}

	err := json.NewDecoder(r.Body).Decode(command)

	if err != nil {
		log.Errorf("Error decoding process command request : %s", err.Error())
		Respond(w, MessageWithStatus(http.StatusBadRequest, "Invalid request"))
		return
	}

	data, err := u.ProcessCommand(ctx, command)
	if err != nil {
		log.Errorf("Error processing command request : %s", err.Error())
		Respond(w, MessageWithStatus(http.StatusConflict, err.Error()))
		return
	}
	if data != nil {
		log.Info("Process Command result success: OutType=%v", data.OutType())
		Respond(w, MessageWithData(http.StatusOK, data))
		return
	} else {
		log.Infof("Command to queue: SessionId=%v, CommandId=%v, CommandType=%v", command.SessionId, command.CommandId, command.CommandType)

		Respond(w, MessageWithStatus(http.StatusCreated, "success"))
	}

}

func (u *HttpUtilityController) HttpGetBalance(w http.ResponseWriter, r *http.Request) {
	res, err := u.GetBookBalance()
	if err != nil {
		log.Errorf("Error processing get balance request : %s", err.Error())
		Respond(w, MessageWithStatus(http.StatusConflict, err.Error()))
		return
	}
	response := &models.GetBalanceResponse{
		Balance:   res.Balance,
		Timestamp: res.Timestamp,
	}

	Respond(w, response)
}

func (u *HttpUtilityController) HttpBookHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commodity := vars["commodity"]
	hours := vars["hours"]
	bins := vars["bins"]
	binsValue, err := strconv.Atoi(bins)

	if err != nil {
		log.Errorf("Error - bad value for bins : %s", vars["bins"])
		Respond(w, common.Error(500, "HISTORY_BINS should be int"))
	}

	hoursValue, err := strconv.Atoi(hours)
	if err != nil {
		log.Errorf("Error - bad value for hours : %s", vars["hours"])
		Respond(w, common.Error(500, "hours should be int"))
	}
	res, err := u.GetBookHistory(commodity, binsValue, hoursValue)

	if err != nil {
		log.Errorf("Error retrieving book history: %s", vars["hours"])
		Respond(w, common.Error(500, err.Error()))
	}
	Respond(w, res)

}

func (u *HttpUtilityController) HttpBookBalance(w http.ResponseWriter, r *http.Request) {
	res, err := u.GetBookBalance()
	if err != nil {
		log.Errorf("Error retrieving book balance: %s", err.Error())
		Respond(w, common.Error(500, err.Error()))
	}
	Respond(w, res)
}

func (u *HttpUtilityController) HttpGetTorRoutingMode(w http.ResponseWriter, r *http.Request) {
	ctx, span := spanFromRequest(r, "requesthandler:HttpGetTorRoutingMode")

	defer span.End()

	state, err := u.GetRuntimeState(ctx)

	if err != nil {
		log.Errorf("Error retrieving book balance: %s", err.Error())
		Respond(w, common.Error(500, err.Error()))
	}

	Respond(w, MessageWithData(http.StatusOK, state))
}

func (u *HttpUtilityController) HttpSetTorRoutingMode(w http.ResponseWriter, r *http.Request) {
	ctx, span := spanFromRequest(r, "requesthandler:HttpSetTorRoutingMode")
	defer span.End()

	state := &entity.RuntimeState{}

	err := json.NewDecoder(r.Body).Decode(state)

	if err != nil {
		log.Errorf("Error decoding state request : %s", err.Error())
		Respond(w, MessageWithStatus(http.StatusBadRequest, "Invalid request"))
		return
	}

	err = u.SetRuntimeState(ctx, state.TorRoutingMode)

	if err != nil {
		log.Errorf("Error setting state: %s", err.Error())
		Respond(w, MessageWithStatus(http.StatusBadRequest, "Invalid request"))
		return
	}

	Respond(w, MessageWithStatus(http.StatusOK, "success"))
}
