package main

import (
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// f, err := os.Create("/tmp/dat")
	// check(err)

	// buf := make([]byte, 1024*1024*128)
	// for i := 0; i < 1024*1024*128; i++ {
	// 	buf[i] = byte(i % 256)
	// }

	// for i := 0; i < 1024; i++ {
	// 	n2, err := f.Write(buf)
	// 	check(err)
	// 	fmt.Printf("%s wrote %d bytes\n", i, n2)
	// }
	// f.Close()

	// fmt.Println("file generated.")

	f, err := os.Open("/tmp/dat")
	check(err)
	b1 := make([]byte, 1024*1024*1024)

	i := 0
	for {
		t1 := time.Now()
		n1, err := f.Read(b1)
		check(err)
		if n1 == 0 {
			break
		}
		t2 := time.Now()
		fmt.Printf("%d: %d bytes %v\n", i, n1, t2.Sub(t1))
		i++
	}
}
