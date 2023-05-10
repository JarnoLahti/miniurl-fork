// Package miniurl providers building blocks for url shortener app.
package miniurl

import (
	"crypto/md5"
	"encoding/hex"
)

// Hash generates uses md5 hash algorithm and produces 32 bytes long hex encoded string
func Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
