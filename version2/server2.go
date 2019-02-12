// reference link: https://github.com/nlopes/slack/blob/master/examples/slash/slash.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nicewook/slack_slash_cmd/version2/slash"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome, I'm Timebot\n\nI can convert KST <-> PST/PDT")
}

func main() {

	r := httprouter.New()
	r.GET("/", index)
	r.POST("/slash", slash.Handler)

	fmt.Println("[INFO] Server listening")
	log.Fatal(http.ListenAndServe("", r))

}
