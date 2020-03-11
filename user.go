package gozeplin

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"time"
)

type UserDetails struct {
	ID      string `json:"_id"`
	Members []struct {
		Role       string   `json:"role"`
		Restricted bool     `json:"restricted"`
		Tags       []string `json:"tags"`
		User       struct {
			ID       string    `json:"_id"`
			Email    string    `json:"email"`
			Username string    `json:"username"`
			Status   string    `json:"status"`
			LastSeen time.Time `json:"lastSeen"`
		} `json:"user"`
	} `json:"members"`
}

func GetOrganizationUsers (token string, organizationId string) (*UserDetails,error){
	var response UserDetails
	url := URI + "/v2/organizations/" + organizationId + "/members"
	client := resty.New()
	resp,err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("zeplin-token",token).Get(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp.Body(), &response)
	return &response, err
}

func RemoveOrganizationUser (token string, organizationId string, userId string) error {
	url := URI + "/organizations/"+ organizationId + "/members/"+userId
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("zeplin-token", token).Delete(url)
	if err != nil {
		return err
	}
	return nil
}