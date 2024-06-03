// the package main in Go defines an executable program
package main

// net/http package is used to create a web server, it is a standard library in Go
// gorilla/mux package is used to create a router, it is a third-party package
import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "myapp/api/handlers"
    "github.com/rs/cors"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/expenses", handlers.CreateExpense).Methods("POST")
    r.HandleFunc("/expenses/{orderNumber}", handlers.GetExpense).Methods("GET")
    r.HandleFunc("/expenses", handlers.GetAllExpenses).Methods("GET")
    r.HandleFunc("/expenses/{orderNumber}", handlers.DeleteExpense).Methods("DELETE")

    handler := cors.Default().Handler(r)

    log.Fatal(http.ListenAndServe(":8080", handler))
}