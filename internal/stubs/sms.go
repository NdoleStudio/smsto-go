package stubs

// SmsSendSingle is a dummy json response for the `/api/airtime/v2/` endpoint
func SmsSendSingle() []byte {
	return []byte(`
		{
			"message": "Message is queued for sending! Please check report for update",
			"success": true,
			"message_id": "11ec-832f-a6f3fcfe-9fea-02420a0002ab"
		}
`)
}

// SmsLastMessage is a dummy JSON response for the `/last/message` endpoint
func SmsLastMessage() []byte {
	return []byte(`
		{
			"id": 302050741,
			"_id": "11ec-8330-c25f5870-a614-02420a0002ab",
			"campaign_id": null,
			"user_id": 342847,
			"sender_id": "SMSto",
			"message": "This is a test message",
			"to": "+35799999999999",
			"status": "REJECTED",
			"client_cost": 0,
			"callback_url": "https://example.com/callback/handler",
			"scheduled_for": null,
			"timezone": null,
			"created_at": "2022-02-01T07:29:59.000000Z",
			"updated_at": "2022-02-01T07:30:07.000000Z",
			"sent_at": "2022-02-01 07:30:07",
			"message_id": null,
			"sms_count": 1,
			"final_callback_sent": 0,
			"is_api": 1,
			"a2_code": null,
			"optout": null,
			"failed_reason": "Invalid Number",
			"campaign": null,
			"provider": null
		}
`)
}
