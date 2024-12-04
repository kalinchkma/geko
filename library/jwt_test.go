package library

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateECDSAKeys(t *testing.T) {
	_, _, err := GenerateECDSAKeys()
	if err != nil {
		fmt.Println("Error: ", err)
		t.Failed()
	}
}

func TestGenerateJWT(t *testing.T) {
	priv, _, err := GenerateECDSAKeys()
	if err != nil {
		fmt.Println("Error", err)
		t.Fatalf("Faild ECDSA Generation")
	}

	_, err = GenerateJWT("1", time.Duration(time.Hour*24), priv)
	if err != nil {
		fmt.Println("Error:", err)
		t.Fatalf("Faild geneating JWT")
	}
}
