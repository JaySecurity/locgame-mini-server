package siwe

import (
	"crypto/ecdsa"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// InitMessage creates a Message object with the provided parameters
func InitMessage(address, nonce string, issuedAt, expirationTime time.Time, client string) (*Message, error) {
	return &Message{
		Address:        common.HexToAddress(address),
		Nonce:          nonce,
		Client:         client,
		IssuedAt:       issuedAt,
		ExpirationTime: expirationTime,
	}, nil
}

func (m *Message) eip191Hash() common.Hash {
	// Ref: https://stackoverflow.com/questions/49085737/geth-ecrecover-invalid-signature-recovery-id
	data := []byte(m.String())
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256Hash([]byte(msg))
}

// ValidNow validates the time constraints of the message at current time.
func (m *Message) ValidNow() (bool, error) {
	return m.ValidAt(time.Now().UTC())
}

// ValidAt validates the time constraints of the message at a specific point in time.
func (m *Message) ValidAt(when time.Time) (bool, error) {
	if when.After(m.ExpirationTime) {
		return false, &ExpiredMessage{"Message expired"}
	}
	return true, nil
}

// VerifyEIP191 validates the integrity of the object by matching it's signature.
func (m *Message) VerifyEIP191(signature string) (*ecdsa.PublicKey, error) {
	if isEmpty(&signature) {
		return nil, &InvalidSignature{"Signature cannot be empty"}
	}

	sigBytes, err := hexutil.Decode(signature)
	if err != nil {
		return nil, &InvalidSignature{"Failed to decode signature"}
	}

	// Ref:https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	sigBytes[64] %= 27
	if sigBytes[64] != 0 && sigBytes[64] != 1 {
		return nil, &InvalidSignature{"Invalid signature recovery byte"}
	}

	pkey, err := crypto.SigToPub(m.eip191Hash().Bytes(), sigBytes)
	if err != nil {
		return nil, &InvalidSignature{"Failed to recover public key from signature"}
	}

	address := crypto.PubkeyToAddress(*pkey)

	if address != m.Address {
		return nil, &InvalidSignature{"Signer address must match message address"}
	}

	return pkey, nil
}

// Verify validates time constraints and integrity of the object by matching it's signature.
func (m *Message) Verify(signature string, nonce string) (*ecdsa.PublicKey, error) {
	var err error

	_, err = m.ValidNow()
	if err != nil {
		return nil, err
	}

	if m.Nonce != nonce {
		return nil, &InvalidSignature{"Message nonce doesn't match"}
	}

	return m.VerifyEIP191(signature)
}

func (m *Message) prepareMessage() string {
	var header string
	if m.Client == "sale" {
		header = "Welcome to $ACEON SALE portal!\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nWallet address:\n" + m.Address.String() + "\n"
	} else {
		header = "Welcome to LoC Game!\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nWallet address:\n" + m.Address.String() + "\n"
	}
	nonce := "Nonce: " + m.Nonce
	issuedAt := "Issued At: " + m.IssuedAt.Format(time.RFC3339)

	bodyArr := []string{nonce, issuedAt}

	value := "Expiration Time: " + m.ExpirationTime.Format(time.RFC3339)
	bodyArr = append(bodyArr, value)

	body := strings.Join(bodyArr, "\n")

	return strings.Join([]string{header, body}, "\n")
}

func (m *Message) String() string {
	return m.prepareMessage()
}
