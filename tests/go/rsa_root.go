package main

/**
Setup:
$ git clone https://github.com/golang/go
$ cp p3.patch go/
$ cp rsa_root_proof.go go/ ( This file )
$ cd go
$ git apply p3.patch
$ cd src && ./all.bash
$ cd .. && export GOROOT=$(pwd)
$ export PATH="$GOROOT/bin:$PATH"
$ go version 
go version devel go1.25-d31c805535 Thu Feb 27 09:23:42 2025 -0800 linux/amd64
$ go run rsa_root_proof.go
Successfully created SHA3256WithRSA Root -> bin/SHA3_256WithRSA.pem
Successfully created SHA3384WithRSA Root -> bin/SHA3_384WithRSA.pem
Successfully created SHA3512WithRSA Root -> bin/SHA3_512WithRSA.pem
**/

import (
	// "crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	// Generate RSA 4096-bit key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println("Failed to generate RSA key:", err)
		return
	}

	// Define the certificate configurations
	certConfigs := []struct {
		CN       string
		FileName string
		sigAlg   x509.SignatureAlgorithm
	}{
		{"SHA3256WithRSA Root", "SHA3_256WithRSA.pem", x509.SHA3_256WithRSA},
		{"SHA3384WithRSA Root", "SHA3_384WithRSA.pem", x509.SHA3_384WithRSA},
		{"SHA3512WithRSA Root", "SHA3_512WithRSA.pem", x509.SHA3_512WithRSA},
	}

	// Generate and save each certificate
	for _, cfg := range certConfigs {
		err := generateAndSaveCert(cfg.CN, cfg.FileName, cfg.sigAlg, privateKey)
		if err != nil {
			fmt.Printf("Failed to generate %s: %v\n", cfg.CN, err)
		} else {
			fmt.Printf("Successfully created %s -> %s\n", cfg.CN, cfg.FileName)
		}
	}
}

// generateAndSaveCert creates a self-signed root certificate and saves it to a PEM file
func generateAndSaveCert(cn, fileName string, sigAlg x509.SignatureAlgorithm, privateKey *rsa.PrivateKey) error {
	// Define certificate template
	template := &x509.Certificate{
		SerialNumber:          bigIntSerial(),
		Subject:               pkix.Name{CommonName: cn},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(10 * 365 * 24 * time.Hour), // 10 years validity
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		BasicConstraintsValid: true,
		IsCA:                  true,
		SignatureAlgorithm:    sigAlg, // 👈 EXPLICITLY SET SIGNATURE ALGORITHM
	}

	// Create self-signed certificate
	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return fmt.Errorf("failed to create certificate: %w", err)
	}

	// Save the certificate as a PEM file
	return savePEMFile(fileName, "CERTIFICATE", certBytes)
}

// savePEMFile writes data to a PEM-encoded file
func savePEMFile(fileName, pemType string, data []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fileName, err)
	}
	defer file.Close()

	return pem.Encode(file, &pem.Block{Type: pemType, Bytes: data})
}

// bigIntSerial generates a random serial number
func bigIntSerial() *big.Int {
	serial, _ := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	return serial
}