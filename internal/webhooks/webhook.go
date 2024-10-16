package webhooks

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/service/payments"
	"locgame-mini-server/internal/store"
	storeDto "locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Webhook struct {
	config   *config.Config
	payments *payments.Service
	store    *store.Store

	Mux *http.ServeMux
}

type PaypalEventType string

const (
	Approved PaypalEventType = "CHECKOUT.ORDER.APPROVED"
	Captured PaypalEventType = "PAYMENT.CAPTURE.COMPLETED"
)

type PaypalAmount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}
type PaypalPurchaseUnit struct {
	Amount      PaypalAmount `json:"amount"`
	Description string       `json:"description"`
	InvoiceId   string       `json:"invoice_id"`
}
type PaypalResource struct {
	Id            string               `json:"id"`
	PurchaseUnits []PaypalPurchaseUnit `json:"purchase_units,omitempty"`
	Amount        PaypalAmount         `json:"amount,omitempty"`
	Status        string               `json:"status,omitempty"`
	InvoiceId     string               `json:"invoice_id,omitempty"`
}
type PaypalWebhookRequest struct {
	Id        string         `json:"id,omitempty"`
	EventType string         `json:"event_type,omitempty"`
	Resource  PaypalResource `json:"resource"`
}

func NewWebhook(c *config.Config, p *payments.Service, s *store.Store) *Webhook {

	webhookMux := http.NewServeMux()

	webhook := &Webhook{
		config:   c,
		payments: p,
		store:    s,
		Mux:      webhookMux,
	}
	// Register webhook handlers
	webhookMux.HandleFunc("/approve", webhook.approve)
	webhookMux.HandleFunc("/capture", webhook.capture)

	return webhook
}

func (h *Webhook) approve(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Error reading request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		verified, err := payments.VerifySignature(r, h.config.Paypal.ApproveId)
		if err != nil {
			log.Error(err)
		}
		if !verified {
			log.Error("Signature verification failed")
		}
	}

	data := &PaypalWebhookRequest{}
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Error("Error decoding paypal order request", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.store.Orders.Get(context.Background(), data.Resource.PurchaseUnits[0].InvoiceId)
	if err != nil {
		log.Error("Error getting order: ", data.Resource.PurchaseUnits[0].InvoiceId, err)
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if data.EventType != string(Approved) {
		log.Error("Invalid event type: ", data.EventType)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := h.payments.GeneratePaypalToken()
	if err != nil {
		log.Error("Error generating paypal token", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	req, err := http.NewRequest("POST", h.config.Paypal.URL+"/v2/checkout/orders/"+data.Resource.Id+"/capture", nil)
	if err != nil {
		log.Error("Error creating capture request", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Error capturing order", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Debug(resp.StatusCode)
	w.WriteHeader(http.StatusOK)
}

func (h *Webhook) capture(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Error reading request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		verified, err := payments.VerifySignature(r, h.config.Paypal.CaptureId)
		if err != nil {
			log.Error(err)
		}
		if !verified {
			log.Error("Signature verification failed")
		}
		data := &PaypalWebhookRequest{}
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Error("Error decoding paypal order request", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if data.EventType != string(Captured) {
			log.Error("Invalid event type: ", data.EventType)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//TODO: Verify Payment and update order status
		order, err := h.store.Orders.Get(ctx, data.Resource.InvoiceId)
		if err != nil {
			log.Error("Error getting order:", data.Resource.InvoiceId, err)
			if err == mongo.ErrNoDocuments {
				w.WriteHeader(http.StatusOK)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		order.Status = storeDto.OrderStatus_PaymentReceived
		err = h.store.Orders.Update(ctx, order)
		if err != nil {
			log.Error("Unable to update order in database:", err)
		}

		h.payments.OnPaymentSuccess(ctx, order)
		log.Debug("Capture Conplete")
		w.WriteHeader(http.StatusOK)
		return

	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
