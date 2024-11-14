package models

type LetterDocument struct {
	Subject      string `json:"subject" validate:"required"`
	ReceiverName string `json:"receiverName" validate:"required"`
	Content      string `json:"content" validate:"required"`
	SenderName   string `json:"senderName" validate:"required"`
	FileName     string `json:"fileName" validate:"required"`
}
