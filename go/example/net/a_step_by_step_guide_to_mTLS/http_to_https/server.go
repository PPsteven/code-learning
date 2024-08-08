package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "hello, world")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil))
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
