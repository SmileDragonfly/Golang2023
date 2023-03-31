package SCB

type RequestData interface {
}

type RequestBody struct {
	Data            string `json:"Data"`            //1.Encrypt Triple des using secret key
	FunctionName    string `json:"FunctionName"`    //2.Each api have it own functional name
	RequestDateTime string `json:"RequestDateTime"` //3.yyyy-MM-ddTHH:mm:ssZ
	RequestID       string `json:"RequestID"`       //4.Guid
}

// Request OTP: ERequestOTPAcc
type Request_ERequestOTPAcc struct {
	CustomerID     string `json:"CustomerID"`     //5.ID khách hàng tại ĐVCNTT
	SubscriptionID string `json:"SubscriptionID"` //6.Số tài khoản thẻ/token của thẻ/tài khoản đại diện cho số thẻ/tài khoản KH
	RefNumber      string `json:"RefNumber"`      //7.Số tham chiếu (Tiền tố 2 ký tự đầu STB sẽ quy định cho từng ĐVCNTT)
}

// Topup transaction: ETopupByAccount
type Request_ETopupByAccount struct {
	IsRequiredOTPBySTB bool   `json:"IsRequiredOTPBySTB"` //5.True hoặc False: Gửi lại giá trị đã được Sacombank trả về ở trường IsRequiredOTP bước ERequestOTP tương ứng
	IsByPassOTP        bool   `json:"IsByPassOTP"`        //6.True: Hệ thống ĐVCNT đánh giá khách hàng là an toàn, không cần xác thực OTP Và trong trường hợp này ĐVCNT chịu hoàn toàn rủi ro, trách nhiệm nếu bị khách hàng khiếu nại. False: Xác thực theo yêu cầu của Sacombank nếu có
	AuthType           string `json:"AuthType"`           //7.Hình thức xác thực của khách hàng: Gửi lại giá trị đã được Sacombank trả về ở trường AuthType bước ERequestOTP tương ứng: - 1: SMS OTP - 2: HardToken - 3: mCode - 4: mConnected - 5: AdvToken Nếu IsByPassOTP = True thì AuthType =””
	OTP                string `json:"OTP"`                //8.Mã OTP nếu có. Nếu IsByPassOTP = True thì OTP =””
	CustomerID         string `json:"CustomerID"`         //9.ID khách hàng tại ĐVCNTT
	SubscriptionID     string `json:"SubscriptionID"`     //10.Số tài khoản thẻ/token của thẻ/tài khoản đại diện cho số thẻ/tài khoản KH
	Amount             string `json:"Amount"`             //11.Số tiền giao dịch
	Description        string `json:"Description"`        //12.Diễn giải Format: [CIA Diễn giải của ĐVCNTT]
	RefNumber          string `json:"RefNumber"`          //13.Số tham chiếu (Tiền tố 2 ký tự đầu STB sẽ quy định cho từng ĐVCNTT)
	InqRefNumber       string `json:"InqRefNumber"`       //14.Số tham chiếu của giao dịch ERequestOTP tương ứng trước đó
}

// Purchase transaction: EPurchaseByAccount
type Request_EPurchaseByAccount struct {
	IsRequiredOTPBySTB bool   `json:"IsRequiredOTPBySTB"` //5.True hoặc False: Gửi lại giá trị đã được Sacombank trả về ở trường IsRequiredOTP bước ERequestOTP tương ứng
	IsByPassOTP        bool   `json:"IsByPassOTP"`        //6.True: Hệ thống ĐVCNT đánh giá khách hàng là an toàn, không cần xác thực OTP Và trong trường hợp này ĐVCNT chịu hoàn toàn rủi ro, trách nhiệm nếu bị khách hàng khiếu nại. False: Xác thực theo yêu cầu của Sacombank nếu có
	AuthType           string `json:"AuthType"`           //7.Hình thức xác thực của khách hàng: Gửi lại giá trị đã được Sacombank trả về ở trường AuthType bước ERequestOTP tương ứng: - 1: SMS OTP - 2: HardToken - 3: mCode - 4: mConnected - 5: AdvToken Nếu IsByPassOTP = True thì AuthType =””
	OTP                string `json:"OTP"`                //8.Mã OTP nếu có. Nếu IsByPassOTP = True thì OTP =””
	CustomerID         string `json:"CustomerID"`         //9.ID khách hàng tại ĐVCNTT
	SubscriptionID     string `json:"SubscriptionID"`     //10.Số tài khoản thẻ/token của thẻ/tài khoản đại diện cho số thẻ/tài khoản KH
	Amount             string `json:"Amount"`             //11.Số tiền giao dịch
	Description        string `json:"Description"`        //12.Diễn giải Format: [CIA Diễn giải của ĐVCNTT]
	RefNumber          string `json:"RefNumber"`          //13.Số tham chiếu (Tiền tố 2 ký tự đầu STB sẽ quy định cho từng ĐVCNTT)
	InqRefNumber       string `json:"InqRefNumber"`       //14.Số tham chiếu của giao dịch ERequestOTP tương ứng trước đó
}

// Withdraw transaction
// Inquiry: ESubscriptionInquiry
// Cashout: ECashoutSubscription
type Request_ESubscriptionInquiry struct {
	CustomerID     string `json:"CustomerID"`     //5.ID khách hàng tại ĐVCNTT
	SubscriptionID string `json:"SubscriptionID"` //6.Số tài khoản thẻ/token của thẻ/tài khoản đại diện cho số thẻ/tài khoản KH
	RefNumber      string `json:"RefNumber"`      //7.Số tham chiếu (Tiền tố 2 ký tự đầu STB sẽ quy định cho từng ĐVCNTT)
}
type Request_ECashoutSubscription struct {
	SubscriptionID string `json:"SubscriptionID"` //5.Số tài khoản thẻ/token của thẻ/tài khoản đại diện cho số thẻ/tài khoản KH
	CustomerID     string `json:"CustomerID"`     //6.ID khách hàng tại ĐVCNTT
	Amount         string `json:"Amount"`         //7.Số tiền thực hiện giao dịch
	RefNumber      string `json:"RefNumber"`      //8.Số tham chiếu (Tiền tố 2 ký tự đầu STB sẽ quy định cho từng ĐVCNTT)
	Description    string `json:"Description"`    //9.Diễn giải Thẻ: Fomat: [COC Diễn giải của ĐVCNTT] Tài khoản: Fomat: [COA Diễn giải của ĐVCNTT]
	SenderName     string `json:"SenderName"`     //10.Tên khách hàng thực hiện chuyển tiền
}

// Transfer to an account
// Inquery: ESTBAccountInquiry
// Transfer: EFundTransferToSTBAccount
type Request_ESTBAccountInquiry struct {
	ToSTBAccountNo string `json:"ToSTBAccountNo"` //5.Tài khoản STB nhận (nếu tài khoản chưa liên kết). ĐVCNTT không được phép lưu trữ full số tài khoản
	RefNumber      string `json:"RefNumber"`      //6.Số tham chiếu (Tiền tố 2 ký tự đầu STB sẽ quy định cho từng ĐVCNTT)
}

type Request_EFundTransferToSTBAccount struct {
	AccountNo    string `json:"AccountNo"`    //5.Tài khoản STB nhận (nếu tài khoản chưa liên kết). ĐVCNTT không được phép lưu trữ full số tài khoản
	Amount       string `json:"Amount"`       //6.Số tiền
	Description  string `json:"Description"`  //7.Diễn giải Fomat: [ISA Diễn giải của ĐVCNTT]
	RefNumber    string `json:"RefNumber"`    //8.Số tham chiếu (Tiền tố 2 ký tự đầu STB sẽ quy định cho từng ĐVCNTT)
	InqRefNumber string `json:"InqRefNumber"` //9.Số tham chiếu của giao dịch ESTBCardInquiry trước đó
	SenderName   string `json:"SenderName"`   //10.Tên khách hàng thực hiện chuyển tiền
}

// Cancel Link: ECancelSubscription
type Request_ECancelSubscription struct {
	CustomerID     string `json:"CustomerID"`     //5.ID khách hàng tại ĐVCNTT
	SubscriptionID string `json:"SubscriptionID"` //6.Số tài khoản thẻ/token của thẻ/tài khoản đại diện cho số thẻ/tài khoản KH
	RefNumber      string `json:"RefNumber"`      //7.Số tham chiếu (Tiền tố 2 ký tự đầu STB sẽ quy định cho từng ĐVCNTT)
}
