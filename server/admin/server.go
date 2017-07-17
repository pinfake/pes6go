package admin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pinfake/pes6go/storage"
)

type Server struct {
	storage.Storage
}

func (_ Server) account(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		key := req.FormValue("key")
		password := req.FormValue("password")
		fmt.Fprintf(w, "%s %s\n", key, password)
	}
}

func Start() {
	server := Server{storage.Forged{}}
	fmt.Println("Administration Server starting")
	mux := http.NewServeMux()
	mux.Handle("/account", http.HandlerFunc(server.account))
	log.Fatal(http.ListenAndServe("0.0.0.0:19770", mux))
}
