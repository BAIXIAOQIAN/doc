package golang

import (
	"crypto/sha256"
	"encoding/hex"
)

//sha256
func Sum256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
