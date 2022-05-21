package main

import (
	envi "Vault/env"
	"os"
	"testing"
)

func TestSetEnvFile(t *testing.T) {
	envi.SetEnvFile("../.env")
	e := os.Getenv("ADRESS")
	if e != "http://127.0.0.1:8200" {
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

	if env[0] != "ADRESS=http://127.0.0.1:8200" {
		t.Errorf("Wron value")
	}
	if env[1] != "TOKEN=hvs.C5rcwKI673pSTlwell5bRHCG" {
		t.Errorf("Wron value")
	}
}
