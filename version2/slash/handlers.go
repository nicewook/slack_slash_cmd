package slash

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handler deals with slash commands
// Check HTTP Request and ParseForm()
func Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// raw HTTP Request
	fmt.Printf("r\n\n%+v\n\n", r)

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ParseForm() HTTP Request
	fmt.Printf("r ParseForm()\n\n%+v\n\n", r)

	for k, v := range r.PostForm {
		fmt.Printf("%v: %v\n", k, v)
	}

	response := fmt.Sprintf("Request accepted")
	w.Write([]byte(response))
}

// Handler deals with slash commands
// func Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	cmd := r.PostFormValue("command")

// 	// TODO: validateToken

// 	switch cmd {
// 	case "/time":
// 		response := fmt.Sprintf("You requested for KST <-> PST/PDT for %+v", r)
// 		w.Write([]byte(response))

// 	default:
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }
