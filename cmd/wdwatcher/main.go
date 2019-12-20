package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	api "github.com/galihrvanto/wdms-api"
)

const (
	// token refersh in duration
	tokenRefreshPeriod  = 300 * time.Second
	deviceWatcherPeriod = 5 * time.Second
)

func main() {
	var (
		baseurl   string
		username  string
		password  string
		startDate string
	)

	flag.StringVar(&baseurl, "baseurl", "https://wdms.magicsoft-asia.com/api", "wdms login username")
	flag.StringVar(&username, "username", "keke", "wdms login username")
	flag.StringVar(&password, "password", "123456", "wdms login password")
	flag.Parse()

	c, err := api.New(
		nil,
		api.SetBaseURL(baseurl),
	)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	dt, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		dt = time.Now()
	}

	w := NewWatcher(c, &api.WatcherOption{
		Usename: username,
		Password: password,
		TokenRefreshPeriod: 1 * time.Hour,
		DeviceCheckPeriod: 5 * time.Second,
	})

	w.Watch(context.Background())
}
