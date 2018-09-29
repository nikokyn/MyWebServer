package main

import (	
	"fmt"
	"net/http"
	"html/template"
	"bufio"
	"os"
	"log"
)

type Data struct{
	Book1 []string
	Book2 []string
}

func readTXT(path string) ([]string, error){
	file, err :=os.Open(path)
	if err != nil { return nil, err}
	defer file.Close()
	var lines []string
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		lines=append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	 t, err :=template.ParseFiles("Html/Columns.html")
	 if err!=nil{
	 	fmt.Fprintln(w, err.Error())

	 }

	BookOne, err:= readTXT("Books/left.txt")
	if err != nil {
		log.Fatalf("readTXT: %s", err)
	}

	BookTwo, err2:= readTXT("Books/right.txt")
	if err2 != nil {
		log.Fatalf("readTXT: %s", err2)
	}
		data:=&Data{BookOne,BookTwo}
		t.ExecuteTemplate(w, "Columns", data)
}

func main() {
	fmt.Println("Port 1111")
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":1111",nil)
}