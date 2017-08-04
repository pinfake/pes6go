package server

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/blowfish"

	"bytes"

	"crypto/md5"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/pinfake/pes6go/storage"
)

type AdminServer struct {
	storage storage.Storage
}

func (s AdminServer) account(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Hey im receiving a request\n")
	switch req.Method {
	case "POST":
		key := req.FormValue("key")
		password := req.FormValue("password")

		fmt.Fprintf(w, "%s %s\n", key, password)

		var keypadded [36]byte
		copy(keypadded[:], []byte(key))
		var buf bytes.Buffer
		buf.Write(keypadded[:])
		buf.Write([]byte(password))
		var data = buf.Bytes()
		fmt.Fprintf(w, "% x\n", data)
		md5sum := md5.Sum(data)
		block, _ := blowfish.NewCipher(BlowfishKey)
		encrypter := ecb.NewECBEncrypter(block)
		dst := make([]byte, len(md5sum))
		encrypter.CryptBlocks(dst, md5sum[:])
		fmt.Fprintf(w, "% x\n", dst)

		id, err := s.storage.CreateAccount(&storage.Account{
			Key:  key,
			Hash: dst,
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
