package types2

import "fmt"

const (
	DDC = "ddc-sdk-go"
)

var (
	AccountNameError     = register(DDC, "1026", "account name is empty")
	AccountError         = register(DDC, "1027", "account is not a standard address format")
	OwnerAccountError    = register(DDC, "1028", "owner account is not a standard address format")
	SenderAccountError   = register(DDC, "1029", "sender account is not a standard address format")
	OperatorAccountError = register(DDC, "1030", "owner account is not a standard address format")
	AccountListError     = register(DDC, "1031", "accounts must not be empty")
	AccNamesError        = register(DDC, "1032", "accNames must not be empty")
	AccDIDsError         = register(DDC, "1033", "accDIDs must not be empty")
	LeaderDIDsError      = register(DDC, "1034", "LeaderDIDs must not be empty")
	ToListError     = register(DDC, "1035", "toList must not be empty")
	AmountListError = register(DDC, "1036", "amounts must not be empty")
	NameError       = register(DDC, "1037", "name must not be null")
	SymbolError          = register(DDC, "1038", "symbol must not be null")

	FromAccountError     = register(DDC, "1007", "from account is not a standard address format")
	ToAccountError       = register(DDC, "1008", "to account is not a standard address format")
	ContractAddressError = register(DDC, "1009", "contract address is not a standard address format")

	DDCInfoError      = register(DDC, "1010", "the length of amounts and ddcURIs do not match")
	OwnersDDCIdsError = register(DDC, "1011", "the length of owners and ddcIds do not match")
	BatchInfoNumError = register(DDC, "1012", "the number of the input params is inconsistent")

	DDCIdError  = register(DDC, "1016", "ddcId is wrong")
	DDCURIError = register(DDC, "1017", "ddcURI cannot be empty")
	AmountError = register(DDC, "1018", "Amount is wrong")
	DDCIdsError = register(DDC, "1019", "ddcIds must not be empty")

	TransactError = register(DDC, "3001", "failed to transact")
	QueryError    = register(DDC, "3002", "failed to query")
	InternalError = register(DDC, "3003", "internal error")
)

type AppError struct {
	codeSpace string
	code      string
	desc      string
}

func NewAppError(codeSpace string, code string, desc string) *AppError {
	return &AppError{codeSpace: codeSpace, code: code, desc: desc}
}
func (a *AppError) Error() string {
	return a.desc
}

// 记录已注册的error
var appErrors = map[string]*AppError{}

func register(codeSpace, code, description string) *AppError {
	if e := getAppError(codeSpace, code); e != nil {
		panic(fmt.Sprintf("error with code %s is already registered: %q", e.code, e.desc))
	}
	err := NewAppError(codeSpace, code, description)
	appErrors[errorKey(codeSpace, code)] = err
	return err
}

func getAppError(codeSpace, code string) *AppError {
	return appErrors[errorKey(codeSpace, code)]
}

//组装errorKey
func errorKey(codeSpace, code string) string {
	return fmt.Sprintf("%s:%s", codeSpace, code)
}

type SDKError struct {
	codeSpace string
	code      string
	desc      string
}

func NewSDKError(code string, desc string) *SDKError {
	return &SDKError{codeSpace: DDC, code: code, desc: desc}
}
func (a *SDKError) Error() string {
	return "[" + a.codeSpace + "] " + a.code + " : " + a.desc
}
