package Routes

import (
	"fmt"
	"log"
	"main/Controller"
	"main/Utilities"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	r *mux.Router
)

func CreateRouter() {
	r = mux.NewRouter()

}

func InitializeRouters() {
	//Get Methods
	r.HandleFunc("/", Controller.Home).Methods("GET")
	r.HandleFunc("/cars", Controller.Cars).Methods("GET")
	r.HandleFunc("/contact", Controller.Contact).Methods("GET")
	r.HandleFunc("/about", Controller.About).Methods("GET")
	r.HandleFunc("/signup", Controller.Signup).Methods("GET")
	r.HandleFunc("/signin", Controller.Signin).Methods("GET")
	r.HandleFunc("/reset", Controller.Reset).Methods("GET")
	r.HandleFunc("/admin", Controller.AdminLogin).Methods("GET")
	r.Handle("/admindashboard", Utilities.IsAuthorized(Controller.AdminDashboard)).Methods("GET")
	r.Handle("/salesdashboard", Utilities.IsAuthorized(Controller.SalesDashboard)).Methods("GET")

	//POST Methods
	r.HandleFunc("/admin", Controller.AdminLogin).Methods("POST")
}

func StartServer() {
	// Choose the folder to serve
	staticDir := "/static/"
	// Create the route
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	fmt.Println("Server starting at : http://localhost:8045")
	if err := http.ListenAndServe(":8045", r); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
