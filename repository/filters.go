package repository

import (
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func FiltersGet(args string) bson.D{
	var filter bson.D
	searchEqualValue := "holis"

	if args != ""{
		searchEqualValue = args
	}

	filter = bson.D{{
		"content",
		bson.M{
			"$eq" :searchEqualValue,
		},
	}}

	return filter
}

func FiltersListAll() bson.D{
	return bson.D{{
		"content",
		bson.M{
			"$ne" : "",
		},
	}}
}

func FilterByObjectID(oid primitive.ObjectID) bson.D{
	return bson.D{{
		"_id",
		bson.M{
			"$eq" : oid,
		},
	}}
}

func GetAllMemos(MemoCollection *mongo.Collection,memos []*Memo, c *cli.Context)[]*Memo{
	filter := FiltersListAll()

	getResult, err := MemoCollection.Find(c.Context,filter)
	if err != nil{
		log.Println(err)
	}

	for getResult.Next(c.Context) {
		var memo Memo

		err := getResult.Decode(&memo)

		if err != nil {
			log.Fatal(err)
		}

		memos = append(memos, &memo)
	}

	return memos
}
