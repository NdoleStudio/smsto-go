package smsto

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/NdoleStudio/smsto-go/internal/helpers"
	"github.com/NdoleStudio/smsto-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestSmsService_SendSingleRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	apiKey := "apiKey"
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, [][]byte{stubs.SmsSendSingle()}, &requests)
	client := New(WithBaseURL(server.URL), WithAPIKey(apiKey))
	msg := "This is test and \\n this is a new line"
	recipient := "+35799999999999"
	bypassOptOut := true
	senderID := "SMSto"
	callbackURL := "https://example.com/callback/handler"

	params := SmsSendSingleRequest{
		Message:      msg,
		To:           recipient,
		BypassOptOut: &bypassOptOut,
		SenderID:     &senderID,
		CallbackURL:  &callbackURL,
	}

	expectedRequest := map[string]interface{}{
		"message":       msg,
		"to":            recipient,
		"bypass_optout": bypassOptOut,
		"sender_id":     senderID,
		"callback_url":  callbackURL,
	}

	// Act
	_, _, err := client.SMS.SendSingle(context.Background(), &params)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, 1, len(requests))

	buf, err := ioutil.ReadAll(requests[0].Body)
	assert.NoError(t, err)

	assert.Equal(t, "Bearer "+apiKey, requests[0].Header.Get("Authorization"))

	requestBody := map[string]interface{}{}
	err = json.Unmarshal(buf, &requestBody)
	assert.NoError(t, err)

	assert.Equal(t, expectedRequest, requestBody)
}

func TestSmsService_SendSingleResponse(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.SmsSendSingle())
	client := New(WithBaseURL(server.URL))

	// Act
	result, response, err := client.SMS.SendSingle(context.Background(), &SmsSendSingleRequest{})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	assert.Equal(t, &SmsSendSingleResponse{
		Message:   "Message is queued for sending! Please check report for update",
		Success:   true,
		MessageID: "11ec-832f-a6f3fcfe-9fea-02420a0002ab",
	}, result)

	// Teardown
	server.Close()
}
