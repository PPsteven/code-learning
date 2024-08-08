package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "hello, world")
}

func main() {
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// // Create the TLS Config with the CA pool and enable Client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs: caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	// BuildNameToCertificate（建立名称到证书的映射）是一种SSL/TLS优化技术，
	// 用于加速服务器端SSL/TLS握手过程。当客户端与服务器进行SSL/TLS握手时，
	// 服务器需要将自己的证书链发送给客户端进行验证。而在证书链中，通常存在多个证书，每个证书都包含一个公钥和一个主题名称。
	// 为了验证证书链的有效性，客户端需要检查证书中的主题名称是否与服务器的域名匹配。这个过程对于大型网站来说可能会比较耗时。
	// 为了优化这个过程，服务器可以使用BuildNameToCertificate技术，将主题名称和证书之间建立一个映射关系。
	// 这样，在客户端请求到达服务器后，服务器就可以快速地找到相应的证书，并将其发送给客户端进行验证，从而缩短SSL/TLS握手时间，提高HTTPS连接的响应速度。

	// BuildNameToCertificate技术自TLSv1.3开始被废弃。这是因为TLSv1.3在协议设计上进行了改进，使用了更加高效和安全的证书验证机制。
	// 在TLSv1.3中，服务器可以发送无序的证书链给客户端，而不需要显式地建立主题名称到证书的映射关系，从而降低了握手过程中的延迟。
	// 另外，在实际应用中，由于一些TLS实现对BuildNameToCertificate技术支持不完整或存在漏洞，可能会导致安全性问题。
	// 因此，TLSv1.3标准化组织决定废弃BuildNameToCertificate技术，并提倡使用更加安全和高效的证书验证机制。
	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr: ":8080",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}


// Generate and use the Certificates
// openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -out cert.pem -keyout key.pem -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost"
// openssl 详情
// 1. req command: The req command primarily creates and processes certificate requests in PKCS#10 format.
//    It can additionally create self-signed certificates, for use as root CAs, for example.
// 2. "-newkey rsa:2048" 指定了要使用 RSA 算法生成一个 2048 位的密钥对。
// 3. "-nodes" 选项用于生成私钥时不加密私钥。如果使用了该选项，则生成的私钥文件将不需要密码才能打开或使用。
// 4. "-keyout" 选项来指定私钥文件的输出位置，"-out" 选项用于指定 CSR 文件的输出位置
// warning:
// go version >= 1.15 CN 被废弃了，SAN 应该被添加至CA证书
// openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -out cert.pem -keyout key.pem -subj "/C=US/ST=California/L=Mountain View/O=Your Organization/OU=Your Unit/CN=localhost" -addext "subjectAltName=DNS:localhost"
