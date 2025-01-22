package config

import "testing"

func TestConfig(t *testing.T) {
	conf, err := Load("example-config.json")
	if err != nil {
		t.Fatal(err)
	}
	if conf == nil {
		t.Fatal("conf is nil")
	}
	if conf.MAL == nil {
		t.Fatal("conf.MAL is nil")
	}
	if conf.MAL.ClientID != "ID-api" {
		t.Errorf("conf.MAL.ClientID is wrong, expected '%s', got '%s'", "ID-api", conf.MAL.ClientID)
	}
	if conf.MAL.ClientSecret != "shhh" {
		t.Errorf("conf.MAL.ClinetSeecret is wong, expected '%s', got '%s'", "shhh", conf.MAL.ClientSecret)
	}
}
