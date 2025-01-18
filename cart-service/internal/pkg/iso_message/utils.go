package isomessage

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kingstonduy/go-core/logger"
	"github.com/moov-io/iso8583"
)

type Rule struct {
	FieldName  string `db:"FIELDNAME"`
	CheckNull  string `db:"CHECKNULL"`
	Length     int64  `db:"LENGTH"`
	Pattern    string `db:"PATTERN"`
	TransType  string `db:"TRANSTYPE"`
	Route      string `db:"ROUTE"`
	LengthType string `db:"LENGTHTYPE"`
}

func (i *IsoMessage) CheckHmac() bool {
	actual := i.GetField(128)

	expected := GenHMAC(&i.isoMsgLib, i.hmac)

	return actual == expected
}

func GenHMAC(isoMsg *iso8583.Message, hmacKey string) string {
	var macDefine = map[string]string{
		"2":   "LLVAR",
		"32":  "LLVAR",
		"48":  "LLLVAR",
		"63":  "LLLVAR",
		"102": "LLVAR",
		"103": "LLVAR",
		"120": "LLLVAR",
	}

	macList := []string{"2", "3", "4", "5", "6", "7", "11", "32", "37", "38", "39", "41", "42", "48", "63", "66", "90", "102", "103", "120"}

	data4Mac, err := isoMsg.GetMTI()
	if err != nil {
		panic(err)
	}

	fields := isoMsg.GetFields()

	for _, fieldId := range macList {
		fieldID, _ := strconv.Atoi(fieldId)
		field, _ := isoMsg.GetField(fieldID).String()

		_, ok := fields[fieldID]
		if ok {
			l := len(field)
			var len string
			if definition, ok := macDefine[fieldId]; ok {
				if definition == "LLVAR" {
					len = fillZero(l, 2)
					data4Mac = data4Mac + len + field
				} else if definition == "LLLVAR" {
					len = fillZero(l, 3)
					data4Mac = data4Mac + len + field
				}
			} else {
				data4Mac = data4Mac + field
			}
		} else {
			data4Mac = data4Mac + ""
		}
	}

	secretKey := []byte(hmacKey)
	algorithm := sha256.New
	mac := hmac.New(algorithm, secretKey)
	_, err = mac.Write([]byte(data4Mac))
	if err != nil {
		panic(err)
	}
	hmacBytes := mac.Sum(nil)
	hmacHex := hex.EncodeToString(hmacBytes)
	return strings.ToUpper(hmacHex)
}

func CheckRule(ctx context.Context, isoMsg *IsoMessage, ls []Rule) (isValid bool) {
	// check rule
	msg := isoMsg.isoMsgLib
	m := checkFormat(ls, parseIsoMessageToMap(&msg))
	if m["result"] != "1" {
		logger.Errorf(ctx, "Check rule failed: %s", m["resultlog"])
	}
	return m["result"] == "1"
}

func parseIsoMessageToMap(msg *iso8583.Message) map[string]string {
	isomsgMap := make(map[string]string)
	isomsgMap["MTI"], _ = msg.GetField(0).String()

	for i := 2; i <= 128; i++ {
		field, _ := msg.GetField(i).String()
		if i < 10 {
			isomsgMap["F0"+fmt.Sprintf("%d", i)] = field
		} else {
			isomsgMap["F"+fmt.Sprintf("%d", i)] = field
		}

		if i == 48 {
			f48Array := strings.Split(field, "\r")
			count := 0
			for _, f48Item := range f48Array {
				if count == 0 {
					isomsgMap["F481"] = f48Item
				}
				if count == 1 {
					isomsgMap["F482"] = f48Item
				}
			}
		}
	}

	return isomsgMap
}

func checkFormat(ls []Rule, isomsgMap map[string]string) map[string]string {
	transType := ""
	route := ""
	checkResult := ""
	resultMap := make(map[string]string)

	if strings.HasPrefix(isomsgMap["F03"][:2], "43") {
		transType = "QUERY"
	}
	if strings.HasPrefix(isomsgMap["F03"][:2], "91") {
		transType = "PAYMENT"
	}

	if isomsgMap["MTI"] == "0210" {
		route = "NHPL"
	}

	if isomsgMap["MTI"] == "0200" {
		route = "NHTH"
	}

	for _, rule := range ls {
		ruleTransType := rule.TransType
		ruleRoute := rule.Route
		ruleFieldName := rule.FieldName
		checkResult += "-" + ruleFieldName

		if (transType == ruleTransType || ruleTransType == "ALL") && (route == ruleRoute || ruleRoute == "ALL") {
			result := checkFormatByValue(rule, isomsgMap[ruleFieldName])
			checkResult += ":" + result
			if result == "30" {
				resultMap["result"] = ruleFieldName
				resultMap["resultlog"] = checkResult
				return resultMap
			}
		}
	}

	resultMap["result"] = "1"
	resultMap["resultlog"] = checkResult
	return resultMap
}

func checkFormatByValue(napasRule Rule, value string) string {
	checkNull := napasRule.CheckNull
	lengthType := napasRule.LengthType
	length := napasRule.Length
	fieldName := napasRule.FieldName
	pattern := napasRule.Pattern
	route := napasRule.Route

	if checkNull == "1" && value == "" {
		return "30"
	}

	if value != "" {
		if lengthType == "EQUAL" && len(value) != int(length) {
			return "30"
		}
		if lengthType == "MAX" && len(value) > int(length) {
			return "30"
		}
	}

	if fieldName == "F104" && pattern != "" {
		if route == "NHTH" || route == "ALL" {
			match, err := regexp.MatchString(pattern, value)
			if err != nil || match {
				return "30"
			}
		}
	}

	if fieldName == "F07" && pattern == "checkMMddHHmmss" {
		if !checkMMddHHmmss(value) {
			return "30"
		}
	}

	if fieldName == "F13" && pattern == "checkMMDD" {
		if !checkMMDD(value) {
			return "30"
		}
	}

	if fieldName == "F12" && pattern == "checkHHmmss" {
		if !checkHHmmss(value) {
			return "30"
		}
	}

	return "1"
}

// checkMMddHHmmss checks if the input string matches the pattern MMddHHmmss
func checkMMddHHmmss(value string) bool {
	// Define the regular expression pattern for MMddHHmmss
	pattern := `^(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01])(0[0-9]|1[0-9]|2[0-3])([0-5][0-9]){2}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}

// checkMMDD checks if the input string matches the pattern MMDD
func checkMMDD(value string) bool {
	// Define the regular expression pattern for MMDD
	pattern := `^(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01])$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}

// checkHHmmss checks if the input string matches the pattern HHmmss
func checkHHmmss(value string) bool {
	// Define the regular expression pattern for HHmmss
	pattern := `^(0[0-9]|1[0-9]|2[0-3])([0-5][0-9]){2}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}

func padLeft(str string, length int, padChar byte) string {
	for len(str) < length {
		str = string(padChar) + str
	}
	return str
}

func fillZero(len int, length int) string {
	if length == 2 && len < 10 {
		return "0" + strconv.Itoa(len)
	} else if length == 3 && len < 10 {
		return "00" + strconv.Itoa(len)
	} else if length == 3 && len < 100 && len >= 10 {
		return "0" + strconv.Itoa(len)
	} else {
		return strconv.Itoa(len)
	}
}
