package main

import (
	"fmt"

	"io"
	"log"
	"net/http"
	"os"
)

// type Data struct {
// 	Filename string
// 	Size     int64
// 	content   []byte
// }

func main(){
	//ReadFile()
	ProcessData("/data" , ProcessData )
	http.HandleFunc("/" , MainPage )
	http.HandleFunc("/upload" ,getDataFromFrontEnd )
	err := http.ListenAndServe(":8080" , nil )
	if err != nil {
		log.Fatal(err)
	}
}

func MainPage( w http.ResponseWriter , r *http.Request ){
	http.ServeFile(w,r, "templates/index.html")
}

func getDataFromFrontEnd(w http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(10 << 20)
	file , handler , err := r.FormFile("fileInput")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	f , err := os.OpenFile("./uploads/"+handler.Filename,os.O_WRONLY |os.O_CREATE , 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_ , err = io.Copy( f , file )
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w , "File uploaded successfully ! %s\n" , handler.Filename)
}

func ProcessData( w http.ResponseWriter , r *http.Request ){
	file, err := os.Open("dictionary/words.txt")
	
	fi , handler , err := r.FormFile("fileInput")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()
	ufi := handler.Filename 
	u , err := os.Open("uploads/"+ufi)
	fmt.Println(u)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := make([]byte , 2048 )
	data , err := file.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	slurs := string(buf[:data])
	fmt.Println(slurs)
}
