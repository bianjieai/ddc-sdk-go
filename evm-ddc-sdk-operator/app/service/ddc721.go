package service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-operator/app/handler"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-operator/pkg/log"
	types2 "github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-operator/pkg/types"
)

type DDC721Service struct {
	Base
}

func NewDDC721Service() *DDC721Service {
	return &DDC721Service{}
}

// Approve
// @Description: DDC拥有者可以通过调用该方法进行DDC的授权，发起者需要是DDC的拥有者
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 授权者账户
// @param ddcID DDC唯一标识
// @return signedTx 签名好的交易
// @return error
func (d *DDC721Service) Approve(opts *bind.TransactOpts, to string, ddcID int64) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(to) == common.HexToAddress("0") || !common.IsHexAddress(to) {
		return nil, types2.ToAccountError
	}
	if ddcID <= 0 {
		return nil, types2.DDCIdError
	}

	//设置opts
	d.setOpts(opts)
	signedTx, err = handler.GetDDC721().Approve(opts, common.HexToAddress(to), big.NewInt(ddcID))
	if err != nil {
		log.Error.Printf("failed to execute FuncApprove: %v", err.Error())
		return signedTx, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// GetApprove
// @Description: 运营方、平台方或终端用户都可以通过调用该方法查询DDC的授权情况
// @receiver d
// @param ddcID DDC唯一标识
// @return string 授权的账户
// @return error
func (d *DDC721Service) GetApprove(ddcID int64) (string, error) {
	if ddcID <= 0 {
		return "", types2.DDCIdError
	}

	approved, err := handler.GetDDC721().GetApproved(nil, big.NewInt(ddcID))
	if err != nil {
		log.Error.Printf("failed to execute GetApprove: %v", err.Error())
		return "", types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return approved.String(), nil
}

// SetApprovalForAll
// @Description: DDC拥有者可以通过调用该方法进行账户授权，发起者需要是DDC的拥有者
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param operator 授权者账户
// @param approved 授权标识
// @return signedTx 签名好的交易
// @return error
func (d *DDC721Service) SetApprovalForAll(opts *bind.TransactOpts, operator string, approved bool) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(operator) == common.HexToAddress("0") || !common.IsHexAddress(operator) {
		return nil, types2.OperatorAccountError
	}

	//设置opts
	d.setOpts(opts)
	signedTx, err = handler.GetDDC721().SetApprovalForAll(opts, common.HexToAddress(operator), approved)
	if err != nil {
		log.Error.Printf("failed to execute SetApprovalForAll: %v", err.Error())
		return signedTx, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// IsApprovedForAll
// @Description: 运营方、平台方或终端用户可以通过调用该方法进行账户授权查询
// @receiver d
// @param owner 拥有者账户
// @param operator 授权者账户
// @return bool 授权标识
// @return error
func (d *DDC721Service) IsApprovedForAll(owner, operator string) (bool, error) {
	if common.HexToAddress(owner) == common.HexToAddress("0") || !common.IsHexAddress(owner) {
		return false, types2.OwnerAccountError
	}
	if common.HexToAddress(operator) == common.HexToAddress("0") || !common.IsHexAddress(operator) {
		return false, types2.OperatorAccountError
	}

	isApprovedForAll, err := handler.GetDDC721().IsApprovedForAll(nil, common.HexToAddress(owner), common.HexToAddress(operator))
	if err != nil {
		log.Error.Printf("failed to execute IsApprovedForAll: %v", err.Error())
		return false, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return isApprovedForAll, nil
}

// SafeTransferFrom
// @Description: DDC的拥有者或授权者可以通过调用该方法进行DDC的转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcID DDC唯一标识
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
func (d *DDC721Service) SafeTransferFrom(opts *bind.TransactOpts, from, to string, ddcID int64, data []byte) (signedTx *types.Transaction, err error) {
	if common.HexToAddress(from) == common.HexToAddress("0") || !common.IsHexAddress(from) {
		return nil, types2.FromAccountError
	}
	if common.HexToAddress(to) == common.HexToAddress("0") || !common.IsHexAddress(to) {
		return nil, types2.ToAccountError
	}
	if ddcID <= 0 {
		return nil, types2.DDCIdError
	}

	//设置opts
	d.setOpts(opts)
	signedTx, err = handler.GetDDC721().SafeTransferFrom(opts, common.HexToAddress(from), common.HexToAddress(to), big.NewInt(ddcID), data)
	if err != nil {
		log.Error.Printf("failed to execute SafeTransferFrom: %v", err.Error())
		return signedTx, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// TransferFrom
// @Description: DDC拥有者或授权者可以通过调用该方法进行DDC的转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcID ddc唯一标识
// @return signedTx 签名好的交易
// @return error
func (d *DDC721Service) TransferFrom(opts *bind.TransactOpts, from, to string, ddcID int64) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(from) == common.HexToAddress("0") || !common.IsHexAddress(from) {
		return nil, types2.FromAccountError
	}
	if common.HexToAddress(to) == common.HexToAddress("0") || !common.IsHexAddress(to) {
		return nil, types2.ToAccountError
	}
	if ddcID <= 0 {
		return nil, types2.DDCIdError
	}

	//设置opts
	d.setOpts(opts)
	signedTx, err = handler.GetDDC721().TransferFrom(opts, common.HexToAddress(from), common.HexToAddress(to), big.NewInt(ddcID))
	if err != nil {
		log.Error.Printf("failed to execute TransferFrom: %v", err.Error())
		return signedTx, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// Freeze
// @Description: 运营方可以通过调用该方法进行DDC的冻结
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param ddcID DDC唯一标识
// @return signedTx 签名好的交易
// @return error
func (d *DDC721Service) Freeze(opts *bind.TransactOpts, ddcID int64) (signedTx *types.Transaction, err error) {

	if ddcID <= 0 {
		return nil, types2.DDCIdError
	}

	//设置opts
	d.setOpts(opts)
	signedTx, err = handler.GetDDC721().Freeze(opts, big.NewInt(ddcID))
	if err != nil {
		log.Error.Printf("failed to execute Freeze: %v", err.Error())
		return signedTx, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// UnFreeze
// @Description: 运营方可以通过调用该方法进行DDC的解冻
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param ddcID DDC唯一标识
// @return signedTx 签名好的交易
// @return error
func (d *DDC721Service) UnFreeze(opts *bind.TransactOpts, ddcID int64) (signedTx *types.Transaction, err error) {
	if ddcID <= 0 {
		return nil, types2.DDCIdError
	}

	//设置opts
	d.setOpts(opts)
	signedTx, err = handler.GetDDC721().UnFreeze(opts, big.NewInt(ddcID))
	if err != nil {
		log.Error.Printf("failed to execute UnFreeze: %v", err.Error())
		return signedTx, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// Burn
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的销毁
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param ddcID DDC唯一标识
// @return signedTx 签名好的交易
// @return error
func (d *DDC721Service) Burn(opts *bind.TransactOpts, ddcID int64) (signedTx *types.Transaction, err error) {

	if ddcID <= 0 {
		return nil, types2.DDCIdError
	}

	//设置opts
	d.setOpts(opts)
	signedTx, err = handler.GetDDC721().Burn(opts, big.NewInt(ddcID))
	if err != nil {
		log.Error.Printf("failed to execute Burn: %v", err.Error())
		return signedTx, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// BalanceOf
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行查询当前账户拥有的DDC的数量
// @receiver d
// @param owner 拥有者账户
// @return uint64 ddc的数量
// @return error
func (d *DDC721Service) BalanceOf(owner string) (uint64, error) {
	if common.HexToAddress(owner) == common.HexToAddress("0") || !common.IsHexAddress(owner) {
		return 0, types2.OwnerAccountError
	}

	balance, err := handler.GetDDC721().BalanceOf(nil, common.HexToAddress(owner))
	if err != nil {
		log.Error.Printf("failed to execute BalanceOf: %v", err.Error())
		return 0, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return balance.Uint64(), nil
}

// OwnerOf
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的拥有者
// @receiver d
// @param ddcID ddc唯一标识
// @return string 拥有者账户
// @return error
func (d *DDC721Service) OwnerOf(ddcID int64) (string, error) {
	if ddcID <= 0 {
		return "", types2.DDCIdError
	}

	owner, err := handler.GetDDC721().OwnerOf(nil, big.NewInt(ddcID))
	if err != nil {
		log.Error.Printf("failed to execute OwnerOf: %v", err.Error())
		return "", types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return owner.String(), nil
}

// Name
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的名称
// @receiver d
// @return string DDC运营方名称
// @return error
func (d *DDC721Service) Name() (string, error) {
	name, err := handler.GetDDC721().Name(nil)
	if err != nil {
		log.Error.Printf("failed to execute Name: %v", err.Error())
		return "", types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return name, nil
}

// Symbol
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的符号标识
// @receiver d
// @return string DDC运营方符号
// @return error
func (d *DDC721Service) Symbol() (string, error) {
	symbol, err := handler.GetDDC721().Symbol(nil)
	if err != nil {
		log.Error.Printf("failed to execute Symbol: %v", err.Error())
		return "", types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return symbol, nil
}

// DdcURI
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的资源标识符
// @receiver d
// @param ddcID DDC唯一标识符
// @return string DDC资源标识符
// @return error
func (d *DDC721Service) DdcURI(ddcID int64) (string, error) {
	ddcURI, err := handler.GetDDC721().DdcURI(nil, big.NewInt(ddcID))
	if err != nil {
		log.Error.Printf("failed to execute DdcURI: %v", err.Error())
		return "", types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return ddcURI, nil
}

// SetURI
// @Description: DDC拥有者或DDC授权者通过调用该方法对DDC的资源标识符进行设置
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param ddcID DDC唯一标识
// @param ddcURI DDC资源标识符
// @return signedTx 签名好的交易
// @return err
func (d *DDC721Service) SetURI(opts *bind.TransactOpts, ddcID int64, ddcURI string) (signedTx *types.Transaction, err error) {
	if ddcID <= 0 {
		return nil, types2.DDCIdError
	}
	if len(ddcURI) == 0 {
		return nil, types2.DDCURIError
	}

	//设置opts
	d.setOpts(opts)
	signedTx, err = handler.GetDDC721().SetURI(opts, big.NewInt(ddcID), ddcURI)
	if err != nil {
		log.Error.Printf("failed to execute SetURI: %v", err.Error())
		return signedTx, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// SetNameAndSymbol
// @Description: 合约拥有者调用该方法对721的名称及符号进行设置。
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param name 名称
// @param symbol 符号
// @return signedTx 签名好的交易
// @return err
func (d *DDC721Service) SetNameAndSymbol(opts *bind.TransactOpts, name, symbol string) (signedTx *types.Transaction, err error) {

	if len(name) == 0 {
		return nil, types2.NameError
	}
	if len(symbol) == 0 {
		return nil, types2.SymbolError
	}

	//设置opts
	d.setOpts(opts)
	signedTx, err = handler.GetDDC721().SetNameAndSymbol(opts, name, symbol)
	if err != nil {
		log.Error.Printf("failed to execute SetNameAndSymbol: %v", err.Error())
		return signedTx, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}
