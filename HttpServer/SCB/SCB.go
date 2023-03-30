package SCB

import (
	"bytes"
	"encoding/json"
	log "github.com/jeanphorn/log4go"
	"net/http"
)

type SCB struct{}

type Config struct {
	Link       string `json:"Link"`
	User       string `json:"User"`
	Password   string `json:"Password"`
	SecretKey  string `json:"SecretKey"`
	PrivateKey string `json:"PrivateKey"`
	PublicKey  string `json:"PublicKey"`
}

type BodyData interface {
}

var Conf Config

func (r SCB) doRequest(url string, body BodyData) (*http.Response, error) {
	// Marshal body data
	bodyBuf, err := json.Marshal(body)
	if err != nil {
		log.Error("Marshal body data error", err.Error())
		return nil, err
	}
	bodyReader := bytes.NewReader(bodyBuf)
	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		log.Error("Create request error", err.Error())
		return nil, err
	}
	// Create header info
	req.SetBasicAuth(Conf.User, Conf.Password)
	req.Header.Set("Signature", Conf.PrivateKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	// Send request to server
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Do client request error")
		return nil, err
	}
	return resp, err
}

func (r SCB) Link(data []byte) ([]byte, error) {
	log.Info("Link")
	url := Conf.Link + "/Link"
	var body RequestOTP
	resp, err := r.doRequest(url, body)
	if err != nil {
		log.Error("Do client request: ", resp.Status)
		return nil, err
	}
	return nil, nil
}

func (r SCB) RequestOTP(data []byte) ([]byte, error) {
	log.Info("RequestOTP")
	return nil, nil
}

func (r SCB) RequestOTPForLink(data []byte) ([]byte, error) {
	log.Info("RequestOTPForLink")
	return nil, nil
}

func (r SCB) TopupByCard(data []byte) ([]byte, error) {
	log.Info("TopupByCard")
	return nil, nil
}

func (r SCB) TopupByAccount(data []byte) ([]byte, error) {
	log.Info("TopupByAccount")
	return nil, nil
}

func (r SCB) PayByCard(data []byte) ([]byte, error) {
	log.Info("PayByCard")
	return nil, nil
}

func (r SCB) PayByAccount(data []byte) ([]byte, error) {
	log.Info("PayByAccount")
	return nil, nil
}

func (r SCB) Withdraw(data []byte) ([]byte, error) {
	log.Info("Withdraw")
	return nil, nil
}

func (r SCB) TransferToInternalCard(data []byte) ([]byte, error) {
	log.Info("TransferToInternalCard")
	return nil, nil
}

func (r SCB) TransferToInternalAccount(data []byte) ([]byte, error) {
	log.Info("TransferToInternalAccount")
	return nil, nil
}

func (r SCB) CancelLink(data []byte) ([]byte, error) {
	log.Info("CancelLink")
	return nil, nil
}

func (r SCB) TransferToExternalCard(data []byte) ([]byte, error) {
	log.Info("TransferToExternalCard")
	return nil, nil
}

func (r SCB) TransferToExternalAccount(data []byte) ([]byte, error) {
	log.Info("TransferToExternalAccount")
	return nil, nil
}

func (r SCB) TransferCardless(data []byte) ([]byte, error) {
	log.Info("TransferCardless")
	return nil, nil
}

func (r SCB) TransferVisaMasterCard(data []byte) ([]byte, error) {
	log.Info("TransferVisaMasterCard")
	return nil, nil
}

func (r SCB) QueryBalance(data []byte) ([]byte, error) {
	log.Info("QueryBalance")
	return nil, nil
}

func (r SCB) QueryStatus(data []byte) ([]byte, error) {
	log.Info("QueryStatus")
	return nil, nil
}

func (r SCB) CheckGuaranteeBlance(data []byte) ([]byte, error) {
	log.Info("CheckGuaranteeBlance")
	return nil, nil
}

func (r SCB) CheckCustomerInfo(data []byte) ([]byte, error) {
	log.Info("CheckCustomerInfo")
	return nil, nil
}

func (r SCB) PushNotification(data []byte) ([]byte, error) {
	log.Info("PushNotification")
	return nil, nil
}
