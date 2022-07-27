package service

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/models/dto"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/tyler-smith/go-bip39"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/constant"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/config"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/contracts"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/log"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/utils"
)

const gasRatio = 1

type Base struct {
	TxService
	BlockService
}

// setOpts
// @Description: 为opts设置GasPrice和Signer
// @receiver b
// @param opts
func (b Base) setOpts(opts *bind.TransactOpts) {

	if opts.GasPrice == nil {
		opts.GasPrice = big.NewInt(int64(config.Info.GasPrice()))
	}
	if opts.Signer == nil {
		opts.Signer = config.Info.SignEventListener().SignEvent
	}

}

// GetNonce
// @Description:查询对应账户下的最新nonce
// @receiver b
// @param addr 账户地址
// @return uint64 最新nonce
// @return error
func (b Base) GetNonce(addr common.Address) (uint64, error) {
	nonce, err := config.Info.Conn().PendingNonceAt(context.Background(), addr)
	if err != nil {
		log.Error.Printf("failed to GetNonce: %v", err.Error())
		return 0, err
	}

	return nonce, nil
}

// DDCIdByHash
// @Description: 通过交易hash查询相关ddc的ddcId
// @receiver b
// @param hash 交易哈希
// @return []uint64 对应的ddcId的集合
func (b Base) DDCIdByHash(hash string) (res []uint64) {
	// 查出相关事件
	events, err := b.GetTxEvents(common.HexToHash(hash))
	if err != nil {
		log.Error.Printf("failed to GetTxEvents: %v", err.Error())
		return nil
	}
	for _, event := range events {
		switch e := event.(type) {
		case *contracts.DDC721Transfer: // 创建、转让、销毁
			return append(res, e.DdcId.Uint64())
		case *contracts.DDC721SetURI:
			return append(res, e.DdcId.Uint64())
		case *contracts.DDC721Approval:
			return append(res, e.DdcId.Uint64())
		case *contracts.DDC721EnterBlacklist:
			return append(res, e.DdcId.Uint64())
		case *contracts.DDC721ExitBlacklist:
			return append(res, e.DdcId.Uint64())
		case *contracts.DDC1155TransferSingle:
			return append(res, e.DdcId.Uint64())
		case *contracts.DDC1155TransferBatch:
			for _, ddcID := range e.DdcIds {
				res = append(res, ddcID.Uint64())
			}
			return
		case *contracts.DDC1155SetURI:
			return append(res, e.DdcId.Uint64())
		}
	}
	return nil
}

// NewRawTx
// @Description: 组装待签名的交易
// @receiver b
// @param opts
// @param contract 目标合约地址
// @param input 目标方法的字节码
// @return *types.Transaction 待签名的交易
// @return error
func (b Base) NewRawTx(opts *bind.TransactOpts, contract common.Address, input []byte) (*types.Transaction, error) {
	nonce, err := b.GetNonce(opts.From)
	if err != nil {
		return nil, err
	}
	// Normalize value
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	baseTx := &types.LegacyTx{
		To:       &contract,
		Nonce:    nonce,
		GasPrice: opts.GasPrice,
		Gas:      opts.GasLimit,
		Value:    value,
		Data:     input,
	}

	return types.NewTx(baseTx), nil
}

// Bech32ToHex
// @Description: irita链账户转以太坊账户格式
// @receiver b
// @param addr irita链账户
// @return string 以太坊账户
// @return error
func (b Base) Bech32ToHex(addr string) (string, error) {
	bz, err := utils.GetFromBech32(addr, "iaa")
	if err != nil {
		log.Error.Printf("failed to convert address: %v", err.Error())
		return "", err
	}

	return "0x" + hexutils.BytesToHex(bz), nil
}

// HexToBech32
// @Description: 以太坊账户格式转irita链账户
// @receiver b
// @param addr 0x开头的以太坊账户
// @return string irita链账户
// @return error
func (b Base) HexToBech32(addr string) (string, error) {
	bz, err := hexutil.Decode(addr)
	if err != nil {
		return "", err
	}
	t, err := utils.ConvertAndEncode("iaa", bz)
	if err != nil {
		log.Error.Printf("convert address: %v", err.Error())
		return "", err
	}

	return t, nil
}

// StringToPrivateKey
// @Description: 从明文的私钥字符串转换成该类型
// @receiver b
// @param privateKeyStr 字符串形式的私钥
// @return *ecdsa.PrivateKey ECDSA格式的私钥
// @return error
func (b Base) StringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyByte, err := hexutil.Decode(privateKeyStr)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// EstimateGasLimit
// @Description: 根据合约名、方法名以及入参预估gasLimit
// @receiver b
// @param opts
// @param contrName 合约名（见constant/contr_name.go）
// @param funcName 方法名（见constant/func_name.go）
// @param params 方法入参
// @return gasLimit
// @return err
func (b Base) EstimateGasLimit(opts *bind.TransactOpts, contrName, funcName string, params ...interface{}) (gasLimit uint64, err error) {
	var abiTmp *abi.ABI
	var contrAddr common.Address
	//根据不同合约名获取不同abi字节码，以及获取合约地址
	switch contrName {
	case constant.ContrAuthority:
		abiTmp, _ = contracts.AuthorityMetaData.GetAbi()
		contrAddr = config.Info.AuthorityAddress()
	case constant.ContrCharge:
		abiTmp, _ = contracts.ChargeMetaData.GetAbi()
		contrAddr = config.Info.ChargeAddress()
	case constant.ContrDDC721:
		abiTmp, _ = contracts.DDC721MetaData.GetAbi()
		contrAddr = config.Info.Ddc721Address()
	case constant.ContrDDC1155:
		abiTmp, _ = contracts.DDC1155MetaData.GetAbi()
		contrAddr = config.Info.Ddc1155Address()
	}
	input, err := abiTmp.Pack(funcName, params...)
	if err != nil {
		log.Error.Printf("failed to pack abiTmp: %v", err.Error())
		return 0, err
	}
	b.setOpts(opts)
	//预估gas
	gas, err := b.EstimateGas(opts, &contrAddr, input)
	if err != nil {
		log.Error.Printf("failed to EstimateGas: %v", err.Error())
		return 0, err
	}

	gasLimit = gas * gasRatio

	return gasLimit, err
}

// EstimateGas 预估gas
func (b Base) EstimateGas(opts *bind.TransactOpts, contract *common.Address, input []byte) (uint64, error) {
	if contract != nil {
		// Gas estimation cannot succeed without code for method invocations.
		if code, err := config.Info.Conn().PendingCodeAt(context.Background(), *contract); err != nil {
			return 0, err
		} else if len(code) == 0 {
			return 0, bind.ErrNoCode
		}
	}
	msg := ethereum.CallMsg{
		From:      opts.From,
		To:        contract,
		GasPrice:  opts.GasPrice,
		GasTipCap: opts.GasTipCap,
		GasFeeCap: opts.GasFeeCap,
		Value:     opts.Value,
		Data:      input,
	}

	return config.Info.Conn().EstimateGas(context.Background(), msg)
}

// CreateAccount
// @Description:平台方或终端用户可以通过此方法生成离线账户
// @receiver b
// @return *dto.Account 返回的账户信息
// @return error
func (b Base) CreateAccount() (*dto.Account, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	seed := bip39.NewSeed(mnemonic, "")

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		return nil, err
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, err
	}
	privateKeyHex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return nil, err
	}
	publicKeyHex, err := wallet.PublicKeyHex(account)
	if err != nil {
		return nil, err
	}

	return dto.NewAccount(account.Address.Hex(), publicKeyHex, privateKeyHex, mnemonic), nil
}
