package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func sendMail() {

	from := "john.doe@example.com"
	to := "root@core9"
	name := "John Doe"
	subject := "Hello"
	body := "Hello there"

	host := "core9:25"

	con, err := net.Dial("tcp", host)
	checkError(err)

	req := "HELO core9\r\n" +
		"MAIL FROM: " + from + "\r\n" +
		"RCPT TO: " + to + "\r\n" +
		"DATA\r\n" +
		"From: " + name + "\r\n" +
		"Subject: " + subject + "\r\n" +
		body + "\r\n.\r\n" + "QUIT\r\n"

	_, err = con.Write([]byte(req))
	checkError(err)

	res, err := ioutil.ReadAll(con)
	checkError(err)

	fmt.Println(string(res))
}

func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
