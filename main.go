package main

import (
	envi "Vault/Env"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/hashicorp/vault/api"
)

var ErrKeyNotFound = errors.New("Key not found")

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

type Connection interface {
	NewClient() (*Getter, error)
}

type Getter struct {
	Cli *api.Client
}

func (g *Getter) NewClient() (*Getter, error) {
	vaultAddr := os.Getenv("ADRESS")
	var err error

	g.Cli, err = api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		log.Fatalln("errors New client")
	}
	return g, err
}

func callApi(c Connection) (*Getter, error) {
	val, err := c.NewClient()
	return val, err
}

func createFile() *os.File {
	f, err := os.OpenFile("./.envrc", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	return f
}

func retrieveData(sev *api.Secret) (map[string]interface{}, error) {
	s, ok := sev.Data["data"]
	if !ok {
		return nil, errors.New("could not find data field")
	}

	switch t := s.(type) {
	case map[string]interface{}:
		return t, nil
	default:
		return nil, errors.New("Errors occurs on types: not map[string]interface")
	}
}

func writeData(f *os.File, m map[string]interface{}) {
	for k, v := range m {
		f.WriteString(fmt.Sprintf("%s=%s\n", k, v))
	}
}

func concurency(n time.Duration, client *api.Client) error {
	for range time.Tick(n * time.Second) {
		f := createFile()
		data, err := client.Logical().Read(os.Getenv("PATHREAD"))
		if err != nil {
			return ErrKeyNotFound
		}
		m, err := retrieveData(data)
		if err != nil {
			return errors.Wrap(err, "In retrieveData")
		}
		writeData(f, m)
		println("Data saved in .envrc")
	}
	return nil
}

func main() {
	envi.SetEnvFile()
	token := os.Getenv("TOKEN")
	g := &Getter{}
	client, err := callApi(g)

	if err != nil {
		log.Fatalln(err)
	}
	client.Cli.SetToken(token)
	if err = concurency(5, client.Cli); err != nil {
		log.Fatalf("%+v\n", err)
	}
}
