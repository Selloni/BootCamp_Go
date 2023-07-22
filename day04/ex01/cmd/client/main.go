package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"github.com/lizrice/secure-connections/utils"
	"io"
	"log"
	"net/http"
	"os"
)

type Data struct {
	name  string
	count int
	money int
}

func main() {
	k := flag.String("k", "", "flag for name candy")
	c := flag.Int("c", 0, "count of candy")
	m := flag.Int("m", 0, "price of candy")
	flag.Parse()

	if !isFlagPassed("k") || !isFlagPassed("c") || !isFlagPassed("m") {
		log.Fatalln("Wrong arguments")
	}

	var dataFlag Data = Data{*k, *c, *m}
	var buff bytes.Buffer

	err := json.NewEncoder(&buff).Encode(dataFlag)
	if err != nil {
		log.Fatalf("could not JsonMarshal %s", err)
	}

	client := getClient()
	res, err := client.Post("https://localhost:3333/buy_candy", "application/json", &buff)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func getClient() *http.Client {
	data, err := os.ReadFile("../../CA/minica.pem")
	if err != nil {
		log.Fatal(err)
	}
	cp, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	cp.AppendCertsFromPEM(data)

	config := &tls.Config{
		InsecureSkipVerify:    true,
		ClientAuth:            tls.RequireAnyClientCert,
		RootCAs:               cp,
		GetCertificate:        utils.CertReqFunc("../../CA/server/cert.pem", "../../CA/server/key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

func isFlagPassed(kcm string) bool {
	flagHave := false
	flag.Visit(
		func(f *flag.Flag) {
			if f.Name == kcm {
				flagHave = true
			}
		})
	return flagHave
}
