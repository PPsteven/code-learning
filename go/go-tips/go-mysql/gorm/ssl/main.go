package main

// DSN 是什么? 和 Connection String 区别是什么
// An ODBC Data Source Name (DSN) stores information for establishing a connection to a database on a remote database server.
// A system DSN provides access to multiple users, rather than only the user who created it.

// DSN 和 Connection String 经常可以混用
// Gorm 的DSN 参考的是 [](https://github.com/go-sql-driver/mysql#dsn-data-source-name)
// scheme://username:password@host:port/dbname?param1=value1&param2=value2&...

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/sirupsen/logrus"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"time"

	"github.com/go-sql-driver/mysql"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func NewWithDSN() *gorm.DB {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{})
	return db
}

func NewDSN(isTLS bool) (dsn string) {
	tslconfig := "skip-verify"
	if isTLS {
		rootCertPool := x509.NewCertPool()
		pemFile := "/path/to/pem/file"
		pem, err := ioutil.ReadFile(pemFile)
		if err != nil {
			logrus.Fatalf("%v", err)
		}
		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			logrus.Fatalf("Failed to append PEM, %v", pemFile)
		}
		tslconfig = "custom"
		err = mysql.RegisterTLSConfig("custom", &tls.Config{
			RootCAs: rootCertPool,
		})
		if err != nil {
			logrus.Fatalf("%v", err)
		}
	}

	cfg := mysql.Config{
		User:      "user",
		Passwd:    "passwd",
		Addr:      "ip:port",
		Net:       "tcp",
		DBName:    "evaprod",
		Loc:       func() *time.Location { loc, _ := time.LoadLocation("Asia/Shanghai"); return loc }(),
		TLSConfig: tslconfig,
		//Params: o,
		AllowNativePasswords: true,
	}

	dsn = cfg.FormatDSN()
	logrus.Debugf("DSN: %v", cfg.FormatDSN())
	return dsn
}

func NewDB() (*gorm.DB, error) {
	dsn := NewDSN(true)
	conn, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	db, err := conn.DB()
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(86400)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func main() {
	db, err := NewDB()
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	logrus.Infof("%v", db != nil)
}
