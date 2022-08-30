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
	msg := `041366d4eaa27e1ddd3822a3a9bfa00812cd238c352fe648a6f95783c639703a7b3c5115cebfff799e6b28720f61a25d5e038dd55fd15e90cd2656e2e050806bd9c43f9090564de052befe6d2794a1f86a83e31e5353cb78c0264512a1c289d75b129babd5429757c5d76e4eca74443af719c7e8a3eebd4f66c7988168d92335864828fe7b1dfd5b361a8d71606e132c76d61da62b18f161b6fcc9a87870cc73e0a6e0796b59c623f1c3d2554dbc9915f12348078816065b9d68e93f6b51efd4a3b65020400275ada7ee7a3d4238f16240d9338572bf36837469726a9089a4e64c0b4373cb4afd46f35da8833f3db5d357830cca0ae4633f6029cbc357536f53640821e790033da1590d0ae98631a667717c4b560a4569a82f60e01c6838747bf663d201eeeb15b7e8e1ab7fa1f7fb6b3bb78169a9e5bc2e09492ba5598c9c759f17e2ab809407960d2e20755f84b281d8ba576255f2369c079044c5666b71a9a94d5512d0ad60349249c87691e05b2f0842a38662e0cdf376954378840906295b044f85074969f7162e762b091160157afbf4065fb7be6357c5e8d6300fb4baf6a1ded368e24e31bf`
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
	// msg := "notary"
	// msg := "{\"docType\":\"notary\"}"
	// msg := `{
	// 	"keys": "[\"144\",\"C402022072640\"]",
	// 	"content": "{\"notaryOfficeId\":40,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}"
	// }`
	msg := `{"keys":"[\"key115\", \"key225\", \"key335\"]","content":"{\"pCode\":\"ooo\",\"name\":\"石宏伟\",\"phone\":\"iii\",\"meid\":\"ppp\",\"deviceModel\":\"iPhone 7\",\"brand\":\"APPLE\",\"deviceVersion\":\"iOS 13.3.1\",\"txId\":\"0a0d60978ba86518b43fd4b2376f5fe86c1b10742647e05f357ec291b901be45\",\"status\":\"1\",\"createTime\":\"1637543996428\"}"}`

	encMsg, err := Encrypt(public, msg)
	if err != nil {
		t.Error(err)
	}

	t.Log(hex.EncodeToString(encMsg))
}

func TestMockInitChaincodeData(t *testing.T) {
	data := `{"docType":"notary"}`

	public, _ := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")

	encData, err := Encrypt(public, data)
	if err != nil {
		t.Error(err)
	}

	t.Log("data: ", hex.EncodeToString(encData))
}

func TestMockAddData(t *testing.T) {
	// data := `{"keys":["144","C402022072640"],"content":"{\"notaryOfficeId\":40,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}"}`

	data := `{"keys":["121","C402","12312"],"content":"{\"notaryOfficeId\":40,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}"}`

	public, _ := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")

	encData, err := Encrypt(public, data)
	if err != nil {
		t.Error(err)
	}

	t.Log("data: ", hex.EncodeToString(encData))
}

func TestMockUpdateData(t *testing.T) {
	data := `{"keys":["144","C402022072640"],"content":"{\"notaryOfficeId\":40,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}"}`

	public, _ := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")

	encData, err := Encrypt(public, data)
	if err != nil {
		t.Error(err)
	}

	t.Log("data: ", hex.EncodeToString(encData))
}

func TestMockDeleteData(t *testing.T) {
	data := `{"keys":["144","C402022072640"]}`

	public, _ := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")

	encData, err := Encrypt(public, data)
	if err != nil {
		t.Error(err)
	}

	t.Log("data: ", hex.EncodeToString(encData))
}

func TestMockQueryByPageData(t *testing.T) {
	data := "{\"keys\":[],\"pagesize\":\"4\",\"nextmark\":\"\\u0000notary\\u0000key1155\\u0000key2266\\u0000key3377\\u0000\"}"

	public, _ := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")

	encData, err := Encrypt(public, data)
	if err != nil {
		t.Error(err)
	}

	t.Log("data: ", hex.EncodeToString(encData))
}

func TestMockCheckData(t *testing.T) {
	data := `{"keys":["121","C402","12312"],"compares":["notaryOfficeId"],"checkType":"test","content":"{\"notaryOfficeId\":23,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}"}`

	public, _ := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")

	encData, err := Encrypt(public, data)
	if err != nil {
		t.Error(err)
	}

	t.Log("data: ", hex.EncodeToString(encData))
}
