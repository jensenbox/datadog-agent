package testlib

type MyStruct struct {
	Size int
}

//go:noinline
func (m *MyStruct) MyFunc(b []byte) []byte {
	for i := 0; i < m.Size; i++ {
		b = append(b, byte(i))
	}

	return b
}
