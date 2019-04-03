package rsa

import (
	"fmt"
	"testing"
)

var publicKey = `-----BEGIN PUBLICK KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzf1PAgsMpCAHwtInjiVD
P8P5XIi44qIv78RHZgbifJSoxyn8KwcbXXQ46fjTc7R+8rXXC2GKZm0gNqGFivjh
ICVWe/s5TjsyQAiszKdjDsWVqcasbAk2sOMepthghB7J0BlF3ImLbrXr5DdOjgV8
GXF/u3lBkL6EpRuKGJ3A9hLHHXX7xEj4NTJB+hXNL4n5TskfpmfK6FO5S+IX/bPb
yfi9u1QlAmrjWA/nRmMA2lYGiV/Iu15dvDiFVuV6RcwnkS5sEvV4rwDpXvHtALVP
FW+maIrFu63kfCkoWFD6k6OJZ7ac63OWyqgUQ8LGTdpsIsRl3Fjx1jsnf1AvZniK
IwIDAQAB
-----END PUBLICK KEY-----
`
var privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzf1PAgsMpCAHwtInjiVDP8P5XIi44qIv78RHZgbifJSoxyn8
KwcbXXQ46fjTc7R+8rXXC2GKZm0gNqGFivjhICVWe/s5TjsyQAiszKdjDsWVqcas
bAk2sOMepthghB7J0BlF3ImLbrXr5DdOjgV8GXF/u3lBkL6EpRuKGJ3A9hLHHXX7
xEj4NTJB+hXNL4n5TskfpmfK6FO5S+IX/bPbyfi9u1QlAmrjWA/nRmMA2lYGiV/I
u15dvDiFVuV6RcwnkS5sEvV4rwDpXvHtALVPFW+maIrFu63kfCkoWFD6k6OJZ7ac
63OWyqgUQ8LGTdpsIsRl3Fjx1jsnf1AvZniKIwIDAQABAoIBAF5P2jq/QFxs5eDH
Q6ELYozxd1xspOnvBC+tGFuYT15m2hfGG1S//kF9kwldCcNyIqWRqBHn6OwYjqU0
suX8KWcEcWJ2lbhB1jIHRsa05N/Ai1OeSRd16Hp4onIVo6saRaBn0dq1hwd7d7OX
tvnDfRxRIQB+9gwFjsnhoArBP7Xp9HXDC38hgz41zhL8mhkXLH15w5Wh6LB7ZfEV
CoOnff9GPjD80y6iUKEuOGZdrNhlt5ejfOVXxW4r06FAbZsyfB7QgC74OT3Bn2Bn
cRv0xriPCX7DFhUcTHVFbhfB1/fiDFUpE7sW2dVWBTxvdnP6Mq2NTYrKAxqrW25w
NL6xIxECgYEA8mXD1sADBinBhWxGYwV4p1Y2yHu1ZT6i8G0ABpFd67Ny397PptW6
co5yoYAcj/VJyzRWyBxvbrut2oSVt71jm47Ue/o6014yiEpjT1jWLn9z9a+5cCgU
2Wz2FmLKaoAQVgZG1LKDoJnLOzkbGeZSPz/mmdx8IhDl+RHP+3IA0oUCgYEA2YyD
VFCFMh6cS+4Bz2/3B4Bb3ZRFg4UsqQbJIAGfPptF4k3tpvo4bIaq/rfWySmNTF3m
Gn+q8m6sm5eHNUlTI2rbw39t6RdQDN93YT1qOJ3Z9K/E7A5DJoSPvwElcOvaAukg
8TlXhxuEjsLqQ4V47qic0+YijM8OCzX/Q03hTocCgYAWsO4XUVw3HHuQOOWR6XY0
+/4e4G3Hr4hjR2DkPIF1ShQr3tjDfmh4xtr2QV7rHwQscJbbHsZTsGAC/7xgOet0
KXe8r10IYl4KoiQPznomWioJysxqMmZZQDj9OXxHYfulgerpiFiIFT2QKrVdxFDY
uOumZZx9N+W/XiqBGFf5XQKBgFlzODC75viDs5pY46+qhvUzpxJFvAtm/8UEQECD
yjmQ2LxDdJs6uWaOPNL9Pjh6eFnVJh7xLZX/QcO8G2EipCek8XXB+Kxl5IlWfRyj
hwB+jzbX6u1ws9WrnftYek+i+oJF30CwZjfsbaXRqUdYIzBmbg2Gy9//5vNfIXVI
GYMTAoGBAIVgpUTeWaS5Tp/P1Big6gTIgVcHdSngcbbbkHnfHjA0jI0zwHCT/F2k
hvDP70P3dO7zzVq/eYTK93uDpL6Io4O0NHdL1K/tD/ZvF4xmTHOdaExX6ki5TfMy
GXLv/s+9FhZAVsphSlqqv50TsrhC0iUuqRDVROQatFRn8qXrMgol
-----END RSA PRIVATE KEY-----
`

func TestDecrypt(t *testing.T) {

}
func TestEncryptPKCS1v15(t *testing.T) {
	key, err := LoadPublicKey([]byte(publicKey))
	if err != nil {
		t.Fatalf("加载私钥失败:%s", err.Error())
	}
	orgData := "fsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfsfsdfsafsafsfssfsfsfsdfsdfsdfs"
	data, err := EncryptPKCS1v15(key, []byte(orgData))
	if err != nil {
		t.Fatalf("加密失败:%s", err.Error())
	}
	pk, err := LoadPrivateKey([]byte(privateKey))
	if err != nil {
		t.Fatalf("加载私钥失败:%s", err.Error())
	}
	udata, err := Decrypt(pk, data)
	if err != nil {
		t.Fatalf("解密失败:%s", err.Error())
	}
	fmt.Printf("解密后的内容：%s", string(udata))
	if string(udata) != orgData {
		t.Fatalf("加密和解密不一致")
	}

}
func TestGenkey(t *testing.T) {

}
func TestLoadPrivateKey(t *testing.T) {

}
func TestLoadPublicKey(t *testing.T) {

}
