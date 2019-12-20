package wdmsapi

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/galihrivanto/runner"
)

// DeviceStatusChanged is callback whenever
// device status changed from offline to online
// or otherwise
type DeviceStatusChanged func(IClock, bool)

// AttendanceReceived is callback when watcher
// received new attendance record
type AttendanceReceived func(Transaction)

// Device represent biometric device
// which currently being watched
type Device struct {
	IClock

	// last attendance records
	LastAttDate Time
}

// WatcherOption provide watcher setting
type WatcherOption struct {
	// auth
	Username           string
	Password           string
	TokenRefreshPeriod time.Duration

	// data watching
	DeviceCheckPeriod time.Duration

	OnDeviceStatusChanged DeviceStatusChanged
	OnAttendanceReceived  AttendanceReceived
}

// Watcher watch WDMS api for data changes
type Watcher struct {
	// lock
	mu *sync.Mutex

	client  *Client
	options *WatcherOption

	// current token
	token string

	// internal list of registered devices
	devices sync.Map
}

func (w *Watcher) setToken(token string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.client.SetAuthToken(token)
	w.token = token
}

func (w *Watcher) getToken() string {
	w.mu.Lock()
	defer w.mu.Unlock()

	return w.token
}

// autoRefreshToken periodically try to refresh current token
func (w *Watcher) autoRefreshToken(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(w.options.TokenRefreshPeriod):
			newToken, _, err := w.client.TokenService.Refresh(ctx, &RefreshTokenRequest{Token: w.getToken()})
			if err != nil {
				log.Println(err)
				break
			}
			w.setToken(newToken.Token)
			w.client.SetAuthToken(newToken.Token)
		}
	}
}

func (w *Watcher) fetchDeviceAttRecords(ctx context.Context, iclock IClock) error {
	// get from device list
	// if not exists then put new information
	device := Device{
		IClock: iclock,

		// set default att time, from device last sync
		LastAttDate: Time{Time: time.Now(), Offset: 8},
	}

	v, loaded := w.devices.LoadOrStore(iclock.SN, device)
	if loaded {
		// cast v to device
		device, _ = v.(Device)
	}

	// take 10 records at one time
	result, _, err := w.client.TransactionService.List(ctx, &TransactionListRequest{
		ListRequest: ListRequest{
			Limit: 10,
		},
		SN:        iclock.SN,
		StartDate: device.LastAttDate,
		EndDate:   Time{Time: time.Now(), Offset: 8},
	})
	if err != nil {
		return err
	}

	// record last att
	if len(result.Data) > 0 {
		for _, att := range result.Data {
			fmt.Println(att.Pin, att.SN, att.Verify, att.CreateTime.Time)
		}

		device.LastAttDate = result.Data[len(result.Data)-1].Time
		w.devices.Store(iclock.SN, device)
	}

	return nil
}

func (w *Watcher) checkDevices(ctx context.Context) error {
	result, _, err := w.client.IClockService.List(ctx, &IClockListRequest{
		ListRequest: ListRequest{
			Limit: 1000, // considering all
		},
	})
	if err != nil {
		return err
	}

	if result.Data != nil && len(result.Data) > 0 {
		for _, iclock := range result.Data {
			if err := w.fetchDeviceAttRecords(ctx, iclock); err != nil {
				log.Println("err:", err)
			}
		}
	}

	return nil
}

// Watch start polling data from wdms server
func (w *Watcher) Watch(root context.Context) {
	ctx, cancel := context.WithCancel(root)
	defer cancel()

	token, _, err := w.client.TokenService.Create(ctx, &TokenRequest{
		Username: w.options.Username,
		Password: w.options.Password,
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	w.setToken(token.Token)

	// run token refresh job
	go w.autoRefreshToken(ctx)

	runner.
		Run(ctx, func(ctx context.Context) error {
			fmt.Println("Start Watching...")

			for {
				select {
				case <-ctx.Done():
					return nil
				case <-time.After(w.options.DeviceCheckPeriod):
					if err := w.checkDevices(ctx); err != nil {
						log.Println("err:", err)
					}
				}
			}
		}).
		Handle(func(sig os.Signal) {
			if sig == syscall.SIGHUP {
				return
			}

			log.Println("Shutting down...")
			cancel()
		})
}

// NewWatcher initialize new wdms watcher
func NewWatcher(client *Client, options *WatcherOption) *Watcher {
	if options == nil {
		options = &WatcherOption{}
	}

	return &Watcher{
		client:  client,
		options: options,
	}
}
