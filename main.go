package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/vault/api"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func getEnvFile() []string {
	ftm, err := os.ReadFile("./.env")
	if err != nil {
		log.Fatalln(err)
	}
	ret := string(ftm)
	return strings.Split(ret, "\n")
}

func setEnvFile() {
	tmp := getEnvFile()
	for i := 0; i != len(tmp); i++ {
		val := strings.Split(tmp[i], "=")
		err := os.Setenv(val[0], val[1])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func parseReader(str string) string {
	res := strings.Replace(str, "map[", "", 1)
	res = strings.Replace(res, "]", "", 1)
	res = strings.Replace(res, ":", "=", 100)
	return res
}

func strToWordArray(str string) []string {
	p := parseReader(str)
	array := strings.Split(p, " ")
	return array
}

func cutToEqual(str string) []string {
	p := strings.Split(str, "=")
	return p
}

func createFile() *os.File {
	f, err := os.OpenFile("./.envrc", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	return f
}

func formatData(sev *api.Secret) []string {
	s := sev.Data["data"]
	se := fmt.Sprintf("%v", s)
	m := strToWordArray(se)
	return m
}

func writeData(f *os.File, arg []string, m []string) {
	if len(arg) == 2 {
		for i := 0; i != len(m); i++ {
			p := cutToEqual(m[i])
			if p[0] == arg[1] {
				f.WriteString(m[i] + "\n")
			}
		}
	} else {
		for i := 0; i != len(m); i++ {
			f.WriteString(m[i] + "\n")
		}
	}
}

func concurency(n time.Duration, client *api.Client, arg []string) {
	for range time.Tick(n * time.Second) {
		f := createFile()
		data, err := client.Logical().Read(os.Getenv("PATHREAD"))
		if err != nil {
			panic(err)
		}
		m := formatData(data)
		writeData(f, arg, m)
		println("Data saved in .envrc")
	}
}

func main() {

	arg := os.Args
	setEnvFile()
	token := os.Getenv("TOKEN")
	vaultAddr := os.Getenv("ADRESS")

	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		log.Fatalln(err)
	}
	client.SetToken(token)
	go concurency(10, client, arg)
	select {}
}
