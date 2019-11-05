package temp

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Howie struct {
	//struct里面获取ObjectID
	HowieId     primitive.ObjectID `bson:"id"`
	Name        string
	Pwd         string
	Age         int64
	CreateTime  int64
	ExpiredTime time.Time
}

const uri = "mongodb://192.168.5.25:27017"

//connection
func ConnectMongoDb() *mongo.Collection {
	opt := options.Client().ApplyURI(uri)
	out, _ := json.Marshal(opt)
	fmt.Printf("options:%v\n", string(out))
	opt.SetLocalThreshold(time.Second * 3)
	opt.SetMaxConnIdleTime(time.Second * 3)
	opt.SetMaxPoolSize(100)
	ctx, concel := context.WithTimeout(context.Background(), time.Second*10)
	defer concel()
	cli, err := mongo.Connect(ctx, opt)
	if err != nil {
		fmt.Printf("mongo.Connect,error(%v)\n", err)
		return nil
	}
	out, _ = json.Marshal(cli)
	fmt.Printf("client:%v\n", string(out))
	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("ping:error(%v)\n", err)
		return nil
	}
	fmt.Printf("ping ok... \n")
	//选择数据库和集合
	collection := cli.Database("testing_base").Collection("howie_info")
	fmt.Printf("collection:%v,name:%v\n", &collection, collection.Name())
	return collection
}

func InsertOne(document interface{}) {
	coll := ConnectMongoDb()
	inserRes, err := coll.InsertOne(context.Background(), document)
	if err != nil {
		fmt.Printf("insertone error(%v)\n", err)
		return
	}
	out, _ := json.Marshal(inserRes)
	fmt.Printf("insertRes:%v\n", string(out))
}
