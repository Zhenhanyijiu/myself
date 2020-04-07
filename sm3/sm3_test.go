package sm3

import (
	"encoding/binary"
	"testing"
)

func TestTp(t *testing.T) {
	var v uint32 = 0xffccddaa
	buf := make([]byte, 8)
	binary.BigEndian.PutUint32(buf, v)
	t.Logf("buf:%x\n", buf)
	n := copy(buf[4:], []byte{0, 1, 2, 3, 4, 5, 6})
	t.Logf("buf:%x,n:%v\n", buf, n)
}
