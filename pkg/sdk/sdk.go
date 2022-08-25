// 链码sdk初始化包
//
package sdk

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hxx258456/ccgo/sm2"
	"github.com/hxx258456/fabric-sdk-go-gm/pkg/core/config"
	"github.com/hxx258456/fabric-sdk-go-gm/pkg/gateway"
)

var (
	wallet *gateway.Wallet
)

type Sdk struct {
	GW               *gateway.Gateway
	Netowrk          *gateway.Network
	Contract         *gateway.Contract
	ChaincodeName    string
	ChannelName      string
	Public           *sm2.PublicKey
	Private          *sm2.PrivateKey
	OrganizationsDir string
	WalletLabel      string //钱包身份信息
}

func (s *Sdk) InitSdk() {
	log.Printf("============ %s sdk initing ============", s.ChaincodeName)
	// 清理钱包，确保获取最新的客户端信息
	wallet.Remove(s.WalletLabel)
	if !wallet.Exists(s.WalletLabel) {
		// 注册钱包sdk身份信息
		err := s.populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		s.OrganizationsDir,
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)
	var err error
	s.GW, err = gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, s.WalletLabel),
		// gateway.WithTimeout(2*time.Hour),
	)
	if err != nil {
		log.Fatal(err)
	}

	s.Netowrk, err = s.GW.GetNetwork(s.ChannelName)
	if err != nil {
		log.Fatal(err)
	}

	s.Contract = s.Netowrk.GetContract(s.ChaincodeName)
}

func (s *Sdk) populateWallet(wallet *gateway.Wallet) error {
	log.Printf("============ 初始化sdk身份信息:%s ============\n", s.ChaincodeName)
	credPath := filepath.Join(
		s.OrganizationsDir,
		"peerOrganizations",
		"org1.example.com",
		"users",
		"Admin@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put(s.WalletLabel, identity)
}

func init() {
	log.Println("============ generalapp wallet initing ============")
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}
	wallet, err = gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}
}
