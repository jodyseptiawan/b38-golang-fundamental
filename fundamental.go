// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type Todos struct {
// 	Id string `json:"id"`
// 	Title string `json:"title"`
// 	IsDone bool `json:"isDone"`
// }

// func main() {

// 	r := mux.NewRouter()

// 	r.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
// 		w.Header().Set("Content-Type", "application/json")

// 		var dataTodos []Todos

// 		dataTodos[0].Id = "27"
// 		dataTodos[0].Title = "Sarapan"
// 		dataTodos[0].IsDone = false

// 		json.NewEncoder(w).Encode(dataTodos)
// 	})

// 	fmt.Println("server running localhost:5000")
// 	http.ListenAndServe("localhost:5000", r)
// }