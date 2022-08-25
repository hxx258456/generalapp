package utils

import (
	"encoding/hex"
	"testing"
)

func TestInitSm2(t *testing.T) {
	public, private := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")
	encByte, err := Encrypt(public, "test")
	if err != nil {
		t.Error(err)
	}
	t.Log("加密结果: ", hex.EncodeToString(encByte))
	msg, err := Decrypt(private, encByte)
	if err != nil {
		t.Error(err)
	}
	if msg != "test" {
		t.Fail()
	}
}

func TestDecrypt(t *testing.T) {
	_, private := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")
	msg := "049cd4de7d6e05fcbab23dbf70358f7f26929805aeedb0c587bd7b7ca731d17ccd922569fc5d1667fe744ad8ba706a1c872820dcb30c9ecc01e845797a117a7db8a6341d69229212654dbd73de814ed5cd97e5f71d1005a673795d585d7d2d2508272540c2c1ecad9a0ab7ce73ae5b2540ecaada5f680d86f3e68a03ecf0b85bf348f541cc817e"
	hex_byte, err := hex.DecodeString(msg)
	if err != nil {
		t.Error(err)
	}

	encMsg, err := Decrypt(private, hex_byte)
	if err != nil {
		t.Error(err)
	}
	t.Log(encMsg)
}
