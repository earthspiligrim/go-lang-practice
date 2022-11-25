package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

/*

	SECTION: CHANNELS AND GO ROUTINES
	  channels and go routines are both structures used to handeling concurrent programming

	What is concurrency?


	69. Website Status Checker
			func main() {
				links := []string{
					"http://google.com",
					"http://facebook.com",
					"http://stackoverflow.com",
					"http://golang.org",
					"http://amazon.com",
				}
				for _, link := range links {

				}
			}


	70. Printing Site Status
		func main() {
			links := []string{
				"http://google.com",
				"http://facebook.com",
				"http://stackoverflow.com",
				"http://golang.org",
				"http://amazon.com",
			}

			for _, link := range links {
				checkLink(link)
			}
		}

		func checkLink(link string) {
			_, err := http.Get(link)
			if err != nil {
				fmt.Println(link, "might be down!")
				return
			}
			fmt.Println(link, "is up!")
		}

		upon running the print output will be printing it out in the same sequence as defined in the splice

	71. Serial Link Checking
		when you actually run the above program upon some calls there is a slight delay (time to hit the req and respond back) while printing the output
		
		Additionally you are hitting that link and hanging in for the response and then hitting the second link and waiting on for the response and then the third and so on....,
		-- a sequential flow

		alternatively the better approach of dealing with this would be parallel execution of the requests rather than waiting on the response of each request.

	72. Go Routines

		go routine ante simple ga unna go code line by line execute cheyadam. so code anta compile ayyaka go routine anedi flow of execution ni line by line execute chestu untundi
		and mana code lo http.Get(link)--> e line response ravadaniki koncham time padutundi so dani "Blocking Call" ani pilustaru technically

		so manam ekkada aite manaki blocking call vastundo akada inkoka go routine run chestam which will run that piece and akada kanaka blocking call oste 
		control back main ki ichestundi to go run the rest and this new sub go routine can handle the current scenario of execution

		
	73. Theory of Go Routines
			Rules and side effects and edge cases of go routines.
			



*/
