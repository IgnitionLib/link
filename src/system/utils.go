package system

import (
	"crypto/rand"
	"encoding/binary"
)

// https://stackoverflow.com/questions/39756133/how-do-you-generate-a-random-uint64-in-go 🙏
func RandomUint64() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf) // Always succeeds, no need to check error
	return binary.LittleEndian.Uint64(buf)
}
