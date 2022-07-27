package service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/handler"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/log"
	types2 "github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/types"
)

type DDC1155Service struct {
	Base
}

func NewDDC1155Service() *DDC1155Service {
	return &DDC1155Service{}
}

// SafeMint
// @Description: 平台方或终端用户可以通过调用该方法进行DDC的安全生成
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 接收者账户
// @param amount DDC数量
// @param ddcURI DDC资源
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
func (d *DDC1155Service) SafeMint(opts *bind.TransactOpts, to string, amount int64, ddcURI string, data []byte) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(to) == common.HexToAddress("0") || !common.IsHexAddress(to) {
		return nil, types2.ToAccountError
	}
	if amount <= 0 {
		return nil, types2.AmountError
	}

	d.setOpts(opts)
	signedTx, err = handler.GetDDC1155().SafeMint(opts, common.HexToAddress(to), big.NewInt(amount), ddcURI, data)
	if err != nil {
		log.Error.Printf("failed to execute SafeMint: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// SafeMintBatch
// @Description: 平台方或终端用户可以通过调用该方法进行DDC的批量安全生成
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 接收者账户
// @param amounts DDC数量集合
// @param ddcURIs 对应的DDC资源链接集合
// @param data 附加数据
// @return signedTx 交易结果
// @return err
func (d *DDC1155Service) SafeMintBatch(opts *bind.TransactOpts, to string, amounts []*big.Int, ddcURIs []string, data []byte) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(to) == common.HexToAddress("0") || !common.IsHexAddress(to) {
		return nil, types2.ToAccountError
	}
	if len(amounts) != len(ddcURIs) {
		return nil, types2.DDCInfoError
	}

	d.setOpts(opts)
	signedTx, err = handler.GetDDC1155().SafeMintBatch(opts, common.HexToAddress(to), amounts, ddcURIs, data)
	if err != nil {
		log.Error.Printf("failed to execute SafeMintBatch: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// SetApprovalForAll
// @Description: DDC拥有者可以通过调用该方法进行账户授权，发起者需要是DDC的拥有者
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param operator 授权者账户
// @param approved 授权标识
// @return signedTx 签名好的交易
// @return error
func (d *DDC1155Service) SetApprovalForAll(opts *bind.TransactOpts, operator string, approved bool) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(operator) == common.HexToAddress("0") || !common.IsHexAddress(operator) {
		return nil, types2.OperatorAccountError
	}

	d.setOpts(opts)
	signedTx, err = handler.GetDDC1155().SetApprovalForAll(opts, common.HexToAddress(operator), approved)
	if err != nil {
		log.Error.Printf("failed to execute SetApprovalForAll: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
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
func (d *DDC1155Service) IsApprovedForAll(owner, operator string) (bool, error) {
	if common.HexToAddress(owner) == common.HexToAddress("0") || !common.IsHexAddress(owner) {
		return false, types2.OwnerAccountError
	}
	if common.HexToAddress(operator) == common.HexToAddress("0") || !common.IsHexAddress(operator) {
		return false, types2.OperatorAccountError
	}

	isApprovedForAll, err := handler.GetDDC1155().IsApprovedForAll(nil, common.HexToAddress(owner), common.HexToAddress(operator))
	if err != nil {
		log.Error.Printf("failed to execute IsApprovedForAll: %v", err.Error())
		return false, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return isApprovedForAll, nil
}

// SafeTransferFrom
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcID DDC唯一标识
// @param amount 需要转移的DDC数量
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
func (d *DDC1155Service) SafeTransferFrom(opts *bind.TransactOpts, from, to string, ddcID, amount int64, data []byte) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(from) == common.HexToAddress("0") || !common.IsHexAddress(from) {
		return nil, types2.FromAccountError
	}
	if common.HexToAddress(to) == common.HexToAddress("0") || !common.IsHexAddress(to) {
		return nil, types2.ToAccountError
	}
	if ddcID <= 0 {
		return nil, types2.DDCIdError
	}
	if amount <= 0 {
		return nil, types2.AmountError
	}

	d.setOpts(opts)
	signedTx, err = handler.GetDDC1155().SafeTransferFrom(opts, common.HexToAddress(from), common.HexToAddress(to), big.NewInt(ddcID), big.NewInt(amount), data)
	if err != nil {
		log.Error.Printf("failed to execute SafeTransferFrom: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// SafeBatchTransferFrom
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的批量转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcInfo 要转移的ddc集合（ddcID->amount）
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
func (d *DDC1155Service) SafeBatchTransferFrom(opts *bind.TransactOpts, from, to string, ddcInfo map[int64]int64, data []byte) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(from) == common.HexToAddress("0") || !common.IsHexAddress(from) {
		return nil, types2.FromAccountError
	}
	if common.HexToAddress(to) == common.HexToAddress("0") || !common.IsHexAddress(to) {
		return nil, types2.ToAccountError
	}
	var ddcIds []*big.Int
	var amounts []*big.Int

	for ddcID, amount := range ddcInfo {
		if ddcID <= 0 {
			return nil, types2.DDCIdError
		}
		if amount <= 0 {
			return nil, types2.AmountError
		}
		ddcIds = append(ddcIds, big.NewInt(ddcID))
		amounts = append(amounts, big.NewInt(amount))
	}

	d.setOpts(opts)
	signedTx, err = handler.GetDDC1155().SafeBatchTransferFrom(opts, common.HexToAddress(from), common.HexToAddress(to), ddcIds, amounts, data)
	if err != nil {
		log.Error.Printf("failed to execute SafeBatchTransferFrom: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
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
func (d *DDC1155Service) Burn(opts *bind.TransactOpts, owner string, ddcID int64) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(owner) == common.HexToAddress("0") || !common.IsHexAddress(owner) {
		return nil, types2.OwnerAccountError
	}
	if ddcID <= 0 {
		return nil, types2.DDCIdError
	}

	d.setOpts(opts)
	signedTx, err = handler.GetDDC1155().Burn(opts, common.HexToAddress(owner), big.NewInt(ddcID))
	if err != nil {
		log.Error.Printf("failed to execute Burn: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// BurnBatch
// @Description: DDC拥用者可以通过调用该方法进行DDC的批量销毁
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param owner 拥有者账户
// @param ddcIds 要销毁的DDC的唯一标识的集合
// @return signedTx 签名好的交易
// @return error
func (d *DDC1155Service) BurnBatch(opts *bind.TransactOpts, owner string, ddcIds []int64) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(owner) == common.HexToAddress("0") || !common.IsHexAddress(owner) {
		return nil, types2.OwnerAccountError
	}
	if len(ddcIds) == 0 {
		return nil, types2.DDCIdError
	}
	var ddcIDs []*big.Int
	for _, ddcID := range ddcIds {
		if ddcID <= 0 {
			return nil, types2.DDCIdError
		}
		ddcIDs = append(ddcIDs, big.NewInt(int64(ddcID)))
	}

	d.setOpts(opts)
	signedTx, err = handler.GetDDC1155().BurnBatch(opts, common.HexToAddress(owner), ddcIDs)
	if err != nil {
		log.Error.Printf("failed to execute BurnBatch: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// BalanceOf
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行查询当前账户拥有的对应ddcId的数量
// @receiver d
// @param owner 拥有者账户
// @param ddcID DDC唯一标识
// @return uint64 ddcId对应的数量
// @return error
func (d *DDC1155Service) BalanceOf(owner string, ddcID int64) (uint64, error) {
	if common.HexToAddress(owner) == common.HexToAddress("0") || !common.IsHexAddress(owner) {
		return 0, types2.OwnerAccountError
	}
	if ddcID <= 0 {
		return 0, types2.DDCIdError
	}

	balance, err := handler.GetDDC1155().BalanceOf(nil, common.HexToAddress(owner), big.NewInt(ddcID))
	if err != nil {
		log.Error.Printf("failed to execute BalanceOf: %v", err.Error())
		return 0, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return balance.Uint64(), nil
}

// BalanceOfBatch
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行批量查询账户拥有的DDC的数量
// @receiver d
// @param owners 要查询的账户地址
// @param ddcIds 要查询的ddcId
// @return []*big.Int ddc数量的集合
// @return error
func (d *DDC1155Service) BalanceOfBatch(owners []common.Address, ddcIds []*big.Int) ([]*big.Int, error) {

	if len(owners) != len(ddcIds) {
		return nil, types2.OwnersDDCIdsError
	}
	balance, err := handler.GetDDC1155().BalanceOfBatch(nil, owners, ddcIds)
	if err != nil {
		log.Error.Printf("failed to execute BalanceOfBatch: %v", err.Error())
		return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return balance, nil
}

// DdcURI
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的资源标识符
// @receiver d
// @param ddcID DDC唯一标识符
// @return string DDC资源标识符
// @return error
func (d *DDC1155Service) DdcURI(ddcID int64) (ddcURI string, err error) {
	ddcURI, err = handler.GetDDC1155().DdcURI(nil, big.NewInt(ddcID))
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
// @param owner DDC拥有者
// @param ddcID DDC唯一标识
// @param ddcURI DDC资源标识符
// @return signedTx 签名好的交易
// @return error
func (d *DDC1155Service) SetURI(opts *bind.TransactOpts, owner string, ddcID int64, ddcURI string) (signedTx *types.Transaction, err error) {

	if common.HexToAddress(owner) == common.HexToAddress("0") || !common.IsHexAddress(owner) {
		return nil, types2.OwnerAccountError
	}
	if ddcID <= 0 {
		return nil, types2.DDCIdError
	}
	if len(ddcURI) == 0 {
		return nil, types2.DDCURIError
	}
	d.setOpts(opts)
	signedTx, err = handler.GetDDC1155().SetURI(opts, common.HexToAddress(owner), big.NewInt(ddcID), ddcURI)
	if err != nil {
		log.Error.Printf("failed to execute SetURI: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}
