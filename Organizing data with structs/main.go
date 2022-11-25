package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

type person2 struct {
	firstName string
	lastName  string
	contactInfo2
}

type contactInfo2 struct {
	email string
	zip   int
}

func main() {
	//one of the multiple ways of creating a persion is below
	//the way go interprets this is the sequence of strings go by the
	//sequence in which the variables are defined in the struct
	//so like wise Alex goes with the firstName and Anderson goees to the lastName
	//Not a recommended way since this could hit if the order of the definition changes

	// alex := person{"Alex", "Anderson", {"abc@gmail.com", 3444} }
	// fmt.Println(alex.firstName)

	//SECONDWAY
	//an efficient way to define the struct
	alex2 := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex2.firstName)
	fmt.Println(alex2.lastName)

	//******VIDEO 40 updating the values of the struct
	//say you declare a struct and do not define the values to the variables
	//go usually defines the values with the Zero(default) values
	/*
	   Type    Zero Value
	   string  ""
	   int     0
	   float   0
	   bool    false
	*/

	alex2.firstName = "Alexandria"
	fmt.Println(alex2.firstName)
	fmt.Println("%+v", alex2)

	//video 41: embedding structs

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "abc@gmail.com",
			zip:   3443,
		},
	}

	fmt.Printf("%+v", jim)

	jim.print()
	jim.updatedFirstname("Jimmy")

}

func (p person) updatedFirstname(updatedFirstname string) {
	p.firstName = updatedFirstname
}
func (p person) print() {
	fmt.Println("%+v", p)
}

/*
   Struct is nothing but a structure. in simeple a data structure
   holding a few propertites just like a class

   Video 38 Defining Struct
   Wheenver we make a struct we need to define all the properties that a struct can take place
   Struct of person below

   type person struct {
   firstName string
   lastName  string
   }

   Video 39 Declaring Structs
       //one of the multiple ways of creating a persion is below
   //the way go interprets this is the sequence of strings go by the
   //sequence in which the variables are defined in the struct
   //so like wise Alex goes with the firstName and Anderson goees to the lastName
   //Not a recommended way since this could hit if the order of the definition changes

   alex := person{"Alex", "Anderson"}

   //SECONDWAY
   //an efficient way to define the struct
   alex2 := person{firstName: "Alex", lastName: "Anderson"}
   fmt.Println(alex2.firstName)
   fmt.Println(alex2.lastName)

   Video 40 Updating Struct Values

           //******VIDEO 40 updating the values of the struct
           //say you declare a struct and do not define the values to the variables
           //go usually defines the values with the Zero(default) values
               Type    Zero Value
               string  ""
               int     0
               float   0
               bool    false

       type persion struct {
           firstName string
           lastName string
       }

       func main(){
           var alex person
           // first ida oka default values set chestundundi
           alex.firstName = "Alex"
           alex.lastName = "Anderson"
       }

   VIDEO 41 EMBEDDING STRUCTS
       Ability to embedd one struct inside another

       type persion                    type ContactInfo
       firstName -> string                 email -> string
       lastName -> string                  zip -> int
       contact -> contactInfo

           type person struct {
           firstName string
           lastName  string
           contact   contactInfo
           }

           type contactInfo struct {
               email string
               zip   int
           }
           //remember after every property declation you need to add ,
           jim := person{
               firstName: "Jim",
               lastName:  "Party",
               contact: contactInfo{
                   email: "abc@gmail.com",
                   zip:   3443,
               },
           }
           fmt.Printf("%+v", jim)


   VIDEO 42: STRUCTS WITH RECEIVER FUNCTION
           type person struct {
           firstName string
           lastName  string
           contactInfo  //->creating a receiver to the contactInfo type which will make it work the same way
           }

           type contactInfo struct {
               email string
               zip   int
           }

           //works the same way as before
               jim := person{
               firstName: "Jim",
               lastName:  "Party",
               contactInfo: contactInfo{
                   email: "abc@gmail.com",
                   zip:   3443,
               },
               }
           fmt.Printf("%+v", jim)

           //additionally creating a receiver
           func (p person) print() {
               fmt.Printf("%+v", p)
           }
           func (p person) updateFirstName(updatedFirstname string) {
               p.firstName = updatedFirstname
           }
           func main(){
               jim := person{
               firstName: "Jim",
               lastName:  "Party",
               contactInfo: contactInfo{
                   email: "abc@gmail.com",
                   zip:   3443,
               },
               }
               jim.print()

               jim.updateFirstName("Jimmy")
			   //the update will not happen like above
			   //go to the next lecture to know how to update it
           }

   VIDEO 43: Pass By Value
		Pointers in Go:
		   	Go is a pass by value language. heres how the flow goes in go
			1. when you create a new peron (jim) go will allocate the object in a memory location say 001
			2. when you call update firstname like in 205 then go does the fllowing
		   		i. makes a copy all the person type into another memeory locaiton say 0004
				ii. updates the firstname to the record at the address location 0004
		   		iii. thats the reason why the update of the firstname is not being reflecting
		   how to update the older one go to the next lecture

   VIDEO 44: Structs with Pointers (*******COMEBACK*******)

		   updating our code can make it happen


           func main(){
               jim := person{
               firstName: "Jim",
               lastName:  "Party",
               contactInfo: contactInfo{
                   email: "abc@gmail.com",
                   zip:   3443,
               },
               }
               jim.print()

			   //CHANGE 1
			   //when you say &jim --> here we turn the &jim to a memory pointer and assign it to the jimpointer
			   jimPointer := &jim
			  		 //jim.updateFirstName("Jimmy")
			   //instead of jim use the pointer to the jim
			   jimPointer.updateFirstName("Jimmy")

           }

		   //CHANGE 2
		   //*person is saying that type of person pointing to a pointer
		   //*person : is a type description. it means were working with a pointer to a person

			func (pointertoPerson *person) updateFirstName(updatedFirstname string) {
				//CHANGE 3:
               //p.firstName = updatedFirstname
			   //*pointerToPerson : this is an operator. it means we want to manipulate the value the pointer is referencing
			   (*pointertoPerson).firstName = updatedFirstName
           }

   VIDEO 45: Pointer Operations (*******COMEBACK*******)
   //explanation of the before code
		   two type of ways
		   &variable : give me memory address of the value this variable is pointing at
		   *pointer : give me the value this memory address is pointing at
		   KeyNote(important):
		   Turn address into value with *address
		   Turn value into address with &value
		   * before type --> i need address of the value
		   * before variable --> fetch me the value of the address to the value where i need to either read or update it


   VIDEO 46: Pointer Shortcut (*******COMEBACK*******)
		// go has a shortcut to implement the below code
			jimPointer := &jim
			jimPointer.updateFirstName("Jimmy")


			func (pointertoPerson *person) updateFirstName(updatedFirstname string) {
				(*pointerToPerson).firstName = updatedFirstName
			}
		//shortcut implementation
		   the code even works now
		   jim.updateFirstName("Jimmy")
		   //as long as the receiver is *person the call made by
		   //jim.update..... will work the same way as we do it by
		   //using the pointers
		   func (pointertoPerson *person) updateFirstName(updatedFirstname string) {
				(*pointerToPerson).firstName = updatedFirstName
			}


   VIDEO 47: Gotchas with pointers

   package main
   import ("fmt")

   func main(){
		mySlice := []string{"Hi", "There", "How", "Are", "You"}

		updateSlice(mySlice)
		fmt.Println(mySlice)
   }

   func updateSlice(s []string) {
		s[0] = "Bye"
   }
   //without using pointers the above code will update the myslice array
   //to know why go to the next lecture


   VIDEO 48: Reference vs Value Types (*******COMEBACK*******)

   	Basic differences between arrays and slices
   			Arrays								Slice
	1. Primitive data structure					1. Can grow and shrink
	2. Can't be resized							2. Used 99% of the time for lists of elements
	3. Rarely used directly

	Upon running this line (mySlice := []string{"Hi", "There", "How", "Are", "You"})
	Go creates three lements for the slice
	1. pointer to head --> points to an array storing ["Hi", "There", "How", "Are", "You"]
	2. capacity
	3. length
   THE array is store in an address say 00034
   another address is used to store the slice information consisting of belo
   		pointer -> 00034
		capacity
		length
   	all the above 3 are stored into one address location

   so when you modify the slice here is what happens
   	1. clones up the entire slice data into a different address
	2. updates the arrray, where in both the original and the clone of the slice points to the same array location
   		the update is reflected to both the slices

		So all the reference type of elements behave in the same way

   		Value Types							Reference Types
			int									slices
			float								maps
			string								channels
			bool								pointers
			structs								functions

		Value types: (use pointers to change these things in a function)
   		Reference Types: (dont worry about the pointers with these)

*/
