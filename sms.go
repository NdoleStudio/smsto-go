package smsto

import "time"

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

// SmsMessage is the details of an SMS message that was sent out
type SmsMessage struct {
	ID                int         `json:"id"`
	SecondaryID       string      `json:"_id"`
	CampaignID        interface{} `json:"campaign_id"`
	UserID            int         `json:"user_id"`
	SenderID          string      `json:"sender_id"`
	Message           string      `json:"message"`
	To                string      `json:"to"`
	Status            string      `json:"status"`
	ClientCost        float64     `json:"client_cost"`
	CallbackURL       string      `json:"callback_url"`
	ScheduledFor      interface{} `json:"scheduled_for"`
	Timezone          interface{} `json:"timezone"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	SentAt            string      `json:"sent_at"`
	MessageID         interface{} `json:"message_id"`
	SmsCount          int         `json:"sms_count"`
	FinalCallbackSent int         `json:"final_callback_sent"`
	IsAPI             int         `json:"is_api"`
	A2Code            interface{} `json:"a2_code"`
	OptOut            interface{} `json:"optout"`
	FailedReason      string      `json:"failed_reason"`
	Campaign          interface{} `json:"campaign"`
	Provider          interface{} `json:"provider"`
}
