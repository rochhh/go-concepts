package main

import (
	"fmt"
	"io"
	"mime/multipart"

	// "html/template"
	"log"
	"net/http"
	"os"
)

type Data struct {
	FileBody *multipart.FileHeader
}

func main(){
	//ReadFile()
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
	fmt.Printf("the file is %v and the handler is %v \n" , file , handler)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	f , err := os.OpenFile("./uploads/"+handler.Filename,os.O_WRONLY |os.O_CREATE , 0666)
	fmt.Printf("the f value is %v  \n" , f )
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	n , err := io.Copy( f , file )
	fmt.Printf("the n value is %v  \n" , n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w , "File uploaded successfully ! %s\n" , handler.Filename)
}

func ReadFile(){
	f , err := os.Open("dictionary/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte , 2048)
	n , err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	fmt.Println(string(buf))
}

