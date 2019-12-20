package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	api "github.com/galihrivanto/wdms-api"
	"github.com/galihrivanto/wdms-api/store"
)

const (
	// token refersh in duration
	tokenRefreshPeriod  = 300 * time.Second
	deviceWatcherPeriod = 5 * time.Second
)

func main() {
	var (
		baseurl  string
		username string
		password string
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

	w := api.NewWatcher(c,
		store.NewInMemoryStore(),
		&api.WatcherOption{
			Username:           username,
			Password:           password,
			TokenRefreshPeriod: 1 * time.Hour,
			DeviceCheckPeriod:  5 * time.Second,
			OnAttendanceReceived: func(att api.Transaction, iclock api.IClock) {
				log.Println(att)
			},
		})

	w.Watch(context.Background())
}
