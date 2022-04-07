package database

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

func setup_cassandra() *gocql.Session {
	env := godotenv.Load(".env")
	if env != nil {
		log.Fatal("Failed to load environment variables")
	}
	config := gocql.NewCluster(os.Getenv("cassandra_uri"))
	config.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("cassandra_username"),
		Password: os.Getenv("cassandra_password"),
	}
	config.Hosts = []string{os.Getenv("cassandra_uri") + ":" + os.Getenv("cassandra_port")}
	certificate_path, _ := filepath.Abs("././static/bundle/cert")
	key_path, _ := filepath.Abs("././static/bundle/key")
	ca_path, _ := filepath.Abs("././static/bundle.ca.crt")

	cert, _ := tls.LoadX509KeyPair(certificate_path, key_path)
	ca_cert, _ := ioutil.ReadFile(ca_path)
	ca_cert_pool := x509.NewCertPool()
	ca_cert_pool.AppendCertsFromPEM(ca_cert)
	tls_config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      ca_cert_pool,
	}
	config.SslOpts = &gocql.SslOptions{
		Config:                 tls_config,
		EnableHostVerification: false,
	}

	session, _ := config.CreateSession()

	return session
}
