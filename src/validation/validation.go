package validation

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
)

// ValidateHMac : Safe validate a given hmac with the (secret,text) pair
func ValidateHMac(h string, s string, t string) bool {
	return SafeCompareStrings(h, ComputeSha1HMac(s, t))
}

// SafeCompareStrings : Time safe compare of 2 strings
func SafeCompareStrings(s1 string, s2 string) bool {
	return subtle.ConstantTimeCompare([]byte(s1), []byte(s2)) == 1
}

// ComputeSha1HMac : Compute Sha1 HMac from a (secret,text) pair
func ComputeSha1HMac(s string, t string) string {
	mac := hmac.New(sha1.New, []byte(s))
	mac.Write([]byte(t))
	r := mac.Sum(nil)
	return fmt.Sprintf("sha1=%s", hex.EncodeToString(r))
}
