package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	args, err := getArgs()
	if err != nil {
		log.Fatal(err)
	}
	res, err := handler(os.Args[1], args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func getArgs() ([]string, error) {
	args := os.Args[2:]
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		args = append(args, scan.Text())
		fmt.Println(args)
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	return args, nil
}

func handler(command string, args []string) (string, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", errors.New(stderr.String())
	}
	return out.String(), nil
}
