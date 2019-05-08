package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

type PrivateKey struct {
	priv string
	PublicKey
}
type PublicKey struct {
	curve string
	pk    string
}

const (
	P256 = "P256"
	P224 = "P224"
	P384 = "P384"
	P521 = "P521"
)

var CurveType = "P256"

//type curveTab struct {
//	tab map[int]
//}
func initCurve(ct string) {
	CurveType = ct
}

func GetCurve(curveType string) elliptic.Curve {
	switch curveType {
	case P256:
		return elliptic.P256()
	case P224:
		return elliptic.P224()
	case P384:
		return elliptic.P384()
	case P521:
		return elliptic.P521()
	default:
		return elliptic.P256()
	}
}

var ErrPublicKeyStringFormat = errors.New("publickey string format error")
var ErrPublicKeyNull = errors.New("publickey is null error")
var ErrPublicKeyFormat = errors.New("publickey format error")

//pkstring==curve+xStr+yStr
func (pk *PublicKey) SetString(pkStr string) error {
	length := len(pkStr)
	if length <= 6 {
		return ErrPublicKeyStringFormat
	}
	size, err := strconv.Atoi(pkStr[4:6])
	if err != nil {
		return err
	}
	if size >= length-6 {
		return ErrPublicKeyStringFormat
	}
	pk.curve = pkStr[:4]
	pk.pk = pkStr[4:]
	return nil
}
func (pk *PublicKey) GenerateKey(curve string) (prvk *PrivateKey, err error) {
	c := GetCurve(curve)
	prvkesdsa, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		return nil, err
	}
	switch curve {
	case P256, P224, P384, P521:
		prvk.PublicKey.curve = curve
	default:
		//todo
		prvk.PublicKey.curve = P256
	}
	prvk.SetInt(prvkesdsa.D)
	prvk.PublicKey.SetInt(prvkesdsa.PublicKey.X, prvkesdsa.PublicKey.Y)
	return prvk, nil
}
func (pk *PublicKey) GetString() (string, error) {
	if pk.curve == "" || pk.pk == "" || len(pk.pk) <= 2 {
		return "", ErrPublicKeyNull
	}
	//leng:=len(pk.pk)
	leng, err := strconv.Atoi(pk.pk[:2])
	if err != nil {
		return "", err
	}
	if leng >= len(pk.pk)-2 {
		return "", ErrPublicKeyFormat
	}
	switch pk.curve {
	case "P256":
		return pk.curve + pk.pk, nil
	case "P224":
		return pk.curve + pk.pk, nil
	case "P384":
		return pk.curve + pk.pk, nil
	case "P521":
		return pk.curve + pk.pk, nil
	default:
		return "", ErrPublicKeyFormat
	}
}
func (pk *PublicKey) SetInt(x *big.Int, y *big.Int) {
	xBase64 := base64.StdEncoding.EncodeToString(x.Bytes())
	yBase64 := base64.StdEncoding.EncodeToString(y.Bytes())
}
func (priv *PrivateKey) SetBase64(base64 string) {
	priv.priv = base64
	//priv.curve = CurveType
	//priv.pk=
}
func (priv *PrivateKey) GetBase64() (base64 string) {
	return priv.priv
}
func (priv *PrivateKey) SetInt(i *big.Int) {
	//intByte := i.Bytes()
	priv.SetBase64(base64.StdEncoding.EncodeToString(i.Bytes()))
}
func (priv *PrivateKey) GetInt() (*big.Int, error) {
	intByte, err := base64.StdEncoding.DecodeString(priv.GetBase64())
	if err != nil {
		return nil, err
	}
	bigI := new(big.Int)
	bigI.SetBytes(intByte)
	return bigI, nil
}
func (priv *PrivateKey) SetHexString(hexStr string) {
	var bigI big.Int
	bigI.SetString(hexStr, 16)
	priv.SetInt(&bigI)
}
func (priv *PrivateKey) GetHexString() (hexStr string, err error) {
	bigByte, err := base64.StdEncoding.DecodeString(priv.GetBase64())
	if err != nil {
		return "", err
	}
	var bigI big.Int
	bigI.SetBytes(bigByte)
	hexStr = bigI.Text(16)
	return hexStr, nil
}
func (priv *PrivateKey) SetEcdsaPrivKey(prv ecdsa.PrivateKey) {
	priv.SetInt(prv.D)
}
func (priv *PrivateKey) GetEcdsaPrivKey() (*ecdsa.PrivateKey, error) {
	bigI, err := priv.GetInt()
	if err != nil {
		return nil, err
	}
	prv := new(ecdsa.PrivateKey)
	prv.D = bigI
	return prv, nil
}
func main() {
	ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	prv := ecdsa.PrivateKey{}
	//pk := ecdsa.PublicKey{}
	//prv.Sign()
	elliptic.P224()
	//var bigI big.Int
	var p PrivateKey
	p.SetHexString("1111")
	bigI, _ := p.GetInt()
	fmt.Println(bigI, bigI.Bytes())
	//bigI.SetString("1FF", 16)
	//fmt.Println("bigInt:", bigI)
	//fmt.Println("bigI.Bytes():", bigI.Bytes())
	//fmt.Println("bigI.String()", len(bigI.Text(16)))
	//bigI.SetBytes([]byte{1, 255})
	//fmt.Println("bigInt:", bigI)
	//
	//fmt.Print("nihao,", p)
}
