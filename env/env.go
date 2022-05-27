package Env

import (
	"log"
	"os"
	"strings"
)

func GetEnvFile(path string) []string {
	ftm, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	ret := string(ftm)
	return strings.Split(ret, "\n")
}

func SetEnvFile(path string) {
	tmp := GetEnvFile(path)
	for i := 0; i != len(tmp); i++ {
		val := strings.Split(tmp[i], "=")
		err := os.Setenv(val[0], val[1])
		if err != nil {
			log.Fatal(err)
		}
	}
}
