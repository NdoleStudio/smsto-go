package smsto

import (
	"context"
	"encoding/json"
	"net/http"
)

// smsService is the API client for the `/sms` endpoint
type smsService service

// SendSingle Sends a single SMS message to a phone number
//
// API Docs: https://developers.sms.to/#0eff8134-6ad5-4b61-86f3-28ff18145bfc
func (service *smsService) SendSingle(
	ctx context.Context,
	params *SmsSendSingleRequest,
) (*SmsSendSingleResponse, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodPost, "/sms/send", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	status := new(SmsSendSingleResponse)
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}

// LastMessage Gets the last message that you have sent
//
// API Docs: https://developers.sms.to/#80f14590-14cd-4512-9c8c-cf8c7ef07f96
func (service *smsService) LastMessage(ctx context.Context) (*SmsMessage, *Response, error) {
	request, err := service.client.newRequest(ctx, http.MethodGet, "/last/message", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	status := new(SmsMessage)
	if err = json.Unmarshal(*response.Body, status); err != nil {
		return nil, response, err
	}

	return status, response, nil
}
