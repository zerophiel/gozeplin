package gozeplin

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"time"
)

const (
	URI = "https://api.zeplin.io"
)

type Response struct {
	ID                       string    `json:"_id"`
	Email                    string    `json:"email"`
	Username                 string    `json:"username"`
	Status                   string    `json:"status"`
	PaymentPlan              string    `json:"paymentPlan"`
	EmailNotifications       bool      `json:"emailNotifications"`
	NotificationLastReadTime time.Time `json:"notificationLastReadTime"`
	Token                    string    `json:"token"`
	IntercomHash             string    `json:"intercomHash"`
}

func GetLoginToken(username string, password string) (*Response, error) {
	var response Response
	url := URI + "/users/login"
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"handle":"` + username + `", "password":"` + password + `"}`).Post(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body(), &response)
	return &response, err
}
