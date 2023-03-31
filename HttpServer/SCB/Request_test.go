package SCB

import (
	"testing"
)

func TestCreateSignatureDataB64(t *testing.T) {
	src := []byte("This is test data")
	sSignature, err := CreateSignatureDataB64(src, "D:/OnePay/_MyCode/eWallet2023/BankConfigs/SCB/PrivateKey.pem")
	if err != nil {
		t.Error("Create signature failed", err.Error())
	}
	t.Log("Signature: ", sSignature)
	err = VerifySignatureB64(sSignature, src, "D:/OnePay/_MyCode/eWallet2023/BankConfigs/SCB/PrivateKey.pem")
	if err != nil {
		t.Error("Verify signature failed", err.Error())
	}
}
