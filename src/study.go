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

func getCurveStr(curve elliptic.Curve) (string, error) {
	switch curve {
	case elliptic.P256():
		return P256, nil
	case elliptic.P224():
		return P224, nil
	case elliptic.P384():
		return P384, nil
	case elliptic.P521():
		return P521, nil
	}
	return "", errors.New("No Type Curve")
}

var ErrPublicKeyStringFormat = errors.New("publickey string format error")
var ErrPublicKeyNull = errors.New("publickey is null error")
var ErrPublicKeyFormat = errors.New("publickey format error")
var ErrPrivKeyFormat = errors.New("private key format error")

//pkstring==curve+xStr+yStr
func (pk *PublicKey) SetString(pkStr string) error {
	length := len(pkStr)
	if length <= 6 {
		return ErrPublicKeyStringFormat
	}
	size, err := strconv.ParseUint(pkStr[4:6], 16, 0)
	if err != nil {
		return err
	}
	if int(size) >= length-6 {
		return ErrPublicKeyStringFormat
	}
	pk.curve = pkStr[:4]
	pk.pk = pkStr[4:]
	return nil
}
func GenerateKey(curve string) (prvk *PrivateKey, err error) {
	c := GetCurve(curve)
	prvkesdsa, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		return nil, err
	}
	prvk = new(PrivateKey)
	err = prvk.SetEcdsaPrivKey(prvkesdsa)
	if err != nil {
		return nil, err
	}
	return prvk, nil
}
func (pk *PublicKey) GetString() (string, error) {
	if pk.curve == "" || pk.pk == "" || len(pk.pk) <= 2 {
		return "", ErrPublicKeyNull
	}
	//leng:=len(pk.pk)
	leng, err := strconv.ParseUint(pk.pk[:2], 16, 0)
	if err != nil {
		return "", err
	}
	if int(leng) >= len(pk.pk)-2 {
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
func (pk *PublicKey) setInt(x *big.Int, y *big.Int) {
	//todo
	xBase64 := base64.StdEncoding.EncodeToString(x.Bytes())
	yBase64 := base64.StdEncoding.EncodeToString(y.Bytes())
	length := len(xBase64)
	pk.pk = IntToHexString(length) + xBase64 + yBase64
}
func (pk *PublicKey) getInt() (x *big.Int, y *big.Int, err error) {
	if len(pk.pk) <= 2 {
		return nil, nil, ErrPublicKeyFormat
	}
	size, err := strconv.ParseUint(pk.pk[:2], 16, 0)
	if err != nil {
		return nil, nil, err
	}
	if int(size) >= len(pk.pk)-2 {
		return nil, nil, ErrPublicKeyFormat
	}
	xBase64 := pk.pk[2 : size+2]
	yBase64 := pk.pk[size+2:]
	xB, err := base64.StdEncoding.DecodeString(xBase64)
	if err != nil {
		return nil, nil, err
	}
	yB, err := base64.StdEncoding.DecodeString(yBase64)
	if err != nil {
		return nil, nil, err
	}
	x = new(big.Int)
	y = new(big.Int)
	x.SetBytes(xB)
	y.SetBytes(yB)
	return x, y, nil
}
func IntToHexString(i int) string {
	out := strconv.FormatInt(int64(i)%256, 16)
	if i < 16 {
		return "0" + out
	} else {
		return out
	}
}
func (pk *PublicKey) check() error {
	if len(pk.pk) <= 2 {
		return ErrPublicKeyFormat
	}
	if (pk.curve != P256) && (pk.curve != P521) && (pk.curve != P384) && (pk.curve != P224) {
		return ErrPublicKeyFormat
	}
	size, err := strconv.ParseUint(pk.pk[:2], 16, 0)
	if err != nil {
		return err
	}
	if int(size) >= len(pk.pk)-2 {
		return ErrPublicKeyFormat
	}
	return nil
}

func (pk *PublicKey) SetEcdsaPubKey(ecpk *ecdsa.PublicKey) error {
	if ecpk == nil {
		return ErrPublicKeyNull
	}
	c, err := getCurveStr(ecpk.Curve)
	if err != nil {
		return err
	}
	pk.curve = c
	pk.setInt(ecpk.X, ecpk.Y)
	return nil
}
func (pk *PublicKey) GetEcdsaPubKey() (ecpk *ecdsa.PublicKey, err error) {
	err = pk.check()
	if err != nil {
		return nil, err
	}
	ecpk = new(ecdsa.PublicKey)
	ecpk.Curve = GetCurve(pk.curve)
	ecpk.X, ecpk.Y, err = pk.getInt()
	if err != nil {
		return nil, err
	}
	return ecpk, nil
}
func (priv *PrivateKey) setBase64(base64 string) {
	priv.priv = base64
	//priv.curve = CurveType
	//priv.pk=
}
func (priv *PrivateKey) getBase64() (base64 string) {
	return priv.priv
}
func (priv *PrivateKey) setInt(i *big.Int) {
	//intByte := i.Bytes()
	priv.setBase64(base64.StdEncoding.EncodeToString(i.Bytes()))
}
func (priv *PrivateKey) getInt() (*big.Int, error) {
	intByte, err := base64.StdEncoding.DecodeString(priv.getBase64())
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
	priv.setInt(&bigI)
}
func (priv *PrivateKey) GetHexString() (hexStr string, err error) {
	bigByte, err := base64.StdEncoding.DecodeString(priv.getBase64())
	if err != nil {
		return "", err
	}
	var bigI big.Int
	bigI.SetBytes(bigByte)
	hexStr = bigI.Text(16)
	return hexStr, nil
}

//////////////////////////////////////////////
func (priv *PrivateKey) check() error {
	err := priv.PublicKey.check()
	if err != nil {
		return err
	}
	if len(priv.priv) == 0 {
		return ErrPrivKeyFormat
	}
	return nil
}
func (pri *PrivateKey) SetEcdsaPrivKey(ecprivk *ecdsa.PrivateKey) error {
	c, err := getCurveStr(ecprivk.Curve)
	if err != nil {
		return err
	}
	pri.curve = c
	pri.setInt(ecprivk.D)
	pri.PublicKey.setInt(ecprivk.X, ecprivk.Y)
	return nil
}
func (priv *PrivateKey) GetEcdsaPrivKey() (*ecdsa.PrivateKey, error) {
	err := priv.check()
	if err != nil {
		return nil, err
	}
	bigI, err := priv.getInt()
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
	bigI, _ := p.getInt()
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
