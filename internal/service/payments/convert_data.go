package payments

import "time"

type convertData struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data struct {
		Id          int       `json:"id"`
		Symbol      string    `json:"symbol"`
		Name        string    `json:"name"`
		Amount      int       `json:"amount"`
		LastUpdated time.Time `json:"last_updated"`
		Quote       struct {
			LOCG struct {
				Price       float64   `json:"price"`
				LastUpdated time.Time `json:"last_updated"`
			} `json:"LOCG"`
			ETH struct {
				Price       float64   `json:"price"`
				LastUpdated time.Time `json:"last_updated"`
			} `json:"ETH"`
		} `json:"quote"`
	} `json:"data"`
}
