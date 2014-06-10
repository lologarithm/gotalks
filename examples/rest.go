package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BEGIN EXAMPLE OMIT
type Thing struct {
	SomeValue string
}

func handleThing(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t := &Thing{SomeValue: "The thing value you care about."}
		jsonString, _ := json.Marshal(t)
		fmt.Fprintf(w, string(jsonString))
	case "DELETE":
		fmt.Fprintf(w, "YOUR THING WAS DELETED.")
	}
}

func main() {
	http.HandleFunc("/thing", handleThing)
	http.ListenAndServe(":24816", nil)
}

// END EXAMPLE OMIT
