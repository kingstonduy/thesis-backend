### Example tao isomessage chuyen tien

```
func genExecuteIsoMessage[T any](ctx context.Context) *isomessage.IsoMessage {

    isoMSg := isomessage.NewIsoMessage(hmacKey)

	// 0200 - request; 0210 – response
	isoMSg.SetMTI("0200")

	//DE #2: Primary Account Number
	if string_utils.IsEmpty(debitAccount) {
		isoMSg.SetField(2, "00"+time_utils.GetTimeYYYYMMDDHHMMSS(now))
	} else {
		isoMSg.SetField(2, debitAccount)
	}

	//DE #3: Processing Code
	//DE #3: Processing Code
	var processingCode string = "9120"
	if !string_utils.IsEmpty(paymentMethod) {
		if paymentMethod == "ACCOUNT" {
			processingCode = processingCode + "20"
		} else {
			processingCode = processingCode + "00"
		}
	} else {
		processingCode = processingCode + "00"
	}
	isoMSg.SetField(3, processingCode)

	//DE #4: Transaction Amount
	if amount != 0 {
		isoMSg.SetField(4, string_utils.PadLeft(fmt.Sprintf("%d00", int(amount)), 12, '0'))
	} else { // default amount = 0
		isoMSg.SetField(4, string_utils.PadLeft(fmt.Sprintf("%d00", 0), 12, '0'))
	}

	//DE #7: Transmission Date and Time
	isoMSg.SetField(7, time_utils.GetTimeGmtMMDDHHMMSS(now))

	//DE #11: System Trace Audit Number
	f11 := getSystemTrace()
	isoMSg.SetField(11, f11)

	//DE #12: Local Transaction Time
	isoMSg.SetField(12, time_utils.GetTimeHHMMSS(now))

	//DE #13 Local Transaction Date
	isoMSg.SetField(13, time_utils.GetTimeMMDD(now))

	//DE #15 Local Transaction Date
	isoMSg.SetField(15, time_utils.GetTimeMMDD(now))

	//DE #18: Merchant Category Code
	if !string_utils.IsEmpty(merchanCatergoryCode) {
		isoMSg.SetField(18, merchanCatergoryCode)
	} else {
		// default value
		isoMSg.SetField(18, "7399")

	}
	//DE #22: PAN mode
	isoMSg.SetField(22, "000") //(Không biết chế độ PAN được nhập vào + Không biết được khả năng nhập số PIN của thiết bị đầu cuối)

	//DE #25: service condition
	isoMSg.SetField(25, "05") //(Khách hàng có mặt tại nơi giao dịch nhưng không có thẻ)

	//DE #32: Acquiring Instititution Code
	isoMSg.SetField(32, "970448") //970448 for OCB

	//DE #37: Retrieval reference number
	isoMSg.SetField(37, time_utils.GetTimeYDDDHHNNNNNN(now, f11))

	//DE #38: Authorization identification response
	isoMSg.SetField(38, "103698") //Mã chuẩn chi của tổ chức nhận lệnh, ghi nhận response thì sẽ có 1 số khác.

	//DE #41 Card Acceptor Terminal Identification
	isoMSg.SetField(41, "00000001") //Giá trị xác định thiết bị chấp nhận thẻ

	//DE #42 Card Acceptor Identification Code
	isoMSg.SetField(42, "000000000000001") //Mã xác định đơn vị chấp nhận thẻ

	//DE #43 Card Acceptor Name/Location
	//Địa điểm thực hiện
	if !string_utils.IsEmpty(acceptorNameAndLocation) {
		isoMSg.SetField(43, string_utils.PadRight(acceptorNameAndLocation, 40, ' '))
	} else { //default
		isoMSg.SetField(43, string_utils.PadRight("OCB EBANKING", 22, ' ')+" "+string_utils.PadRight("TP HCM", 13, ' ')+" VNM")
	}

	//DE #48: Additional private data
	//Dữ liệu cá nhân bổ sung, lấy ra tên người gửi từ fromAccountNumber
	if !string_utils.IsEmpty(sender) {
		isoMSg.SetField(48, fmt.Sprintf("%s\r", sender))
	} else {
		isoMSg.SetField(48, "\r")
	}

	//DE #49 (Currency Code) = 704 (VND)
	//Mã tiền tệ giao dịch, VN là 704
	isoMSg.SetField(49, "704")

	//DE #60: Self – defined Field
	//Kênh thực hiện giao dịch
	// - 00: Không xác định
	// - 01: ATM
	// - 02: Counter (Quầy giao dịch)
	// - 03: POS
	// - 04: Internet Banking
	// - 05: Mobile Application
	// - 06: SMS Banking
	// - 07: Kênh khác
	if !string_utils.IsEmpty(paymentChannel) {
		if paymentChannel == "IB" {
			isoMSg.SetField(60, "04")
		} else if paymentChannel == "QR" {
			isoMSg.SetField(60, "99")
		} else { // default
			isoMSg.SetField(60, "04")
		}
	} else {
		isoMSg.SetField(60, "04")
	}

	//DE #62 Service Code
	// Mã dịch vụ của NAPAS
	// IF_INQ:Truy vấn thông tin chủ
	// thẻ/ tài khoản thụ hưởng
	// IF_DEP:Chuyển tiền tới chủ thẻ/
	// tài khoản thụ hưởng
	// BL_INQ:Truy vấn thông tin hóa
	// đơn
	// BL_DEP: Thanh toán hóa đơn
	// Default: IF_DEP
	isoMSg.SetField(62, "IF_DEP")

	//DE #63 Transaction Reference Number
	//Mã tham chiếu của napas trả ra khi tra cứu ở bước /get-napas-recipient-name
	if !string_utils.IsEmpty(napasRefNum) {
		isoMSg.SetField(63, napasRefNum)
	}

	//DE #100: Receiving Institution Identification Code
	//Mã ngân hàng nhận lệnh
	if !string_utils.IsEmpty(benBankID) {
		isoMSg.SetField(100, benBankID)
	}

	//DE #102: From Account Identificati
	//Tài khoản nguồn
	if !string_utils.IsEmpty(debitAccount) {
		isoMSg.SetField(102, debitAccount)
	}

	//DE #103: To Account Identification
	//Tài khoản đích
	if !string_utils.IsEmpty(benAccount) {
		isoMSg.SetField(103, benAccount)
	}

	//DE #104 : Content transfer
	//Nội dung chuyển tiền
	if !string_utils.IsEmpty(paymentDetail) {
		isoMSg.SetField(104, paymentDetail)
	}

	//DE #120 : Tên người thụ hưởng trả ra khi tra cứu ở bước /get-napas-recipient-name
	if !string_utils.IsEmpty(benCustomerName) {
		isoMSg.SetField(120, benCustomerName)
	}

	//DE #128 : HMAC
	isoMSg.GenHMAC()

	return isoMSg
```

### parse from string to struct (gui tu napas adapter or napas) sang struct

```
	hmac := `7235CFDE06BD96B8D4811B655C36594D`
	s := "04150200F23A44810CE1801600000000170001011600371000232980019120000000040000001106073323363836143323110611067399000050697044843110736383610369800000001000000000000001OCB EBANKING           TP HCM        VNM033ACCOUNT.TITLE.1-0037100023298001\r7040020406IF_DEP0184129OCBB880430610606970406160037100023298001165218xxxxxxxx9999015chuyentiennapas013NGUYENVANTEST535CFB4570535085643435A3A9156F927A0F7EEC8F8EF777577688909EE14A51"
	isoMessage, err := isomessage.ToIsoMessage(s, hmac)
```

### parse from struct to string (tu service cua ocb gui di napas)

```
	hmac := `7235CFDE06BD96B8D4811B655C36594D`
    isoMSg := isomessage.NewIsoMessage(hmac)
    return isoMSg.ToString()


```
