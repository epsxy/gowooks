package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*===================== SafeCompareString Tests =====================*/
func TestShouldSafeCompareTwoSameStrings(t *testing.T) {
	isEqual := SafeCompareStrings("str1", "str1")

	assert.True(t, isEqual, "Should compare two same strings")
}
func TestShouldSafeCompareTwoDifferentStrings(t *testing.T) {
	isEqual := SafeCompareStrings("str1", "str2")

	assert.False(t, isEqual, "Should compare two different strings")
}

/*===================== ComputeSha1HMac Tests =====================*/
func TestShouldDropKthElement(t *testing.T) {
	sample := "this text will be hashed"
	secret := "secret"

	pruned := ComputeSha1HMac(secret, sample)
	expected := "sha1=7608a419d6585d1e15e3d501ea01bd692383454d"

	assert.Equal(t, expected, pruned, "Should hash text with secret")
}

/*===================== ValidateHMac Tests =====================*/
func TestShouldValidateHMacTest(t *testing.T) {
	isEqual := ValidateHMac(
		"sha1=7608a419d6585d1e15e3d501ea01bd692383454d",
		"secret",
		"this text will be hashed",
	)

	assert.True(t, isEqual, "Should validate hmac")
}

func TestShouldInvalidateHMacTest(t *testing.T) {
	isEqual := ValidateHMac(
		"sha1=7608a419d6585d1e15e3d501ea01bd692383454d",
		"wrong secret",
		"this text will be hashed",
	)

	assert.False(t, isEqual, "Should validate hmac")
}
