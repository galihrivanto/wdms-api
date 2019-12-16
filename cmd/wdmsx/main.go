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

	// set token
	log.Println("token", res1.Token)
	c.SetAuthToken(res1.Token)

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

	res3, _, err := c.DepartmentService.List(context.Background(),
		&api.DepartmentListRequest{
			ListRequest: api.ListRequest{
				Limit: 5,
			},
		},
	)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(res3)

	res4, _, err := c.AreaService.List(context.Background(),
		&api.AreaListRequest{
			ListRequest: api.ListRequest{
				Limit: 5,
			},
		},
	)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(res4)

	res5, _, err := c.EmployeeService.List(context.Background(),
		&api.EmployeeListRequest{
			ListRequest: api.ListRequest{
				Limit: 10,
			},
		},
	)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(res5)

	res6, _, err := c.IClockService.List(context.Background(),
		&api.IClockListRequest{
			ListRequest: api.ListRequest{
				Limit: 10,
			},
		},
	)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(res6)
}
