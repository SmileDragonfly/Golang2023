package SCB

type RequestOTP struct {
	FunctionName    string `json:"FunctionName"`
	RequestDateTime string `json:"RequestDateTime"`
	RequestID       string `json:"RequestID"`
	Data            string `json:"Data"`
	CustomerID      string `json:"CustomerID"`
	SubscriptionID  string `json:"SubscriptionID"`
	RefNumber       string `json:"RefNumber"`
}
