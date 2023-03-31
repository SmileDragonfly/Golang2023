package SCB

import (
	"HttpServer/Utils"
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"github.com/google/uuid"
	log "github.com/jeanphorn/log4go"
	"net/http"
	"os"
	"time"
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

func (r SCB) Link(data []byte) ([]byte, error) {
	log.Info("Link")
	return nil, nil
}

func (r SCB) RequestOTP(data []byte) ([]byte, error) {
	log.Info("RequestOTP")
	url := Conf.Link + "/RequestOTPAcc"
	var body RequestBody
	var bodyData Request_ERequestOTPAcc
	err := CreateRequestBody("ERequestOTPAcc", bodyData, &body)
	if err != nil {
		log.Error("Create request error: ", err.Error())
		return nil, err
	}
	resp, err := DoRequest(url, body)
	if err != nil {
		log.Error("Do client request error: ", err.Error())
		return nil, err
	}
	log.Error("Do client request successfully: ", resp.Status)
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

func CreateRequestBody(funcName string, reqData RequestData, reqBody *RequestBody) error {
	// Encrypt request data using secret key
	reqDataByte, err := json.Marshal(reqData)
	if err != nil {
		return err
	}
	// Encrypt data
	sEncryptData, err := Utils.EncryptTripleDESB64(string(reqDataByte), Conf.SecretKey)
	if err != nil {
		return err
	}
	reqBody.Data = sEncryptData
	// Assign function name
	reqBody.FunctionName = funcName
	// Get current time: yyyy-MM-ddTHH:mm:ssZ
	now := time.Now()
	strNow := now.UTC().Format("2006-01-02T15:04:05Z")
	reqBody.RequestDateTime = strNow
	// Generate request id: GUID
	guid := uuid.New()
	reqBody.RequestID = guid.String()
	return nil
}

func DoRequest(url string, body BodyData) (*http.Response, error) {
	// Marshal body data
	bodyBuf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(bodyBuf)
	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, err
	}
	// Create header info
	req.SetBasicAuth(Conf.User, Conf.Password)
	sSignature, err := CreateSignatureDataB64(bodyBuf, Conf.PrivateKey)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Signature", sSignature)
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

func CreateSignatureDataB64(src []byte, pathPrivateKey string) (string, error) {
	bytePrivateKey, err := os.ReadFile(pathPrivateKey)
	if err != nil {
		return "", err
	}
	// Load private key
	pemBlock, _ := pem.Decode(bytePrivateKey)
	if pemBlock == nil || pemBlock.Type != "RSA PRIVATE KEY" {
		return "", errors.New("Decode pem file failed")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return "", err
	}
	// MD5 body data
	byteMD5 := md5.Sum(src)
	// SHA1 MD5 data
	byteSha1 := sha1.Sum(byteMD5[:])
	// Sign data using private key
	byteSignedData, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, byteSha1[:])
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(byteSignedData), nil
}

func VerifySignatureB64(sSignature string, src []byte, pathPublicKey string) error {
	// Decode base64 signature
	byteSignature, err := base64.StdEncoding.DecodeString(sSignature)
	if err != nil {
		return err
	}
	// MD5 body data
	byteMD5 := md5.Sum(src)
	// SHA1 MD5 data
	byteSha1 := sha1.Sum(byteMD5[:])
	// Load public key
	pubKeyByte, err := os.ReadFile(pathPublicKey)
	if err != nil {
		return err
	}
	pemBlock, _ := pem.Decode(pubKeyByte)
	if pemBlock == nil || pemBlock.Type != "PUBLIC KEY" {
		return err
	}
	publicKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return err
	}
	rsaPublicKey, _ := publicKey.(*rsa.PublicKey)
	return rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA1, byteSha1[:], byteSignature)
}
