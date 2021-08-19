package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	NameDBMongo = "memoDB"

	MemoCollectionName = "memo"
)

type Memo struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`
	Content        string `json:"content"`
	CreationDate   string `json:"creation_date"`
	ExpirationDate string `json:"expiration_date"`
}

func (m *Memo) Create(content string) error {
	m.CreationDate, err = getParsedTime()
	if err != nil {
		return err
	}
	m.Content = content

	return nil
}

func getParsedTime() (string, error) {
	loc, err := time.LoadLocation("America/Buenos_Aires")
	if err != nil {
		return "", fmt.Errorf("error loading time location")
	}

	return time.Now().In(loc).Format(time.RFC3339), err
}
