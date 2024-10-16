package payments

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"hash/crc32"
	"io"
	"locgame-mini-server/pkg/dto/store"
	"locgame-mini-server/pkg/log"
	"net/http"
	"net/url"
	"strings"
)

type PaypalAuthResponse struct {
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	AppId       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
	Nonce       string `json:"nonce"`
}

type PaypalAmount struct {
	CurrencyCode string `json:"currency_code" default:"USD"`
	Value        string `json:"value"`
}

type PaypalPurchaseUnit struct {
	Amount      PaypalAmount `json:"amount"`
	Description string       `json:"description"`
	InvoiceId   string       `json:"invoice_id"`
}
type PaypalOrderRequest struct {
	Intent        string               `json:"intent"`
	PurchaseUnits []PaypalPurchaseUnit `json:"purchase_units"`
}
type PaypalLinks struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}
type PaypalOrderResponse struct {
	ID    string        `json:"id"`
	Links []PaypalLinks `json:"links"`
}

func (s *Service) GeneratePaypalToken() (string, error) {
	// Generate the token
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/oauth2/token", s.config.Paypal.URL), strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(s.config.Paypal.ClientId, s.config.Paypal.Secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Error getting paypal auth token", err)
		return "", err
	}
	var authResponse PaypalAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		log.Error("Error decoding paypal auth response", err)
		return "", err
	}
	return authResponse.AccessToken, nil
}

func (s *Service) CreatePaypalPayment(order *store.Order) (string, error) {
	// Create the order
	token, err := s.GeneratePaypalToken()
	if err != nil {
		return "", err
	}

	data := &PaypalOrderRequest{
		Intent: "CAPTURE",
		PurchaseUnits: []PaypalPurchaseUnit{
			{
				Amount: PaypalAmount{
					Value:        order.Price,
					CurrencyCode: "USD",
				},
				Description: "LOCGame Card Purchase",
				InvoiceId:   order.ID.Value,
			},
		},
	}

	reqBody, err := json.Marshal(&data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v2/checkout/orders", s.config.Paypal.URL), bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Error creating paypal order", err)
		return "", err
	}
	var orderResponse PaypalOrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&orderResponse); err != nil {
		log.Error("Error decoding paypal order response", err)
		return "", err
	}

	return orderResponse.Links[1].Href, nil
}

func VerifySignature(r *http.Request, hookId string) (bool, error) {
	headers := r.Header
	rawData, err := io.ReadAll(r.Body)
	if err != nil {
		return false, err
	}
	transmissionId := headers.Get("paypal-transmission-id")
	timeStamp := headers.Get("paypal-transmission-time")
	certUrl := headers.Get("paypal-cert-url")
	crc := crc32.ChecksumIEEE(rawData)
	message := fmt.Sprintf("%s|%s|%s|%d", transmissionId, timeStamp, hookId, crc)
	resp, err := http.Get(certUrl)
	if err != nil {
		return false, err
	}
	certPem, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	signature, err := base64.StdEncoding.DecodeString(headers.Get("paypal-transmission-sig"))
	if err != nil {
		return false, err
	}
	verifier := sha256.New()
	verifier.Write([]byte(message))
	hashed := verifier.Sum(nil)
	block, _ := pem.Decode([]byte(certPem))
	if block == nil || block.Type != "CERTIFICATE" {
		return false, fmt.Errorf("failed to decode PEM block containing the certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return false, err
	}

	err = rsa.VerifyPKCS1v15(cert.PublicKey.(*rsa.PublicKey), crypto.SHA256, hashed, signature)
	return err == nil, err
}
