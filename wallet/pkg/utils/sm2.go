package utils

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/hxx258456/ccgo/sm2"
)

// 公钥加密
func Encrypt(public *sm2.PublicKey, msg string) ([]byte, error) {
	encByte, err := sm2.EncryptDefault(public, []byte(msg), rand.Reader)
	if err != nil {
		return nil, err
	}
	return encByte, err
}

// 私钥解密
func Decrypt(private *sm2.PrivateKey, encByte []byte) (string, error) {
	decMsg, err := sm2.DecryptDefault(private, encByte)
	if err != nil {
		return "", err
	}
	return string(decMsg), err
}

func InitSm2(publicKeyX, publicKeyY, privateKey string) (*sm2.PublicKey, *sm2.PrivateKey) {
	c := sm2.P256Sm2()

	x_byte, err := hex.DecodeString(publicKeyX)
	if err != nil {
		log.Panic(err)
	}
	y_byte, err := hex.DecodeString(publicKeyY)
	if err != nil {
		log.Panic(err)
	}
	x := big.Int{}
	y := big.Int{}
	// 公钥
	public := &sm2.PublicKey{
		X:     x.SetBytes(x_byte),
		Y:     y.SetBytes(y_byte),
		Curve: c,
	}

	private_byte, err := hex.DecodeString(privateKey)
	if err != nil {
		log.Panic(err)
	}
	d := big.Int{}
	// 私钥
	private := &sm2.PrivateKey{
		PublicKey: *public,
		D:         d.SetBytes(private_byte),
	}
	return public, private
}
