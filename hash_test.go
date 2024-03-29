package hash

import (
	"github.com/flxxyz/hash"
	"testing"
)

func TestSHA1(t *testing.T) {
	t.Log(hash.SHA1("0"))
}

func TestSHA256(t *testing.T) {
	t.Log(hash.SHA256("0"))
}

func TestSHA512(t *testing.T) {
	t.Log(hash.SHA512("0"))
}

func TestMD5(t *testing.T) {
	t.Log(hash.MD5("0"))
}

func TestHMAC(t *testing.T) {
	t.Log(hash.HMAC("md5", "123456", "0"))
	t.Log(hash.HMAC("sha1", "123456", "0"))
	t.Log(hash.HMAC("sha256", "123456", "0"))
	t.Log(hash.HMAC("sha512", "123456", "0"))
}

func TestBase64Encode(t *testing.T) {
	str := "hello world"
	t.Logf("old str = %s, new str = %s", str, hash.Base64Encode(str))
}

func TestBase64Decode(t *testing.T) {
	str := "aGVsbG8gd29ybGQ="
	t.Logf("old str = %s, new str = %s", str, hash.Base64Decode(str))
}
