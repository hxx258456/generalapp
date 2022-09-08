package sdkpool

import (
	"generalapp/config"
	"generalapp/pkg/sdk"
	"generalapp/pkg/utils"
	"log"
)

var SdkPoll map[string]sdk.Sdk = map[string]sdk.Sdk{}

func init() {
	log.Println("============ sdkpool initing ============")
	for _, v := range config.Get().Chaincodes {
		sdk_ := sdk.Sdk{
			ChaincodeName:    v.Chaincode,
			ChannelName:      v.Channel,
			OrganizationsDir: v.OrganizationsDir,
			WalletLabel:      v.WalletLabel,
			CerdPath:         v.CerdPath,
			CcpPath:          v.CcpPath,
		}
		sdk_.InitSdk()
		log.Printf("============ init sdk[%s] sm2 ============", v.Chaincode)
		sdk_.Public, sdk_.Private = utils.InitSm2(v.Public.PublicKeyX, v.Public.PublicKeyY, v.Private.PrivateKey)
		SdkPoll[v.Chaincode] = sdk_
	}
}
