package rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	_path "path/filepath"
)

func Genkey(path string, size int) error {
	pf, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		} else {
			fmt.Printf("路径错误:%s", err.Error())
			return err
		}
	} else {
		if !pf.IsDir() {
			fmt.Printf("参数path不是目录")
			return errors.New("参数path不是目录")
		}
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateKeyPath := _path.Join(path, "rsa.private")
	publicKeyPath := _path.Join(path, "rsa.public")
	file, err := os.OpenFile(privateKeyPath, os.O_TRUNC|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	publicKey := &privateKey.PublicKey
	derPki, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLICK KEY",
		Bytes: derPki,
	}
	file, err = os.OpenFile(publicKeyPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	err = pem.Encode(file, block)
	return err
}
func LoadPublicKey(data []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("加密密钥失败")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("加密密钥失败")
	}

	return pub.(*rsa.PublicKey), nil
}
func LoadPrivateKey(data []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("加载密钥失败")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("加载密钥失败")
	}
	return priv, nil
}

func EncryptPKCS1v15(key *rsa.PublicKey, data []byte) ([]byte, error) {

	buf := bytes.NewBuffer(make([]byte, 0, int(float32(len(data))*1.5)))
	mod := key.Size() - 11
	l := len(data)
	for k := 0; k < l; k += mod {
		var bys []byte
		if k+mod >= l {
			bys = data[k:]
		} else {
			bys = data[k : k+mod]
		}
		by, err := rsa.EncryptPKCS1v15(rand.Reader, key, bys)
		if err != nil {
			return nil, err
		}
		buf.Write(by)
	}
	return buf.Bytes(), nil
}

func Decrypt(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, len(data)))
	mod := key.PublicKey.Size()
	for k := 0; k < len(data); k += mod {
		bys := data[k : k+mod]
		bys, err := rsa.DecryptPKCS1v15(rand.Reader, key, bys)
		if err != nil {
			return nil, err
		}
		buf.Write(bys)
	}
	return buf.Bytes(), nil
}
