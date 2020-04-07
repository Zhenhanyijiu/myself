package sm3

const (
	blockSize  = 64
	sm3Size256 = 32
)

var IV = []uint32{0x7380166f, 0x4914b2b9, 0x172442d7, 0xda8a0600,
	0xa96f30bc, 0x163138aa, 0xe38dee4d, 0xb0fb0e4e}

type sm3Text struct {
	h   [8]uint32
	x   [blockSize]byte
	nx  int
	len uint64
}

func (s *sm3Text) Size() int {
	return sm3Size256
}
func (s *sm3Text) BlockSize() int {
	return blockSize
}
func (s *sm3Text) Reset() {
	for i := 0; i < 8; i++ {
		s.h[i] = IV[i]
	}
	s.len, s.nx = 0, 0
}

func (s *sm3Text) Sum(b []byte) []byte {
	return nil
}
func (s *sm3Text) Write(p []byte) (nn int, err error) {
	nn = len(p)
	s.len += uint64(nn)
	if s.nx > 0 {
		n := copy(s.x[s.nx:], p)
		s.nx += n
		if s.nx == blockSize {
			computeBlockSize(s, s.x[:])
			s.nx = 0
		}
		p = p[n:]
	}
	if len(p) >= blockSize {
		n := len(p) &^ (blockSize - 1)
		computeBlockSize(s, p[:n])
		p = p[n:]
	}
	if len(p) > 0 {
		s.nx = copy(s.x[:], p)
	}
	return nn, err
}
func computeBlockSize(s *sm3Text, p []byte) {

}
