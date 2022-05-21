package main

import (
	envi "Vault/env"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	envi.SetEnvFile("../.env")
	e := os.Getenv("ADRESS")
	if e != "http://127.0.0.1:8200" {
		t.Errorf("Incorect read of env")
	}
}

func TestEnv2(t *testing.T) {
	envi.SetEnvFile("../.env")
	e := os.Getenv("PATHREAD")
	if e != "secret/data/abd" {
		t.Errorf("Incorect read of env")
	}
}
func TestEnv3(t *testing.T) {
	envi.SetEnvFile("../.env")
	e := os.Getenv("TOKEN")
	if e != "hvs.DPCpBiRQdMUMMoZdrzDvIceQ" {
		t.Errorf("Incorect read of env")
	}
}
