// the package main in Go defines an executable program
package main

// net/http package is used to create a web server, it is a standard library in Go
// gorilla/mux package is used to create a router, it is a third-party package
import (
	"log"
	"myapp/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/expenses", handlers.CreateExpense).Methods("POST")
	// r.HandleFunc("/expenses/{orderNumber}", handlers.GetExpense).Methods("GET")
	r.HandleFunc("/expenses", handlers.GetAllExpenses).Methods("GET")
	r.HandleFunc("/expenses/{orderNumber}", handlers.DeleteExpense).Methods("DELETE")
	r.HandleFunc("/expenses/search", handlers.SearchExpenses).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}).Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
