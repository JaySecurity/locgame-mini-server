package token

type Info struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	ExternalUrl string `json:"external_url"`
	Attributes  []Attribute
}
