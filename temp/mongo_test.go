package temp

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestConnectMongoDb(t *testing.T) {
	ConnectMongoDb()
}

func TestInsertOne(t *testing.T) {
	tm := time.Now()
	data := Howie{
		HowieId:     primitive.NewObjectID(),
		Name:        "yzs",
		Pwd:         "123",
		Age:         7,
		CreateTime:  tm.UnixNano(),
		ExpiredTime: tm.Add(time.Hour * 7),
	}
	InsertOne(&data)
}
