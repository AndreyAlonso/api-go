package handler

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

const (
	Error   = "error"
	Message = "message"
)

func newResponse(messageType string, messsage string, data interface{}) response {
	return response{
		messageType,
		messsage,
		data,
	}
}
