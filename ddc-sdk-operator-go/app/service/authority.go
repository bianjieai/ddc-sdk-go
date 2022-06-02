package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/app/handler"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/app/models/dto"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/pkg/log"
	types2 "github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/pkg/types"
)

type AuthorityService struct {
	Base
}

func NewAuthorityService() *AuthorityService {
	return &AuthorityService{}
}

// AddAccountByOperator
// @Description: 运营方可以通过调用该方法直接对平台方或平台方的终端用户进行创建
// @receiver a
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param account DDC链账户地址
// @param accName DDC账户对应的账户名称
// @param accDID DDC账户对应的DID信息
// @param leaderDID 该账户对应的上级账户的DID
// @return signedTx 签名好的交易
// @return error
func (a *AuthorityService) AddAccountByOperator(opts *bind.TransactOpts, account, accName, accDID, leaderDID string) (signedTx *types.Transaction, err error) {

	if !common.IsHexAddress(account) {
		return nil, types2.AccountError
	}
	if len(accName) == 0 {
		return nil, types2.AccountNameError
	}

	a.SetOpts(opts)
	signedTx, err = handler.GetAuthority().AddAccountByOperator(opts, common.HexToAddress(account), accName, accDID, leaderDID)
	if err != nil {
		log.Error.Printf("failed to execute AddAccountByOperator: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// AddBatchAccountByOperator
// @Description: 运营方可以通过调用该方法直接对平台方或平台方的终端用户进行批量创建
// @receiver a
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param accounts DDC链账户地址集合
// @param accNames DDC账户对应的账户名称集合
// @param accDIDs DDC账户对应的DID信息集合
// @param leaderDIDs 该账户对应的上级账户的DID集合
// @return signedTx 签名好的交易
// @return error
func (a *AuthorityService) AddBatchAccountByOperator(opts *bind.TransactOpts, accounts []common.Address, accNames, accDIDs, leaderDIDs []string) (signedTx *types.Transaction, err error) {

	if len(accounts)^len(accNames)^len(accNames)^len(leaderDIDs) != 0 {
		return nil, types2.BatchInfoNumError
	}

	a.SetOpts(opts)
	signedTx, err = handler.GetAuthority().AddBatchAccountByOperator(opts, accounts, accNames, accDIDs, leaderDIDs)
	if err != nil {
		log.Error.Printf("failed to execute AddAccountByOperator: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// GetAccount
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行DDC账户信息的查询
// @receiver a
// @param account DDC用户的账户地址
// @return *dto.AccountInfo DDC账户信息实体
// @return error
func (a *AuthorityService) GetAccount(account string) (*dto.AccountInfo, error) {

	if !common.IsHexAddress(account) {
		return nil, types2.AccountError
	}

	accDID, accName, accRole, leaderDID, platformState, operatorState, field, err := handler.GetAuthority().GetAccount(nil, common.HexToAddress(account))
	if err != nil {
		log.Error.Printf("failed to execute GetAccount: %v", err.Error())
		return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return dto.NewAccountInfo(accDID, accName, accRole, leaderDID, platformState, operatorState, field), nil
}

// UpdateAccState
// @Description: 运营方或平台方可以通过调用该方法对终端用户进行DDC账户信息状态的更改
// @receiver a
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param account 要更新的账户地址
// @param state 状态 ：Frozen 0 - 冻结状态 ； Active 1 - 活跃状态
// @param changePlatformState 修改平台方状态标识
// @return signedTx 签名好的交易
// @return error
func (a *AuthorityService) UpdateAccState(opts *bind.TransactOpts, account string, state uint8, changePlatformState bool) (signedTx *types.Transaction, err error) {

	if !common.IsHexAddress(account) {
		return nil, types2.AccountError
	}

	a.SetOpts(opts)
	signedTx, err = handler.GetAuthority().UpdateAccountState(opts, common.HexToAddress(account), state, changePlatformState)
	if err != nil {
		log.Error.Printf("failed to execute UpdateAccountState: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// CrossPlatformApproval
// @Description: 运营方可以通过调用该方法对DDC的跨平台操作进行授权
// @receiver a
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 授权者账户
// @param to 接收者地址账户
// @param approved 授权标识
// @return signedTx 签名好的交易
// @return error
func (a *AuthorityService) CrossPlatformApproval(opts *bind.TransactOpts, from, to string, approved bool) (signedTx *types.Transaction, err error) {

	if !common.IsHexAddress(from) {
		return nil, types2.FromAccountError
	}
	if !common.IsHexAddress(to) {
		return nil, types2.ToAccountError
	}

	a.SetOpts(opts)
	signedTx, err = handler.GetAuthority().CrossPlatformApproval(opts, common.HexToAddress(from), common.HexToAddress(to), approved)
	if err != nil {
		log.Error.Printf("failed to execute CrossPlatformApproval: %v", err)
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// AddFunction
// @Description:运营方调用该方法为平台方和终端用户设置调用对应方法的权限
// @receiver a
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param role 目标角色
// @param ctrAddr 目标合约地址
// @param sig 目标方法对应的sig编码
// @return signedTx 签名好的交易
// @return error
func (a *AuthorityService) AddFunction(opts *bind.TransactOpts, role uint8, ctrAddr string, sig [4]byte) (signedTx *types.Transaction, err error) {

	a.SetOpts(opts)
	signedTx, err = handler.GetAuthority().AddFunction(opts, role, common.HexToAddress(ctrAddr), sig)
	if err != nil {
		log.Error.Printf("failed to execute UpdateAccState: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}
