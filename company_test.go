package wdmsapi

import (
	"context"
	"log"
	"os"
	"testing"
)

func getEnv(key string, def ...string) string {
	v := os.Getenv(key)
	if v == "" {
		if len(def) > 0 && def[0] != "" {
			v = def[0]
		}
	}

	return v
}

func TestCompanyCrud(t *testing.T) {
	// TODO: this testing mean to create, check and delete company
	baseURL := getEnv("WDMS_URL", "https://wdms.magicsoft-asia.com/api")
	user := getEnv("WDMS_USER", "keke")
	password := getEnv("WDMS_PASSWORD", "123456")

	c, err := New(
		nil,
		SetBaseURL(baseURL),
	)
	if err != nil {
		t.Fatal(err)
	}

	tokenResult, _, err := c.TokenService.Create(context.Background(), &TokenRequest{Username: user, Password: password})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// set token
	log.Println("token", tokenResult.Token)
	c.SetAuthToken(tokenResult.Token)

	data := &Company{
		CompanyID: 99,
		Name:      "Wulong",
	}

	// create company
	id, _, err := c.CompanyService.Create(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}

	// get company back
	dataResult, _, err := c.CompanyService.Get(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}

	// check if equal
	if data.CompanyID != dataResult.CompanyID || data.Name != dataResult.Name {
		t.Error("Expected result doesn't match with inserted data")
	}

	// remove company
	_, err = c.CompanyService.Delete(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}

}
