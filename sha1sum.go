// sha1sum.go
// Just a naive implementation of a sha1sum using golang
// but not using stdin when no file is passed as argument
// meik / at / moulinette.org
//

package main

import(
	"fmt"
	"os"
	"crypto/sha1"
)

func main() {
     if len(os.Args) < 2 {
     	fmt.Println("Need at least 1 file")
	return
     }

     i := 1
     for i < len(os.Args) {
     	 file,err := os.Open(os.Args[i])
     	 if err != nil {
     	    return
     	 }
     	 defer file.Close()

     	 stat,err := file.Stat()
     	 if err != nil {
     	    return
	 }

     	 bs := make([]byte, stat.Size())
     	 _, err = file.Read(bs)
     	 if err != nil {
     	    return
         }

     	 h := sha1.New()
     	 h.Write(bs)
     	 ck := h.Sum(nil)

         fmt.Printf("%x  %s\n", ck, os.Args[i])
     	 i = i + 1
     }
}
