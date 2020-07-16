package test

import (
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestPrint2(t *testing.T) {
	sh := sha256.New()
	n, err := sh.Write([]byte("hello"))
	assert.NoError(t, err)

	t.Logf("n:%v\n", n)
	res := sh.Sum([]byte{0xff, 0xee, 0xcc, 0xdd})
	t.Logf("size:%v,blocsize:%v\n", sh.Size(), sh.BlockSize())
	t.Logf("res:%v,len:%v\n", new(big.Int).SetBytes(res).Text(16), len(res))
	/////
	sh = sha256.New()
	n, err = sh.Write([]byte("hello"))
	n, err = sh.Write([]byte("golang"))
	assert.NoError(t, err)
	t.Logf("n:%v\n", n)
	res = sh.Sum(nil)
	t.Logf("size:%v,blocsize:%v\n", sh.Size(), sh.BlockSize())
	t.Logf("res:%v,len:%v\n", new(big.Int).SetBytes(res).Text(16), len(res))
	/////
	sh = sha256.New()
	n, err = sh.Write([]byte("hello"))

	assert.NoError(t, err)
	t.Logf("n:%v\n", n)
	res = sh.Sum(nil)
	t.Logf("size:%v,blocsize:%v\n", sh.Size(), sh.BlockSize())
	t.Logf("res:%v,len:%v\n", new(big.Int).SetBytes(res).Text(16), len(res))
	n, err = sh.Write([]byte("golang"))
	res = sh.Sum(nil)
	t.Logf("size:%v,blocsize:%v\n", sh.Size(), sh.BlockSize())
	t.Logf("res:%v,len:%v\n", new(big.Int).SetBytes(res).Text(16), len(res))
	////0011 1111
	ch := 0
	length := 7
	r := length &^ ch
	t.Logf("r:%v\n", r)
}
func TestPrint3(t *testing.T) {
	//i, j := 10, 15
	//for k := 0; k < 10; k++ {
	//	f1(t, i+k, j+k)
	//	f2(t, i+k, j+k)
	//	f3(t, i+k, j+k)
	//}
	//0110 &^ 1011 = 0100
	//1011 &^ 1101 = 0010
	//i, j = 0x6, 11
	//f1(t, i, j)
	//i, j = 11, 13
	//f1(t, i, j)
	//&^ 二元运算符的操作结果是“bit clear”
	//a &^ b 的意思就是 将b中为1的位 对应于a的位清0， a中其他位不变
	j := 0x3f //0011 1111

	for i := 0; i < 513; i++ {
		//t.Logf("i:%v, ", i)
		f1(t, i, j)
	}
}
func f1(t *testing.T, i, j int) {
	r := i &^ j
	t.Logf("index:%v, f1,r:0x%x, %v\n", i, r, r)
}
func f2(t *testing.T, i, j int) {
	r := i & j
	t.Logf("f2,r:%x\n", r)
}
func f3(t *testing.T, i, j int) {
	r := i & j
	t.Logf("f3,r:%x\n", r)
}
