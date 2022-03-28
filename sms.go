package smsto

// SmsSendSingleResponse is the response after sending a single SMS message
type SmsSendSingleResponse struct {
	Message   string `json:"message"`
	Success   bool   `json:"success"`
	MessageID string `json:"message_id"`
}

// SmsSendSingleRequest are parameters for sending a single SMS message
type SmsSendSingleRequest struct {
	Message      string  `json:"message"`
	To           string  `json:"to"`
	BypassOptOut *bool   `json:"bypass_optout,omitempty"`
	SenderID     *string `json:"sender_id,omitempty"`
	CallbackURL  *string `json:"callback_url,omitempty"`
}
