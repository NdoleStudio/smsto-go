package stubs

// SmsSendSingle is a dummy json response for the `/api/airtime/v2/` endpoint with an error
func SmsSendSingle() []byte {
	return []byte(`
		{
			"message": "Message is queued for sending! Please check report for update",
			"success": true,
			"message_id": "11ec-832f-a6f3fcfe-9fea-02420a0002ab"
		}
`)
}
