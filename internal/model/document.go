package model

import "time"

// Document представляет модель документа
type Document struct {
	ID             uint            `gorm:"primaryKey" json:"id"`
	AcceptID	   uint			   `json:"acceptID`
	Title          string         `json:"title"`
	Owner          string         `json:"owner"`
	ReceivedTime   *time.Time     `json:"receivedTime"`
	SentTime       *time.Time     `json:"sentTime"`
	CreatedAt      time.Time      `json:"createdAt"`
	DeliveryStatus *DeliveryStatus `json:"deliveryStatus"`
	Status         Status         `json:"status"`
	Payload        string         `json:"payload" gorm:"type:text"`
	Files          []File         `json:"files"`
}

// Document который приходит
type DocumentRequest struct {
	ID             uint            `gorm:"primaryKey" json:"id"`
	Title          string         `json:"title"`
	Owner          string         `json:"owner"`
	SentTime       *time.Time     `json:"sentTime"`
	CreatedAt      time.Time      `json:"createdAt"`
	Payload        string         `json:"payload" gorm:"type:text"`
	Files         [][]byte         `json:"files"`
}