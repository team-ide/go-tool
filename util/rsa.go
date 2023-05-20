package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// 私钥生成
//openssl genrsa -out rsa_private_key.pem 1024
// 公钥: 根据私钥生成
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem

// RsaEncryptByKey RSA加密，返回 base64 字符
// RsaEncryptByKey("这是需要加密的文本", "这是密钥")
func RsaEncryptByKey(origData string, publicKey string) (res string, err error) {
	bs, err := RsaEncrypt([]byte(origData), []byte(publicKey))
	if err != nil {
		return
	}
	// 经过一次base64 否则 直接转字符串乱码
	res = base64.StdEncoding.EncodeToString(bs)
	return
}

// RsaDecryptByKey RSA解密
// RsaDecryptByKey("这是加密后的文本", "这是密钥")
func RsaDecryptByKey(decrypt string, privateKey string) (res string, err error) {
	// 经过一次base64 否则 直接转字符串乱码
	bs, err := base64.StdEncoding.DecodeString(decrypt)
	if err != nil {
		return
	}
	bs, err = RsaDecrypt(bs, []byte(privateKey))
	if err != nil {
		return
	}
	res = string(bs)
	return
}

func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
