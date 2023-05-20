package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// AesEncryptCBCByKey AES加密,CBC模式，返回 base64 字符
// AesEncryptCBCByKey("这是需要加密的文本", "这是密钥")
func AesEncryptCBCByKey(origData string, key string) (res string, err error) {
	bs, err := AesCBCEncrypt([]byte(origData), []byte(key))
	if err != nil {
		return
	}
	// 经过一次 base64 否则 直接转字符串乱码
	res = base64.StdEncoding.EncodeToString(bs)
	return
}

// AesDecryptCBCByKey AES解密,CBC模式
// AesDecryptCBCByKey("这是加密后的文本", "这是密钥")
func AesDecryptCBCByKey(encrypt string, key string) (res string, err error) {
	// 经过一次base64 否则 直接转字符串乱码
	bs, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return
	}
	bs, err = AesCBCDecrypt(bs, []byte(key))
	if err != nil {
		return
	}
	res = string(bs)
	return
}

// AesEncryptECBByKey AES加密,ECB模式，返回 base64 字符
// AesEncryptECBByKey("这是需要加密的文本", "这是密钥")
func AesEncryptECBByKey(origData string, key string) (res string, err error) {
	bs, err := AesECBEncrypt([]byte(origData), []byte(key))
	if err != nil {
		return
	}
	// 经过一次base64 否则 直接转字符串乱码
	res = base64.StdEncoding.EncodeToString(bs)
	return
}

// AesDecryptECBByKey AES解密,ECB模式
// AesDecryptECBByKey("这是加密后的文本", "这是密钥")
func AesDecryptECBByKey(encrypt string, key string) (res string, err error) {
	// 经过一次base64 否则 直接转字符串乱码
	bs, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return
	}
	bs, err = AesECBDecrypt(bs, []byte(key))
	if err != nil {
		return
	}
	res = string(bs)
	return
}

func AesECBEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil, errors.New("need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func AesCBCEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AesECBDecrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return out, nil
}

func AesCBCDecrypt(encrypt, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	//origData := make([]byte, len(encrypt))
	origData := encrypt
	blockMode.CryptBlocks(origData, encrypt)
	//origData = PKCS5UnPadding(origData)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	l := int(origData[length-1])
	return origData[:(length - l)]
}
