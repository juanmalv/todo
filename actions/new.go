package actions

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"todo/console"
	"todo/repository"
	"strconv"
)

var MemoCollection = repository.Client.Database(repository.NameDBMongo).Collection(repository.MemoCollectionName)

func New(c *cli.Context) error{
	req := console.ReadFromConsole()

	memo := repository.Memo{}
	err := memo.Create(req)
	if err != nil{
		return err
	}

	insertResult, err := MemoCollection.InsertOne(c.Context, memo)

	if err != nil{
		log.Println(err)
		fmt.Println("NOOOO")
	}

	if insertResult != nil{
		log.Println("Memo saved correctly with ID: ", insertResult.InsertedID)
	}

	return nil
}

func Get(c *cli.Context) error{
	var memos []*repository.Memo

	filter := repository.FiltersGet(c.Args().First())

	getResult, err := MemoCollection.Find(c.Context,filter)
	if err != nil{
		log.Println(err)
	}

	for getResult.Next(c.Context) {
		var memo repository.Memo
		err := getResult.Decode(&memo)

		if err != nil {
			log.Fatal(err)
		}

		memos = append(memos, &memo)
	}

	if len(memos) > 0{
		log.Println(memos[0].Content)
	}

	return nil
}

func List(c *cli.Context) error{
	var memos []*repository.Memo

	memos = repository.GetAllMemos(MemoCollection,memos,c)

	for i, memo := range memos {
		index := strconv.Itoa(i+1)
		fmt.Println(index + ". " + memo.Content)
	}

	fmt.Println("--------------")
	return nil
}

func Delete(c *cli.Context) error{
	var memos []*repository.Memo

	memos = repository.GetAllMemos(MemoCollection,memos,c)

	fmt.Println("Which one do you wish to delete? Type the number of the memos listed above")
	for i, memo := range memos {
		index := strconv.Itoa(i+1)
		fmt.Println(index + ". " + memo.Content)
	}

	id := console.ReadFromConsole()
	println(id)
	indexNumber, err := strconv.Atoi(id)
	if err != nil{
		return fmt.Errorf("error: id provided is not a number. %v", err)
	}

	objectID := memos[indexNumber-1].Id

	result, err := MemoCollection.DeleteOne(c.Context,repository.FilterByObjectID(objectID))

	if result != nil{
		fmt.Println(result.DeletedCount)
	}

	return nil
}
