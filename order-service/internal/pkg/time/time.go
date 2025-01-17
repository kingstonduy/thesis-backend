package time_utils

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/kingstonduy/go-core/logger"
)

func padLeft(str string, length int, padChar byte) string {
	for len(str) < length {
		str = string(padChar) + str
	}
	return str
}

func GetTimeYYYYMMDDHHMMSS(t time.Time) string {
	return t.Format("20060102150405") // YYYYMMDDHHMMSS
}

func GetTimeGmtMMDDHHMMSS(t time.Time) string {
	return t.In(time.FixedZone("GMT", 0)).Format("0102150405") // MMDDHHMMSS
}

func GetTimeHHMMSS(t time.Time) string {
	return t.Format("150405") // HHMMSS
}

func GetTimeMMDD(t time.Time) string {
	return t.Format("0102") // MMDD
}

func GetTimeYDDDHHNNNNNN(t time.Time, f11 string) string {
	lastLetterYear := strconv.Itoa(t.Year() % 10)
	dayOfYear := strconv.Itoa(t.YearDay())

	dayOfYear = padLeft(dayOfYear, 3, '0')

	f37 := fmt.Sprintf("%s%s%s%s", lastLetterYear, dayOfYear, t.In(time.FixedZone("GMT", 0)).Format("15"), f11)

	return f37
}

func ConvertTimeToUnixMilli(ctx context.Context, timeStr string) (int64, error) {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		logger.Errorf(ctx, "Failed to parse time to unix mili")
		return 0, err
	}
	timeRes, err := time.ParseInLocation("2006-01-02 15:04:05.000000", timeStr, location)
	if err != nil {
		logger.Errorf(ctx, "Failed to parse time to unix mili")
		return 0, err
	}
	return timeRes.UnixMilli(), nil
}
