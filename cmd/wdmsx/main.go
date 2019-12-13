package main

import (
	"context"
	api "github.com/gophercode/wdmsapi"
	"log"
	"os"
)

func main() {
	c, err := api.New(
		nil,
		api.SetBaseURL("https://wdms.magicsoft-asia.com/api"),
	)
	if err != nil {
		panic(err)
	}

	res1, _, err := c.TokenService.Create(context.Background(), &api.TokenRequest{Username: "keke", Password: "123456"})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(res1)

	// override URL
	res2, _, err := c.CompanyService.List(context.Background(),
		&api.CompanyListRequest{
			ListRequest: api.ListRequest{
				Limit: 5,
			},
		},
	)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(res2)
}
