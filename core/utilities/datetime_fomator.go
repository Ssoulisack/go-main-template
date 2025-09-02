package utilities

import (
	"errors"
	"fmt"
	"kkl-v2/api/middleware"
	"kkl-v2/core/logs"
	"net/http"
	"time"
)

var dateTimeFormat = "2006-01-02 15:04:05"
var layoutDateTime = "2006-01-02 15:04:05.000000 -0700 -07 MST m=+0.000000000"

func NewDateTimeFormatToString(datetime time.Time) string {
	return datetime.Format(dateTimeFormat)
}

func DateTimeFormat(datetime string) (time.Time, bool) {
	// Define the layout of your input string
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, datetime)
	if err != nil {
		fmt.Println("Error:", err)
		return time.Now(), false
	}
	return t, true
}

func DateTimeFormatFromString(datetime string) (*time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, datetime)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func NewStringFormatToDateTime(datetime string) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return time.Time{}, middleware.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("["+datetime+"]"+"INVALID_DATETIME_FORMAT_EX:(yyyy-mm-dd hh:mm:ss)"))
	}
	t, err := time.ParseInLocation(dateTimeFormat, datetime, loc)
	if err != nil {
		logs.Error(err)
		return time.Time{}, middleware.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("["+datetime+"]"+"INVALID_DATETIME_FORMAT_EX:(yyyy-mm-dd hh:mm:ss)"))
	}

	// format the time value
	s := t.Format(layoutDateTime)
	//TODO: format layout := "2006-01-02 15:04:05.000000 -0700 -07 MST m=+0.000000000"
	//TODO: input 2023-04-05 17:28:04 output 2023-04-05 17:28:04 +0700 +07
	parse, err := time.Parse(layoutDateTime, s)
	if err != nil {
		return time.Time{}, middleware.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("["+datetime+"]"+"INVALID_DATETIME_FORMAT_EX:(yyyy-mm-dd hh:mm:ss)"))
	}
	return parse, nil
}
