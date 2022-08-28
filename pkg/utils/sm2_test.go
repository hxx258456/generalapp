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
	msg := `04728e49f6d4b695d5ff07415a16565eb07a5c5af7d55e452c79816883844f3f56eb003a87a54040c2da0e88ca196aecaa585585ffed50c4451dfcd01e11749fea90d9e294839e301036978b05f0466802dcebfc14beb9cb8f9bcaec3c097574e12ff7f18823ac02a3fbb6e8e57b8e5faeab1f9234eaa43d5f47764cdf774a86fcaca43f0d914a95f9cca4cc971eae6705e23d68a59958aa53a4088706d5a8e09e390b6e68b8c2c8942dcae11522cb05c6bf0190c8844536bbee6b46bbd846b31ba54750157eee0a437b110079222129ab91a52621bfbb05ec8339ea40321fa883e0c7c64462f99e933208ac2342eb0c5110ac201654e47bbc574727d9be13cf75`
	hex_byte, err := hex.DecodeString(msg)
	if err != nil {
		t.Error(err)
	}

	decMsg, err := Decrypt(private, hex_byte)
	if err != nil {
		t.Error(err)
	}

	t.Log(decMsg)
}

func TestENcrypt(t *testing.T) {
	public, _ := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")
	msg := "notary"

	encMsg, err := Encrypt(public, msg)
	if err != nil {
		t.Error(err)
	}

	t.Log(hex.EncodeToString(encMsg))
}

func TestMock(t *testing.T) {
	// keys := "[\"144\",\"C402022072640\"]"
	// keys := "[]"
	keys := ""
	content := "{\"notaryOfficeId\":40,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}"

	public, _ := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")

	encKeys, err := Encrypt(public, keys)
	if err != nil {
		t.Error(err)
	}

	encContent, err := Encrypt(public, content)
	if err != nil {
		t.Error(err)
	}

	t.Log("keys: ", hex.EncodeToString(encKeys))
	t.Log("content: ", hex.EncodeToString(encContent))
}
