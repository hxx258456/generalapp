package config

type Public struct {
	PublicKeyX string `yaml:"publicKeyX"`
	PublicKeyY string `yaml:"publicKeyY"`
}

type Private struct {
	PrivateKey string `yaml:"privateKey"`
}

type Chaincodes struct {
	Channel          string  `yaml:"channel"`
	Chaincode        string  `yaml:"chaincode"`
	Public           Public  `yaml:"public"`
	Private          Private `yaml:"private"`
	OrganizationsDir string  `yaml:"organizationsDir"`
	WalletLabel      string  `yaml:"walletLabel"`
}
