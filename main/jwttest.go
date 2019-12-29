package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type MapInfo map[string]string

func (m MapInfo) String() string {
	out, _ := json.Marshal(m)
	return string(out)
}

func (m MapInfo) Set(arg string) error {
	s := strings.SplitN(arg, "=", 2)
	if len(s) != 2 {
		return fmt.Errorf("error:len(arg)=%v, %v", len(s), "len(arg)!=2")
	}
	m[s[0]] = s[1]
	return nil
}

func main() {
	mapInfo := make(MapInfo)
	flag.Var(mapInfo, "map", "map info ")
	alg := flag.String("alg", "", "(must) algoritnm(ec,rsa,hash,RS256,RS384,RS512..)")
	ty := flag.String("type", "", "(must) action(sign,verify,check)")
	flag.Parse()
	if *alg == "" || *ty == "" {
		flag.Usage()
		return
	}
	switch *alg {
	case "RS256":
		rsaPri, err := rsa.GenerateKey(rand.Reader, 2048) //2048 bits
		if err != nil {
			fmt.Printf("rsa genkey error(%v)\n", err)
			return
		}
		block1 := pem.Block{
			Type:    "RSA PRIVATE KEY",
			Headers: map[string]string{"test": "rsa"},
			Bytes:   x509.MarshalPKCS1PrivateKey(rsaPri),
		}
		rsaPublic := rsaPri.Public().(*rsa.PublicKey)
		block2 := pem.Block{
			Type:    "RAS PUBLIC KEY",
			Headers: map[string]string{"test": "rsa"},
			Bytes:   x509.MarshalPKCS1PublicKey(rsaPublic),
		}
		privPem := pem.EncodeToMemory(&block1)   //pem file
		publicPem := pem.EncodeToMemory(&block2) //pem file
		/////decode
		block3, _ := pem.Decode(privPem)
		pri, err := x509.ParsePKCS1PrivateKey(block3.Bytes)
		if err != nil {
			fmt.Printf("ParsePKCS1PrivateKey error(%v)\n")
			return
		}
		block4, _ := pem.Decode(publicPem)
		pub, err := x509.ParsePKCS1PublicKey(block4.Bytes)
		if err != nil {
			fmt.Printf("ParsePKCS1PublicKey error(%v)\n")
			return
		}
		method := jwt.GetSigningMethod(*alg)
		if method == nil {
			fmt.Printf("get method error\n")
			return
		}
		switch *ty {
		case "sign":

		case "verify":
		case "check":
			claim := jwt.MapClaims{}
			if len(mapInfo) > 0 {
				for k, v := range mapInfo {
					claim[k] = v
				}
			}
			token := jwt.NewWithClaims(method, claim)
			tokenSting, err := token.SignedString(pri)
			if err != nil {
				fmt.Printf("oken signed string error(%v)\n")
				return
			}

			//verify,
			// tokData = regexp.MustCompile(`\s*$`).ReplaceAll(tokData, []byte{})
			tokenParse, err := jwt.Parse(tokenSting, func(token *jwt.Token) (interface{}, error) {
				return pub, nil
			})
			if err != nil {
				fmt.Printf("jwt.parse error(%v)\n")
				return
			}
			out, err := json.Marshal(tokenParse)
			fmt.Printf("verify: %v\n", string(out))
			//claim:=
		}
	case "ec":

	}

	fmt.Printf("++++++++++,%v\n", mapInfo)
	out, err := loadData("go.mod")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	fmt.Printf("out len:%v\n", len(out))
}
func loadData(p string) ([]byte, error) {
	if p == "" {
		return nil, fmt.Errorf("No path specified")
	}

	var rdr io.Reader
	if p == "-" {
		rdr = os.Stdin
	} else if p == "+" {
		return []byte("{}"), nil
	} else {
		if f, err := os.Open(p); err == nil {
			rdr = f
			defer f.Close()
		} else {
			return nil, err
		}
	}
	return ioutil.ReadAll(rdr)
}
