package main

import (
	"fmt"
	"net/http"
)

// func main() {
// 	pool := x509.NewCertPool()
// 	caCertPath := "ca.crt"

// 	caCrt, err := ioutil.ReadFile(caCertPath)
// 	if err != nil {
// 		fmt.Println("ReadFile err:", err)
// 		return
// 	}
// 	pool.AppendCertsFromPEM(caCrt)

// 	http.HandleFunc("/", handler)

// 	s := &http.Server{
// 		Addr: ":4001",
// 		TLSConfig: &tls.Config{
// 			ClientCAs:  pool,
// 			ClientAuth: tls.RequireAndVerifyClientCert,
// 		},
// 	}
// 	err = s.ListenAndServeTLS("server.crt", "server.key")
// 	if err != nil {
// 		fmt.Println("ListenAndServeTLS err:", err)
// 	}
// }

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HTTPS Server.")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":4001",
		"server.crt",
		"server.key", nil)
}
