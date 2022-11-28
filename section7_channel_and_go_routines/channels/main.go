package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	//this is how we create a new channel which takes in of type string and this channel is scopped within the main block of funcion and
	//not available for any other blocks could be checkLink function for an instance

	for _, link := range links {
		go checkLink(link, c) ///-----------> this is the only change (adding go to the ahead of the line) and passing in the channel as the second argument
	}

	//fmt.Println(<-c) //printing out the messages received from the

	//creating the same number of receivers as the same number of the links in the slice

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c) //program execution halts for the first time when it runs <-c and continues creating more only when the first receiver created encounters a response from the the first created channel.
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second) //wait for five seconds
			checkLink(link, c)
		}(l)
	}
}

// adding anothe argument of channel below so that this channel be used for all the communication if a go routine is been executed
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down!" //returning back a message back through the channel saying that the communication has been successful
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link //returning back a message back through the channel saying that the communication has been successful

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

		Parallelism - Multiple threads executed at the exact same time. Required multiple CPU's

		Example of parallelism
							--------------------						--------------------
							|	One CPU Core.   |						|	One CPU Core.   |
							--------------------						---------------------
									|.  ^										|.  ^
									V.  |										V.  |
							------------------------					--------------------------
							|. Pick one Go Routine | 					|	Pick one Go Routine	  |
							------------------------ 					---------------------------
							^		 ^		    ^						^		 ^		    ^
							|		 |		    |						|		 |		    |
							V		 V		    V						V		 V		    V
						----------	---------   ---------			----------	---------   ---------
						| Go	 |  | Go	|  | Go.    |			| Go	 |  | Go	|  | Go.    |
						| Routine|	|Routine|. | Routine|			| Routine|	|Routine|. | Routine|
						----------  ---------  -----------			----------  ---------  -----------


			Main Routine: Main routine is created when we launched our program
			Child go Routine: Child routines created by the 'go' keyword

			Normally when you add "go" infront of the function call (go checkLink(link)) here it what happens. the main go routine kicks of and creates small child go routines
				and exits the code even before the child go routine finishes the execution. in order to tell the main routine about the created child go routines we need **CHANNELS**

	74. Channels
			Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

			Channels: Channels create communication between Main Routine and many different Child go routines

										-------------------
									    | Child go routine |
										--------------------
												^
												|
												V
				-----------------		-----------------		   --------------------
				| Main routine  |. <--> |.  Channel		|   <---> |. Child go routine |
				-----------------		-----------------         ---------------------
												^
												|
												|
												V
										-------------------
										| Child go routine |
										--------------------

				Channel here communicates between all of the child go routine and the main go routine and all of the Child go routine's and is the source for the main routine to know if all of the child go routines
				have finished the blocking call within it. Additionally,

						(https://gobyexample.com/channels)
						Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
						Create a new channel with make(chan val-type). Channels are typed by the values they convey. Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
						(https://gobyexample.com/channels)

							package main
							import "fmt"

							func main() {
								messages := make(chan string)

								go func() { messages <- "ping" }()

								msg := <- messages
								fmt.Println(msg)
							}

						Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine. The <- channel syntax receivces a value from the channel.
						Here we'll receive the "ping" message we sent above and print it out. When we run the program the "ping" message is successfully passed from one goroutine to another via our channel. By default sends and receives block until both the sender and receiver are ready.
						This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.

	75. Channel Implementation

		func main() {
			links := []string{
				"http://google.com",
				"http://facebook.com",
				"http://stackoverflow.com",
				"http://golang.org",
				"http://amazon.com",
			}

			c := make(chan string)
			//this is how we create a new channel which takes in of type string and this channel is scopped within the main block of funcion and
			//not available for any other blocks could be checkLink function for an instance

			for _, link := range links {
				go checkLink(link, c) ///-----------> this is the only change (adding go to the ahead of the line) and passing in the channel as the second argument
			}

			fmt.Println(<-c) //printing out the messages received from the channel
		}

		// adding anothe argument of channel below so that this channel be used for all the communication if a go routine is been executed
		func checkLink(link string, c chan string) {
			_, err := http.Get(link)
			if err != nil {
				fmt.Println(link, "might be down!")
				c <- "Might be down!" //returning back a message back through the channel saying that the communication has been successful
				return
			}

			c <- "is up!" //returning back a message back through the channel saying that the communication has been successful

			fmt.Println(link, "is up!")
		}


		//Output here is that this is just going to print "http://google.com" is up for any number of times we execute this
		its just priting only one communication and executing

	76. Blocking Channels

		a little reglance is that http.Get(..) is a blocking io call similarly in the main () function post initialization of links
		when u say so checkLink() --> a new go routine is been called off from the main (default) routine.

		and when the main routine sees fmt.Println(<-c) // here at <-c the main routine find this as a blocking and makes the main routine sleep till it hears back from the channel

		back at the child go routine from the old. and the child routine goes sleep when it hits a blocking call at http.get(...)
		and when the child go routine resolves (after receiving a response from the get call) the child awakes and when it returns i mean returns a routine c <- "..." the main wakes up and resumes the program and terminates.

		******TIP: Watch Video 76 from 4:30 mins ******************

		A hacky solution for this is since the program is terminating upon receiving the first link response and when we have only 1 fmt.Println(<-c) link out there
		we can add a few more (<-c) lines to tell the main routine to hang in till all the other responses are back. so say in the slice we have about
		5 links and for which if we write 5 fmt.Println(<-c) statements then all the program would only terminate only when all the 5 links responses are been fetched

		But say we have about 5 links in the slice and we write the fmt.Println(<-c) for 6 times or any number of times which is more than 5 then the main routine would never terminate and keep hanging on for the
		6th one where we have just created 5 child go routines and resolve only 5 of them

		Another solution in the next video

	77. Receiving Messages

		Takeaway fromt he last video: Receiving from a channel (<-c) is a blocking action.
			simple hack is to use for loop to write the fmt.Println(<-c) for the same number of links we we got
				for i:= 0; i<len(links); i++ {
					fmt.Println(<-c)
				}

				//here in the for loop a point to observe is everytime the go run time sees "<-c" the program halts. and likewise the program halts when it sees
				//<-c while running the for loop for the first time and creates another <-c only after the first channel reception is successful
				// and only then it creates a receiver for the second one and from there on...


			func main() {
				links := []string{
					"http://google.com",
					"http://facebook.com",
					"http://stackoverflow.com",
					"http://golang.org",
					"http://amazon.com",
				}

				c := make(chan string)
				//this is how we create a new channel which takes in of type string and this channel is scopped within the main block of funcion and
				//not available for any other blocks could be checkLink function for an instance

				for _, link := range links {
					go checkLink(link, c) ///-----------> this is the only change (adding go to the ahead of the line) and passing in the channel as the second argument
				}

				//fmt.Println(<-c) //printing out the messages received from the

				//creating the same number of receivers as the same number of the links in the slice

				for i := 0; i < len(links); i++ {
					fmt.Println(<-c) //program execution halts for the first time when it runs <-c and continues creating more only when the first receiver created encounters a response from the the first created channel.
				}
			}

			// adding anothe argument of channel below so that this channel be used for all the communication if a go routine is been executed
			func checkLink(link string, c chan string) {
				_, err := http.Get(link)
				if err != nil {
					fmt.Println(link, "might be down!")
					c <- "Might be down!" //returning back a message back through the channel saying that the communication has been successful
					return
				}

				c <- "is up!" //returning back a message back through the channel saying that the communication has been successful

				fmt.Println(link, "is up!")
			}


	78. Repeating Routines

		Objective: say google.com ki request aipoyindi so apudu immediately manam inkoka channel create cheyali hitting the same google.com similarly parallel ga
		vere links ki kuda back to back hits ni ivvali one after the other.

		So manam mana checkLink function lo channel return ping response message to patu link kuda return cheyali adi mana main lo unna for loop lo catch chesi inkoka channel ni initiate cheyali

			func main() {
				links := []string{
					"http://google.com",
					"http://facebook.com",
					"http://stackoverflow.com",
					"http://golang.org",
					"http://amazon.com",
				}

				c := make(chan string)
				//this is how we create a new channel which takes in of type string and this channel is scopped within the main block of funcion and
				//not available for any other blocks could be checkLink function for an instance

				for _, link := range links {
					go checkLink(link, c) ///-----------> this is the only change (adding go to the ahead of the line) and passing in the channel as the second argument
				}
				for {   //(for {} --> is an infinite loop)
					go checkLink(<-c, c)
				}
			}
			func checkLink(link string, c chan string) {
				_, err := http.Get(link)
				if err != nil {
					fmt.Println(link, "might be down!")
					c <- link
					return
				}
				c <- link
				fmt.Println(link, "is up!")
			}

	79. Alternative Loop Syntax -- foreach

			for l := range c {
				go checkLink(l, c)
			}

			Notes from : https://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go

			A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, or values received on a channel. For each entry it assigns iteration values to corresponding iteration variables and then executes the block.

			for index, element := range someSlice {
				// index is the index where we are
				// element is the element from someSlice for where we are
			}

			//If you don't care about the index, you can use _:
			//The underscore, _, is the blank identifier, an anonymous placeholder.
			//https://go.dev/ref/spec#Blank_identifier
			for _, element := range someSlice {
    			// element is the element from someSlice for where we are
			}

	80. Sleeping a Routine
		Lets put a pause after every successful fetch call been made. ante a time of pause between any two google.com calls
			Time package in golang is a good option to use where we have sleep function out there, https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/time/sleep.go;l=9


			for l := range c {
				time.Sleep(time.Second)// pauses for 1 second and to pause the code for 5 seconds time.Sleep(5 * time.Second)
				go checkLink(l, c)
			}
		// but putting the sleep in the main function pauses up the entire execution for 5 seconds say your have just received a response from the channel of google.com and the main sleeps for 5 seconds and even if within this 5 seconds if stackoverflow.com finishes up its execution the main is stil asleep causing adding a pause over the blocking call which is made.


	81. Function Literals

		Function literals in go are nothing but Anonymous function in JS.
			Brush of JS
			//annonymous
				const increment = function() {
					counter+= this.val;
				}
			// arrow
				const increment = () => {
					counter+= this.val;
				}

				obj = new Obj()
				obj.val = 2
				obj.increment(2)

		instead of
		for l := range c {
			go checkLink()
		}

		use

		for l := range c {
			go func() {
				time.Sleep(5* time.Second)
				checkLink(l, c) // comes with a warning (range variable l captured by func literal)
			}()
		}

	82. Channels Gotchas

		upon running the above code some only facebook.com is been called most of the time this is because the scoping issue
		the l anedi main routine lo define ayi undi and danini inkoka go routine lo vadutunapudu value changes anedi avachu
		so simple ga aa l ni argument ga function literal loki pampiste saripotundi

		Always Avoid: Avoid using the variable from the parent routine inside a child go routine.

		for l := range c {
			go func(link string) {
				time.Sleep(5* time.Second)
				checkLink(link, c) // comes with a warning (range variable l captured by func literal)
			}(l)
		}










*/
