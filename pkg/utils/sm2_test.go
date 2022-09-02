package utils

import (
	"encoding/hex"
	"encoding/json"
	"generalapp/internal/model"
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
	msg := `04526f2c9ddf634b4446e044d36a4a87855c06dc6019d5e59949e1049c1e58ad0a57445c5df8d7ff789f3b66812427544268580db5dfe2e882cd1ff0c6e0cb54128e5a245d134c7f07c389315d17f68042cf1e676d0e72a3a3b4a73c3a43a6fdf9ec2c72b96999e4cb77b2e3a7f57eb0185af21b143d18cdb166636d114ba94955f95a68bfa9da069d47cfe4c4c7998bc7924ceb62acf5fd0b98c5f7505319ed3b6fd9b2e2c53532eeed6f4f6f14da44eae989a50009f8733a53fd60e30218900d305067a1`
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
	// data := `{"keys":["121","C402","12312"],"compares":"{\"notaryOfficeId\":\"notaryOfficeId\"}","checkType":"test","content":"{\"notaryOfficeId\":40,\"serviceType\":3,\"notarizationNumber\":\"KFGZS003\",\"payTime\":\"2021-12-12 10:17:07\"}"}`
	data := new(model.CheckData)
	data.CheckType = "operatorCheck"
	data.Compares = map[string]string{"pCodeqq": "pCode", "phone": "phone", "meid": "meid"}
	data.Content = "{\"pCodeqq\":\"ooo\",\"name\":\"石宏伟3333\",\"phone\":\"iii\",\"meid\":\"ppp\",\"deviceModel\":\"iPhone 7\",\"brand\":\"APPLE\",\"deviceVersion\":\"iOS 13.3.1\",\"txId\":\"0a0d60978ba86518b43fd4b2376f5fe86c1b10742647e05f357ec291b901be45\",\"status\":\"1\",\"createTime\":\"1637543996428\"}"
	data.Keys = []string{"key1111", "key2222", "key3333"}

	byteData, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	public, _ := InitSm2("0ce2fa6e66521155f780573beb0e5f18d0aeea6b9a145f54e5c8c442efd15ecf", "fa332850bffd6e06cbbd6e29ac851fe12da302c74550c3d75e24db54a2a1fdd7", "0b15a775077e438bce6ebcb7b30c3e61d9909ee861568723661d4728ee701068")

	encData, err := Encrypt(public, string(byteData))
	if err != nil {
		t.Error(err)
	}

	t.Log("data: ", hex.EncodeToString(encData))
}
