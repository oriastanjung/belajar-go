package main // ini bisa diganti namanya

// std library seperti iostream pada c++
import (
	"fmt"
)

func main() { // harus ada function main di code
	fmt.Println("Hello Dunia!!!!")

	// variables syntax
	// var namaVar tipedata
	var whatToSay string
	whatToSay = "Orias tanjung nihhh"

	var myAge int = 29

	fmt.Println("\n" + whatToSay)
	fmt.Println(myAge)
	fmt.Println(saySomething())

	// cara untuk copy tipe dari suatu function tanpa di declare lagi menggunakan :=
	whatSomeSaid := saySomething()

	fmt.Println("dengan := ", whatSomeSaid)

	someSaid, othersSaid := sayNew()

	fmt.Println(someSaid, othersSaid)

	var angka1, angka2 int
	angka1 = 15
	angka2 = 30
	sum := sum(angka1, angka2)

	fmt.Println("hasil dari sum = ")
	fmt.Println(sum)
	fmt.Println("\n")
	x := 0
	for i := 1; i <= 10; i++ {
		x += i
	}

	fmt.Println(x)
}

// function synntax
// func namaFunction tipeReturn
func saySomething() string {
	return "Something"
}

func sayNew() (string, string) {
	return "Something new", "else"
}

func sum(x int, y int) int {
	fmt.Println(x)
	fmt.Println(y)
	return x + y
}
