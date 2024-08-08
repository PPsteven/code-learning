package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// cert 是客户端用以请求服务器，供给服务端验证的证书
	// caCert 是用以验证服务端的证书
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// create a ca certificate pool
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
			Certificates: []tls.Certificate{cert},
		},
	}}

	r, err := client.Get("https://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("body:%v\n", string(body))
}