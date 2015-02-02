// portscan.go
// naive implementation of a basic tcp ports scanner
// I did that to learn Go ...
// meik

package main

import (
       "fmt"
       "strconv"
       "net"
       "os"
       )

func main() {
     if len(os.Args) < 2 {
     	fmt.Println("No host given")
	return
     }

     i := 1
     for i < 65535 {
     	 s :=  os.Args[1] + ":" + strconv.Itoa(i)
	 conn, err := net.Dial("tcp", s)
	 if err != nil {
	 } else {
	      fmt.Printf("Port %d open\n", i)
	 }
	 _ = conn
	 i = i + 1
     }
}
