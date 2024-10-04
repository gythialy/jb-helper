package power

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/gythialy/jb-helper/internal/cert"
	"github.com/gythialy/jb-helper/internal/utils"
)

func GenerateEqualResult() error {
	crt, err := cert.ParseCertPem()
	if err != nil {
		return err
	}
	x := new(big.Int).SetBytes(crt.Signature)
	y := 65537
	block, _ := pem.Decode([]byte(RootCert))
	rootCertificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}
	p, _ := rootCertificate.PublicKey.(*rsa.PublicKey)
	z := p.N
	zp, _ := crt.PublicKey.(*rsa.PublicKey)
	r := new(big.Int)
	r.Exp(x, big.NewInt(int64(y)), zp.N)
	equal := fmt.Sprintf("EQUAL,%d,%d,%d->%d", x, y, z, r)
	powerConf := filepath.Join(utils.GetStaticPath(), "conf", "power.conf")
	writerStr := `[Result]
;active code 
` + equal
	return os.WriteFile(powerConf, []byte(writerStr), 0o644)
}
