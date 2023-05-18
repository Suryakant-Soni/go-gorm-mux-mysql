package main

import (
	"go-gorm-mux-mysql/pkg/routes"
	"log"
	"net/http"
	"os"
	"runtime/trace"

	"github.com/gorilla/mux"
)

func main() {
	f, err1 := os.Create("trace.out")
	if err1 != nil {
		log.Fatal("Error", err1)
	}
	defer f.Close()
	err1 = trace.Start(f)
	if err1 != nil {
		log.Fatal("Error", err1)
	}
	defer trace.Stop()

	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	err := http.ListenAndServe("localhost:9000", r)
	if err != nil {
		log.Println("Error", err)
	}
}
