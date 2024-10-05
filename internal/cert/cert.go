package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/gythialy/jb-helper/internal/utils"
)

const (
	// OneDay represents a time duration of 1 day.
	OneDay = time.Hour * 24
	// TenYears represents a time duration of 10 years.
	TenYears     = OneDay * 3650
	certPem      = "jb-active.pem"
	certKey      = "jb-active.key"
	certPKCS8Key = "jb-active-pkcs8.key"
)

// CreateCert 创建证书
func CreateCert() error {
	today := time.Now()
	yesterday := today.Add(-OneDay)
	tomorrow := today.Add(TenYears)
	// 生成RSA私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}
	serialNumber, err := randomSerialNumber()
	if err != nil {
		return err
	}
	subject := pkix.Name{
		CommonName: "JetProfile CA",
	}
	template := &x509.Certificate{
		SerialNumber:       serialNumber,
		Subject:            subject,
		Issuer:             subject,
		NotBefore:          yesterday,
		NotAfter:           tomorrow,
		PublicKeyAlgorithm: x509.RSA,
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKey:          &privateKey.PublicKey,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return err
	}
	// 保存 CA 证书
	privateKeyPEM := encodePrivateKeyToPEM(privateKey)
	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	})
	privateKeyPKCS8PEM := encodePrivateKeyToPKCS(privateKey)

	if err := writeToFile(filepath.Join(utils.GetCurrentPath(), "cert", certPem), certPEM); err != nil {
		return err
	}
	if err := writeToFile(filepath.Join(utils.GetCurrentPath(), "cert", certKey), privateKeyPEM); err != nil {
		return err
	}
	if err := writeToFile(filepath.Join(utils.GetCurrentPath(), "cert", certPKCS8Key), privateKeyPKCS8PEM); err != nil {
		return err
	}
	return nil
}

func encodePrivateKeyToPEM(key *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
}

func encodePrivateKeyToPKCS(key *rsa.PrivateKey) []byte {
	if privateKey, err := x509.MarshalPKCS8PrivateKey(key); err == nil {
		return pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKey,
		})
	}
	return nil
}

func writeToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0o644)
}

func randomSerialNumber() (*big.Int, error) {
	maxSerialNumber := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1))
	serialNumber, err := rand.Int(rand.Reader, maxSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random serial number: %w", err)
	}
	return serialNumber, nil
}

// ParseCertPem 解析PEM
func ParseCertPem() (*x509.Certificate, error) {
	data, err := os.ReadFile(filepath.Join(utils.CertPath(), certPem))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the certificate")
	}
	crt, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return crt, nil
}

// ParseCertKey 解析key
func ParseCertKey() (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(filepath.Join(utils.CertPath(), certKey))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}
