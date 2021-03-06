package main
import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	http.HandleFunc("/configmap",ConfigMap)
	http.HandleFunc("/secret",Secret)
	http.HandleFunc("/",Hello)
	http.ListenAndServe(":8000",nil)
}

func Secret(w http.ResponseWriter, r *http.Request){
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w,"User: %s. Password: %s",user,password)
}

func Hello(w http.ResponseWriter, r *http.Request){
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w,"Hello I'm %s. I'm %s.",name,age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request){
	
	data, err := ioutil.ReadFile("myfamily/family.txt")

	if err != nil {
		log.Fatalf("Error reading file: ",err)
	}
	fmt.Fprintf(w,"My family: %s.",string(data))
}