package main

import (
	envi "Vault/env"
	"fmt"
	"os"
	"testing"
)

func TestSetEnvFile(t *testing.T) {
	envi.SetEnvFile("../.env")
	e := os.Getenv("ADRESS")
	if e != "http://secureenv.poc-innovation.com" {
		t.Errorf("Unable to read [ADRESS] env variable")
	}
}

func TestSetEnvFile2(t *testing.T) {
	envi.SetEnvFile("../.env")
	e := os.Getenv("PATHREAD")
	if e != "secret/data/db" {
		t.Errorf("Unable to read [PATHREAD] env variable")
	}
}
func TestSetEnvFile3(t *testing.T) {
	envi.SetEnvFile("../.env")
	e := os.Getenv("TOKEN")
	if e != "hvs.C5rcwKI673pSTlwell5bRHCG" {
		t.Errorf("Unable to read [TOKEN] env variable")
	}
}

func TestGetEnvFile(t *testing.T) {
	env := envi.GetEnvFile("../.env")
	fmt.Println(env[0])
	if env[0] != "ADRESS=http://secureenv.poc-innovation.com" {
		t.Errorf("Unable to read [ADRESS] env variable")
	}
	if env[1] != "TOKEN=hvs.C5rcwKI673pSTlwell5bRHCG" {
		t.Errorf("Unable to read [TOKEN] env variable")
	}
}
