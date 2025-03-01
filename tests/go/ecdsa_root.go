package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	// Generate ECDSA key pair using P-384 curve
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Println("Failed to generate ECDSA key:", err)
		return
	}

	// Define the certificate configurations
	certConfigs := []struct {
		CN       string
		FileName string
		sigAlg   x509.SignatureAlgorithm
	}{
		{"ECDSAWithSHA3_256 Root", "ECDSAWithSHA3_256.pem", x509.ECDSAWithSHA3_256},
		{"ECDSAWithSHA3_384 Root", "ECDSAWithSHA3_384.pem", x509.ECDSAWithSHA3_384},
		{"ECDSAWithSHA3_512 Root", "ECDSAWithSHA3_512.pem", x509.ECDSAWithSHA3_512},
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
func generateAndSaveCert(cn, fileName string, sigAlg x509.SignatureAlgorithm, privateKey *ecdsa.PrivateKey) error {
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
