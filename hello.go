package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"math/rand"
	"math"
)

func main() {
	fmt.Printf("hello, world\n")
	
	fmt.Println("The time is", time.Now());
	
	fmt.Println("And if you try to open a file:")
	fmt.Println(os.Open("filename"))

	fmt.Println("Or access the network:")
	fmt.Println(net.Dial("tcp", "google.com"))

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("Now you have %g problems.", math.Nextafter(2, 3))
	fmt.Println(math.Pi)
}
