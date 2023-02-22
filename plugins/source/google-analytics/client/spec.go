package client

type Spec struct {
	AccountIDs []string `json:"account_ids,omitempty"`
	APIKey     string   `json:"api_key,omitempty"`
}
