package SCB

type ResponseOTP struct {
	FunctionName     string `json:"FunctionName"`
	RequestDateTime  string `json:"RequestDateTime"`
	RequestID        string `json:"RequestID"`
	Description      string `json:"Description"`
	RespCode         string `json:"RespCode"`
	Data             string `json:"Data"`
	IsRequiredOTP    string `json:"IsRequiredOTP"`
	AuthInfo         string `json:"AuthInfo"`
	AuthType         string `json:"AuthType"`
	AuthDesVN        string `json:"AuthDesVN"`
	AuthDesEN        string `json:"AuthDesEN"`
	AuthDesTimeoutVN string `json:"AuthDesTimeoutVN"`
	AuthDesTimeoutEN string `json:"AuthDesTimeoutEN"`
	RefNumber        string `json:"RefNumber"`
}
