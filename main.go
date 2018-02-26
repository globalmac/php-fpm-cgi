package main

import (
	"log"
	"io/ioutil"
	"strings"
	"fmt"
	"time"
	"github.com/tomasen/fcgi_client"
)

var (
	phpVersion   = "7.2"
	viaSocket    = false
	host	       = "127.0.0.1"
	port	       = 9000
	timeout	     = int(1000)
	responseType = "full&json"
)

func main() {
	status := GetStatus()
	fmt.Println(status)
}

func GetStatus() string {

	connect := fmt.Sprintf("%s:%d", host, port)
	env := make(map[string]string)
	env["SCRIPT_FILENAME"] = "/status"
	env["SCRIPT_NAME"] = "/status"
	env["SERVER_SOFTWARE"] = "go"
	env["REMOTE_ADDR"] = connect
	env["QUERY_STRING"] = responseType

	if viaSocket {

		return RunCgi(env)

	} else {

		return RunTcp(connect, env)

	}

}

func RunCgi(env map[string]string) string {

	fscgi, err := fcgiclient.Dial("unix", "/var/run/php/php"+phpVersion+"-fpm.sock")

	if err != nil {
		log.Fatal("Error! Cannot connect via PHP-FPM socket!\n", err)
	}

	resp, err := fscgi.Get(env)

	if err != nil {
		log.Fatal("Error! Cannot send request via PHP-FPM socket!\n", err)
	}

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error! Cannot read response via PHP-FPM socket!\n", err)
	}

	return strings.TrimSpace(string(content))

}

func RunTcp(connect string, env map[string]string) string {

	fscgi, err := fcgiclient.DialTimeout("tcp", connect, time.Duration(timeout)*time.Millisecond)

	if err != nil {
		log.Fatal("Error! Cannot connect via TCP!\n", err)
	}

	resp, err := fscgi.Get(env)

	if err != nil {
		log.Fatal("Error! Cannot send request via TCP!\n", err)
	}

	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(content))

}
