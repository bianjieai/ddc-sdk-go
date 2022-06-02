# DDC-SDK-GoLang

- [DDC-SDK-GoLang](#ddc-sdk-go)
  - [平台方可调用的如下方法：](#平台方可调用的如下方法)
    - [1.初始化Client (连接测试网)](#1初始化client-连接测试网-1)
    - [2.BSN-DDC-权限管理](#2bsn-ddc-权限管理-1)
    - [3.BSN-DDC-费用管理](#3bsn-ddc-费用管理-1)
    - [4.BSN-DDC-721](#4bsn-ddc-721-1)
    - [5.BSN-DDC-1155](#5bsn-ddc-1155-1)
    - [6.BSN-DDC-交易查询](#6bsn-ddc-交易查询-1)
    - [7.BSN-DDC-区块查询](#7bsn-ddc-区块查询-1)
    - [8.BSN-DDC-数据解析](#8bsn-ddc-数据解析-1)
    - [9.离线账户创建](#9离线账户创建-1)
  - [终端用户可调用的如下方法：](#终端用户可调用的如下方法)
    - [1.初始化Client (连接测试网)](#1初始化client-连接测试网-2)
    - [2.BSN-DDC-权限管理](#2bsn-ddc-权限管理-2)
    - [3.BSN-DDC-费用管理](#3bsn-ddc-费用管理-2)
    - [4.BSN-DDC-721](#4bsn-ddc-721-2)
    - [5.BSN-DDC-1155](#5bsn-ddc-1155-2)
    - [6.BSN-DDC-交易查询](#6bsn-ddc-交易查询-2)
    - [7.BSN-DDC-区块查询](#7bsn-ddc-区块查询-2)
    - [8.BSN-DDC-数据解析](#8bsn-ddc-数据解析-2)
    - [9.离线账户创建](#9离线账户创建-2)
  - [测试用例](#测试用例)
  
## 平台方可调用的如下方法：

### 1.初始化Client (连接测试网)

初始化client时，可设置：签名方法、GasPrice、GatewayUrl以及相关的合约地址等，具体可设置项见以下示例：
```go
    //初始化clientBuilder
    clientBuilder = app.DDCSdkClientBuilder{}
    //可设置项如下：
    client        = clientBuilder.
        SetSignEventListener(new(SignListener)).//注册实现了签名接口的结构体
        SetGasPrice(1e10).//建议设置gasPrice>=1e10
        SetGatewayUrl("https://opbtest.bsngate.com:18602/api/0e346e1fb134477cafb6c6c2583ce3c4/evmrpc").
        SetAuthorityAddress("0xa7FC5B0F4A0085c5Ce689b919a866675Ce37B66b").
        SetChargeAddress("0x3BBb01B38958d4dbF1e004611EbB3c65979B0511").
        SetDDC721Address("0x3B09b7A00271C5d9AE84593850dE3A526b8BF96e").
        SetDDC1155Address("0xe5d3b9E7D16E03A4A1060c72b5D1cb7806DD9070").
        RegisterLog("./log.log").//设置日志输出的文件路径
        Build()
```
其中实现签名接口的示例如下：
```go
type SignListener struct {
}

// SignEvent
// @Description: 自定义的签名方法
// @receiver s
// @param sender 调用者的账户地址
// @param tx 待签名的交易
// @return *types.Transaction 签名好的交易
// @return error
func (s *SignListener) SignEvent(sender common.Address, tx *types.Transaction) (*types.Transaction, error) {

	// 提取私钥
	privateKey, err := StringToPrivateKey("sender的私钥")
	if err != nil {
		log.Fatalf("StringToPrivateKey failed:%v", err)
	}
	// 签名
	signTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)

	return signTx, err
}

// StringToPrivateKey
// @Description: 从明文的私钥字符串转换成ECDSA格式的私钥
// @param privateKeyStr 明文的私钥字符串
// @return *ecdsa.PrivateKey
// @return error
func StringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
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
```

### 2.BSN-DDC-权限管理

在调用权限管理的方法前需要先获取对应的service：
```go
var authorityService = client.GetAuthorityService()
```
接下来可通过service调用的方法如下：
```go
// GetAccount
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行DDC账户信息的查询
// @receiver a
// @param account DDC用户的账户地址
// @return *dto.AccountInfo DDC账户信息实体
// @return error
GetAccount(account string) (*dto.AccountInfo, error)

// UpdateAccState
// @Description: 运营方或平台方可以通过调用该方法对终端用户进行DDC账户信息状态的更改
// @receiver a
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param account 要更新的账户地址
// @param state 状态 ：Frozen 0 - 冻结状态 ； Active 1 - 活跃状态
// @param changePlatformState 修改平台方状态标识
// @return signedTx 签名好的交易
// @return error
UpdateAccState(opts *bind.TransactOpts, account string, state uint8, changePlatformState bool) (signedTx *types.Transaction, err error)
```

### 3.BSN-DDC-费用管理

在调用费用管理的方法前需要先获取对应的service：
```go
var chargeService = client.GetChargeService()
```
接下来可通过service调用的方法如下：
```go
// Recharge
// @Description:运营方、平台方调用该接口为所属同一方的同一级别账户或者下级账户充值
// @receiver c
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 接收者账户
// @param amount 充值金额
// @return signedTx 签名好的交易
// @return error
Recharge(opts *bind.TransactOpts, to string, amount int64) (signedTx *types.Transaction, err error)

// RechargeBatch
// @Description: 运营方或平台方可以通过调用该方法为所属同一方的同一级别账户或者下级账户进行批量充值
// @receiver c
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 被充值者的账户列表
// @param amount 每个账户充值的金额，与被充值者一一对应
// @return signedTx 签名好的交易
// @return err
RechargeBatch(opts *bind.TransactOpts, to []string, amount []int64)

// BalanceOf
// @Description: 查询指定账户的余额
// @receiver c
// @param accAddr 要查询的账户地址
// @return uint64 账户所对应的业务费余额
// @return error
BalanceOf(accAddr string) (uint64, error)

// BalanceOfBatch
// @Description: 运营方、平台方或终端用户可以调用方法批量查询账户的余额
// @receiver c
// @param accAddr 要查询的账户地址
// @return uint64 账户所对应的业务费余额
// @return error
BalanceOfBatch(accAddrs []string)

// QueryFee
// @Description: 查询指定的DDC业务方法所对应的业务费用
// @receiver c
// @param contrAddr DDC业务合约地址
// @param sig DDC业务方法对应的sig编码
// @return uint32 业务费用
// @return error
QueryFee(contrAddr string, sig [4]byte) (uint32, error)
```

### 4.BSN-DDC-721


在调用ddc721的方法前需要先获取对应的service：
```go
var ddc721Service = client.GetDDC721Service()
```
接下来可通过service调用的方法如下：
```go

// Mint
// @Description: 平台方或终端用户可以通过调用该方法进行DDC的生成
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 接收者账户
// @param ddcURI DDC资源标识符
// @return signedTx 签名好的交易
// @return error
Mint(opts *bind.TransactOpts, to, ddcURI string) (signedTx *types.Transaction, err error)

// SafeMint
// @Description: 平台方或终端用户可以通过调用该方法进行DDC的安全生成
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 接收者账户
// @param ddcURI DDC资源标识符
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
SafeMint(opts *bind.TransactOpts, to, ddcURI string, data []byte) (signedTx *types.Transaction, err error)

// Approve
// @Description: DDC拥有者可以通过调用该方法进行DDC的授权，发起者需要是DDC的拥有者
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 授权者账户
// @param ddcId DDC唯一标识
// @return signedTx 签名好的交易
// @return error
Approve(opts *bind.TransactOpts, to string, ddcId int64) (signedTx *types.Transaction, err error)

// GetApprove
// @Description: 运营方、平台方或终端用户都可以通过调用该方法查询DDC的授权情况
// @receiver d
// @param ddcId DDC唯一标识
// @return string 授权的账户
// @return error
GetApprove(ddcId int64) (string, error) 

// SetApprovalForAll
// @Description: DDC拥有者可以通过调用该方法进行账户授权，发起者需要是DDC的拥有者
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param operator 授权者账户
// @param approved 授权标识
// @return signedTx 签名好的交易
// @return error
SetApprovalForAll(opts *bind.TransactOpts, operator string, approved bool) (signedTx *types.Transaction, err error)

// IsApprovedForAll
// @Description: 运营方、平台方或终端用户可以通过调用该方法进行账户授权查询
// @receiver d
// @param owner 拥有者账户
// @param operator 授权者账户
// @return bool 授权标识
// @return error
IsApprovedForAll(owner, operator string) (bool, error)

// SafeTransferFrom
// @Description: DDC的拥有者或授权者可以通过调用该方法进行DDC的转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcId DDC唯一标识
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
SafeTransferFrom(opts *bind.TransactOpts, from, to string, ddcId int64, data []byte) (signedTx *types.Transaction, err error)

// TransferFrom
// @Description: DDC拥有者或授权者可以通过调用该方法进行DDC的转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcId ddc唯一标识
// @return signedTx 签名好的交易
// @return error
TransferFrom(opts *bind.TransactOpts, from, to string, ddcId int64) (signedTx *types.Transaction, err error) 

// Burn
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的销毁
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param ddcId DDC唯一标识
// @return signedTx 签名好的交易
// @return error
Burn(opts *bind.TransactOpts, ddcId int64) (signedTx *types.Transaction, err error)

// BalanceOf
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行查询当前账户拥有的DDC的数量
// @receiver d
// @param owner 拥有者账户
// @return uint64 ddc的数量
// @return error
BalanceOf(owner string) (uint64, error)

// OwnerOf
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的拥有者
// @receiver d
// @param ddcId ddc唯一标识
// @return string 拥有者账户
// @return error
OwnerOf(ddcId int64) (string, error)

// Name
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的名称
// @receiver d
// @return string DDC运营方名称
// @return error
Name() (string, error)

// Symbol
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的符号标识
// @receiver d
// @return string DDC运营方符号
// @return error
Symbol() (string, error)

// DdcURI
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的资源标识符
// @receiver d
// @param ddcId DDC唯一标识符
// @return string DDC资源标识符
// @return error
DdcURI(ddcId int64) (string, error)

// SetURI
// @Description: DDC拥有者或DDC授权者通过调用该方法对DDC的资源标识符进行设置
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param ddcId DDC唯一标识
// @param ddcURI DDC资源标识符
// @return signedTx 签名好的交易
// @return err
SetURI(opts *bind.TransactOpts, ddcId int64, ddcURI string) (signedTx *types.Transaction, err error)
```

### 5.BSN-DDC-1155

在调用ddc1155的方法前需要先获取对应的service：
```go
var ddc1155Service = client.GetDDC1155Service()
```
接下来可通过service调用的方法如下：
```go
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
SafeMint(opts *bind.TransactOpts, to string, amount int64, ddcURI string, data []byte) (signedTx *types.Transaction, err error)

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
SafeMintBatch(opts *bind.TransactOpts, to string, amounts []*big.Int, ddcURIs []string, data []byte) (signedTx *types.Transaction, err error) 

// SetApprovalForAll
// @Description: DDC拥有者可以通过调用该方法进行账户授权，发起者需要是DDC的拥有者
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param operator 授权者账户
// @param approved 授权标识
// @return signedTx 签名好的交易
// @return error
SetApprovalForAll(opts *bind.TransactOpts, operator string, approved bool) (signedTx *types.Transaction, err error)

// IsApprovedForAll
// @Description: 运营方、平台方或终端用户可以通过调用该方法进行账户授权查询
// @receiver d
// @param owner 拥有者账户
// @param operator 授权者账户
// @return bool 授权标识
// @return error
IsApprovedForAll(owner, operator string) (bool, error)

// SafeTransferFrom
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcId DDC唯一标识
// @param amount 需要转移的DDC数量
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
SafeTransferFrom(opts *bind.TransactOpts, from, to string, ddcId, amount int64, data []byte) (signedTx *types.Transaction, err error)

// SafeBatchTransferFrom
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的批量转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcInfo 要转移的ddc集合（ddcId->amount）
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
SafeBatchTransferFrom(opts *bind.TransactOpts, from, to string, ddcInfo map[int64]int64, data []byte) (signedTx *types.Transaction, err error)

// Burn
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的销毁
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param ddcId DDC唯一标识
// @return signedTx 签名好的交易
// @return error
Burn(opts *bind.TransactOpts, owner string, ddcId int64) (signedTx *types.Transaction, err error)

// BurnBatch
// @Description: DDC拥用者可以通过调用该方法进行DDC的批量销毁
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param owner 拥有者账户
// @param ddcIds 要销毁的DDC的唯一标识的集合
// @return signedTx 签名好的交易
// @return error
BurnBatch(opts *bind.TransactOpts, owner string, ddcIds []int64) (signedTx *types.Transaction, err error)

// BalanceOf
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行查询当前账户拥有的对应ddcId的数量
// @receiver d
// @param owner 拥有者账户
// @param ddcId DDC唯一标识
// @return uint64 ddcId对应的数量
// @return error
BalanceOf(owner string, ddcId int64) (uint64, error)

// BalanceOfBatch
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行批量查询账户拥有的DDC的数量
// @receiver d
// @param owners 要查询的账户地址
// @param ddcIds 要查询的ddcId
// @return []*big.Int ddc数量的集合
// @return error
BalanceOfBatch(owners []common.Address, ddcIds []*big.Int) ([]*big.Int, error)

// DdcURI
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的资源标识符
// @receiver d
// @param ddcId DDC唯一标识符
// @return string DDC资源标识符
// @return error
DdcURI(ddcId int64) (ddcURI string, err error)

// SetURI
// @Description: DDC拥有者或DDC授权者通过调用该方法对DDC的资源标识符进行设置
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param owner DDC拥有者
// @param ddcId DDC唯一标识
// @param ddcURI DDC资源标识符
// @return signedTx 签名好的交易
// @return error
SetURI(opts *bind.TransactOpts, owner string, ddcId int64, ddcURI string) (signedTx *types.Transaction, err error)
```

### 6.BSN-DDC-交易查询

使用client或任意service即可调用以下方法：
```go
// GetTransByHash
// @Description: 运营方或平台方根据交易哈希对交易信息进行查询
// @receiver t
// @param txHash: 交易哈希
// @return *types.Transaction: 交易信息
// @return bool： pending状态
// @return error
GetTransByHash(txHash string) (*types.Transaction, bool, error)

// GetTransReceipt
// @Description: 运营方或平台方根据交易哈希对交易回执信息进行查询。
// @receiver t
// @param txHash: 交易哈希
// @return string： 交易回执
// @return error
GetTransReceipt(txHash string) (*types.Receipt, error)

// GetTransStatus
// @Description: 运营方或平台方根据交易哈希查询交易状态是否成功
// @receiver t
// @param txHash: 交易哈希
// @return bool：交易是否成功
// @return error
GetTransStatus(txHash string) (bool, error)

// GetTimeByTxHash
// @Description: 通过txHash查询所在块的时间
// @receiver t
// @param txHash 交易哈希
// @return uint64 所在块的时间
// @return error
GetTimeByTxHash(txHash string) (uint64, error)
```

### 7.BSN-DDC-区块查询

使用client或任意service即可调用以下方法：
```go
// GetBlockByNumber
// @Description: 运营方或平台方根据区块高度对区块信息进行查询，并解析区块数据返回给运营方或平台方
// @receiver b
// @param blockNumber： 区块高度
// @return *types2.Block： 区块信息
// @return error
GetBlockByNumber(blockNumber int64) (*types.Block, error)
```

### 8.BSN-DDC-数据解析

使用client或任意service即可调用以下方法：
```go
// GetTxEvents
// @Description: 获取指定交易中的所有events
// @receiver b
// @param txHash 交易哈希
// @return events 查询出的事件
// @return err
GetTxEvents(txHash common.Hash) (events []interface{}, err error)

// GetBlockEvents
// @Description: 获取指定区块内的所有events
// @receiver b
// @param blockNumber 块高
// @return *dto.BlockEventBean 事件和时间戳的实体
// @return error
GetBlockEvents(blockNumber int64) (*dto.BlockEventBean, error)
```
### 9.离线账户创建

使用client或任意service即可调用以下方法：
```go
// CreateAccount
// @Description:可以通过此方法生成离线账户
// @receiver b
// @return *dto.Account 返回的账户信息
// @return error
CreateAccount() (*dto.Account, error)
```

## 终端用户可调用的如下方法：

### 1.初始化Client (连接测试网)

初始化client时，可设置：签名方法、GasPrice、GatewayUrl以及相关的合约地址等，具体可设置项见以下示例：
```go
    //初始化clientBuilder
    clientBuilder = app.DDCSdkClientBuilder{}
    //可设置项如下：
    client        = clientBuilder.
        SetSignEventListener(new(SignListener)).//注册实现了签名接口的结构体
        SetGasPrice(1e10).//建议设置gasPrice>=1e10
        SetGatewayUrl("https://opbtest.bsngate.com:18602/api/0e346e1fb134477cafb6c6c2583ce3c4/evmrpc").
        SetAuthorityAddress("0xa7FC5B0F4A0085c5Ce689b919a866675Ce37B66b").
        SetChargeAddress("0x3BBb01B38958d4dbF1e004611EbB3c65979B0511").
        SetDDC721Address("0x3B09b7A00271C5d9AE84593850dE3A526b8BF96e").
        SetDDC1155Address("0xe5d3b9E7D16E03A4A1060c72b5D1cb7806DD9070").
        RegisterLog("./log.log").//设置日志输出的文件路径
        Build()
```
其中实现签名接口的示例如下：
```go
type SignListener struct {
}

// SignEvent
// @Description: 自定义的签名方法
// @receiver s
// @param sender 调用者的账户地址
// @param tx 待签名的交易
// @return *types.Transaction 签名好的交易
// @return error
func (s *SignListener) SignEvent(sender common.Address, tx *types.Transaction) (*types.Transaction, error) {

	// 提取私钥
	privateKey, err := StringToPrivateKey("sender的私钥")
	if err != nil {
		log.Fatalf("StringToPrivateKey failed:%v", err)
	}
	// 签名
	signTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)

	return signTx, err
}

// StringToPrivateKey
// @Description: 从明文的私钥字符串转换成ECDSA格式的私钥
// @param privateKeyStr 明文的私钥字符串
// @return *ecdsa.PrivateKey
// @return error
func StringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
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
```

### 2.BSN-DDC-权限管理

在调用权限管理的方法前需要先获取对应的service：
```go
var authorityService = client.GetAuthorityService()
```
接下来可通过service调用的方法如下：
```go
// GetAccount
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行DDC账户信息的查询
// @receiver a
// @param account DDC用户的账户地址
// @return *dto.AccountInfo DDC账户信息实体
// @return error
GetAccount(account string) (*dto.AccountInfo, error)
```

### 3.BSN-DDC-费用管理

在调用费用管理的方法前需要先获取对应的service：
```go
var chargeService = client.GetChargeService()
```
接下来可通过service调用的方法如下：
```go
// BalanceOf
// @Description: 查询指定账户的余额
// @receiver c
// @param accAddr 要查询的账户地址
// @return uint64 账户所对应的业务费余额
// @return error
BalanceOf(accAddr string) (uint64, error)

// BalanceOfBatch
// @Description: 运营方、平台方或终端用户可以调用方法批量查询账户的余额
// @receiver c
// @param accAddr 要查询的账户地址
// @return uint64 账户所对应的业务费余额
// @return error
BalanceOfBatch(accAddrs []string)

// QueryFee
// @Description: 查询指定的DDC业务方法所对应的业务费用
// @receiver c
// @param contrAddr DDC业务合约地址
// @param sig DDC业务方法对应的sig编码
// @return uint32 业务费用
// @return error
QueryFee(contrAddr string, sig [4]byte) (uint32, error)
```

### 4.BSN-DDC-721

在调用ddc721的方法前需要先获取对应的service：
```go
var ddc721Service = client.GetDDC721Service()
```
接下来可通过service调用的方法如下：
```go

// Mint
// @Description: 平台方或终端用户可以通过调用该方法进行DDC的生成
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 接收者账户
// @param ddcURI DDC资源标识符
// @return signedTx 签名好的交易
// @return error
Mint(opts *bind.TransactOpts, to, ddcURI string) (signedTx *types.Transaction, err error)

// SafeMint
// @Description: 平台方或终端用户可以通过调用该方法进行DDC的安全生成
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 接收者账户
// @param ddcURI DDC资源标识符
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
SafeMint(opts *bind.TransactOpts, to, ddcURI string, data []byte) (signedTx *types.Transaction, err error)

// Approve
// @Description: DDC拥有者可以通过调用该方法进行DDC的授权，发起者需要是DDC的拥有者
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param to 授权者账户
// @param ddcId DDC唯一标识
// @return signedTx 签名好的交易
// @return error
Approve(opts *bind.TransactOpts, to string, ddcId int64) (signedTx *types.Transaction, err error)

// GetApprove
// @Description: 运营方、平台方或终端用户都可以通过调用该方法查询DDC的授权情况
// @receiver d
// @param ddcId DDC唯一标识
// @return string 授权的账户
// @return error
GetApprove(ddcId int64) (string, error) 

// SetApprovalForAll
// @Description: DDC拥有者可以通过调用该方法进行账户授权，发起者需要是DDC的拥有者
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param operator 授权者账户
// @param approved 授权标识
// @return signedTx 签名好的交易
// @return error
SetApprovalForAll(opts *bind.TransactOpts, operator string, approved bool) (signedTx *types.Transaction, err error)

// IsApprovedForAll
// @Description: 运营方、平台方或终端用户可以通过调用该方法进行账户授权查询
// @receiver d
// @param owner 拥有者账户
// @param operator 授权者账户
// @return bool 授权标识
// @return error
IsApprovedForAll(owner, operator string) (bool, error)

// SafeTransferFrom
// @Description: DDC的拥有者或授权者可以通过调用该方法进行DDC的转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcId DDC唯一标识
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
SafeTransferFrom(opts *bind.TransactOpts, from, to string, ddcId int64, data []byte) (signedTx *types.Transaction, err error)

// TransferFrom
// @Description: DDC拥有者或授权者可以通过调用该方法进行DDC的转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcId ddc唯一标识
// @return signedTx 签名好的交易
// @return error
TransferFrom(opts *bind.TransactOpts, from, to string, ddcId int64) (signedTx *types.Transaction, err error) 

// Burn
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的销毁
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param ddcId DDC唯一标识
// @return signedTx 签名好的交易
// @return error
Burn(opts *bind.TransactOpts, ddcId int64) (signedTx *types.Transaction, err error)

// BalanceOf
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行查询当前账户拥有的DDC的数量
// @receiver d
// @param owner 拥有者账户
// @return uint64 ddc的数量
// @return error
BalanceOf(owner string) (uint64, error)

// OwnerOf
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的拥有者
// @receiver d
// @param ddcId ddc唯一标识
// @return string 拥有者账户
// @return error
OwnerOf(ddcId int64) (string, error)

// Name
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的名称
// @receiver d
// @return string DDC运营方名称
// @return error
Name() (string, error)

// Symbol
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的符号标识
// @receiver d
// @return string DDC运营方符号
// @return error
Symbol() (string, error)

// DdcURI
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的资源标识符
// @receiver d
// @param ddcId DDC唯一标识符
// @return string DDC资源标识符
// @return error
DdcURI(ddcId int64) (string, error)

// SetURI
// @Description: DDC拥有者或DDC授权者通过调用该方法对DDC的资源标识符进行设置
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param ddcId DDC唯一标识
// @param ddcURI DDC资源标识符
// @return signedTx 签名好的交易
// @return err
SetURI(opts *bind.TransactOpts, ddcId int64, ddcURI string) (signedTx *types.Transaction, err error)
```

### 5.BSN-DDC-1155

在调用ddc1155的方法前需要先获取对应的service：
```go
var ddc1155Service = client.GetDDC1155Service()
```
接下来可通过service调用的方法如下：
```go
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
SafeMint(opts *bind.TransactOpts, to string, amount int64, ddcURI string, data []byte) (signedTx *types.Transaction, err error)

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
SafeMintBatch(opts *bind.TransactOpts, to string, amounts []*big.Int, ddcURIs []string, data []byte) (signedTx *types.Transaction, err error) 

// SetApprovalForAll
// @Description: DDC拥有者可以通过调用该方法进行账户授权，发起者需要是DDC的拥有者
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param operator 授权者账户
// @param approved 授权标识
// @return signedTx 签名好的交易
// @return error
SetApprovalForAll(opts *bind.TransactOpts, operator string, approved bool) (signedTx *types.Transaction, err error)

// IsApprovedForAll
// @Description: 运营方、平台方或终端用户可以通过调用该方法进行账户授权查询
// @receiver d
// @param owner 拥有者账户
// @param operator 授权者账户
// @return bool 授权标识
// @return error
IsApprovedForAll(owner, operator string) (bool, error)

// SafeTransferFrom
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcId DDC唯一标识
// @param amount 需要转移的DDC数量
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
SafeTransferFrom(opts *bind.TransactOpts, from, to string, ddcId, amount int64, data []byte) (signedTx *types.Transaction, err error)

// SafeBatchTransferFrom
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的批量转移
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param from 拥有者账户
// @param to 接收者账户
// @param ddcInfo 要转移的ddc集合（ddcId->amount）
// @param data 附加数据
// @return signedTx 签名好的交易
// @return error
SafeBatchTransferFrom(opts *bind.TransactOpts, from, to string, ddcInfo map[int64]int64, data []byte) (signedTx *types.Transaction, err error)

// Burn
// @Description: DDC拥有者或DDC授权者可以通过调用该方法进行DDC的销毁
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param ddcId DDC唯一标识
// @return signedTx 签名好的交易
// @return error
Burn(opts *bind.TransactOpts, owner string, ddcId int64) (signedTx *types.Transaction, err error)

// BurnBatch
// @Description: DDC拥用者可以通过调用该方法进行DDC的批量销毁
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param owner 拥有者账户
// @param ddcIds 要销毁的DDC的唯一标识的集合
// @return signedTx 签名好的交易
// @return error
BurnBatch(opts *bind.TransactOpts, owner string, ddcIds []int64) (signedTx *types.Transaction, err error)

// BalanceOf
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行查询当前账户拥有的对应ddcId的数量
// @receiver d
// @param owner 拥有者账户
// @param ddcId DDC唯一标识
// @return uint64 ddcId对应的数量
// @return error
BalanceOf(owner string, ddcId int64) (uint64, error)

// BalanceOfBatch
// @Description: 运营方、平台方以及终端用户可以通过调用该方法进行批量查询账户拥有的DDC的数量
// @receiver d
// @param owners 要查询的账户地址
// @param ddcIds 要查询的ddcId
// @return []*big.Int ddc数量的集合
// @return error
BalanceOfBatch(owners []common.Address, ddcIds []*big.Int) ([]*big.Int, error)

// DdcURI
// @Description: 运营方、平台方以及终端用户可以通过调用该方法查询当前DDC的资源标识符
// @receiver d
// @param ddcId DDC唯一标识符
// @return string DDC资源标识符
// @return error
DdcURI(ddcId int64) (ddcURI string, err error)

// SetURI
// @Description: DDC拥有者或DDC授权者通过调用该方法对DDC的资源标识符进行设置
// @receiver d
// @param opts opts.Price和opts.Signer为空时，默认使用初始化client时set的price和signer
// @param owner DDC拥有者
// @param ddcId DDC唯一标识
// @param ddcURI DDC资源标识符
// @return signedTx 签名好的交易
// @return error
SetURI(opts *bind.TransactOpts, owner string, ddcId int64, ddcURI string) (signedTx *types.Transaction, err error)
```

### 6.BSN-DDC-交易查询

使用client或任意service即可调用以下方法：
```go
// GetTransByHash
// @Description: 运营方或平台方根据交易哈希对交易信息进行查询
// @receiver t
// @param txHash: 交易哈希
// @return *types.Transaction: 交易信息
// @return bool： pending状态
// @return error
GetTransByHash(txHash string) (*types.Transaction, bool, error)

// GetTransReceipt
// @Description: 运营方或平台方根据交易哈希对交易回执信息进行查询。
// @receiver t
// @param txHash: 交易哈希
// @return string： 交易回执
// @return error
GetTransReceipt(txHash string) (*types.Receipt, error)

// GetTransStatus
// @Description: 运营方或平台方根据交易哈希查询交易状态是否成功
// @receiver t
// @param txHash: 交易哈希
// @return bool：交易是否成功
// @return error
GetTransStatus(txHash string) (bool, error)

// GetTimeByTxHash
// @Description: 通过txHash查询所在块的时间
// @receiver t
// @param txHash 交易哈希
// @return uint64 所在块的时间
// @return error
GetTimeByTxHash(txHash string) (uint64, error)
```

### 7.BSN-DDC-区块查询

使用client或任意service即可调用以下方法：
```go
// GetBlockByNumber
// @Description: 运营方或平台方根据区块高度对区块信息进行查询，并解析区块数据返回给运营方或平台方
// @receiver b
// @param blockNumber： 区块高度
// @return *types2.Block： 区块信息
// @return error
GetBlockByNumber(blockNumber int64) (*types.Block, error)
```


### 8.BSN-DDC-数据解析

```go
// GetTxEvents
// @Description: 获取指定交易中的所有events
// @receiver b
// @param txHash 交易哈希
// @return events 查询出的事件
// @return err
GetTxEvents(txHash common.Hash) (events []interface{}, err error)

// GetBlockEvents
// @Description: 获取指定区块内的所有events
// @receiver b
// @param blockNumber 块高
// @return *dto.BlockEventBean 事件和时间戳的实体
// @return error
GetBlockEvents(blockNumber int64) (*dto.BlockEventBean, error)
```
### 9.离线账户创建

使用client或任意service即可调用以下方法：
```go
// CreateAccount
// @Description:可以通过此方法生成离线账户
// @receiver b
// @return *dto.Account 返回的账户信息
// @return error
CreateAccount() (*dto.Account, error)
```

## 测试用例

[AuthorityServiceTest.java](test/authority_test.go)

[ChargeServiceTest.java](test/charge_test.go)

[BaseServiceTest.java](test/base_test.go)

[BlockEventServiceTest.java](test/block_test.go)

[DDC721ServiceTest.java](test/ddc721_test.go)

[DDC1155ServiceTest.java](test/ddc1155_test.go)

[SignEventTest.java](test/sign_test.go)
