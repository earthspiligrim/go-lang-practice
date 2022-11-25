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

			You can create any number of "go routines" in the flow of execution of yours. Even though you have any number of go routines at any point of time
			only one go routine will execute and which go routine to execute and when is decided by "Go Scheduler"

							--------------------
							|	One CPU Core.   |. **Even in the case of a dual core by *default* Go tries to use one core!
							--------------------
									|.  ^
									V.  |
							----------------------	***Note: Scheduler runs one routine until it finishes
							|	Go Scheduler	  | *** or makes a blocking call(like an HTTP request)
							----------------------- *** Meaning below
							^		 ^		    ^
							|		 |		    |
							V		 V		    V
						----------	---------   ---------
						| Go	 |  | Go	|  | Go.    |
						| Routine|	|Routine|. | Routine|
						----------  ---------  -----------

			Explanation: Say the go scheduler is running "Go Routine 1" where "Go Scheduler" will run till the "Go Ruoutine1" is attained either finished state or
			a state of blocking call, then once either of the states are attained then it would pause the "Go Routine 1" and starts executing the second or the next in line


			By default Go runs on the one core yet if we change that behaviour of that setting(explanation below)


			--------------------		--------------------		---------------------
			| One CPU Core		|		| One CPU Core		|		| One CPU Core		|
			|					|		|					|		|					|
			--------------------		--------------------		---------------------
				|		^					|		^					|		^
				V		|					V		|					V		|
			--------------------------------------------------------------------------------
			|							GO SCHEDULER										|.  ** Scheduler runs one thread on each "logical" core
			---------------------------------------------------------------------------------
				^							^								^
				|							|								|
				V							V								V
			-----------------			-----------------				-----------------
			| Go Routine	|			| Go Routine	|				| Go Routine	|
			-----------------			-----------------				-----------------

		here we have multiple cpus and multiple go routines so here
		By default Go runs on the one core yet if we change that behaviour of that settings to run on multiple cpu cores
		then the go scheduler instead of monitoring the go routines and running them one after the other go scheduler can run multiple go routines parallelly with the help
		of multiple cpu core's

		There is a lot of discussions happening in the go world weather to run the go scheduler on one cpu core instance or the multiple. (come back and check it out)

		Concurrency (my knemonic: Synchronization) - We can have multiple threads executing code. If one thread blocks another one is picked up and worked on
		Parallelism
		--stopped at 6:03


*/
