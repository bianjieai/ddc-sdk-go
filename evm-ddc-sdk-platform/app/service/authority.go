package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/handler"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/models/dto"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/log"
	types2 "github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/types"
)

type AuthorityService struct {
	Base
}

func NewAuthorityService() *AuthorityService {
	return &AuthorityService{}
}

// GetAccount
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行DDC账户信息的查询
// @receiver a
// @param account DDC用户的账户地址
// @return *dto.AccountInfo DDC账户信息实体
// @return error
func (a *AuthorityService) GetAccount(account string) (*dto.AccountInfo, error) {

	if common.HexToAddress(account) == common.HexToAddress("0") || !common.IsHexAddress(account) {
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

	if common.HexToAddress(account) == common.HexToAddress("0") || !common.IsHexAddress(account) {
		return nil, types2.AccountError
	}

	a.setOpts(opts)
	signedTx, err = handler.GetAuthority().UpdateAccountState(opts, common.HexToAddress(account), state, changePlatformState)
	if err != nil {
		log.Error.Printf("failed to execute UpdateAccountState: %v", err.Error())
		return nil, types2.NewSDKError(types2.TransactError.Error(), err.Error())
	}

	return signedTx, nil
}

// HasFunctionPermission
// @Description: 运营方调用该方法查询平台方和终端用户是否有调用对应方法的权限
// @receiver a
// @param accAddr 账户地址
// @param ctrAddr 合约地址
// @param sig 目标方法对应的sig编码
// @return bool 是否已被授权
// @return error
func (a *AuthorityService) HasFunctionPermission(accAddr, ctrAddr string, sig [4]byte) (bool, error) {

	IsPermission, err := handler.GetAuthority().HasFunctionPermission(nil, common.HexToAddress(accAddr), common.HexToAddress(ctrAddr), sig)
	if err != nil {
		log.Error.Printf("failed to execute HasFunctionPermission: %v", err.Error())
		return false, err
	}

	return IsPermission, nil
}
