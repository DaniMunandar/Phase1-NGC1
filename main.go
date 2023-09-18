package main

import "fmt"

func main() {

	// name := "Dani"
	// age := 30

	// fmt.Printf("nama nya adalah: %v  \n", name)
	// fmt.Printf("umur nya adalah: %v  \n", age)
	// fmt.Printf("%T, %T \n", name, age)

	// const address string = "jakarta"
	// fmt.Println(address)

	//Non Graded Challenges 1 Variable 1
	var myNum int32 = 50
	fmt.Printf("%T, %v \n", myNum, myNum)

	var myNum2 float32 = 51
	fmt.Printf("%T, %v, \n", myNum2, myNum2)

	var myNumStr = "50"
	fmt.Printf("%T, %v, \n", myNumStr, myNumStr)

	//Non Graded Challenges 1 Variable 2
	var x int32 = 5
	var y int32 = 10
	var z = x + y

	fmt.Println("Nilai Z : ", z)

	//NGC 1 - CLI
	var name string
	fmt.Println("Masukan Nama")
	fmt.Scanln(&name)

	fmt.Println("hello", name)

	//NGC 1 - Array & Slice 1
	people := []string{"Walt", "Jasse", "Skyler", "Saul"}
	fmt.Println(people)
	fmt.Println(len(people))
	people = append(people, "Hank", "Marie")
	fmt.Println(people)
	fmt.Println(len(people))

	//NGC 1 Arary & Slice 2
	people2 := [3]map[string]string{
		{
			"name":   "Hank",
			"gender": "M",
		},
		{
			"name":   "Hank",
			"gender": "M",
		},
		{
			"name":   "Hank",
			"gender": "M",
		},
	}

	fmt.Println("people :", people2)

	// Mengganti gender jadi "F" dan menambahkan "Mrs"
	for i := 0; i < len(people2); i++ {
		if people2[i]["gender"] == "F" {
			people2[i]["name"] = "Mrs. " + people2[i]["name"]
		} else {
			people2[i]["name"] = "Mr. " + people2[i]["name"]
		}
	}
	fmt.Println("people Baru :", people2)
}
