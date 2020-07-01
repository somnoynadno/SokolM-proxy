package api

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

const URL = "https://sokolmeteo.com"
var loginPayload string

func init() {
	loginPayload = os.Getenv("login_payload")
	if loginPayload == "" {
		panic("no login payload")
	}
}

func Login() (*http.Response, error) {
	data := []byte(loginPayload)

	req, err := http.NewRequest("POST", URL + "/platform/api/user/login", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// set content type
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 30}
	return client.Do(req)
}

func GetDailyData(deviceID string, cookies []*http.Cookie) (*http.Response, error) {
	endDate   := time.Now().AddDate(0, 0, 1)
	startDate := time.Now().AddDate(0, 0, -1)

	endDateFormatted := fmt.Sprintf("%d.%d.%d",
		endDate.Day(), endDate.Month(), endDate.Year())
	startDateFormatted := fmt.Sprintf("%d.%d.%d",
		startDate.Day(), startDate.Month(), startDate.Year())

	postfix := "/api/analytics/record?" +
		"deviceId=" + deviceID + "&" +
		"startDate=" + startDateFormatted + "&" +
		"endDate=" + endDateFormatted + "&" +
		"parameters=EVS,UVI,L,LI,RSSI,RN,TRR,TR2,t,WD,HM,WV,WM,UV,Upow,PR1,PR,KS,V,TP,TR,AN9,WV2,td"

	req, err := http.NewRequest("POST", URL + postfix, nil)
	if err != nil {
		return nil, err
	}

	// attach access cookies
	for _, c := range cookies {
		req.AddCookie(c)
	}

	client := &http.Client{Timeout: time.Second * 60}
	return client.Do(req)
}