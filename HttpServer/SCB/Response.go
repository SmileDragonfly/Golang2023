package SCB

type ResponseBody struct {
	FunctionName    string `json:"FunctionName"`    //1.Echoed
	RequestDateTime string `json:"RequestDateTime"` //2.Echoed
	RequestID       string `json:"RequestID"`       //3.Echoed
	Description     string `json:"Description"`     //4.Diễn giải
	RespCode        string `json:"RespCode"`        //5.Kết quả thực hiện giao dịch
	Data            string `json:"Data"`            //6.Thông tin dữ liệu L2: encrypted data
}

// Request OTP: ERequestOTPAcc
type Response_ERequestOTPAcc struct {
	IsRequiredOTP bool                             `json:"IsRequiredOTP"` //7.Giá trị yêu cầu: - “True”: Sacombank yêu cầu xác thực OTP. - “False”: Sacombank không yêu cầu xác thực OTP.
	AuthInfo      Response_ERequestOTPAcc_AuthInfo `json:"AuthInfo"`      //8.Thông tin dữ liệu L3
}

type Response_ERequestOTPAcc_AuthInfo struct {
	IsRequiredOTP    bool   `json:"IsRequiredOTP"`    //9.Hình thức xác thực của khách hàng khi IsRequiredOTP = True: - 1: SMS OTPTài liệu tích hợp Sacombank Ecom - 2: HardToken - 3: mCode - 4: mConnected (*) - 5: AdvToken
	AuthDesVN        string `json:"AuthDesVN"`        //10.Diễn giải loại hình xác thực cần hiển thị cho khách hàng (Tiếng Việt) khi IsRequiredOTP = True
	AuthDesEN        string `json:"AuthDesEN"`        //11.Diễn giải loại hình xác thực cần hiển thị cho khách hàng (Tiếng Anh) khi IsRequiredOTP = True
	AuthDesTimeoutVN string `json:"AuthDesTimeoutVN"` //12.Diễn giải loại hình xác thực cần hiển thị cho khách hàng (Tiếng Việt) khi IsRequiredOTP = True, AuthType = 4 Xác thực mConnected bị timeout
	AuthDesTimeoutEN string `json:"AuthDesTimeoutEN"` //13.Diễn giải loại hình xác thực cần hiển thị cho khách hàng (Tiếng Anh) khi IsRequiredOTP = True, AuthType = 4 Xác thực mConnected bị timeout
	RefNumber        string `json:"RefNumber"`        //14.Echoed
}

// Topup transaction: ETopupByAccount
type Response_ETopupByAccount struct{} // Empty data

// Purchase transaction: EPurchaseByAccount
type Response_EPurchaseByAccount struct{} // Empty data

// Withdraw transaction
// Inquiry: ESubscriptionInquiry
// Cashout: ECashoutSubscription
type Response_ESubscriptionInquiry struct {
	SubscriptionID string `json:"SubscriptionID"` //7.Echoed
	CardNo         string `json:"CardNo"`         //8.Nếu liên kết bằng thẻ sẽ trả về số thẻ chỉ hiển thị 6 đầu – 4 cuối thẻ đã liên kết.
	AccountNo      string `json:"AccountNo"`      //9.Nếu liên kết bằng tài khoản sẽ trả về số tài khoản chỉ hiển thị 4 đầu – 4 cuối tài khoản đã liên kết.
	CustomerID     string `json:"CustomerID"`     //10.Echoed
	RefNumber      string `json:"RefNumber"`      //11.Echoed
}
type Response_ECashoutSubscription struct{} // Empty data

// Transfer to an account
// Inquery: ESTBAccountInquiry
// Transfer: EFundTransferToSTBAccount
type Response_ESTBAccountInquiry struct {
	FullName  string `json:"FullName"`  //7.Tên khách hàng
	RefNumber string `json:"RefNumber"` //8.Echoed
}
type Response_EFundTransferToSTBAccount struct{}

// Cancel Link: ECancelSubscription
type Response_ECancelSubscription struct{}
