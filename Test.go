package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"todo/actions"
	"os"
)

func main() {
	app := &cli.App{
		Name: "Todo",
		Usage: "Wait...i can't remember the usage of this app!",
		Action: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "new",
				Action: func(c *cli.Context) error {
					err := actions.New(c)
					if err != nil{
						fmt.Println("Fallo...exitosamente???")
					}
					return nil
				},
			},
			{
				Name: "get",
				Action: func(c *cli.Context) error{
					err := actions.Get(c)
					if err != nil{
						fmt.Println("Fallo...exitosamente???")
					}
					return nil
				},
			},
			{
				Name: "list",
				Action: func(c *cli.Context) error{
					err := actions.List(c)
					if err != nil{
						fmt.Println("Fallo...exitosamente???")
					}
					return nil
				},
			},
			{
				Name: "delete",
				Action: func(c *cli.Context) error{
					err := actions.Delete(c)
					if err != nil{
						fmt.Println(err.Error())
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}