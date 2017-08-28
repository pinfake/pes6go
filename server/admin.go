package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pinfake/pes6go/data/block"
	"github.com/pinfake/pes6go/storage"
)

type AdminServer struct {
	storage storage.Storage
}

func (s AdminServer) account(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		key := req.FormValue("key")
		password := req.FormValue("password")

		auth := block.Authentication{
			Key:      key,
			Password: password,
		}

		passwordHash := auth.GetPasswordHash()

		fmt.Fprintf(w, "%s %s\n", key, password)
		fmt.Fprintf(w, "% x\n", passwordHash)

		id, err := s.storage.CreateAccount(&storage.Account{
			Key:  key,
			Hash: passwordHash,
		})
		if err != nil {
			fmt.Fprintf(w, "Cannot store account: %s", err.Error())
			return
		}
		fmt.Fprintf(w, "id: %d\n", id)
	}
}

func StartAdmin(stor storage.Storage) {
	s := AdminServer{stor}
	fmt.Println("Administration Server starting")
	mux := http.NewServeMux()
	mux.Handle("/account", http.HandlerFunc(s.account))
	log.Fatal(http.ListenAndServe("0.0.0.0:19770", mux))
}
