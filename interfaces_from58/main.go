package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, error := http.Get("https://google.com")

	if error != nil {
		fmt.Println("Error", error)
		os.Exit(1)
	} else {
		//not a good way to handle responses but jsut codint it
		//this way to test
		fmt.Println(resp)
	}
}

/*
	VIDEO 58 THE HTTP PACKAGE
		Net package: package net provides portable interface for network I/O, including TCP/IP,UDP domain name resolution and Unix domain sockets

		function def below

		https://pkg.go.dev/net/http#Get
		func Get(url string) (resp *Response, err error) {}

			resp, error := http.Get("https://google.com")

		if error != nil {
			fmt.Println("Error", error)
			os.Exit(1)
		} else {
			//not a good way to handle responses but jsut codint it
			//this way to test
			fmt.Println(resp)
		}

	VIDEO 59 READING THE DOCS
		func Get(url string) (resp *Response, err error) {}

		https://pkg.go.dev/net/http#Response
		The *Response is a pointer and is a type of struct which contains fields like
			Status string, StatusCode int, Proto string, ProtoMajor int, ProtoMinor int
			Header header,
			Body io.ReadCloser,
			ContentLength int64



	VIDEO 60 MORE INTERFACE SYNTAX

	VIDEO 61: INTERFACE REVIEW

	VIDEO 62: THE READER INTERFACE


	60-62 summary:

	say oka requirement undi like the following
			i. reading from a file and returning a string[]
			ii. reading data from a disk and returning a byte[]
			iii. reading data from xyz source and retrun some abc[] type
		ilanti requirmenent apudu generic ga manam individual sources ga implementaiton rastamu ala kakunda interfaces anedi enti ante
		oka common source which can implement all the different scenarios and return something like [] byte array (or any other common medium)
		such thins i reader interface

		you can technically say that you are building abstraction around a few implementation as a type known to be an interface
		so that all of the implementation differing the multiple scenarios can implement the common interface returning a common return type like []byte


		type Reader interface {
			Read(p []byte) (n int, err error)
		}

	VIDEO 63: More on the Reader Interface


*/
