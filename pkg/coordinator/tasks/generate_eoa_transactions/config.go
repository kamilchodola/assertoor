package generateeoatransactions

import (
	"errors"
	"math/big"
)

type Config struct {
	LimitPerBlock int    `yaml:"limitPerBlock" json:"limitPerBlock"`
	LimitTotal    int    `yaml:"limitTotal" json:"limitTotal"`
	LimitPending  int    `yaml:"limitPending" json:"limitPending"`
	PrivateKey    string `yaml:"privateKey" json:"privateKey"`
	ChildWallets  uint64 `yaml:"childWallets" json:"childWallets"`
	WalletSeed    string `yaml:"walletSeed" json:"walletSeed"`

	RefillPendingLimit uint64   `yaml:"refillPendingLimit" json:"refillPendingLimit"`
	RefillFeeCap       *big.Int `yaml:"refillFeeCap" json:"refillFeeCap"`
	RefillTipCap       *big.Int `yaml:"refillTipCap" json:"refillTipCap"`
	RefillAmount       *big.Int `yaml:"refillAmount" json:"refillAmount"`
	RefillMinBalance   *big.Int `yaml:"refillMinBalance" json:"refillMinBalance"`

	LegacyTransactions  bool     `yaml:"legacyTransactions" json:"legacyTransactions"`
	TransactionFeeCap   *big.Int `yaml:"transactionFeeCap" json:"transactionFeeCap"`
	TransactionTipCap   *big.Int `yaml:"transactionTipCap" json:"transactionTipCap"`
	TransactionGasLimit uint64   `yaml:"transactionGasLimit" json:"transactionGasLimit"`
	TargetAddress       string   `yaml:"targetAddress" json:"targetAddress"`
	RandomTarget        bool     `yaml:"randomTarget" json:"randomTarget"`
	ContractDeployment  bool     `yaml:"contractDeployment" json:"contractDeployment"`
	TransactionData     string   `yaml:"transactionData" json:"transactionData"`
	RandomAmount        bool     `yaml:"randomAmount" json:"randomAmount"`
	TransactionAmount   *big.Int `yaml:"transactionAmount" json:"transactionAmount"`

	ClientPattern string `yaml:"clientPattern" json:"clientPattern"`
}

func DefaultConfig() Config {
	return Config{
		RefillPendingLimit:  200,
		RefillFeeCap:        big.NewInt(500000000000),        // 500 Gwei
		RefillTipCap:        big.NewInt(1000000000),          // 1 Gwei
		RefillAmount:        big.NewInt(1000000000000000000), // 1 ETH
		RefillMinBalance:    big.NewInt(500000000000000000),  // 0.5 ETH
		TransactionFeeCap:   big.NewInt(100000000000),        // 100 Gwei
		TransactionTipCap:   big.NewInt(1000000000),          // 1 Gwei
		TransactionGasLimit: 50000,
		TransactionAmount:   big.NewInt(0),
	}
}

func (c *Config) Validate() error {
	if c.LimitPerBlock == 0 && c.LimitTotal == 0 && c.LimitPending == 0 {
		return errors.New("either limitPerBlock or limitTotal or limitPending must be set")
	}

	if c.PrivateKey == "" {
		return errors.New("privateKey must be set")
	}

	return nil
}
