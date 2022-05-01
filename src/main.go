package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func getFileContent(fileName string) ([]string, error) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	arrayVar := SplitLines(string(fileContent))
	return arrayVar, err
}

func printEnvVariables(contentFile []string) {
	for _, i := range contentFile {
		fmt.Println(i)
	}
}

func argsManagement(args []string) {
	if len(args) < 2 || len(args) > 5 {
		log.Fatalln("Precise your command: \n\t[get] for read all variables present in .envrc\n\t[set] for set variables")
	}
}

func secureEnv(args []string, envVar []string) {
	if args[1] == "get" {
		if len(args) == 2 {
			printEnvVariables(envVar)
		} else if len(args) == 3 {
			for i := 0; i != len(envVar); i++ {
				s := strings.Split(envVar[i], "=")
				if s[0] == args[2] {
					println(s[0] + "=" + s[1])
				}
			}
		}
	} else if args[1] == "set" && len(args) == 4 {
		toSet := concatArgs(args)
		os.Setenv(toSet[0], toSet[1])
	}
}

func concatArgs(args []string) []string {
	saveVar := args[2]
	saveVar += "=" + args[3]
	saveCuttedVar := strings.Split(saveVar, "=")
	return saveCuttedVar
}

func main() {
	args := os.Args
	argsManagement(args)

	content, err := getFileContent("./.envrc")
	if err != nil {
		log.Fatalln(".envrc not found !")
	}
	secureEnv(args, content)
	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
}
