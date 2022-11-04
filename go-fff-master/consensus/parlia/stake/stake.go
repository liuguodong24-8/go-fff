// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake

import (
	ethereum "github.com/fff-chain/go-fff"
	"github.com/fff-chain/go-fff/accounts/abi"
	"github.com/fff-chain/go-fff/accounts/abi/bind"
	"github.com/fff-chain/go-fff/common"
	"github.com/fff-chain/go-fff/core/types"
	"github.com/fff-chain/go-fff/event"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BSCValidatorSetStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type BSCValidatorSetStakeInfo struct {
	StakeAddress common.Address
	StakeCount   *big.Int
}

// BSCValidatorSetABI is the input ABI used to generate the binding from.
const BSCValidatorSetABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"DistributeBlockRewardFunc\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UserStakeFFF\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UserUnStakeFFF\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"batchTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"batchTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"batchTransferLowerFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deprecatedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"addresspayable\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"directTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"addresspayable\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"directTransferFail\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"failReasonWithStr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"feeBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"systemTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"unexpectedPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorEmptyJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorFelony\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"validatorMisdemeanor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"validatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_RATIO_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DUSTY_INCOMING\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_CHECK_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_LEN_OF_VAL_MISMATCH\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_RELAYFEE_TOO_LARGE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_UNKNOWN_PACKAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXPIRE_TIME_SECOND_GAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAllStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"stakeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeCount\",\"type\":\"uint256\"}],\"internalType\":\"structBSCValidatorSet.StakeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_BURN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_VALIDATORSET_BYTES\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"JAIL_MESSAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_OF_VALIDATORS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STAKE_FFF_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StakeFFF\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MANAGER_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UnStakeFFF\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORS_UPDATE_MESSAGE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnRatioInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currStakeFFF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currStakePeople\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currStakePeopleIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BBCFeeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"votingPower\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"incoming\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentValidatorSetMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expireTimeSecondGap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"felony\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getIncoming\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleFailAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"responsePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakeFFFCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakePeopleCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miniStakeFFFCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"misdemeanor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfJailed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeInfoIndexMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeInfoMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInComing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BSCValidatorSetFuncSigs maps the 4-byte function signature to its string representation.
var BSCValidatorSetFuncSigs = map[string]string{
	"3dffc387": "BIND_CHANNELID()",
	"fccc2813": "BURN_ADDRESS()",
	"3de0f0d8": "BURN_RATIO_SCALE()",
	"ab51bb96": "CODE_OK()",
	"51e80672": "CROSS_CHAIN_CONTRACT_ADDR()",
	"d86222d5": "DUSTY_INCOMING()",
	"81650b62": "ERROR_FAIL_CHECK_VALIDATORS()",
	"0bee7a67": "ERROR_FAIL_DECODE()",
	"5d77156c": "ERROR_LEN_OF_VAL_MISMATCH()",
	"219f22d5": "ERROR_RELAYFEE_TOO_LARGE()",
	"eda5868c": "ERROR_UNKNOWN_PACKAGE_TYPE()",
	"853230aa": "EXPIRE_TIME_SECOND_GAP()",
	"96713da9": "GOV_CHANNELID()",
	"9dc09262": "GOV_HUB_ADDR()",
	"54b1a671": "GetAllStakeInfo()",
	"6e47b482": "INCENTIVIZE_ADDR()",
	"78dfed4a": "INIT_BURN_RATIO()",
	"a5422d5c": "INIT_VALIDATORSET_BYTES()",
	"bf9f4995": "JAIL_MESSAGE_TYPE()",
	"dc927faf": "LIGHT_CLIENT_ADDR()",
	"e086c7b1": "MAX_NUM_OF_VALIDATORS()",
	"8a10febe": "MAX_STAKE_FFF_COUNT()",
	"aaf5eb68": "PRECISION()",
	"a1a11bf5": "RELAYERHUB_CONTRACT_ADDR()",
	"7942fd05": "SLASH_CHANNELID()",
	"43756e5c": "SLASH_CONTRACT_ADDR()",
	"4bf6c882": "STAKING_CHANNELID()",
	"c81b1662": "SYSTEM_REWARD_ADDR()",
	"0db6cccb": "StakeFFF()",
	"fd6a6879": "TOKEN_HUB_ADDR()",
	"75d47a0a": "TOKEN_MANAGER_ADDR()",
	"70fd5bad": "TRANSFER_IN_CHANNELID()",
	"fc3e5908": "TRANSFER_OUT_CHANNELID()",
	"30c62a13": "UnStakeFFF()",
	"5667515a": "VALIDATORS_UPDATE_MESSAGE_TYPE()",
	"f9a2bbc7": "VALIDATOR_CONTRACT_ADDR()",
	"a78abc16": "alreadyInit()",
	"493279b1": "bscChainID()",
	"5192c82c": "burnRatio()",
	"152ad3b8": "burnRatioInitialized()",
	"101a875f": "currStakeFFF()",
	"93a101f3": "currStakePeople()",
	"e3260b57": "currStakePeopleIndex()",
	"6969a25c": "currentValidatorSet(uint256)",
	"ad3c9da6": "currentValidatorSetMap(address)",
	"f340fa01": "deposit(address)",
	"86249882": "expireTimeSecondGap()",
	"35409f7f": "felony(address)",
	"565c56b3": "getIncoming(address)",
	"b7ab4db5": "getValidators()",
	"831d65d1": "handleAckPackage(uint8,bytes)",
	"c8509d81": "handleFailAckPackage(uint8,bytes)",
	"1182b875": "handleSynPackage(uint8,bytes)",
	"e1c7392a": "init()",
	"b2ccf3a4": "maxStakeFFFCount()",
	"f78ffcbc": "maxStakePeopleCount()",
	"1faea20b": "miniStakeFFFCount()",
	"eb57e202": "misdemeanor(address)",
	"c9116b69": "myBalance()",
	"daacdb66": "numOfJailed()",
	"8da5cb5b": "owner()",
	"7dedf367": "stakeInfoIndexMap(uint256)",
	"c5f30db5": "stakeInfoMap(address)",
	"1ff18069": "totalInComing()",
	"ac431751": "updateParam(string,bytes)",
}

// BSCValidatorSetBin is the compiled bytecode used for deploying new contracts.
var BSCValidatorSetBin = "0x608060405234801561001057600080fd5b506148c2806100206000396000f3fe6080604052600436106103ce5760003560e01c80638a10febe116101fd578063c81b166211610118578063e3260b57116100ab578063f78ffcbc1161007a578063f78ffcbc14610979578063f9a2bbc71461098e578063fc3e5908146109a3578063fccc2813146109b8578063fd6a6879146109cd576103ce565b8063e3260b571461091c578063eb57e20214610931578063eda5868c14610951578063f340fa0114610966576103ce565b8063daacdb66116100e7578063daacdb66146108c8578063dc927faf146108dd578063e086c7b1146108f2578063e1c7392a14610907576103ce565b8063c81b166214610889578063c8509d81146106d6578063c9116b691461089e578063d86222d5146108b3576103ce565b8063aaf5eb6811610190578063b2ccf3a41161015f578063b2ccf3a414610832578063b7ab4db514610847578063bf9f4995146104f5578063c5f30db514610869576103ce565b8063aaf5eb68146107c8578063ab51bb96146107dd578063ac431751146107f2578063ad3c9da614610812576103ce565b80639dc09262116101cc5780639dc0926214610774578063a1a11bf514610789578063a5422d5c1461079e578063a78abc16146107b3576103ce565b80638a10febe146107205780638da5cb5b1461073557806393a101f31461074a57806396713da91461075f576103ce565b806351e80672116102ed57806375d47a0a1161028057806381650b621161024f57806381650b62146106c1578063831d65d1146106d6578063853230aa146106f6578063862498821461070b576103ce565b806375d47a0a1461066257806378dfed4a146106775780637942fd051461068c5780637dedf367146106a1576103ce565b80635d77156c116102bc5780635d77156c146105f15780636969a25c146106065780636e47b4821461063857806370fd5bad1461064d576103ce565b806351e806721461058557806354b1a6711461059a578063565c56b3146105bc5780635667515a146105dc576103ce565b806330c62a131161036557806343756e5c1161033457806343756e5c14610517578063493279b1146105395780634bf6c8821461055b5780635192c82c14610570576103ce565b806330c62a13146104b857806335409f7f146104c05780633de0f0d8146104e05780633dffc387146104f5576103ce565b8063152ad3b8116103a1578063152ad3b8146104575780631faea20b146104795780631ff180691461048e578063219f22d5146104a3576103ce565b80630bee7a67146103d35780630db6cccb146103fe578063101a875f146104085780631182b8751461042a575b600080fd5b3480156103df57600080fd5b506103e86109e2565b6040516103f59190614789565b60405180910390f35b6104066109e7565b005b34801561041457600080fd5b5061041d6109f3565b6040516103f59190614780565b34801561043657600080fd5b5061044a610445366004613f2f565b6109f9565b6040516103f5919061419a565b34801561046357600080fd5b5061046c610b89565b6040516103f5919061418f565b34801561048557600080fd5b5061041d610b92565b34801561049a57600080fd5b5061041d610b9e565b3480156104af57600080fd5b506103e8610ba4565b610406610ba9565b3480156104cc57600080fd5b506104066104db366004613e50565b610c81565b3480156104ec57600080fd5b5061041d610f79565b34801561050157600080fd5b5061050a610f7f565b6040516103f5919061479a565b34801561052357600080fd5b5061052c610f84565b6040516103f5919061403b565b34801561054557600080fd5b5061054e610f8a565b6040516103f59190614771565b34801561056757600080fd5b5061050a610f8f565b34801561057c57600080fd5b5061041d610f94565b34801561059157600080fd5b5061052c610f9a565b3480156105a657600080fd5b506105af610fa0565b6040516103f59190614137565b3480156105c857600080fd5b5061041d6105d7366004613e50565b6110c8565b3480156105e857600080fd5b5061050a61111b565b3480156105fd57600080fd5b506103e8611120565b34801561061257600080fd5b50610626610621366004613eff565b611125565b6040516103f596959493929190614068565b34801561064457600080fd5b5061052c611189565b34801561065957600080fd5b5061050a61118f565b34801561066e57600080fd5b5061052c611194565b34801561068357600080fd5b5061041d61111b565b34801561069857600080fd5b5061050a61119a565b3480156106ad57600080fd5b5061052c6106bc366004613eff565b61119f565b3480156106cd57600080fd5b506103e86111ba565b3480156106e257600080fd5b506104066106f1366004613f2f565b6111bf565b34801561070257600080fd5b5061041d611220565b34801561071757600080fd5b5061041d611226565b34801561072c57600080fd5b5061041d61122c565b34801561074157600080fd5b5061052c611237565b34801561075657600080fd5b5061041d611246565b34801561076b57600080fd5b5061050a61124c565b34801561078057600080fd5b5061052c611251565b34801561079557600080fd5b5061052c611257565b3480156107aa57600080fd5b5061044a61125d565b3480156107bf57600080fd5b5061046c611279565b3480156107d457600080fd5b5061041d611282565b3480156107e957600080fd5b506103e861111b565b3480156107fe57600080fd5b5061040661080d366004613e97565b61128b565b34801561081e57600080fd5b5061041d61082d366004613e50565b61152c565b34801561083e57600080fd5b5061041d61153e565b34801561085357600080fd5b5061085c61154c565b6040516103f591906140aa565b34801561087557600080fd5b5061041d610884366004613e50565b611672565b34801561089557600080fd5b5061052c611684565b3480156108aa57600080fd5b5061041d61168a565b3480156108bf57600080fd5b5061041d61168e565b3480156108d457600080fd5b5061041d61169a565b3480156108e957600080fd5b5061052c6116a0565b3480156108fe57600080fd5b5061041d6116a6565b34801561091357600080fd5b506104066116ab565b34801561092857600080fd5b5061041d611857565b34801561093d57600080fd5b5061040661094c366004613e50565b61185d565b34801561095d57600080fd5b506103e8611a0c565b610406610974366004613e50565b611a11565b34801561098557600080fd5b5061041d611ca2565b34801561099a57600080fd5b5061052c611ca9565b3480156109af57600080fd5b5061050a611caf565b3480156109c457600080fd5b5061052c611cb4565b3480156109d957600080fd5b5061052c611cba565b606481565b6109f13334611cc0565b565b60035481565b60005460609060ff16610a275760405162461bcd60e51b8152600401610a1e9061424d565b60405180910390fd5b3361200014610a485760405162461bcd60e51b8152600401610a1e90614665565b610a50613d6c565b6000610a9185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611e4d92505050565b9150915080610aad57610aa46064611fa6565b92505050610b82565b815160009060ff16610acd57610ac68360200151612007565b9050610b4e565b825160ff1660011415610b4a57826020015151600114610b24577f70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2604051610b149061442b565b60405180910390a1506067610b45565b610ac68360200151600081518110610b3857fe5b6020026020010151612c04565b610b4e565b5060655b63ffffffff8116610b735750506040805160008152602081019091529150610b829050565b610b7c81611fa6565b93505050505b9392505050565b600d5460ff1681565b670de0b6b3a764000081565b60075481565b606881565b33600081815260096020526040902054610bd55760405162461bcd60e51b8152600401610a1e906142bb565b3360009081526009602052604080822054905190916001600160a01b0384169183156108fc0291849190818181858888f19350505050158015610c1c573d6000803e3d6000fd5b50336000818152600960205260408082208054600380549190910390559190915560018054600019019055517f433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f91610c7591849061404f565b60405180910390a15050565b3361100114610ca25760405162461bcd60e51b8152600401610a1e90614728565b6001600160a01b03811660009081526008602052604090205480610cc65750610f76565b600181039050600060058281548110610cdb57fe5b60009182526020909120600360049092020101546005549091506000190180610d2a57600060058481548110610d0d57fe5b906000526020600020906004020160030181905550505050610f76565b836001600160a01b03167f3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a7083604051610d639190614780565b60405180910390a26001600160a01b038416600090815260086020526040812055600554600019018314610eaf57600580546000198101908110610da357fe5b906000526020600020906004020160058481548110610dbe57fe5b600091825260208220835460049092020180546001600160a01b03199081166001600160a01b0393841617825560018086015481840180548416918616919091179055600280870180549185018054909416919095161780835584546001600160401b03600160a01b91829004160267ffffffffffffffff60a01b1990911617808355935460ff600160e01b918290041615150260ff60e01b199094169390931790556003938401549301929092556005805492860192600892919087908110610e8457fe5b600091825260208083206004909202909101546001600160a01b031683528201929092526040019020555b6005805480610eba57fe5b60008281526020812060046000199093019283020180546001600160a01b0319908116825560018201805490911690556002810180546001600160e81b03191690556003018190559155818381610f0d57fe5b0490508015610f715760055460005b81811015610f6e578260058281548110610f3257fe5b9060005260206000209060040201600301540160058281548110610f5257fe5b6000918252602090912060036004909202010155600101610f1c565b50505b505050505b50565b61271081565b600181565b61100181565b606081565b600881565b600c5481565b61200081565b60606000805b600254811015610fee576000818152600a60209081526040808320546001600160a01b03168084526009909252909120548015610fe4576001909301925b5050600101610fa6565b5060608160405190808252806020026020018201604052801561102b57816020015b611018613d84565b8152602001906001900390816110105790505b50905060005b6002548110156110c1576000818152600a60209081526040808320546001600160a01b031680845260099092529091205480156110b7578184848151811061107557fe5b6020026020010151600001906001600160a01b031690816001600160a01b031681525050808484815181106110a657fe5b602002602001015160200181815250505b5050600101611031565b5091505090565b6001600160a01b038116600090815260086020526040812054806110f0576000915050611116565b6005600182038154811061110057fe5b9060005260206000209060040201600301549150505b919050565b600081565b606781565b6005818154811061113257fe5b600091825260209091206004909102018054600182015460028301546003909301546001600160a01b0392831694509082169291821691600160a01b81046001600160401b031691600160e01b90910460ff169086565b61100581565b600281565b61100881565b600b81565b600a602052600090815260409020546001600160a01b031681565b606681565b33612000146111e05760405162461bcd60e51b8152600401610a1e90614665565b7f41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605838383604051611213939291906147a8565b60405180910390a1505050565b61045f81565b60065481565b6603f2a69c702cdf81565b6004546001600160a01b031681565b60015481565b600981565b61100781565b61100681565b6040518060800160405280604781526020016148466047913981565b60005460ff1681565b6402540be40081565b60005460ff166112ad5760405162461bcd60e51b8152600401610a1e9061424d565b33611007146112ce5760405162461bcd60e51b8152600401610a1e906144b1565b61133884848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080518082019091526013815272065787069726554696d655365636f6e6447617606c1b60208201529150612d7c9050565b156113d5576020811461135d5760405162461bcd60e51b8152600401610a1e906145de565b604080516020601f840181900481028201810190925282815260009161139b91858580838501838280828437600092019190915250612dd692505050565b9050606481101580156113b15750620186a08111155b6113cd5760405162461bcd60e51b8152600401610a1e9061438a565b6006556114e9565b61143584848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040805180820190915260098152686275726e526174696f60b81b60208201529150612d7c9050565b156114d1576020811461145a5760405162461bcd60e51b8152600401610a1e906141df565b604080516020601f840181900481028201810190925282815260009161149891858580838501838280828437600092019190915250612dd692505050565b90506127108111156114bc5760405162461bcd60e51b8152600401610a1e90614317565b600c55600d805460ff191660011790556114e9565b60405162461bcd60e51b8152600401610a1e90614701565b7f6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a8484848460405161151e94939291906141ad565b60405180910390a150505050565b60086020526000908152604090205481565b69152d02c7e14af680000081565b6005546060906000805b8281101561159d576005818154811061156b57fe5b9060005260206000209060040201600201601c9054906101000a900460ff16611595576001909101905b600101611556565b506060816040519080825280602002602001820160405280156115ca578160200160208202803683370190505b50600092509050815b8381101561166a57600581815481106115e857fe5b9060005260206000209060040201600201601c9054906101000a900460ff16611662576005818154811061161857fe5b600091825260209091206004909102015482516001600160a01b039091169083908590811061164357fe5b6001600160a01b03909216602092830291909101909101526001909201915b6001016115d3565b509250505090565b60096020526000908152604090205481565b61100281565b4790565b67016345785d8a000081565b600b5481565b61100381565b602981565b60005460ff16156116ce5760405162461bcd60e51b8152600401610a1e90614540565b6116d6613d6c565b60006116f960405180608001604052806047815260200161484660479139611e4d565b915091508061171a5760405162461bcd60e51b8152600401610a1e90614624565b60005b82602001515181101561183f5760058360200151828151811061173c57fe5b6020908102919091018101518254600181810185556000948552838520835160049093020180546001600160a01b039384166001600160a01b03199182161782558486015182840180549186169183169190911790556040850151600283018054606088015160808901511515600160e01b0260ff60e01b196001600160401b03909216600160a01b0267ffffffffffffffff60a01b199590991692909516919091179290921695909517161790925560a090920151600390910155908501518051918401926008929091908590811061181257fe5b602090810291909101810151516001600160a01b031682528101919091526040016000205560010161171d565b505061045f600655506000805460ff19166001179055565b60025481565b336110011461187e5760405162461bcd60e51b8152600401610a1e90614728565b6001600160a01b038116600090815260086020526040902054806118a25750610f76565b6001810390506000600582815481106118b757fe5b90600052602060002090600402016003015490506000600583815481106118da57fe5b6000918252602090912060036004909202010155600554604051600019909101906001600160a01b038516907f8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d90611933908590614780565b60405180910390a28061194857505050610f76565b600081838161195357fe5b0490508015610f715760005b848110156119b157816005828154811061197557fe5b906000526020600020906004020160030154016005828154811061199557fe5b600091825260209091206003600490920201015560010161195f565b50600554600185015b81811015610f6e5782600582815481106119d057fe5b90600052602060002090600402016003015401600582815481106119f057fe5b60009182526020909120600360049092020101556001016119ba565b606581565b334114611a305760405162461bcd60e51b8152600401610a1e906146b4565b60005460ff16611a525760405162461bcd60e51b8152600401610a1e9061424d565b60003411611a725760405162461bcd60e51b8152600401610a1e906143d1565b6001600160a01b038116600090815260086020526040812054600d5434929060ff1615611a9e5750600c545b600083118015611aae5750600081115b15611b5b576000611ad7612710611acb868563ffffffff612ddb16565b9063ffffffff612e1516565b90508015611b595760405161dead9082156108fc029083906000818181858888f19350505050158015611b0e573d6000803e3d6000fd5b507f627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee581604051611b3e9190614780565b60405180910390a1611b56848263ffffffff612e5716565b93505b505b8115611c5a57600060056001840381548110611b7357fe5b9060005260206000209060040201905080600201601c9054906101000a900460ff1615611be057846001600160a01b03167ff177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b485604051611bd39190614780565b60405180910390a2611c54565b600754611bf3908563ffffffff612e9916565b6007556003810154611c0b908563ffffffff612e9916565b60038201556040516001600160a01b038616907f93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc05590611c4b908790614780565b60405180910390a25b50611c9c565b836001600160a01b03167ff177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b484604051611c939190614780565b60405180910390a25b50505050565b620186a081565b61100081565b600381565b61dead81565b61100481565b80670de0b6b3a76400008110801590611cfe57506001600160a01b03831660009081526009602052604090205469152d02c7e14af680000090820111155b611d1a5760405162461bcd60e51b8152600401610a1e90614577565b6001600160a01b038316600090815260096020526040902054158015611d485750620186a060015460010111155b80611d6a57506001600160a01b03831660009081526009602052604090205415155b611d865760405162461bcd60e51b8152600401610a1e906142e0565b6001600160a01b038316600090815260096020526040902054611df457600180546000908152600a6020908152604080832080546001600160a01b0319166001600160a01b038916908117909155835260099091529020829055805481018155600280549091019055611e13565b6001600160a01b03831660009081526009602052604090208054820190555b60038054820190556040517f26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f90611213908590849061404f565b611e55613d6c565b6000611e5f613d6c565b611e67613d9b565b611e78611e7386612ebe565b612ee3565b90506000805b611e8783612f2d565b15611f985780611eac57611ea2611e9d84612f4e565b612f9c565b60ff168452611f90565b8060011415611f8b576060611ec8611ec385612f4e565b61301e565b90508051604051908082528060200260200182016040528015611f0557816020015b611ef2613dbb565b815260200190600190039081611eea5790505b50602086015260005b8151811015611f8057611f1f613dbb565b6000611f3d848481518110611f3057fe5b60200260200101516130ef565b9150915080611f5a57876000995099505050505050505050611fa1565b8188602001518481518110611f6b57fe5b60209081029190910101525050600101611f0e565b506001925050611f90565b611f98565b600101611e7e565b50919350909150505b915091565b604080516001808252818301909252606091829190816020015b6060815260200190600190039081611fc0579050509050611fe68363ffffffff166131cc565b81600081518110611ff357fe5b6020026020010181905250610b82816131df565b600080606061201584613269565b9150915081612060577f70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb28160405161204d919061419a565b60405180910390a1606692505050611116565b600080805b6005548110156120dd5767016345785d8a00006005828154811061208557fe5b906000526020600020906004020160030154106120a7576001909201916120d5565b6000600582815481106120b657fe5b90600052602060002090600402016003015411156120d5576001909101905b600101612065565b5060608260405190808252806020026020018201604052801561210a578160200160208202803683370190505b509050606083604051908082528060200260200182016040528015612139578160200160208202803683370190505b509050606084604051908082528060200260200182016040528015612168578160200160208202803683370190505b509050606085604051908082528060200260200182016040528015612197578160200160208202803683370190505b50905060006060866040519080825280602002602001820160405280156121c8578160200160208202803683370190505b5090506060876040519080825280602002602001820160405280156121f7578160200160208202803683370190505b509050600098506000975060608d905060006110046001600160a01b031663149d14d96040518163ffffffff1660e01b815260040160206040518083038186803b15801561224457600080fd5b505afa158015612258573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061227c9190613f17565b905067016345785d8a00008111156122d9577f70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb26040516122bb906144ff565b60405180910390a160689d5050505050505050505050505050611116565b60005b60055481101561254c5767016345785d8a0000600582815481106122fc57fe5b90600052602060002090600402016003015410612482576005818154811061232057fe5b906000526020600020906004020160020160009054906101000a90046001600160a01b03168a8d8151811061235157fe5b60200260200101906001600160a01b031690816001600160a01b03168152505060006402540be4006005838154811061238657fe5b9060005260206000209060040201600301548161239f57fe5b06600583815481106123ad57fe5b9060005260206000209060040201600301540390506123d58382612e5790919063ffffffff16565b8a8e815181106123e157fe5b602002602001018181525050600582815481106123fa57fe5b906000526020600020906004020160020160009054906101000a90046001600160a01b0316888e8151811061242b57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505081898e8151811061245857fe5b6020908102919091010152612473878263ffffffff612e9916565b6001909d019c96506125449050565b60006005828154811061249157fe5b906000526020600020906004020160030154111561254457600581815481106124b657fe5b906000526020600020906004020160010160009054906101000a90046001600160a01b0316858c815181106124e757fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506005818154811061251457fe5b906000526020600020906004020160030154848c8151811061253257fe5b60209081029190910101526001909a01995b6001016122dc565b50600085156127c2576006546040516303702b2960e51b815261100491636e056520918991612586918f918f918e914201906004016140bd565b6020604051808303818588803b15801561259f57600080fd5b505af1935050505080156125d0575060408051601f3d908101601f191682019092526125cd91810190613e77565b60015b612747576040516000815260443d10156125ec57506000612687565b60046000803e60005160e01c6308c379a0811461260d576000915050612687565b60043d036004833e81513d60248201116001600160401b038211171561263857600092505050612687565b80830180516001600160401b03811115612659576000945050505050612687565b8060208301013d860181111561267757600095505050505050612687565b601f01601f191660405250925050505b8061269257506126d4565b60019150867fa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280826040516126c6919061419a565b60405180910390a250612742565b3d8080156126fe576040519150601f19603f3d011682016040523d82523d6000602084013e612703565b606091505b5060019150867fbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a4582604051612738919061419a565b60405180910390a2505b6127c2565b8015612789577fa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b708760405161277c9190614780565b60405180910390a16127c0565b867fa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf2806040516127b790614216565b60405180910390a25b505b801561299f5760005b885181101561299d5760008982815181106127e257fe5b602002602001015190506000600582815481106127fb57fe5b906000526020600020906004020160010160009054906101000a90046001600160a01b03166001600160a01b03166108fc6005848154811061283957fe5b9060005260206000209060040201600301549081150290604051600060405180830381858888f1935050505090508015612902576005828154811061287a57fe5b906000526020600020906004020160010160009054906101000a90046001600160a01b03166001600160a01b03167f6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d600584815481106128d657fe5b9060005260206000209060040201600301546040516128f59190614780565b60405180910390a2612993565b6005828154811061290f57fe5b906000526020600020906004020160010160009054906101000a90046001600160a01b03166001600160a01b03167f25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d6005848154811061296b57fe5b90600052602060002090600402016003015460405161298a9190614780565b60405180910390a25b50506001016127cb565b505b845115612ae95760005b8551811015612ae75760008682815181106129c057fe5b60200260200101516001600160a01b03166108fc8784815181106129e057fe5b60200260200101519081150290604051600060405180830381858888f1935050505090508015612a7657868281518110612a1657fe5b60200260200101516001600160a01b03167f6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d878481518110612a5457fe5b6020026020010151604051612a699190614780565b60405180910390a2612ade565b868281518110612a8257fe5b60200260200101516001600160a01b03167f25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d878481518110612ac057fe5b6020026020010151604051612ad59190614780565b60405180910390a25b506001016129a9565b505b4715612b56577f6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d47604051612b1e9190614780565b60405180910390a1604051611002904780156108fc02916000818181858888f19350505050158015612b54573d6000803e3d6000fd5b505b60006007819055600b55825115612b7057612b708361334b565b6110016001600160a01b031663fc4333cd6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015612bad57600080fd5b505af1158015612bc1573d6000803e3d6000fd5b50506040517fedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf925060009150a15060009f9e505050505050505050505050505050565b80516001600160a01b0316600090815260086020526040812054801580612c56575060056001820381548110612c3657fe5b9060005260206000209060040201600201601c9054906101000a900460ff165b15612c9c5782516040516001600160a01b03909116907fe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be90600090a26000915050611116565b600554600b54600019820111801590612cf25784516040516001600160a01b03909116907fe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be90600090a260009350505050611116565b600b80546001908101909155600580546000198601908110612d1057fe5b6000918252602082206002600490920201018054921515600160e01b0260ff60e01b199093169290921790915585516040516001600160a01b03909116917ff226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a91a2506000949350505050565b600081604051602001612d8f919061401f565b6040516020818303038152906040528051906020012083604051602001612db6919061401f565b604051602081830303815290604052805190602001201490505b92915050565b015190565b600082612dea57506000612dd0565b82820282848281612df757fe5b0414610b825760405162461bcd60e51b8152600401610a1e90614470565b6000610b8283836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061380d565b6000610b8283836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250613844565b600082820183811015610b825760405162461bcd60e51b8152600401610a1e90614284565b612ec6613df0565b506040805180820190915281518152602082810190820152919050565b612eeb613d9b565b612ef482613870565b612efd57600080fd5b6000612f0c83602001516138aa565b60208085015160408051808201909152868152920190820152915050919050565b6000612f37613df0565b505080518051602091820151919092015191011190565b612f56613df0565b612f5f82612f2d565b612f6857600080fd5b60208201516000612f788261390d565b80830160209586015260408051808201909152908152938401919091525090919050565b805160009015801590612fb157508151602110155b612fba57600080fd5b6000612fc983602001516138aa565b90508083600001511015612fef5760405162461bcd60e51b8152600401610a1e906145a7565b82516020808501518301805192849003929183101561301557826020036101000a820491505b50949350505050565b606061302982613870565b61303257600080fd5b600061303d836139ee565b905060608160405190808252806020026020018201604052801561307b57816020015b613068613df0565b8152602001906001900390816130605790505b509050600061308d85602001516138aa565b60208601510190506000805b848110156130e4576130aa8361390d565b91506040518060400160405280838152602001848152508482815181106130cd57fe5b602090810291909101015291810191600101613099565b509195945050505050565b6130f7613dbb565b6000613101613dbb565b613109613d9b565b61311285612ee3565b90506000805b61312183612f2d565b15611f98578061314c5761313c61313784612f4e565b613a4a565b6001600160a01b031684526131c4565b80600114156131745761316161313784612f4e565b6001600160a01b031660208501526131c4565b806002141561319c5761318961313784612f4e565b6001600160a01b031660408501526131c4565b8060031415611f8b576131b1611e9d84612f4e565b6001600160401b03166060850152600191505b600101613118565b6060612dd06131da83613a64565b613b4a565b60608151600014156132005750604080516000815260208101909152611116565b60608260008151811061320f57fe5b602002602001015190506000600190505b8351811015613250576132468285838151811061323957fe5b6020026020010151613b9c565b9150600101613220565b50610b82613263825160c060ff16613c19565b82613b9c565b6000606060298351111561329b5760006040518060600160405280602981526020016147f26029913991509150611fa1565b60005b83518110156133315760005b81811015613328578481815181106132be57fe5b6020026020010151600001516001600160a01b03168583815181106132df57fe5b6020026020010151600001516001600160a01b031614156133205760006040518060600160405280602b815260200161481b602b9139935093505050611fa1565b6001016132aa565b5060010161329e565b505060408051602081019091526000815260019150915091565b600554815160005b82811015613468576001613365613dbb565b6005838154811061337257fe5b600091825260208083206040805160c08101825260049490940290910180546001600160a01b0390811685526001820154811693850193909352600281015492831691840191909152600160a01b82046001600160401b03166060840152600160e01b90910460ff16151560808301526003015460a082015291505b8481101561343c5786818151811061340257fe5b6020026020010151600001516001600160a01b031682600001516001600160a01b03161415613434576000925061343c565b6001016133ee565b50811561345e5780516001600160a01b03166000908152600860205260408120555b5050600101613353565b50808211156134dd57805b828110156134db57600580548061348657fe5b60008281526020812060046000199093019283020180546001600160a01b03199081168255600182810180549092169091556002820180546001600160e81b0319169055600390910191909155915501613473565b505b60008183106134ec57816134ee565b825b905060005b818110156136e8576135a085828151811061350a57fe5b60200260200101516005838154811061351f57fe5b60009182526020918290206040805160c08101825260049390930290910180546001600160a01b0390811684526001820154811694840194909452600281015493841691830191909152600160a01b83046001600160401b03166060830152600160e01b90920460ff161515608082015260039091015460a0820152613ceb565b6136bb5780600101600860008784815181106135b857fe5b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020819055508481815181106135f457fe5b60200260200101516005828154811061360957fe5b6000918252602091829020835160049092020180546001600160a01b039283166001600160a01b0319918216178255928401516001820180549184169185169190911790556040840151600282018054606087015160808801511515600160e01b0260ff60e01b196001600160401b03909216600160a01b0267ffffffffffffffff60a01b1995909716929097169190911792909216939093171692909217905560a0909101516003909101556136e0565b6000600582815481106136ca57fe5b9060005260206000209060040201600301819055505b6001016134f3565b5082821115611c9c57825b82811015610f7157600585828151811061370957fe5b6020908102919091018101518254600181810185556000948552838520835160049093020180546001600160a01b039384166001600160a01b03199182161782559484015181830180549185169187169190911790556040840151600282018054606087015160808801511515600160e01b0260ff60e01b196001600160401b03909216600160a01b0267ffffffffffffffff60a01b199590981692909916919091179290921694909417169490941790915560a0909101516003909201919091558651908301916008918890859081106137e057fe5b602090810291909101810151516001600160a01b03168252810191909152604001600020556001016136f3565b6000818361382e5760405162461bcd60e51b8152600401610a1e919061419a565b50600083858161383a57fe5b0495945050505050565b600081848411156138685760405162461bcd60e51b8152600401610a1e919061419a565b505050900390565b805160009061388157506000611116565b6020820151805160001a9060c08210156138a057600092505050611116565b5060019392505050565b8051600090811a60808110156138c4576000915050611116565b60b88110806138df575060c081108015906138df575060f881105b156138ee576001915050611116565b60c08110156139025760b519019050611116565b60f519019050611116565b80516000908190811a608081101561392857600191506139e7565b60b881101561393d57607e19810191506139e7565b60c081101561398e57600060b78203600186019550806020036101000a8651049150600181018201935050808310156139885760405162461bcd60e51b8152600401610a1e90614400565b506139e7565b60f88110156139a35760be19810191506139e7565b600060f78203600186019550806020036101000a8651049150600181018201935050808310156139e55760405162461bcd60e51b8152600401610a1e90614400565b505b5092915050565b80516000906139ff57506000611116565b60008090506000613a1384602001516138aa565b602085015185519181019250015b80821015613a4157613a328261390d565b60019093019290910190613a21565b50909392505050565b8051600090601514613a5b57600080fd5b612dd082612f9c565b604080516020808252818301909252606091829190602082018180368337505050602081018490529050600067ffffffffffffffff198416613aa857506018613acc565b6fffffffffffffffffffffffffffffffff198416613ac857506010613acc565b5060005b6020811015613b0257818181518110613ae157fe5b01602001516001600160f81b03191615613afa57613b02565b600101613acc565b60008160200390506060816040519080825280601f01601f191660200182016040528015613b37576020820181803683370190505b5080830196909652508452509192915050565b606081516001148015613b7c5750607f60f81b82600081518110613b6a57fe5b01602001516001600160f81b03191611155b15613b88575080611116565b612dd0613b9a8351608060ff16613c19565b835b6060806040519050835180825260208201818101602087015b81831015613bcd578051835260209283019201613bb5565b50855184518101855292509050808201602086015b81831015613bfa578051835260209283019201613be2565b508651929092011591909101601f01601f191660405250905092915050565b6060680100000000000000008310613c435760405162461bcd60e51b8152600401610a1e90614362565b60408051600180825281830190925260609160208201818036833701905050905060378411613c9d5782840160f81b81600081518110613c7f57fe5b60200101906001600160f81b031916908160001a9053509050612dd0565b6060613ca885613a64565b90508381510160370160f81b82600081518110613cc157fe5b60200101906001600160f81b031916908160001a905350613ce28282613b9c565b95945050505050565b805182516000916001600160a01b039182169116148015613d25575081602001516001600160a01b031683602001516001600160a01b0316145b8015613d4a575081604001516001600160a01b031683604001516001600160a01b0316145b8015610b825750506060908101519101516001600160401b0390811691161490565b60408051808201909152600081526060602082015290565b604080518082019091526000808252602082015290565b6040518060400160405280613dae613df0565b8152602001600081525090565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b604051806040016040528060008152602001600081525090565b60008083601f840112613e1b578182fd5b5081356001600160401b03811115613e31578182fd5b602083019150836020828501011115613e4957600080fd5b9250929050565b600060208284031215613e61578081fd5b81356001600160a01b0381168114610b82578182fd5b600060208284031215613e88578081fd5b81518015158114610b82578182fd5b60008060008060408587031215613eac578283fd5b84356001600160401b0380821115613ec2578485fd5b613ece88838901613e0a565b90965094506020870135915080821115613ee6578384fd5b50613ef387828801613e0a565b95989497509550505050565b600060208284031215613f10578081fd5b5035919050565b600060208284031215613f28578081fd5b5051919050565b600080600060408486031215613f43578283fd5b833560ff81168114613f53578384fd5b925060208401356001600160401b03811115613f6d578283fd5b613f7986828701613e0a565b9497909650939450505050565b6000815180845260208085019450808401835b83811015613fbe5781516001600160a01b031687529582019590820190600101613f99565b509495945050505050565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b6000815180845261400b8160208601602086016147c5565b601f01601f19169290920160200192915050565b600082516140318184602087016147c5565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b03968716815294861660208601529290941660408401526001600160401b03166060830152911515608082015260a081019190915260c00190565b600060208252610b826020830184613f86565b6000608082526140d06080830187613f86565b828103602084810191909152865180835287820192820190845b81811015614106578451835293830193918301916001016140ea565b5050848103604086015261411a8188613f86565b93505050506001600160401b038316606083015295945050505050565b602080825282518282018190526000919060409081850190868401855b8281101561418257815180516001600160a01b03168552860151868501529284019290850190600101614154565b5091979650505050505050565b901515815260200190565b600060208252610b826020830184613ff3565b6000604082526141c1604083018688613fc9565b82810360208401526141d4818587613fc9565b979650505050505050565b6020808252601c908201527f6c656e677468206f66206275726e526174696f206d69736d6174636800000000604082015260600190565b6020808252601b908201527f6261746368207472616e736665722072657475726e2066616c73650000000000604082015260600190565b60208082526019908201527f74686520636f6e7472616374206e6f7420696e69742079657400000000000000604082015260600190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b6020808252600b908201526a1b5d5cdd081cdd185ad95960aa1b604082015260600190565b60208082526019908201527f7374616b652070656f706c65206f7574206f662072616e676500000000000000604082015260600190565b6020808252602b908201527f746865206275726e526174696f206d757374206265206e6f206772656174657260408201526a0207468616e2031303030360ac1b606082015260800190565b6020808252600e908201526d696e70757420746f6f206c6f6e6760901b604082015260600190565b60208082526027908201527f7468652065787069726554696d655365636f6e64476170206973206f7574206f604082015266662072616e676560c81b606082015260800190565b6020808252601590820152746465706f7369742076616c7565206973207a65726f60581b604082015260600190565b6020808252601190820152706164646974696f6e206f766572666c6f7760781b604082015260600190565b60208082526025908201527f6c656e677468206f66206a61696c2076616c696461746f7273206d757374206260408201526465206f6e6560d81b606082015260800190565b60208082526021908201527f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f6040820152607760f81b606082015260800190565b6020808252602e908201527f746865206d6573736167652073656e646572206d75737420626520676f76657260408201526d1b985b98d94818dbdb9d1c9858dd60921b606082015260800190565b60208082526021908201527f666565206973206c6172676572207468616e2044555354595f494e434f4d494e6040820152604760f81b606082015260800190565b60208082526019908201527f74686520636f6e747261637420616c726561647920696e697400000000000000604082015260600190565b6020808252601690820152751cdd185ad94818dbdd5b9d081b5d5cdd081d985b1a5960521b604082015260600190565b6020808252601a908201527f6c656e677468206973206c657373207468616e206f6666736574000000000000604082015260600190565b60208082526026908201527f6c656e677468206f662065787069726554696d655365636f6e64476170206d696040820152650e6dac2e8c6d60d31b606082015260800190565b60208082526021908201527f6661696c656420746f20706172736520696e69742076616c696461746f7253656040820152601d60fa1b606082015260800190565b6020808252602f908201527f746865206d6573736167652073656e646572206d7573742062652063726f737360408201526e0818da185a5b8818dbdb9d1c9858dd608a1b606082015260800190565b6020808252602d908201527f746865206d6573736167652073656e646572206d75737420626520746865206260408201526c3637b1b590383937b23ab1b2b960991b606082015260800190565b6020808252600d908201526c756e6b6e6f776e20706172616d60981b604082015260600190565b60208082526029908201527f746865206d6573736167652073656e646572206d75737420626520736c6173686040820152680818dbdb9d1c9858dd60ba1b606082015260800190565b61ffff91909116815260200190565b90815260200190565b63ffffffff91909116815260200190565b60ff91909116815260200190565b600060ff8516825260406020830152613ce2604083018486613fc9565b60005b838110156147e05781810151838201526020016147c8565b83811115611c9c575050600091015256fe746865206e756d626572206f662076616c696461746f72732065786365656420746865206c696d69746475706c696361746520636f6e73656e7375732061646472657373206f662076616c696461746f72536574f84580f842f8409485245585de1768e6d40b46c5b4d0c17785ab7a9e9485245585de1768e6d40b46c5b4d0c17785ab7a9e9485245585de1768e6d40b46c5b4d0c17785ab7a9e64a2646970667358221220ef67818951d955e4e7058b33107f57496ec3e4eb35ce5bcdf12a5db94d8bf9d564736f6c63430006040033"

// DeployBSCValidatorSet deploys a new Ethereum contract, binding an instance of BSCValidatorSet to it.
func DeployBSCValidatorSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BSCValidatorSet, error) {
	parsed, err := abi.JSON(strings.NewReader(BSCValidatorSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BSCValidatorSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BSCValidatorSet{BSCValidatorSetCaller: BSCValidatorSetCaller{contract: contract}, BSCValidatorSetTransactor: BSCValidatorSetTransactor{contract: contract}, BSCValidatorSetFilterer: BSCValidatorSetFilterer{contract: contract}}, nil
}

// BSCValidatorSet is an auto generated Go binding around an Ethereum contract.
type BSCValidatorSet struct {
	BSCValidatorSetCaller     // Read-only binding to the contract
	BSCValidatorSetTransactor // Write-only binding to the contract
	BSCValidatorSetFilterer   // Log filterer for contract events
}

// BSCValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BSCValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BSCValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BSCValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSCValidatorSetSession struct {
	Contract     *BSCValidatorSet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BSCValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BSCValidatorSetCallerSession struct {
	Contract *BSCValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BSCValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BSCValidatorSetTransactorSession struct {
	Contract     *BSCValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BSCValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BSCValidatorSetRaw struct {
	Contract *BSCValidatorSet // Generic contract binding to access the raw methods on
}

// BSCValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BSCValidatorSetCallerRaw struct {
	Contract *BSCValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// BSCValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BSCValidatorSetTransactorRaw struct {
	Contract *BSCValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBSCValidatorSet creates a new instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSet(address common.Address, backend bind.ContractBackend) (*BSCValidatorSet, error) {
	contract, err := bindBSCValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSet{BSCValidatorSetCaller: BSCValidatorSetCaller{contract: contract}, BSCValidatorSetTransactor: BSCValidatorSetTransactor{contract: contract}, BSCValidatorSetFilterer: BSCValidatorSetFilterer{contract: contract}}, nil
}

// NewBSCValidatorSetCaller creates a new read-only instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*BSCValidatorSetCaller, error) {
	contract, err := bindBSCValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetCaller{contract: contract}, nil
}

// NewBSCValidatorSetTransactor creates a new write-only instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*BSCValidatorSetTransactor, error) {
	contract, err := bindBSCValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetTransactor{contract: contract}, nil
}

// NewBSCValidatorSetFilterer creates a new log filterer instance of BSCValidatorSet, bound to a specific deployed contract.
func NewBSCValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*BSCValidatorSetFilterer, error) {
	contract, err := bindBSCValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetFilterer{contract: contract}, nil
}

// bindBSCValidatorSet binds a generic wrapper to an already deployed contract.
func bindBSCValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BSCValidatorSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCValidatorSet *BSCValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCValidatorSet.Contract.BSCValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCValidatorSet *BSCValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.BSCValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCValidatorSet *BSCValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.BSCValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCValidatorSet *BSCValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCValidatorSet *BSCValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCValidatorSet *BSCValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "BIND_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) BINDCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.BINDCHANNELID(&_BSCValidatorSet.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BINDCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.BINDCHANNELID(&_BSCValidatorSet.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) BURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) BURNADDRESS() (common.Address, error) {
	return _BSCValidatorSet.Contract.BURNADDRESS(&_BSCValidatorSet.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BURNADDRESS() (common.Address, error) {
	return _BSCValidatorSet.Contract.BURNADDRESS(&_BSCValidatorSet.CallOpts)
}

// BURNRATIOSCALE is a free data retrieval call binding the contract method 0x3de0f0d8.
//
// Solidity: function BURN_RATIO_SCALE() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) BURNRATIOSCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "BURN_RATIO_SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BURNRATIOSCALE is a free data retrieval call binding the contract method 0x3de0f0d8.
//
// Solidity: function BURN_RATIO_SCALE() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) BURNRATIOSCALE() (*big.Int, error) {
	return _BSCValidatorSet.Contract.BURNRATIOSCALE(&_BSCValidatorSet.CallOpts)
}

// BURNRATIOSCALE is a free data retrieval call binding the contract method 0x3de0f0d8.
//
// Solidity: function BURN_RATIO_SCALE() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BURNRATIOSCALE() (*big.Int, error) {
	return _BSCValidatorSet.Contract.BURNRATIOSCALE(&_BSCValidatorSet.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) CODEOK() (uint32, error) {
	return _BSCValidatorSet.Contract.CODEOK(&_BSCValidatorSet.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CODEOK() (uint32, error) {
	return _BSCValidatorSet.Contract.CODEOK(&_BSCValidatorSet.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.CROSSCHAINCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.CROSSCHAINCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) DUSTYINCOMING(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "DUSTY_INCOMING")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) DUSTYINCOMING() (*big.Int, error) {
	return _BSCValidatorSet.Contract.DUSTYINCOMING(&_BSCValidatorSet.CallOpts)
}

// DUSTYINCOMING is a free data retrieval call binding the contract method 0xd86222d5.
//
// Solidity: function DUSTY_INCOMING() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) DUSTYINCOMING() (*big.Int, error) {
	return _BSCValidatorSet.Contract.DUSTYINCOMING(&_BSCValidatorSet.CallOpts)
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORFAILCHECKVALIDATORS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_FAIL_CHECK_VALIDATORS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORFAILCHECKVALIDATORS() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORFAILCHECKVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// ERRORFAILCHECKVALIDATORS is a free data retrieval call binding the contract method 0x81650b62.
//
// Solidity: function ERROR_FAIL_CHECK_VALIDATORS() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORFAILCHECKVALIDATORS() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORFAILCHECKVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORFAILDECODE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORFAILDECODE(&_BSCValidatorSet.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORFAILDECODE(&_BSCValidatorSet.CallOpts)
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORLENOFVALMISMATCH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_LEN_OF_VAL_MISMATCH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORLENOFVALMISMATCH() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORLENOFVALMISMATCH(&_BSCValidatorSet.CallOpts)
}

// ERRORLENOFVALMISMATCH is a free data retrieval call binding the contract method 0x5d77156c.
//
// Solidity: function ERROR_LEN_OF_VAL_MISMATCH() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORLENOFVALMISMATCH() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORLENOFVALMISMATCH(&_BSCValidatorSet.CallOpts)
}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORRELAYFEETOOLARGE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_RELAYFEE_TOO_LARGE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORRELAYFEETOOLARGE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORRELAYFEETOOLARGE(&_BSCValidatorSet.CallOpts)
}

// ERRORRELAYFEETOOLARGE is a free data retrieval call binding the contract method 0x219f22d5.
//
// Solidity: function ERROR_RELAYFEE_TOO_LARGE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORRELAYFEETOOLARGE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORRELAYFEETOOLARGE(&_BSCValidatorSet.CallOpts)
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCaller) ERRORUNKNOWNPACKAGETYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "ERROR_UNKNOWN_PACKAGE_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetSession) ERRORUNKNOWNPACKAGETYPE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORUNKNOWNPACKAGETYPE(&_BSCValidatorSet.CallOpts)
}

// ERRORUNKNOWNPACKAGETYPE is a free data retrieval call binding the contract method 0xeda5868c.
//
// Solidity: function ERROR_UNKNOWN_PACKAGE_TYPE() view returns(uint32)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ERRORUNKNOWNPACKAGETYPE() (uint32, error) {
	return _BSCValidatorSet.Contract.ERRORUNKNOWNPACKAGETYPE(&_BSCValidatorSet.CallOpts)
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) EXPIRETIMESECONDGAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "EXPIRE_TIME_SECOND_GAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) EXPIRETIMESECONDGAP() (*big.Int, error) {
	return _BSCValidatorSet.Contract.EXPIRETIMESECONDGAP(&_BSCValidatorSet.CallOpts)
}

// EXPIRETIMESECONDGAP is a free data retrieval call binding the contract method 0x853230aa.
//
// Solidity: function EXPIRE_TIME_SECOND_GAP() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) EXPIRETIMESECONDGAP() (*big.Int, error) {
	return _BSCValidatorSet.Contract.EXPIRETIMESECONDGAP(&_BSCValidatorSet.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) GOVCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.GOVCHANNELID(&_BSCValidatorSet.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) GOVCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.GOVCHANNELID(&_BSCValidatorSet.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "GOV_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) GOVHUBADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.GOVHUBADDR(&_BSCValidatorSet.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) GOVHUBADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.GOVHUBADDR(&_BSCValidatorSet.CallOpts)
}

// GetAllStakeInfo is a free data retrieval call binding the contract method 0x54b1a671.
//
// Solidity: function GetAllStakeInfo() view returns((address,uint256)[])
func (_BSCValidatorSet *BSCValidatorSetCaller) GetAllStakeInfo(opts *bind.CallOpts) ([]BSCValidatorSetStakeInfo, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "GetAllStakeInfo")

	if err != nil {
		return *new([]BSCValidatorSetStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]BSCValidatorSetStakeInfo)).(*[]BSCValidatorSetStakeInfo)

	return out0, err

}

// GetAllStakeInfo is a free data retrieval call binding the contract method 0x54b1a671.
//
// Solidity: function GetAllStakeInfo() view returns((address,uint256)[])
func (_BSCValidatorSet *BSCValidatorSetSession) GetAllStakeInfo() ([]BSCValidatorSetStakeInfo, error) {
	return _BSCValidatorSet.Contract.GetAllStakeInfo(&_BSCValidatorSet.CallOpts)
}

// GetAllStakeInfo is a free data retrieval call binding the contract method 0x54b1a671.
//
// Solidity: function GetAllStakeInfo() view returns((address,uint256)[])
func (_BSCValidatorSet *BSCValidatorSetCallerSession) GetAllStakeInfo() ([]BSCValidatorSetStakeInfo, error) {
	return _BSCValidatorSet.Contract.GetAllStakeInfo(&_BSCValidatorSet.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "INCENTIVIZE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) INCENTIVIZEADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.INCENTIVIZEADDR(&_BSCValidatorSet.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.INCENTIVIZEADDR(&_BSCValidatorSet.CallOpts)
}

// INITBURNRATIO is a free data retrieval call binding the contract method 0x78dfed4a.
//
// Solidity: function INIT_BURN_RATIO() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) INITBURNRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "INIT_BURN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITBURNRATIO is a free data retrieval call binding the contract method 0x78dfed4a.
//
// Solidity: function INIT_BURN_RATIO() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) INITBURNRATIO() (*big.Int, error) {
	return _BSCValidatorSet.Contract.INITBURNRATIO(&_BSCValidatorSet.CallOpts)
}

// INITBURNRATIO is a free data retrieval call binding the contract method 0x78dfed4a.
//
// Solidity: function INIT_BURN_RATIO() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) INITBURNRATIO() (*big.Int, error) {
	return _BSCValidatorSet.Contract.INITBURNRATIO(&_BSCValidatorSet.CallOpts)
}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() view returns(bytes)
func (_BSCValidatorSet *BSCValidatorSetCaller) INITVALIDATORSETBYTES(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "INIT_VALIDATORSET_BYTES")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() view returns(bytes)
func (_BSCValidatorSet *BSCValidatorSetSession) INITVALIDATORSETBYTES() ([]byte, error) {
	return _BSCValidatorSet.Contract.INITVALIDATORSETBYTES(&_BSCValidatorSet.CallOpts)
}

// INITVALIDATORSETBYTES is a free data retrieval call binding the contract method 0xa5422d5c.
//
// Solidity: function INIT_VALIDATORSET_BYTES() view returns(bytes)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) INITVALIDATORSETBYTES() ([]byte, error) {
	return _BSCValidatorSet.Contract.INITVALIDATORSETBYTES(&_BSCValidatorSet.CallOpts)
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) JAILMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "JAIL_MESSAGE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) JAILMESSAGETYPE() (uint8, error) {
	return _BSCValidatorSet.Contract.JAILMESSAGETYPE(&_BSCValidatorSet.CallOpts)
}

// JAILMESSAGETYPE is a free data retrieval call binding the contract method 0xbf9f4995.
//
// Solidity: function JAIL_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) JAILMESSAGETYPE() (uint8, error) {
	return _BSCValidatorSet.Contract.JAILMESSAGETYPE(&_BSCValidatorSet.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.LIGHTCLIENTADDR(&_BSCValidatorSet.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.LIGHTCLIENTADDR(&_BSCValidatorSet.CallOpts)
}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MAXNUMOFVALIDATORS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "MAX_NUM_OF_VALIDATORS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MAXNUMOFVALIDATORS() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MAXNUMOFVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// MAXNUMOFVALIDATORS is a free data retrieval call binding the contract method 0xe086c7b1.
//
// Solidity: function MAX_NUM_OF_VALIDATORS() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MAXNUMOFVALIDATORS() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MAXNUMOFVALIDATORS(&_BSCValidatorSet.CallOpts)
}

// MAXSTAKEFFFCOUNT is a free data retrieval call binding the contract method 0x8a10febe.
//
// Solidity: function MAX_STAKE_FFF_COUNT() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MAXSTAKEFFFCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "MAX_STAKE_FFF_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTAKEFFFCOUNT is a free data retrieval call binding the contract method 0x8a10febe.
//
// Solidity: function MAX_STAKE_FFF_COUNT() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MAXSTAKEFFFCOUNT() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MAXSTAKEFFFCOUNT(&_BSCValidatorSet.CallOpts)
}

// MAXSTAKEFFFCOUNT is a free data retrieval call binding the contract method 0x8a10febe.
//
// Solidity: function MAX_STAKE_FFF_COUNT() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MAXSTAKEFFFCOUNT() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MAXSTAKEFFFCOUNT(&_BSCValidatorSet.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) PRECISION() (*big.Int, error) {
	return _BSCValidatorSet.Contract.PRECISION(&_BSCValidatorSet.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) PRECISION() (*big.Int, error) {
	return _BSCValidatorSet.Contract.PRECISION(&_BSCValidatorSet.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.RELAYERHUBCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.RELAYERHUBCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "SLASH_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) SLASHCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.SLASHCHANNELID(&_BSCValidatorSet.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) SLASHCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.SLASHCHANNELID(&_BSCValidatorSet.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "SLASH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.SLASHCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.SLASHCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) STAKINGCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.STAKINGCHANNELID(&_BSCValidatorSet.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.STAKINGCHANNELID(&_BSCValidatorSet.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.SYSTEMREWARDADDR(&_BSCValidatorSet.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.SYSTEMREWARDADDR(&_BSCValidatorSet.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) TOKENHUBADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.TOKENHUBADDR(&_BSCValidatorSet.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.TOKENHUBADDR(&_BSCValidatorSet.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "TOKEN_MANAGER_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) TOKENMANAGERADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.TOKENMANAGERADDR(&_BSCValidatorSet.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.TOKENMANAGERADDR(&_BSCValidatorSet.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) TRANSFERINCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.TRANSFERINCHANNELID(&_BSCValidatorSet.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.TRANSFERINCHANNELID(&_BSCValidatorSet.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.TRANSFEROUTCHANNELID(&_BSCValidatorSet.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _BSCValidatorSet.Contract.TRANSFEROUTCHANNELID(&_BSCValidatorSet.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCaller) VALIDATORSUPDATEMESSAGETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "VALIDATORS_UPDATE_MESSAGE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _BSCValidatorSet.Contract.VALIDATORSUPDATEMESSAGETYPE(&_BSCValidatorSet.CallOpts)
}

// VALIDATORSUPDATEMESSAGETYPE is a free data retrieval call binding the contract method 0x5667515a.
//
// Solidity: function VALIDATORS_UPDATE_MESSAGE_TYPE() view returns(uint8)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) VALIDATORSUPDATEMESSAGETYPE() (uint8, error) {
	return _BSCValidatorSet.Contract.VALIDATORSUPDATEMESSAGETYPE(&_BSCValidatorSet.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.VALIDATORCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _BSCValidatorSet.Contract.VALIDATORCONTRACTADDR(&_BSCValidatorSet.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetSession) AlreadyInit() (bool, error) {
	return _BSCValidatorSet.Contract.AlreadyInit(&_BSCValidatorSet.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) AlreadyInit() (bool, error) {
	return _BSCValidatorSet.Contract.AlreadyInit(&_BSCValidatorSet.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_BSCValidatorSet *BSCValidatorSetCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_BSCValidatorSet *BSCValidatorSetSession) BscChainID() (uint16, error) {
	return _BSCValidatorSet.Contract.BscChainID(&_BSCValidatorSet.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BscChainID() (uint16, error) {
	return _BSCValidatorSet.Contract.BscChainID(&_BSCValidatorSet.CallOpts)
}

// BurnRatio is a free data retrieval call binding the contract method 0x5192c82c.
//
// Solidity: function burnRatio() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) BurnRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "burnRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnRatio is a free data retrieval call binding the contract method 0x5192c82c.
//
// Solidity: function burnRatio() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) BurnRatio() (*big.Int, error) {
	return _BSCValidatorSet.Contract.BurnRatio(&_BSCValidatorSet.CallOpts)
}

// BurnRatio is a free data retrieval call binding the contract method 0x5192c82c.
//
// Solidity: function burnRatio() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BurnRatio() (*big.Int, error) {
	return _BSCValidatorSet.Contract.BurnRatio(&_BSCValidatorSet.CallOpts)
}

// BurnRatioInitialized is a free data retrieval call binding the contract method 0x152ad3b8.
//
// Solidity: function burnRatioInitialized() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetCaller) BurnRatioInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "burnRatioInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnRatioInitialized is a free data retrieval call binding the contract method 0x152ad3b8.
//
// Solidity: function burnRatioInitialized() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetSession) BurnRatioInitialized() (bool, error) {
	return _BSCValidatorSet.Contract.BurnRatioInitialized(&_BSCValidatorSet.CallOpts)
}

// BurnRatioInitialized is a free data retrieval call binding the contract method 0x152ad3b8.
//
// Solidity: function burnRatioInitialized() view returns(bool)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) BurnRatioInitialized() (bool, error) {
	return _BSCValidatorSet.Contract.BurnRatioInitialized(&_BSCValidatorSet.CallOpts)
}

// CurrStakeFFF is a free data retrieval call binding the contract method 0x101a875f.
//
// Solidity: function currStakeFFF() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrStakeFFF(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currStakeFFF")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrStakeFFF is a free data retrieval call binding the contract method 0x101a875f.
//
// Solidity: function currStakeFFF() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrStakeFFF() (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrStakeFFF(&_BSCValidatorSet.CallOpts)
}

// CurrStakeFFF is a free data retrieval call binding the contract method 0x101a875f.
//
// Solidity: function currStakeFFF() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrStakeFFF() (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrStakeFFF(&_BSCValidatorSet.CallOpts)
}

// CurrStakePeople is a free data retrieval call binding the contract method 0x93a101f3.
//
// Solidity: function currStakePeople() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrStakePeople(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currStakePeople")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrStakePeople is a free data retrieval call binding the contract method 0x93a101f3.
//
// Solidity: function currStakePeople() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrStakePeople() (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrStakePeople(&_BSCValidatorSet.CallOpts)
}

// CurrStakePeople is a free data retrieval call binding the contract method 0x93a101f3.
//
// Solidity: function currStakePeople() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrStakePeople() (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrStakePeople(&_BSCValidatorSet.CallOpts)
}

// CurrStakePeopleIndex is a free data retrieval call binding the contract method 0xe3260b57.
//
// Solidity: function currStakePeopleIndex() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrStakePeopleIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currStakePeopleIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrStakePeopleIndex is a free data retrieval call binding the contract method 0xe3260b57.
//
// Solidity: function currStakePeopleIndex() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrStakePeopleIndex() (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrStakePeopleIndex(&_BSCValidatorSet.CallOpts)
}

// CurrStakePeopleIndex is a free data retrieval call binding the contract method 0xe3260b57.
//
// Solidity: function currStakePeopleIndex() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrStakePeopleIndex() (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrStakePeopleIndex(&_BSCValidatorSet.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currentValidatorSet", arg0)

	outstruct := new(struct {
		ConsensusAddress common.Address
		FeeAddress       common.Address
		BBCFeeAddress    common.Address
		VotingPower      uint64
		Jailed           bool
		Incoming         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.FeeAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BBCFeeAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.VotingPower = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Jailed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Incoming = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSet(&_BSCValidatorSet.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, address feeAddress, address BBCFeeAddress, uint64 votingPower, bool jailed, uint256 incoming)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	FeeAddress       common.Address
	BBCFeeAddress    common.Address
	VotingPower      uint64
	Jailed           bool
	Incoming         *big.Int
}, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSet(&_BSCValidatorSet.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) CurrentValidatorSetMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "currentValidatorSetMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSetMap(&_BSCValidatorSet.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.CurrentValidatorSetMap(&_BSCValidatorSet.CallOpts, arg0)
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) ExpireTimeSecondGap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "expireTimeSecondGap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) ExpireTimeSecondGap() (*big.Int, error) {
	return _BSCValidatorSet.Contract.ExpireTimeSecondGap(&_BSCValidatorSet.CallOpts)
}

// ExpireTimeSecondGap is a free data retrieval call binding the contract method 0x86249882.
//
// Solidity: function expireTimeSecondGap() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) ExpireTimeSecondGap() (*big.Int, error) {
	return _BSCValidatorSet.Contract.ExpireTimeSecondGap(&_BSCValidatorSet.CallOpts)
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) GetIncoming(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "getIncoming", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) GetIncoming(validator common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.GetIncoming(&_BSCValidatorSet.CallOpts, validator)
}

// GetIncoming is a free data retrieval call binding the contract method 0x565c56b3.
//
// Solidity: function getIncoming(address validator) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) GetIncoming(validator common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.GetIncoming(&_BSCValidatorSet.CallOpts, validator)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_BSCValidatorSet *BSCValidatorSetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_BSCValidatorSet *BSCValidatorSetSession) GetValidators() ([]common.Address, error) {
	return _BSCValidatorSet.Contract.GetValidators(&_BSCValidatorSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_BSCValidatorSet *BSCValidatorSetCallerSession) GetValidators() ([]common.Address, error) {
	return _BSCValidatorSet.Contract.GetValidators(&_BSCValidatorSet.CallOpts)
}

// MaxStakeFFFCount is a free data retrieval call binding the contract method 0xb2ccf3a4.
//
// Solidity: function maxStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MaxStakeFFFCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "maxStakeFFFCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakeFFFCount is a free data retrieval call binding the contract method 0xb2ccf3a4.
//
// Solidity: function maxStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MaxStakeFFFCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MaxStakeFFFCount(&_BSCValidatorSet.CallOpts)
}

// MaxStakeFFFCount is a free data retrieval call binding the contract method 0xb2ccf3a4.
//
// Solidity: function maxStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MaxStakeFFFCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MaxStakeFFFCount(&_BSCValidatorSet.CallOpts)
}

// MaxStakePeopleCount is a free data retrieval call binding the contract method 0xf78ffcbc.
//
// Solidity: function maxStakePeopleCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MaxStakePeopleCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "maxStakePeopleCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakePeopleCount is a free data retrieval call binding the contract method 0xf78ffcbc.
//
// Solidity: function maxStakePeopleCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MaxStakePeopleCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MaxStakePeopleCount(&_BSCValidatorSet.CallOpts)
}

// MaxStakePeopleCount is a free data retrieval call binding the contract method 0xf78ffcbc.
//
// Solidity: function maxStakePeopleCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MaxStakePeopleCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MaxStakePeopleCount(&_BSCValidatorSet.CallOpts)
}

// MiniStakeFFFCount is a free data retrieval call binding the contract method 0x1faea20b.
//
// Solidity: function miniStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MiniStakeFFFCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "miniStakeFFFCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiniStakeFFFCount is a free data retrieval call binding the contract method 0x1faea20b.
//
// Solidity: function miniStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MiniStakeFFFCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MiniStakeFFFCount(&_BSCValidatorSet.CallOpts)
}

// MiniStakeFFFCount is a free data retrieval call binding the contract method 0x1faea20b.
//
// Solidity: function miniStakeFFFCount() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MiniStakeFFFCount() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MiniStakeFFFCount(&_BSCValidatorSet.CallOpts)
}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) MyBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "myBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) MyBalance() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MyBalance(&_BSCValidatorSet.CallOpts)
}

// MyBalance is a free data retrieval call binding the contract method 0xc9116b69.
//
// Solidity: function myBalance() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) MyBalance() (*big.Int, error) {
	return _BSCValidatorSet.Contract.MyBalance(&_BSCValidatorSet.CallOpts)
}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) NumOfJailed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "numOfJailed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) NumOfJailed() (*big.Int, error) {
	return _BSCValidatorSet.Contract.NumOfJailed(&_BSCValidatorSet.CallOpts)
}

// NumOfJailed is a free data retrieval call binding the contract method 0xdaacdb66.
//
// Solidity: function numOfJailed() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) NumOfJailed() (*big.Int, error) {
	return _BSCValidatorSet.Contract.NumOfJailed(&_BSCValidatorSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) Owner() (common.Address, error) {
	return _BSCValidatorSet.Contract.Owner(&_BSCValidatorSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) Owner() (common.Address, error) {
	return _BSCValidatorSet.Contract.Owner(&_BSCValidatorSet.CallOpts)
}

// StakeInfoIndexMap is a free data retrieval call binding the contract method 0x7dedf367.
//
// Solidity: function stakeInfoIndexMap(uint256 ) view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCaller) StakeInfoIndexMap(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "stakeInfoIndexMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeInfoIndexMap is a free data retrieval call binding the contract method 0x7dedf367.
//
// Solidity: function stakeInfoIndexMap(uint256 ) view returns(address)
func (_BSCValidatorSet *BSCValidatorSetSession) StakeInfoIndexMap(arg0 *big.Int) (common.Address, error) {
	return _BSCValidatorSet.Contract.StakeInfoIndexMap(&_BSCValidatorSet.CallOpts, arg0)
}

// StakeInfoIndexMap is a free data retrieval call binding the contract method 0x7dedf367.
//
// Solidity: function stakeInfoIndexMap(uint256 ) view returns(address)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) StakeInfoIndexMap(arg0 *big.Int) (common.Address, error) {
	return _BSCValidatorSet.Contract.StakeInfoIndexMap(&_BSCValidatorSet.CallOpts, arg0)
}

// StakeInfoMap is a free data retrieval call binding the contract method 0xc5f30db5.
//
// Solidity: function stakeInfoMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) StakeInfoMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "stakeInfoMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeInfoMap is a free data retrieval call binding the contract method 0xc5f30db5.
//
// Solidity: function stakeInfoMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) StakeInfoMap(arg0 common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.StakeInfoMap(&_BSCValidatorSet.CallOpts, arg0)
}

// StakeInfoMap is a free data retrieval call binding the contract method 0xc5f30db5.
//
// Solidity: function stakeInfoMap(address ) view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) StakeInfoMap(arg0 common.Address) (*big.Int, error) {
	return _BSCValidatorSet.Contract.StakeInfoMap(&_BSCValidatorSet.CallOpts, arg0)
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCaller) TotalInComing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCValidatorSet.contract.Call(opts, &out, "totalInComing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetSession) TotalInComing() (*big.Int, error) {
	return _BSCValidatorSet.Contract.TotalInComing(&_BSCValidatorSet.CallOpts)
}

// TotalInComing is a free data retrieval call binding the contract method 0x1ff18069.
//
// Solidity: function totalInComing() view returns(uint256)
func (_BSCValidatorSet *BSCValidatorSetCallerSession) TotalInComing() (*big.Int, error) {
	return _BSCValidatorSet.Contract.TotalInComing(&_BSCValidatorSet.CallOpts)
}

// StakeFFF is a paid mutator transaction binding the contract method 0x0db6cccb.
//
// Solidity: function StakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) StakeFFF(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "StakeFFF")
}

// StakeFFF is a paid mutator transaction binding the contract method 0x0db6cccb.
//
// Solidity: function StakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetSession) StakeFFF() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.StakeFFF(&_BSCValidatorSet.TransactOpts)
}

// StakeFFF is a paid mutator transaction binding the contract method 0x0db6cccb.
//
// Solidity: function StakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) StakeFFF() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.StakeFFF(&_BSCValidatorSet.TransactOpts)
}

// UnStakeFFF is a paid mutator transaction binding the contract method 0x30c62a13.
//
// Solidity: function UnStakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) UnStakeFFF(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "UnStakeFFF")
}

// UnStakeFFF is a paid mutator transaction binding the contract method 0x30c62a13.
//
// Solidity: function UnStakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetSession) UnStakeFFF() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.UnStakeFFF(&_BSCValidatorSet.TransactOpts)
}

// UnStakeFFF is a paid mutator transaction binding the contract method 0x30c62a13.
//
// Solidity: function UnStakeFFF() payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) UnStakeFFF() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.UnStakeFFF(&_BSCValidatorSet.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) Deposit(opts *bind.TransactOpts, valAddr common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "deposit", valAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) payable returns()
func (_BSCValidatorSet *BSCValidatorSetSession) Deposit(valAddr common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Deposit(&_BSCValidatorSet.TransactOpts, valAddr)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address valAddr) payable returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) Deposit(valAddr common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Deposit(&_BSCValidatorSet.TransactOpts, valAddr)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) Felony(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "felony", validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Felony(&_BSCValidatorSet.TransactOpts, validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Felony(&_BSCValidatorSet.TransactOpts, validator)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleFailAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleFailAckPackage(&_BSCValidatorSet.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_BSCValidatorSet *BSCValidatorSetTransactor) HandleSynPackage(opts *bind.TransactOpts, arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "handleSynPackage", arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_BSCValidatorSet *BSCValidatorSetSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleSynPackage(&_BSCValidatorSet.TransactOpts, arg0, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 , bytes msgBytes) returns(bytes responsePayload)
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) HandleSynPackage(arg0 uint8, msgBytes []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.HandleSynPackage(&_BSCValidatorSet.TransactOpts, arg0, msgBytes)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_BSCValidatorSet *BSCValidatorSetSession) Init() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Init(&_BSCValidatorSet.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) Init() (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Init(&_BSCValidatorSet.TransactOpts)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) Misdemeanor(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "misdemeanor", validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Misdemeanor(&_BSCValidatorSet.TransactOpts, validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.Misdemeanor(&_BSCValidatorSet.TransactOpts, validator)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_BSCValidatorSet *BSCValidatorSetSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.UpdateParam(&_BSCValidatorSet.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_BSCValidatorSet *BSCValidatorSetTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _BSCValidatorSet.Contract.UpdateParam(&_BSCValidatorSet.TransactOpts, key, value)
}

// BSCValidatorSetDistributeBlockRewardFuncIterator is returned from FilterDistributeBlockRewardFunc and is used to iterate over the raw logs and unpacked data for DistributeBlockRewardFunc events raised by the BSCValidatorSet contract.
type BSCValidatorSetDistributeBlockRewardFuncIterator struct {
	Event *BSCValidatorSetDistributeBlockRewardFunc // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetDistributeBlockRewardFuncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetDistributeBlockRewardFunc)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetDistributeBlockRewardFunc)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetDistributeBlockRewardFuncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetDistributeBlockRewardFuncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetDistributeBlockRewardFunc represents a DistributeBlockRewardFunc event raised by the BSCValidatorSet contract.
type BSCValidatorSetDistributeBlockRewardFunc struct {
	UserCount *big.Int
	Count     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributeBlockRewardFunc is a free log retrieval operation binding the contract event 0x77affbec109ff6a3609c802050db936c55bc48fb23db5c88f4a4818922e65239.
//
// Solidity: event DistributeBlockRewardFunc(uint256 userCount, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterDistributeBlockRewardFunc(opts *bind.FilterOpts) (*BSCValidatorSetDistributeBlockRewardFuncIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "DistributeBlockRewardFunc")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetDistributeBlockRewardFuncIterator{contract: _BSCValidatorSet.contract, event: "DistributeBlockRewardFunc", logs: logs, sub: sub}, nil
}

// WatchDistributeBlockRewardFunc is a free log subscription operation binding the contract event 0x77affbec109ff6a3609c802050db936c55bc48fb23db5c88f4a4818922e65239.
//
// Solidity: event DistributeBlockRewardFunc(uint256 userCount, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchDistributeBlockRewardFunc(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetDistributeBlockRewardFunc) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "DistributeBlockRewardFunc")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetDistributeBlockRewardFunc)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "DistributeBlockRewardFunc", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributeBlockRewardFunc is a log parse operation binding the contract event 0x77affbec109ff6a3609c802050db936c55bc48fb23db5c88f4a4818922e65239.
//
// Solidity: event DistributeBlockRewardFunc(uint256 userCount, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseDistributeBlockRewardFunc(log types.Log) (*BSCValidatorSetDistributeBlockRewardFunc, error) {
	event := new(BSCValidatorSetDistributeBlockRewardFunc)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "DistributeBlockRewardFunc", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetUserStakeFFFIterator is returned from FilterUserStakeFFF and is used to iterate over the raw logs and unpacked data for UserStakeFFF events raised by the BSCValidatorSet contract.
type BSCValidatorSetUserStakeFFFIterator struct {
	Event *BSCValidatorSetUserStakeFFF // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetUserStakeFFFIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetUserStakeFFF)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetUserStakeFFF)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetUserStakeFFFIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetUserStakeFFFIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetUserStakeFFF represents a UserStakeFFF event raised by the BSCValidatorSet contract.
type BSCValidatorSetUserStakeFFF struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserStakeFFF is a free log retrieval operation binding the contract event 0x26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f.
//
// Solidity: event UserStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterUserStakeFFF(opts *bind.FilterOpts) (*BSCValidatorSetUserStakeFFFIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "UserStakeFFF")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetUserStakeFFFIterator{contract: _BSCValidatorSet.contract, event: "UserStakeFFF", logs: logs, sub: sub}, nil
}

// WatchUserStakeFFF is a free log subscription operation binding the contract event 0x26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f.
//
// Solidity: event UserStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchUserStakeFFF(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetUserStakeFFF) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "UserStakeFFF")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetUserStakeFFF)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "UserStakeFFF", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserStakeFFF is a log parse operation binding the contract event 0x26a5c0f8e27c8ffa9a8414a94d6c0137a2ec1446b6925835dc0cf5ba74d2967f.
//
// Solidity: event UserStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseUserStakeFFF(log types.Log) (*BSCValidatorSetUserStakeFFF, error) {
	event := new(BSCValidatorSetUserStakeFFF)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "UserStakeFFF", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetUserUnStakeFFFIterator is returned from FilterUserUnStakeFFF and is used to iterate over the raw logs and unpacked data for UserUnStakeFFF events raised by the BSCValidatorSet contract.
type BSCValidatorSetUserUnStakeFFFIterator struct {
	Event *BSCValidatorSetUserUnStakeFFF // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetUserUnStakeFFFIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetUserUnStakeFFF)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetUserUnStakeFFF)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetUserUnStakeFFFIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetUserUnStakeFFFIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetUserUnStakeFFF represents a UserUnStakeFFF event raised by the BSCValidatorSet contract.
type BSCValidatorSetUserUnStakeFFF struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserUnStakeFFF is a free log retrieval operation binding the contract event 0x433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f.
//
// Solidity: event UserUnStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterUserUnStakeFFF(opts *bind.FilterOpts) (*BSCValidatorSetUserUnStakeFFFIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "UserUnStakeFFF")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetUserUnStakeFFFIterator{contract: _BSCValidatorSet.contract, event: "UserUnStakeFFF", logs: logs, sub: sub}, nil
}

// WatchUserUnStakeFFF is a free log subscription operation binding the contract event 0x433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f.
//
// Solidity: event UserUnStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchUserUnStakeFFF(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetUserUnStakeFFF) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "UserUnStakeFFF")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetUserUnStakeFFF)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "UserUnStakeFFF", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserUnStakeFFF is a log parse operation binding the contract event 0x433979128f828c623a175588cd9821124640d4353dd48d812d2b7273f1b76c3f.
//
// Solidity: event UserUnStakeFFF(address user, uint256 count)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseUserUnStakeFFF(log types.Log) (*BSCValidatorSetUserUnStakeFFF, error) {
	event := new(BSCValidatorSetUserUnStakeFFF)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "UserUnStakeFFF", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetBatchTransferIterator is returned from FilterBatchTransfer and is used to iterate over the raw logs and unpacked data for BatchTransfer events raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransferIterator struct {
	Event *BSCValidatorSetBatchTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetBatchTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetBatchTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetBatchTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetBatchTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetBatchTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetBatchTransfer represents a BatchTransfer event raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransfer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransfer is a free log retrieval operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterBatchTransfer(opts *bind.FilterOpts) (*BSCValidatorSetBatchTransferIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "batchTransfer")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetBatchTransferIterator{contract: _BSCValidatorSet.contract, event: "batchTransfer", logs: logs, sub: sub}, nil
}

// WatchBatchTransfer is a free log subscription operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchBatchTransfer(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetBatchTransfer) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "batchTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetBatchTransfer)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchTransfer is a log parse operation binding the contract event 0xa217d08e65f80c73121cd9db834d81652d544bfbf452f6d04922b16c90a37b70.
//
// Solidity: event batchTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseBatchTransfer(log types.Log) (*BSCValidatorSetBatchTransfer, error) {
	event := new(BSCValidatorSetBatchTransfer)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetBatchTransferFailedIterator is returned from FilterBatchTransferFailed and is used to iterate over the raw logs and unpacked data for BatchTransferFailed events raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransferFailedIterator struct {
	Event *BSCValidatorSetBatchTransferFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetBatchTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetBatchTransferFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetBatchTransferFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetBatchTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetBatchTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetBatchTransferFailed represents a BatchTransferFailed event raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransferFailed struct {
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransferFailed is a free log retrieval operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterBatchTransferFailed(opts *bind.FilterOpts, amount []*big.Int) (*BSCValidatorSetBatchTransferFailedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "batchTransferFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetBatchTransferFailedIterator{contract: _BSCValidatorSet.contract, event: "batchTransferFailed", logs: logs, sub: sub}, nil
}

// WatchBatchTransferFailed is a free log subscription operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchBatchTransferFailed(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetBatchTransferFailed, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "batchTransferFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetBatchTransferFailed)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransferFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchTransferFailed is a log parse operation binding the contract event 0xa7cdeed7d0db45e3219a6e5d60838824c16f1d39991fcfe3f963029c844bf280.
//
// Solidity: event batchTransferFailed(uint256 indexed amount, string reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseBatchTransferFailed(log types.Log) (*BSCValidatorSetBatchTransferFailed, error) {
	event := new(BSCValidatorSetBatchTransferFailed)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetBatchTransferLowerFailedIterator is returned from FilterBatchTransferLowerFailed and is used to iterate over the raw logs and unpacked data for BatchTransferLowerFailed events raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransferLowerFailedIterator struct {
	Event *BSCValidatorSetBatchTransferLowerFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetBatchTransferLowerFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetBatchTransferLowerFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetBatchTransferLowerFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetBatchTransferLowerFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetBatchTransferLowerFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetBatchTransferLowerFailed represents a BatchTransferLowerFailed event raised by the BSCValidatorSet contract.
type BSCValidatorSetBatchTransferLowerFailed struct {
	Amount *big.Int
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchTransferLowerFailed is a free log retrieval operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterBatchTransferLowerFailed(opts *bind.FilterOpts, amount []*big.Int) (*BSCValidatorSetBatchTransferLowerFailedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "batchTransferLowerFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetBatchTransferLowerFailedIterator{contract: _BSCValidatorSet.contract, event: "batchTransferLowerFailed", logs: logs, sub: sub}, nil
}

// WatchBatchTransferLowerFailed is a free log subscription operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchBatchTransferLowerFailed(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetBatchTransferLowerFailed, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "batchTransferLowerFailed", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetBatchTransferLowerFailed)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransferLowerFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchTransferLowerFailed is a log parse operation binding the contract event 0xbfa884552dd8921b6ce90bfe906952ae5b3b29be0cc1a951d4f62697635a3a45.
//
// Solidity: event batchTransferLowerFailed(uint256 indexed amount, bytes reason)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseBatchTransferLowerFailed(log types.Log) (*BSCValidatorSetBatchTransferLowerFailed, error) {
	event := new(BSCValidatorSetBatchTransferLowerFailed)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "batchTransferLowerFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetDeprecatedDepositIterator is returned from FilterDeprecatedDeposit and is used to iterate over the raw logs and unpacked data for DeprecatedDeposit events raised by the BSCValidatorSet contract.
type BSCValidatorSetDeprecatedDepositIterator struct {
	Event *BSCValidatorSetDeprecatedDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetDeprecatedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetDeprecatedDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetDeprecatedDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetDeprecatedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetDeprecatedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetDeprecatedDeposit represents a DeprecatedDeposit event raised by the BSCValidatorSet contract.
type BSCValidatorSetDeprecatedDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedDeposit is a free log retrieval operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterDeprecatedDeposit(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetDeprecatedDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "deprecatedDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetDeprecatedDepositIterator{contract: _BSCValidatorSet.contract, event: "deprecatedDeposit", logs: logs, sub: sub}, nil
}

// WatchDeprecatedDeposit is a free log subscription operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchDeprecatedDeposit(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetDeprecatedDeposit, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "deprecatedDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetDeprecatedDeposit)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "deprecatedDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeprecatedDeposit is a log parse operation binding the contract event 0xf177e5d6c5764d79c32883ed824111d9b13f5668cf6ab1cc12dd36791dd955b4.
//
// Solidity: event deprecatedDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseDeprecatedDeposit(log types.Log) (*BSCValidatorSetDeprecatedDeposit, error) {
	event := new(BSCValidatorSetDeprecatedDeposit)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "deprecatedDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetDirectTransferIterator is returned from FilterDirectTransfer and is used to iterate over the raw logs and unpacked data for DirectTransfer events raised by the BSCValidatorSet contract.
type BSCValidatorSetDirectTransferIterator struct {
	Event *BSCValidatorSetDirectTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetDirectTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetDirectTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetDirectTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetDirectTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetDirectTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetDirectTransfer represents a DirectTransfer event raised by the BSCValidatorSet contract.
type BSCValidatorSetDirectTransfer struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDirectTransfer is a free log retrieval operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterDirectTransfer(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetDirectTransferIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "directTransfer", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetDirectTransferIterator{contract: _BSCValidatorSet.contract, event: "directTransfer", logs: logs, sub: sub}, nil
}

// WatchDirectTransfer is a free log subscription operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchDirectTransfer(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetDirectTransfer, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "directTransfer", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetDirectTransfer)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "directTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDirectTransfer is a log parse operation binding the contract event 0x6c61d60f69a7beb3e1c80db7f39f37b208537cbb19da3174511b477812b2fc7d.
//
// Solidity: event directTransfer(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseDirectTransfer(log types.Log) (*BSCValidatorSetDirectTransfer, error) {
	event := new(BSCValidatorSetDirectTransfer)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "directTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetDirectTransferFailIterator is returned from FilterDirectTransferFail and is used to iterate over the raw logs and unpacked data for DirectTransferFail events raised by the BSCValidatorSet contract.
type BSCValidatorSetDirectTransferFailIterator struct {
	Event *BSCValidatorSetDirectTransferFail // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetDirectTransferFailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetDirectTransferFail)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetDirectTransferFail)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetDirectTransferFailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetDirectTransferFailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetDirectTransferFail represents a DirectTransferFail event raised by the BSCValidatorSet contract.
type BSCValidatorSetDirectTransferFail struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDirectTransferFail is a free log retrieval operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterDirectTransferFail(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetDirectTransferFailIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "directTransferFail", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetDirectTransferFailIterator{contract: _BSCValidatorSet.contract, event: "directTransferFail", logs: logs, sub: sub}, nil
}

// WatchDirectTransferFail is a free log subscription operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchDirectTransferFail(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetDirectTransferFail, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "directTransferFail", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetDirectTransferFail)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "directTransferFail", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDirectTransferFail is a log parse operation binding the contract event 0x25d0ce7d2f0cec669a8c17efe49d195c13455bb8872b65fa610ac7f53fe4ca7d.
//
// Solidity: event directTransferFail(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseDirectTransferFail(log types.Log) (*BSCValidatorSetDirectTransferFail, error) {
	event := new(BSCValidatorSetDirectTransferFail)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "directTransferFail", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetFailReasonWithStrIterator is returned from FilterFailReasonWithStr and is used to iterate over the raw logs and unpacked data for FailReasonWithStr events raised by the BSCValidatorSet contract.
type BSCValidatorSetFailReasonWithStrIterator struct {
	Event *BSCValidatorSetFailReasonWithStr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetFailReasonWithStrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetFailReasonWithStr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetFailReasonWithStr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetFailReasonWithStrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetFailReasonWithStrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetFailReasonWithStr represents a FailReasonWithStr event raised by the BSCValidatorSet contract.
type BSCValidatorSetFailReasonWithStr struct {
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailReasonWithStr is a free log retrieval operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterFailReasonWithStr(opts *bind.FilterOpts) (*BSCValidatorSetFailReasonWithStrIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetFailReasonWithStrIterator{contract: _BSCValidatorSet.contract, event: "failReasonWithStr", logs: logs, sub: sub}, nil
}

// WatchFailReasonWithStr is a free log subscription operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchFailReasonWithStr(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetFailReasonWithStr) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "failReasonWithStr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetFailReasonWithStr)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailReasonWithStr is a log parse operation binding the contract event 0x70e72399380dcfb0338abc03dc8d47f9f470ada8e769c9a78d644ea97385ecb2.
//
// Solidity: event failReasonWithStr(string message)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseFailReasonWithStr(log types.Log) (*BSCValidatorSetFailReasonWithStr, error) {
	event := new(BSCValidatorSetFailReasonWithStr)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "failReasonWithStr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetFeeBurnedIterator is returned from FilterFeeBurned and is used to iterate over the raw logs and unpacked data for FeeBurned events raised by the BSCValidatorSet contract.
type BSCValidatorSetFeeBurnedIterator struct {
	Event *BSCValidatorSetFeeBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetFeeBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetFeeBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetFeeBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetFeeBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetFeeBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetFeeBurned represents a FeeBurned event raised by the BSCValidatorSet contract.
type BSCValidatorSetFeeBurned struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeBurned is a free log retrieval operation binding the contract event 0x627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee5.
//
// Solidity: event feeBurned(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterFeeBurned(opts *bind.FilterOpts) (*BSCValidatorSetFeeBurnedIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "feeBurned")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetFeeBurnedIterator{contract: _BSCValidatorSet.contract, event: "feeBurned", logs: logs, sub: sub}, nil
}

// WatchFeeBurned is a free log subscription operation binding the contract event 0x627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee5.
//
// Solidity: event feeBurned(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchFeeBurned(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetFeeBurned) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "feeBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetFeeBurned)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "feeBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeBurned is a log parse operation binding the contract event 0x627059660ea01c4733a328effb2294d2f86905bf806da763a89cee254de8bee5.
//
// Solidity: event feeBurned(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseFeeBurned(log types.Log) (*BSCValidatorSetFeeBurned, error) {
	event := new(BSCValidatorSetFeeBurned)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "feeBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the BSCValidatorSet contract.
type BSCValidatorSetParamChangeIterator struct {
	Event *BSCValidatorSetParamChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetParamChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetParamChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetParamChange represents a ParamChange event raised by the BSCValidatorSet contract.
type BSCValidatorSetParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterParamChange(opts *bind.FilterOpts) (*BSCValidatorSetParamChangeIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetParamChangeIterator{contract: _BSCValidatorSet.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetParamChange) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetParamChange)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "paramChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParamChange is a log parse operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseParamChange(log types.Log) (*BSCValidatorSetParamChange, error) {
	event := new(BSCValidatorSetParamChange)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetSystemTransferIterator is returned from FilterSystemTransfer and is used to iterate over the raw logs and unpacked data for SystemTransfer events raised by the BSCValidatorSet contract.
type BSCValidatorSetSystemTransferIterator struct {
	Event *BSCValidatorSetSystemTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetSystemTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetSystemTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetSystemTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetSystemTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetSystemTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetSystemTransfer represents a SystemTransfer event raised by the BSCValidatorSet contract.
type BSCValidatorSetSystemTransfer struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSystemTransfer is a free log retrieval operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterSystemTransfer(opts *bind.FilterOpts) (*BSCValidatorSetSystemTransferIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "systemTransfer")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetSystemTransferIterator{contract: _BSCValidatorSet.contract, event: "systemTransfer", logs: logs, sub: sub}, nil
}

// WatchSystemTransfer is a free log subscription operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchSystemTransfer(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetSystemTransfer) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "systemTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetSystemTransfer)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "systemTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSystemTransfer is a log parse operation binding the contract event 0x6ecc855f9440a9282c90913bbc91619fd44f5ec0b462af28d127b116f130aa4d.
//
// Solidity: event systemTransfer(uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseSystemTransfer(log types.Log) (*BSCValidatorSetSystemTransfer, error) {
	event := new(BSCValidatorSetSystemTransfer)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "systemTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetUnexpectedPackageIterator is returned from FilterUnexpectedPackage and is used to iterate over the raw logs and unpacked data for UnexpectedPackage events raised by the BSCValidatorSet contract.
type BSCValidatorSetUnexpectedPackageIterator struct {
	Event *BSCValidatorSetUnexpectedPackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetUnexpectedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetUnexpectedPackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetUnexpectedPackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetUnexpectedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetUnexpectedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetUnexpectedPackage represents a UnexpectedPackage event raised by the BSCValidatorSet contract.
type BSCValidatorSetUnexpectedPackage struct {
	ChannelId uint8
	MsgBytes  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedPackage is a free log retrieval operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterUnexpectedPackage(opts *bind.FilterOpts) (*BSCValidatorSetUnexpectedPackageIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "unexpectedPackage")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetUnexpectedPackageIterator{contract: _BSCValidatorSet.contract, event: "unexpectedPackage", logs: logs, sub: sub}, nil
}

// WatchUnexpectedPackage is a free log subscription operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchUnexpectedPackage(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetUnexpectedPackage) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "unexpectedPackage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetUnexpectedPackage)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "unexpectedPackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnexpectedPackage is a log parse operation binding the contract event 0x41ce201247b6ceb957dcdb217d0b8acb50b9ea0e12af9af4f5e7f38902101605.
//
// Solidity: event unexpectedPackage(uint8 channelId, bytes msgBytes)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseUnexpectedPackage(log types.Log) (*BSCValidatorSetUnexpectedPackage, error) {
	event := new(BSCValidatorSetUnexpectedPackage)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "unexpectedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorDepositIterator is returned from FilterValidatorDeposit and is used to iterate over the raw logs and unpacked data for ValidatorDeposit events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorDepositIterator struct {
	Event *BSCValidatorSetValidatorDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorDeposit represents a ValidatorDeposit event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorDeposit struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeposit is a free log retrieval operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorDeposit(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetValidatorDepositIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorDepositIterator{contract: _BSCValidatorSet.contract, event: "validatorDeposit", logs: logs, sub: sub}, nil
}

// WatchValidatorDeposit is a free log subscription operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorDeposit(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorDeposit, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorDeposit", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorDeposit)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorDeposit is a log parse operation binding the contract event 0x93a090ecc682c002995fad3c85b30c5651d7fd29b0be5da9d784a3302aedc055.
//
// Solidity: event validatorDeposit(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorDeposit(log types.Log) (*BSCValidatorSetValidatorDeposit, error) {
	event := new(BSCValidatorSetValidatorDeposit)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorEmptyJailedIterator is returned from FilterValidatorEmptyJailed and is used to iterate over the raw logs and unpacked data for ValidatorEmptyJailed events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorEmptyJailedIterator struct {
	Event *BSCValidatorSetValidatorEmptyJailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorEmptyJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorEmptyJailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorEmptyJailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorEmptyJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorEmptyJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorEmptyJailed represents a ValidatorEmptyJailed event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorEmptyJailed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorEmptyJailed is a free log retrieval operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorEmptyJailed(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetValidatorEmptyJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorEmptyJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorEmptyJailedIterator{contract: _BSCValidatorSet.contract, event: "validatorEmptyJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorEmptyJailed is a free log subscription operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorEmptyJailed(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorEmptyJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorEmptyJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorEmptyJailed)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorEmptyJailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorEmptyJailed is a log parse operation binding the contract event 0xe209c46bebf57cf265d5d9009a00870e256d9150f3ed5281ab9d9eb3cec6e4be.
//
// Solidity: event validatorEmptyJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorEmptyJailed(log types.Log) (*BSCValidatorSetValidatorEmptyJailed, error) {
	event := new(BSCValidatorSetValidatorEmptyJailed)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorEmptyJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorFelonyIterator is returned from FilterValidatorFelony and is used to iterate over the raw logs and unpacked data for ValidatorFelony events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorFelonyIterator struct {
	Event *BSCValidatorSetValidatorFelony // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorFelonyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorFelony)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorFelony)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorFelonyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorFelonyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorFelony represents a ValidatorFelony event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorFelony struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorFelony is a free log retrieval operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorFelony(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetValidatorFelonyIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorFelony", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorFelonyIterator{contract: _BSCValidatorSet.contract, event: "validatorFelony", logs: logs, sub: sub}, nil
}

// WatchValidatorFelony is a free log subscription operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorFelony(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorFelony, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorFelony", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorFelony)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorFelony", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorFelony is a log parse operation binding the contract event 0x3b6f9ef90462b512a1293ecec018670bf7b7f1876fb727590a8a6d7643130a70.
//
// Solidity: event validatorFelony(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorFelony(log types.Log) (*BSCValidatorSetValidatorFelony, error) {
	event := new(BSCValidatorSetValidatorFelony)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorFelony", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorJailedIterator struct {
	Event *BSCValidatorSetValidatorJailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorJailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorJailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorJailed represents a ValidatorJailed event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorJailed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorJailed(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetValidatorJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorJailedIterator{contract: _BSCValidatorSet.contract, event: "validatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorJailed)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorJailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorJailed is a log parse operation binding the contract event 0xf226e7d8f547ff903d9d419cf5f54e0d7d07efa9584135a53a057c5f1f27f49a.
//
// Solidity: event validatorJailed(address indexed validator)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorJailed(log types.Log) (*BSCValidatorSetValidatorJailed, error) {
	event := new(BSCValidatorSetValidatorJailed)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorMisdemeanorIterator is returned from FilterValidatorMisdemeanor and is used to iterate over the raw logs and unpacked data for ValidatorMisdemeanor events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorMisdemeanorIterator struct {
	Event *BSCValidatorSetValidatorMisdemeanor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorMisdemeanorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorMisdemeanor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorMisdemeanor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorMisdemeanorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorMisdemeanorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorMisdemeanor represents a ValidatorMisdemeanor event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorMisdemeanor struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorMisdemeanor is a free log retrieval operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorMisdemeanor(opts *bind.FilterOpts, validator []common.Address) (*BSCValidatorSetValidatorMisdemeanorIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorMisdemeanor", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorMisdemeanorIterator{contract: _BSCValidatorSet.contract, event: "validatorMisdemeanor", logs: logs, sub: sub}, nil
}

// WatchValidatorMisdemeanor is a free log subscription operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorMisdemeanor(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorMisdemeanor, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorMisdemeanor", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorMisdemeanor)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorMisdemeanor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorMisdemeanor is a log parse operation binding the contract event 0x8cd4e147d8af98a9e3b6724021b8bf6aed2e5dac71c38f2dce8161b82585b25d.
//
// Solidity: event validatorMisdemeanor(address indexed validator, uint256 amount)
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorMisdemeanor(log types.Log) (*BSCValidatorSetValidatorMisdemeanor, error) {
	event := new(BSCValidatorSetValidatorMisdemeanor)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorMisdemeanor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCValidatorSetValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorSetUpdatedIterator struct {
	Event *BSCValidatorSetValidatorSetUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BSCValidatorSetValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCValidatorSetValidatorSetUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BSCValidatorSetValidatorSetUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BSCValidatorSetValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCValidatorSetValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCValidatorSetValidatorSetUpdated represents a ValidatorSetUpdated event raised by the BSCValidatorSet contract.
type BSCValidatorSetValidatorSetUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_BSCValidatorSet *BSCValidatorSetFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts) (*BSCValidatorSetValidatorSetUpdatedIterator, error) {

	logs, sub, err := _BSCValidatorSet.contract.FilterLogs(opts, "validatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return &BSCValidatorSetValidatorSetUpdatedIterator{contract: _BSCValidatorSet.contract, event: "validatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_BSCValidatorSet *BSCValidatorSetFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *BSCValidatorSetValidatorSetUpdated) (event.Subscription, error) {

	logs, sub, err := _BSCValidatorSet.contract.WatchLogs(opts, "validatorSetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCValidatorSetValidatorSetUpdated)
				if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorSetUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0xedd8d7296956dd970ab4de3f2fc03be2b0ffc615d20cd4c72c6e44f928630ebf.
//
// Solidity: event validatorSetUpdated()
func (_BSCValidatorSet *BSCValidatorSetFilterer) ParseValidatorSetUpdated(log types.Log) (*BSCValidatorSetValidatorSetUpdated, error) {
	event := new(BSCValidatorSetValidatorSetUpdated)
	if err := _BSCValidatorSet.contract.UnpackLog(event, "validatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BytesToTypesABI is the input ABI used to generate the binding from.
const BytesToTypesABI = "[]"

// BytesToTypesBin is the compiled bytecode used for deploying new contracts.
var BytesToTypesBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a33734ce4b05e2a4f24fb6a43c76c2a990deec01482ee930645ff1694633109f64736f6c63430006040033"

// DeployBytesToTypes deploys a new Ethereum contract, binding an instance of BytesToTypes to it.
func DeployBytesToTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesToTypes, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesToTypesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesToTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesToTypes{BytesToTypesCaller: BytesToTypesCaller{contract: contract}, BytesToTypesTransactor: BytesToTypesTransactor{contract: contract}, BytesToTypesFilterer: BytesToTypesFilterer{contract: contract}}, nil
}

// BytesToTypes is an auto generated Go binding around an Ethereum contract.
type BytesToTypes struct {
	BytesToTypesCaller     // Read-only binding to the contract
	BytesToTypesTransactor // Write-only binding to the contract
	BytesToTypesFilterer   // Log filterer for contract events
}

// BytesToTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesToTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesToTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesToTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesToTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesToTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesToTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesToTypesSession struct {
	Contract     *BytesToTypes     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesToTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesToTypesCallerSession struct {
	Contract *BytesToTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BytesToTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesToTypesTransactorSession struct {
	Contract     *BytesToTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BytesToTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesToTypesRaw struct {
	Contract *BytesToTypes // Generic contract binding to access the raw methods on
}

// BytesToTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesToTypesCallerRaw struct {
	Contract *BytesToTypesCaller // Generic read-only contract binding to access the raw methods on
}

// BytesToTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesToTypesTransactorRaw struct {
	Contract *BytesToTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesToTypes creates a new instance of BytesToTypes, bound to a specific deployed contract.
func NewBytesToTypes(address common.Address, backend bind.ContractBackend) (*BytesToTypes, error) {
	contract, err := bindBytesToTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesToTypes{BytesToTypesCaller: BytesToTypesCaller{contract: contract}, BytesToTypesTransactor: BytesToTypesTransactor{contract: contract}, BytesToTypesFilterer: BytesToTypesFilterer{contract: contract}}, nil
}

// NewBytesToTypesCaller creates a new read-only instance of BytesToTypes, bound to a specific deployed contract.
func NewBytesToTypesCaller(address common.Address, caller bind.ContractCaller) (*BytesToTypesCaller, error) {
	contract, err := bindBytesToTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesToTypesCaller{contract: contract}, nil
}

// NewBytesToTypesTransactor creates a new write-only instance of BytesToTypes, bound to a specific deployed contract.
func NewBytesToTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesToTypesTransactor, error) {
	contract, err := bindBytesToTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesToTypesTransactor{contract: contract}, nil
}

// NewBytesToTypesFilterer creates a new log filterer instance of BytesToTypes, bound to a specific deployed contract.
func NewBytesToTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesToTypesFilterer, error) {
	contract, err := bindBytesToTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesToTypesFilterer{contract: contract}, nil
}

// bindBytesToTypes binds a generic wrapper to an already deployed contract.
func bindBytesToTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesToTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesToTypes *BytesToTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesToTypes.Contract.BytesToTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesToTypes *BytesToTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesToTypes.Contract.BytesToTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesToTypes *BytesToTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesToTypes.Contract.BytesToTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesToTypes *BytesToTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesToTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesToTypes *BytesToTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesToTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesToTypes *BytesToTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesToTypes.Contract.contract.Transact(opts, method, params...)
}

// CmnPkgABI is the input ABI used to generate the binding from.
const CmnPkgABI = "[]"

// CmnPkgBin is the compiled bytecode used for deploying new contracts.
var CmnPkgBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122013cba6ca7d42b9596ab731c4e5624b92170ac09ff98e457a52610505ed13982a64736f6c63430006040033"

// DeployCmnPkg deploys a new Ethereum contract, binding an instance of CmnPkg to it.
func DeployCmnPkg(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CmnPkg, error) {
	parsed, err := abi.JSON(strings.NewReader(CmnPkgABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CmnPkgBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CmnPkg{CmnPkgCaller: CmnPkgCaller{contract: contract}, CmnPkgTransactor: CmnPkgTransactor{contract: contract}, CmnPkgFilterer: CmnPkgFilterer{contract: contract}}, nil
}

// CmnPkg is an auto generated Go binding around an Ethereum contract.
type CmnPkg struct {
	CmnPkgCaller     // Read-only binding to the contract
	CmnPkgTransactor // Write-only binding to the contract
	CmnPkgFilterer   // Log filterer for contract events
}

// CmnPkgCaller is an auto generated read-only Go binding around an Ethereum contract.
type CmnPkgCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmnPkgTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CmnPkgTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmnPkgFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CmnPkgFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CmnPkgSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CmnPkgSession struct {
	Contract     *CmnPkg           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CmnPkgCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CmnPkgCallerSession struct {
	Contract *CmnPkgCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CmnPkgTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CmnPkgTransactorSession struct {
	Contract     *CmnPkgTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CmnPkgRaw is an auto generated low-level Go binding around an Ethereum contract.
type CmnPkgRaw struct {
	Contract *CmnPkg // Generic contract binding to access the raw methods on
}

// CmnPkgCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CmnPkgCallerRaw struct {
	Contract *CmnPkgCaller // Generic read-only contract binding to access the raw methods on
}

// CmnPkgTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CmnPkgTransactorRaw struct {
	Contract *CmnPkgTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCmnPkg creates a new instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkg(address common.Address, backend bind.ContractBackend) (*CmnPkg, error) {
	contract, err := bindCmnPkg(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CmnPkg{CmnPkgCaller: CmnPkgCaller{contract: contract}, CmnPkgTransactor: CmnPkgTransactor{contract: contract}, CmnPkgFilterer: CmnPkgFilterer{contract: contract}}, nil
}

// NewCmnPkgCaller creates a new read-only instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkgCaller(address common.Address, caller bind.ContractCaller) (*CmnPkgCaller, error) {
	contract, err := bindCmnPkg(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CmnPkgCaller{contract: contract}, nil
}

// NewCmnPkgTransactor creates a new write-only instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkgTransactor(address common.Address, transactor bind.ContractTransactor) (*CmnPkgTransactor, error) {
	contract, err := bindCmnPkg(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CmnPkgTransactor{contract: contract}, nil
}

// NewCmnPkgFilterer creates a new log filterer instance of CmnPkg, bound to a specific deployed contract.
func NewCmnPkgFilterer(address common.Address, filterer bind.ContractFilterer) (*CmnPkgFilterer, error) {
	contract, err := bindCmnPkg(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CmnPkgFilterer{contract: contract}, nil
}

// bindCmnPkg binds a generic wrapper to an already deployed contract.
func bindCmnPkg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CmnPkgABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CmnPkg *CmnPkgRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CmnPkg.Contract.CmnPkgCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CmnPkg *CmnPkgRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CmnPkg.Contract.CmnPkgTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CmnPkg *CmnPkgRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CmnPkg.Contract.CmnPkgTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CmnPkg *CmnPkgCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CmnPkg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CmnPkg *CmnPkgTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CmnPkg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CmnPkg *CmnPkgTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CmnPkg.Contract.contract.Transact(opts, method, params...)
}

// IApplicationABI is the input ABI used to generate the binding from.
const IApplicationABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleFailAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"responsePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IApplicationFuncSigs maps the 4-byte function signature to its string representation.
var IApplicationFuncSigs = map[string]string{
	"831d65d1": "handleAckPackage(uint8,bytes)",
	"c8509d81": "handleFailAckPackage(uint8,bytes)",
	"1182b875": "handleSynPackage(uint8,bytes)",
}

// IApplication is an auto generated Go binding around an Ethereum contract.
type IApplication struct {
	IApplicationCaller     // Read-only binding to the contract
	IApplicationTransactor // Write-only binding to the contract
	IApplicationFilterer   // Log filterer for contract events
}

// IApplicationCaller is an auto generated read-only Go binding around an Ethereum contract.
type IApplicationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApplicationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IApplicationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApplicationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IApplicationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IApplicationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IApplicationSession struct {
	Contract     *IApplication     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IApplicationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IApplicationCallerSession struct {
	Contract *IApplicationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IApplicationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IApplicationTransactorSession struct {
	Contract     *IApplicationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IApplicationRaw is an auto generated low-level Go binding around an Ethereum contract.
type IApplicationRaw struct {
	Contract *IApplication // Generic contract binding to access the raw methods on
}

// IApplicationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IApplicationCallerRaw struct {
	Contract *IApplicationCaller // Generic read-only contract binding to access the raw methods on
}

// IApplicationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IApplicationTransactorRaw struct {
	Contract *IApplicationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIApplication creates a new instance of IApplication, bound to a specific deployed contract.
func NewIApplication(address common.Address, backend bind.ContractBackend) (*IApplication, error) {
	contract, err := bindIApplication(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IApplication{IApplicationCaller: IApplicationCaller{contract: contract}, IApplicationTransactor: IApplicationTransactor{contract: contract}, IApplicationFilterer: IApplicationFilterer{contract: contract}}, nil
}

// NewIApplicationCaller creates a new read-only instance of IApplication, bound to a specific deployed contract.
func NewIApplicationCaller(address common.Address, caller bind.ContractCaller) (*IApplicationCaller, error) {
	contract, err := bindIApplication(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IApplicationCaller{contract: contract}, nil
}

// NewIApplicationTransactor creates a new write-only instance of IApplication, bound to a specific deployed contract.
func NewIApplicationTransactor(address common.Address, transactor bind.ContractTransactor) (*IApplicationTransactor, error) {
	contract, err := bindIApplication(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IApplicationTransactor{contract: contract}, nil
}

// NewIApplicationFilterer creates a new log filterer instance of IApplication, bound to a specific deployed contract.
func NewIApplicationFilterer(address common.Address, filterer bind.ContractFilterer) (*IApplicationFilterer, error) {
	contract, err := bindIApplication(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IApplicationFilterer{contract: contract}, nil
}

// bindIApplication binds a generic wrapper to an already deployed contract.
func bindIApplication(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IApplicationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApplication *IApplicationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApplication.Contract.IApplicationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApplication *IApplicationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApplication.Contract.IApplicationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApplication *IApplicationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApplication.Contract.IApplicationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IApplication *IApplicationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IApplication.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IApplication *IApplicationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IApplication.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IApplication *IApplicationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IApplication.Contract.contract.Transact(opts, method, params...)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleFailAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_IApplication *IApplicationTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleFailAckPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes responsePayload)
func (_IApplication *IApplicationTransactor) HandleSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.contract.Transact(opts, "handleSynPackage", channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes responsePayload)
func (_IApplication *IApplicationSession) HandleSynPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleSynPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes responsePayload)
func (_IApplication *IApplicationTransactorSession) HandleSynPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _IApplication.Contract.HandleSynPackage(&_IApplication.TransactOpts, channelId, msgBytes)
}

// IBSCValidatorSetABI is the input ABI used to generate the binding from.
const IBSCValidatorSetABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"felony\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"misdemeanor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBSCValidatorSetFuncSigs maps the 4-byte function signature to its string representation.
var IBSCValidatorSetFuncSigs = map[string]string{
	"35409f7f": "felony(address)",
	"eb57e202": "misdemeanor(address)",
}

// IBSCValidatorSet is an auto generated Go binding around an Ethereum contract.
type IBSCValidatorSet struct {
	IBSCValidatorSetCaller     // Read-only binding to the contract
	IBSCValidatorSetTransactor // Write-only binding to the contract
	IBSCValidatorSetFilterer   // Log filterer for contract events
}

// IBSCValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBSCValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSCValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBSCValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSCValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBSCValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSCValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBSCValidatorSetSession struct {
	Contract     *IBSCValidatorSet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBSCValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBSCValidatorSetCallerSession struct {
	Contract *IBSCValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IBSCValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBSCValidatorSetTransactorSession struct {
	Contract     *IBSCValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IBSCValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBSCValidatorSetRaw struct {
	Contract *IBSCValidatorSet // Generic contract binding to access the raw methods on
}

// IBSCValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBSCValidatorSetCallerRaw struct {
	Contract *IBSCValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// IBSCValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBSCValidatorSetTransactorRaw struct {
	Contract *IBSCValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBSCValidatorSet creates a new instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSet(address common.Address, backend bind.ContractBackend) (*IBSCValidatorSet, error) {
	contract, err := bindIBSCValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSet{IBSCValidatorSetCaller: IBSCValidatorSetCaller{contract: contract}, IBSCValidatorSetTransactor: IBSCValidatorSetTransactor{contract: contract}, IBSCValidatorSetFilterer: IBSCValidatorSetFilterer{contract: contract}}, nil
}

// NewIBSCValidatorSetCaller creates a new read-only instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*IBSCValidatorSetCaller, error) {
	contract, err := bindIBSCValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSetCaller{contract: contract}, nil
}

// NewIBSCValidatorSetTransactor creates a new write-only instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*IBSCValidatorSetTransactor, error) {
	contract, err := bindIBSCValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSetTransactor{contract: contract}, nil
}

// NewIBSCValidatorSetFilterer creates a new log filterer instance of IBSCValidatorSet, bound to a specific deployed contract.
func NewIBSCValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*IBSCValidatorSetFilterer, error) {
	contract, err := bindIBSCValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBSCValidatorSetFilterer{contract: contract}, nil
}

// bindIBSCValidatorSet binds a generic wrapper to an already deployed contract.
func bindIBSCValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBSCValidatorSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBSCValidatorSet *IBSCValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBSCValidatorSet.Contract.IBSCValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBSCValidatorSet *IBSCValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.IBSCValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBSCValidatorSet *IBSCValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.IBSCValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBSCValidatorSet *IBSCValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBSCValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBSCValidatorSet *IBSCValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBSCValidatorSet *IBSCValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetTransactor) Felony(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.contract.Transact(opts, "felony", validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.Felony(&_IBSCValidatorSet.TransactOpts, validator)
}

// Felony is a paid mutator transaction binding the contract method 0x35409f7f.
//
// Solidity: function felony(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetTransactorSession) Felony(validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.Felony(&_IBSCValidatorSet.TransactOpts, validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetTransactor) Misdemeanor(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.contract.Transact(opts, "misdemeanor", validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.Misdemeanor(&_IBSCValidatorSet.TransactOpts, validator)
}

// Misdemeanor is a paid mutator transaction binding the contract method 0xeb57e202.
//
// Solidity: function misdemeanor(address validator) returns()
func (_IBSCValidatorSet *IBSCValidatorSetTransactorSession) Misdemeanor(validator common.Address) (*types.Transaction, error) {
	return _IBSCValidatorSet.Contract.Misdemeanor(&_IBSCValidatorSet.TransactOpts, validator)
}

// ILightClientABI is the input ABI used to generate the binding from.
const ILightClientABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getAppHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getSubmitter\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"isHeaderSynced\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ILightClientFuncSigs maps the 4-byte function signature to its string representation.
var ILightClientFuncSigs = map[string]string{
	"cba510a9": "getAppHash(uint64)",
	"dda83148": "getSubmitter(uint64)",
	"df5fe704": "isHeaderSynced(uint64)",
}

// ILightClient is an auto generated Go binding around an Ethereum contract.
type ILightClient struct {
	ILightClientCaller     // Read-only binding to the contract
	ILightClientTransactor // Write-only binding to the contract
	ILightClientFilterer   // Log filterer for contract events
}

// ILightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILightClientSession struct {
	Contract     *ILightClient     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILightClientCallerSession struct {
	Contract *ILightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ILightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILightClientTransactorSession struct {
	Contract     *ILightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ILightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILightClientRaw struct {
	Contract *ILightClient // Generic contract binding to access the raw methods on
}

// ILightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILightClientCallerRaw struct {
	Contract *ILightClientCaller // Generic read-only contract binding to access the raw methods on
}

// ILightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILightClientTransactorRaw struct {
	Contract *ILightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILightClient creates a new instance of ILightClient, bound to a specific deployed contract.
func NewILightClient(address common.Address, backend bind.ContractBackend) (*ILightClient, error) {
	contract, err := bindILightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILightClient{ILightClientCaller: ILightClientCaller{contract: contract}, ILightClientTransactor: ILightClientTransactor{contract: contract}, ILightClientFilterer: ILightClientFilterer{contract: contract}}, nil
}

// NewILightClientCaller creates a new read-only instance of ILightClient, bound to a specific deployed contract.
func NewILightClientCaller(address common.Address, caller bind.ContractCaller) (*ILightClientCaller, error) {
	contract, err := bindILightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILightClientCaller{contract: contract}, nil
}

// NewILightClientTransactor creates a new write-only instance of ILightClient, bound to a specific deployed contract.
func NewILightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*ILightClientTransactor, error) {
	contract, err := bindILightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILightClientTransactor{contract: contract}, nil
}

// NewILightClientFilterer creates a new log filterer instance of ILightClient, bound to a specific deployed contract.
func NewILightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*ILightClientFilterer, error) {
	contract, err := bindILightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILightClientFilterer{contract: contract}, nil
}

// bindILightClient binds a generic wrapper to an already deployed contract.
func bindILightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILightClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILightClient *ILightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILightClient.Contract.ILightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILightClient *ILightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILightClient.Contract.ILightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILightClient *ILightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILightClient.Contract.ILightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILightClient *ILightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILightClient *ILightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILightClient *ILightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILightClient.Contract.contract.Transact(opts, method, params...)
}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_ILightClient *ILightClientCaller) GetAppHash(opts *bind.CallOpts, height uint64) ([32]byte, error) {
	var out []interface{}
	err := _ILightClient.contract.Call(opts, &out, "getAppHash", height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_ILightClient *ILightClientSession) GetAppHash(height uint64) ([32]byte, error) {
	return _ILightClient.Contract.GetAppHash(&_ILightClient.CallOpts, height)
}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_ILightClient *ILightClientCallerSession) GetAppHash(height uint64) ([32]byte, error) {
	return _ILightClient.Contract.GetAppHash(&_ILightClient.CallOpts, height)
}

// GetSubmitter is a free data retrieval call binding the contract method 0xdda83148.
//
// Solidity: function getSubmitter(uint64 height) view returns(address)
func (_ILightClient *ILightClientCaller) GetSubmitter(opts *bind.CallOpts, height uint64) (common.Address, error) {
	var out []interface{}
	err := _ILightClient.contract.Call(opts, &out, "getSubmitter", height)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubmitter is a free data retrieval call binding the contract method 0xdda83148.
//
// Solidity: function getSubmitter(uint64 height) view returns(address)
func (_ILightClient *ILightClientSession) GetSubmitter(height uint64) (common.Address, error) {
	return _ILightClient.Contract.GetSubmitter(&_ILightClient.CallOpts, height)
}

// GetSubmitter is a free data retrieval call binding the contract method 0xdda83148.
//
// Solidity: function getSubmitter(uint64 height) view returns(address)
func (_ILightClient *ILightClientCallerSession) GetSubmitter(height uint64) (common.Address, error) {
	return _ILightClient.Contract.GetSubmitter(&_ILightClient.CallOpts, height)
}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_ILightClient *ILightClientCaller) IsHeaderSynced(opts *bind.CallOpts, height uint64) (bool, error) {
	var out []interface{}
	err := _ILightClient.contract.Call(opts, &out, "isHeaderSynced", height)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_ILightClient *ILightClientSession) IsHeaderSynced(height uint64) (bool, error) {
	return _ILightClient.Contract.IsHeaderSynced(&_ILightClient.CallOpts, height)
}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_ILightClient *ILightClientCallerSession) IsHeaderSynced(height uint64) (bool, error) {
	return _ILightClient.Contract.IsHeaderSynced(&_ILightClient.CallOpts, height)
}

// IParamSubscriberABI is the input ABI used to generate the binding from.
const IParamSubscriberABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IParamSubscriberFuncSigs maps the 4-byte function signature to its string representation.
var IParamSubscriberFuncSigs = map[string]string{
	"ac431751": "updateParam(string,bytes)",
}

// IParamSubscriber is an auto generated Go binding around an Ethereum contract.
type IParamSubscriber struct {
	IParamSubscriberCaller     // Read-only binding to the contract
	IParamSubscriberTransactor // Write-only binding to the contract
	IParamSubscriberFilterer   // Log filterer for contract events
}

// IParamSubscriberCaller is an auto generated read-only Go binding around an Ethereum contract.
type IParamSubscriberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParamSubscriberTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IParamSubscriberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParamSubscriberFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IParamSubscriberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParamSubscriberSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IParamSubscriberSession struct {
	Contract     *IParamSubscriber // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IParamSubscriberCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IParamSubscriberCallerSession struct {
	Contract *IParamSubscriberCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IParamSubscriberTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IParamSubscriberTransactorSession struct {
	Contract     *IParamSubscriberTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IParamSubscriberRaw is an auto generated low-level Go binding around an Ethereum contract.
type IParamSubscriberRaw struct {
	Contract *IParamSubscriber // Generic contract binding to access the raw methods on
}

// IParamSubscriberCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IParamSubscriberCallerRaw struct {
	Contract *IParamSubscriberCaller // Generic read-only contract binding to access the raw methods on
}

// IParamSubscriberTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IParamSubscriberTransactorRaw struct {
	Contract *IParamSubscriberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIParamSubscriber creates a new instance of IParamSubscriber, bound to a specific deployed contract.
func NewIParamSubscriber(address common.Address, backend bind.ContractBackend) (*IParamSubscriber, error) {
	contract, err := bindIParamSubscriber(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IParamSubscriber{IParamSubscriberCaller: IParamSubscriberCaller{contract: contract}, IParamSubscriberTransactor: IParamSubscriberTransactor{contract: contract}, IParamSubscriberFilterer: IParamSubscriberFilterer{contract: contract}}, nil
}

// NewIParamSubscriberCaller creates a new read-only instance of IParamSubscriber, bound to a specific deployed contract.
func NewIParamSubscriberCaller(address common.Address, caller bind.ContractCaller) (*IParamSubscriberCaller, error) {
	contract, err := bindIParamSubscriber(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IParamSubscriberCaller{contract: contract}, nil
}

// NewIParamSubscriberTransactor creates a new write-only instance of IParamSubscriber, bound to a specific deployed contract.
func NewIParamSubscriberTransactor(address common.Address, transactor bind.ContractTransactor) (*IParamSubscriberTransactor, error) {
	contract, err := bindIParamSubscriber(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IParamSubscriberTransactor{contract: contract}, nil
}

// NewIParamSubscriberFilterer creates a new log filterer instance of IParamSubscriber, bound to a specific deployed contract.
func NewIParamSubscriberFilterer(address common.Address, filterer bind.ContractFilterer) (*IParamSubscriberFilterer, error) {
	contract, err := bindIParamSubscriber(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IParamSubscriberFilterer{contract: contract}, nil
}

// bindIParamSubscriber binds a generic wrapper to an already deployed contract.
func bindIParamSubscriber(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IParamSubscriberABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IParamSubscriber *IParamSubscriberRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IParamSubscriber.Contract.IParamSubscriberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IParamSubscriber *IParamSubscriberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.IParamSubscriberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IParamSubscriber *IParamSubscriberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.IParamSubscriberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IParamSubscriber *IParamSubscriberCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IParamSubscriber.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IParamSubscriber *IParamSubscriberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IParamSubscriber *IParamSubscriberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.contract.Transact(opts, method, params...)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_IParamSubscriber *IParamSubscriberTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _IParamSubscriber.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_IParamSubscriber *IParamSubscriberSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.UpdateParam(&_IParamSubscriber.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_IParamSubscriber *IParamSubscriberTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _IParamSubscriber.Contract.UpdateParam(&_IParamSubscriber.TransactOpts, key, value)
}

// IRelayerHubABI is the input ABI used to generate the binding from.
const IRelayerHubABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IRelayerHubFuncSigs maps the 4-byte function signature to its string representation.
var IRelayerHubFuncSigs = map[string]string{
	"541d5548": "isRelayer(address)",
}

// IRelayerHub is an auto generated Go binding around an Ethereum contract.
type IRelayerHub struct {
	IRelayerHubCaller     // Read-only binding to the contract
	IRelayerHubTransactor // Write-only binding to the contract
	IRelayerHubFilterer   // Log filterer for contract events
}

// IRelayerHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRelayerHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRelayerHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRelayerHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRelayerHubSession struct {
	Contract     *IRelayerHub      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRelayerHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRelayerHubCallerSession struct {
	Contract *IRelayerHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IRelayerHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRelayerHubTransactorSession struct {
	Contract     *IRelayerHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IRelayerHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRelayerHubRaw struct {
	Contract *IRelayerHub // Generic contract binding to access the raw methods on
}

// IRelayerHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRelayerHubCallerRaw struct {
	Contract *IRelayerHubCaller // Generic read-only contract binding to access the raw methods on
}

// IRelayerHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRelayerHubTransactorRaw struct {
	Contract *IRelayerHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRelayerHub creates a new instance of IRelayerHub, bound to a specific deployed contract.
func NewIRelayerHub(address common.Address, backend bind.ContractBackend) (*IRelayerHub, error) {
	contract, err := bindIRelayerHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRelayerHub{IRelayerHubCaller: IRelayerHubCaller{contract: contract}, IRelayerHubTransactor: IRelayerHubTransactor{contract: contract}, IRelayerHubFilterer: IRelayerHubFilterer{contract: contract}}, nil
}

// NewIRelayerHubCaller creates a new read-only instance of IRelayerHub, bound to a specific deployed contract.
func NewIRelayerHubCaller(address common.Address, caller bind.ContractCaller) (*IRelayerHubCaller, error) {
	contract, err := bindIRelayerHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerHubCaller{contract: contract}, nil
}

// NewIRelayerHubTransactor creates a new write-only instance of IRelayerHub, bound to a specific deployed contract.
func NewIRelayerHubTransactor(address common.Address, transactor bind.ContractTransactor) (*IRelayerHubTransactor, error) {
	contract, err := bindIRelayerHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerHubTransactor{contract: contract}, nil
}

// NewIRelayerHubFilterer creates a new log filterer instance of IRelayerHub, bound to a specific deployed contract.
func NewIRelayerHubFilterer(address common.Address, filterer bind.ContractFilterer) (*IRelayerHubFilterer, error) {
	contract, err := bindIRelayerHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRelayerHubFilterer{contract: contract}, nil
}

// bindIRelayerHub binds a generic wrapper to an already deployed contract.
func bindIRelayerHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRelayerHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayerHub *IRelayerHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayerHub.Contract.IRelayerHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayerHub *IRelayerHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayerHub.Contract.IRelayerHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayerHub *IRelayerHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayerHub.Contract.IRelayerHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayerHub *IRelayerHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayerHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayerHub *IRelayerHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayerHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayerHub *IRelayerHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayerHub.Contract.contract.Transact(opts, method, params...)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_IRelayerHub *IRelayerHubCaller) IsRelayer(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _IRelayerHub.contract.Call(opts, &out, "isRelayer", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_IRelayerHub *IRelayerHubSession) IsRelayer(sender common.Address) (bool, error) {
	return _IRelayerHub.Contract.IsRelayer(&_IRelayerHub.CallOpts, sender)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_IRelayerHub *IRelayerHubCallerSession) IsRelayer(sender common.Address) (bool, error) {
	return _IRelayerHub.Contract.IsRelayer(&_IRelayerHub.CallOpts, sender)
}

// ISlashIndicatorABI is the input ABI used to generate the binding from.
const ISlashIndicatorABI = "[{\"inputs\":[],\"name\":\"clean\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISlashIndicatorFuncSigs maps the 4-byte function signature to its string representation.
var ISlashIndicatorFuncSigs = map[string]string{
	"fc4333cd": "clean()",
}

// ISlashIndicator is an auto generated Go binding around an Ethereum contract.
type ISlashIndicator struct {
	ISlashIndicatorCaller     // Read-only binding to the contract
	ISlashIndicatorTransactor // Write-only binding to the contract
	ISlashIndicatorFilterer   // Log filterer for contract events
}

// ISlashIndicatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISlashIndicatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashIndicatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISlashIndicatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashIndicatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISlashIndicatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashIndicatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISlashIndicatorSession struct {
	Contract     *ISlashIndicator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISlashIndicatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISlashIndicatorCallerSession struct {
	Contract *ISlashIndicatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISlashIndicatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISlashIndicatorTransactorSession struct {
	Contract     *ISlashIndicatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISlashIndicatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISlashIndicatorRaw struct {
	Contract *ISlashIndicator // Generic contract binding to access the raw methods on
}

// ISlashIndicatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISlashIndicatorCallerRaw struct {
	Contract *ISlashIndicatorCaller // Generic read-only contract binding to access the raw methods on
}

// ISlashIndicatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISlashIndicatorTransactorRaw struct {
	Contract *ISlashIndicatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISlashIndicator creates a new instance of ISlashIndicator, bound to a specific deployed contract.
func NewISlashIndicator(address common.Address, backend bind.ContractBackend) (*ISlashIndicator, error) {
	contract, err := bindISlashIndicator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISlashIndicator{ISlashIndicatorCaller: ISlashIndicatorCaller{contract: contract}, ISlashIndicatorTransactor: ISlashIndicatorTransactor{contract: contract}, ISlashIndicatorFilterer: ISlashIndicatorFilterer{contract: contract}}, nil
}

// NewISlashIndicatorCaller creates a new read-only instance of ISlashIndicator, bound to a specific deployed contract.
func NewISlashIndicatorCaller(address common.Address, caller bind.ContractCaller) (*ISlashIndicatorCaller, error) {
	contract, err := bindISlashIndicator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISlashIndicatorCaller{contract: contract}, nil
}

// NewISlashIndicatorTransactor creates a new write-only instance of ISlashIndicator, bound to a specific deployed contract.
func NewISlashIndicatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ISlashIndicatorTransactor, error) {
	contract, err := bindISlashIndicator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISlashIndicatorTransactor{contract: contract}, nil
}

// NewISlashIndicatorFilterer creates a new log filterer instance of ISlashIndicator, bound to a specific deployed contract.
func NewISlashIndicatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ISlashIndicatorFilterer, error) {
	contract, err := bindISlashIndicator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISlashIndicatorFilterer{contract: contract}, nil
}

// bindISlashIndicator binds a generic wrapper to an already deployed contract.
func bindISlashIndicator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISlashIndicatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlashIndicator *ISlashIndicatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlashIndicator.Contract.ISlashIndicatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlashIndicator *ISlashIndicatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashIndicator.Contract.ISlashIndicatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlashIndicator *ISlashIndicatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlashIndicator.Contract.ISlashIndicatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlashIndicator *ISlashIndicatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlashIndicator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlashIndicator *ISlashIndicatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashIndicator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlashIndicator *ISlashIndicatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlashIndicator.Contract.contract.Transact(opts, method, params...)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_ISlashIndicator *ISlashIndicatorTransactor) Clean(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashIndicator.contract.Transact(opts, "clean")
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_ISlashIndicator *ISlashIndicatorSession) Clean() (*types.Transaction, error) {
	return _ISlashIndicator.Contract.Clean(&_ISlashIndicator.TransactOpts)
}

// Clean is a paid mutator transaction binding the contract method 0xfc4333cd.
//
// Solidity: function clean() returns()
func (_ISlashIndicator *ISlashIndicatorTransactorSession) Clean() (*types.Transaction, error) {
	return _ISlashIndicator.Contract.Clean(&_ISlashIndicator.TransactOpts)
}

// ISystemRewardABI is the input ABI used to generate the binding from.
const ISystemRewardABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISystemRewardFuncSigs maps the 4-byte function signature to its string representation.
var ISystemRewardFuncSigs = map[string]string{
	"9a99b4f0": "claimRewards(address,uint256)",
}

// ISystemReward is an auto generated Go binding around an Ethereum contract.
type ISystemReward struct {
	ISystemRewardCaller     // Read-only binding to the contract
	ISystemRewardTransactor // Write-only binding to the contract
	ISystemRewardFilterer   // Log filterer for contract events
}

// ISystemRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISystemRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISystemRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISystemRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISystemRewardSession struct {
	Contract     *ISystemReward    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISystemRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISystemRewardCallerSession struct {
	Contract *ISystemRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISystemRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISystemRewardTransactorSession struct {
	Contract     *ISystemRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISystemRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISystemRewardRaw struct {
	Contract *ISystemReward // Generic contract binding to access the raw methods on
}

// ISystemRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISystemRewardCallerRaw struct {
	Contract *ISystemRewardCaller // Generic read-only contract binding to access the raw methods on
}

// ISystemRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISystemRewardTransactorRaw struct {
	Contract *ISystemRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISystemReward creates a new instance of ISystemReward, bound to a specific deployed contract.
func NewISystemReward(address common.Address, backend bind.ContractBackend) (*ISystemReward, error) {
	contract, err := bindISystemReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISystemReward{ISystemRewardCaller: ISystemRewardCaller{contract: contract}, ISystemRewardTransactor: ISystemRewardTransactor{contract: contract}, ISystemRewardFilterer: ISystemRewardFilterer{contract: contract}}, nil
}

// NewISystemRewardCaller creates a new read-only instance of ISystemReward, bound to a specific deployed contract.
func NewISystemRewardCaller(address common.Address, caller bind.ContractCaller) (*ISystemRewardCaller, error) {
	contract, err := bindISystemReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemRewardCaller{contract: contract}, nil
}

// NewISystemRewardTransactor creates a new write-only instance of ISystemReward, bound to a specific deployed contract.
func NewISystemRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*ISystemRewardTransactor, error) {
	contract, err := bindISystemReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemRewardTransactor{contract: contract}, nil
}

// NewISystemRewardFilterer creates a new log filterer instance of ISystemReward, bound to a specific deployed contract.
func NewISystemRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*ISystemRewardFilterer, error) {
	contract, err := bindISystemReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISystemRewardFilterer{contract: contract}, nil
}

// bindISystemReward binds a generic wrapper to an already deployed contract.
func bindISystemReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISystemRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemReward *ISystemRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemReward.Contract.ISystemRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemReward *ISystemRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemReward.Contract.ISystemRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemReward *ISystemRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemReward.Contract.ISystemRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemReward *ISystemRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemReward *ISystemRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemReward *ISystemRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemReward.Contract.contract.Transact(opts, method, params...)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns(uint256 actualAmount)
func (_ISystemReward *ISystemRewardTransactor) ClaimRewards(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISystemReward.contract.Transact(opts, "claimRewards", to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns(uint256 actualAmount)
func (_ISystemReward *ISystemRewardSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISystemReward.Contract.ClaimRewards(&_ISystemReward.TransactOpts, to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns(uint256 actualAmount)
func (_ISystemReward *ISystemRewardTransactorSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ISystemReward.Contract.ClaimRewards(&_ISystemReward.TransactOpts, to, amount)
}

// ITokenHubABI is the input ABI used to generate the binding from.
const ITokenHubABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipientAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"refundAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"expireTime\",\"type\":\"uint64\"}],\"name\":\"batchTransferOutBNB\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bep2Symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"bindToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"getBep2SymbolByContractAddr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bep2Symbol\",\"type\":\"bytes32\"}],\"name\":\"getContractAddrByBEP2Symbol\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiniRelayFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expireTime\",\"type\":\"uint64\"}],\"name\":\"transferOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bep2Symbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"unbindToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ITokenHubFuncSigs maps the 4-byte function signature to its string representation.
var ITokenHubFuncSigs = map[string]string{
	"6e056520": "batchTransferOutBNB(address[],uint256[],address[],uint64)",
	"8eff336c": "bindToken(bytes32,address,uint256)",
	"bd466461": "getBep2SymbolByContractAddr(address)",
	"59b92789": "getContractAddrByBEP2Symbol(bytes32)",
	"149d14d9": "getMiniRelayFee()",
	"aa7415f5": "transferOut(address,address,uint256,uint64)",
	"b99328c5": "unbindToken(bytes32,address)",
}

// ITokenHub is an auto generated Go binding around an Ethereum contract.
type ITokenHub struct {
	ITokenHubCaller     // Read-only binding to the contract
	ITokenHubTransactor // Write-only binding to the contract
	ITokenHubFilterer   // Log filterer for contract events
}

// ITokenHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenHubSession struct {
	Contract     *ITokenHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenHubCallerSession struct {
	Contract *ITokenHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ITokenHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenHubTransactorSession struct {
	Contract     *ITokenHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ITokenHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenHubRaw struct {
	Contract *ITokenHub // Generic contract binding to access the raw methods on
}

// ITokenHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenHubCallerRaw struct {
	Contract *ITokenHubCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenHubTransactorRaw struct {
	Contract *ITokenHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenHub creates a new instance of ITokenHub, bound to a specific deployed contract.
func NewITokenHub(address common.Address, backend bind.ContractBackend) (*ITokenHub, error) {
	contract, err := bindITokenHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenHub{ITokenHubCaller: ITokenHubCaller{contract: contract}, ITokenHubTransactor: ITokenHubTransactor{contract: contract}, ITokenHubFilterer: ITokenHubFilterer{contract: contract}}, nil
}

// NewITokenHubCaller creates a new read-only instance of ITokenHub, bound to a specific deployed contract.
func NewITokenHubCaller(address common.Address, caller bind.ContractCaller) (*ITokenHubCaller, error) {
	contract, err := bindITokenHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenHubCaller{contract: contract}, nil
}

// NewITokenHubTransactor creates a new write-only instance of ITokenHub, bound to a specific deployed contract.
func NewITokenHubTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenHubTransactor, error) {
	contract, err := bindITokenHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenHubTransactor{contract: contract}, nil
}

// NewITokenHubFilterer creates a new log filterer instance of ITokenHub, bound to a specific deployed contract.
func NewITokenHubFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenHubFilterer, error) {
	contract, err := bindITokenHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenHubFilterer{contract: contract}, nil
}

// bindITokenHub binds a generic wrapper to an already deployed contract.
func bindITokenHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenHub *ITokenHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenHub.Contract.ITokenHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenHub *ITokenHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenHub.Contract.ITokenHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenHub *ITokenHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenHub.Contract.ITokenHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenHub *ITokenHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenHub *ITokenHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenHub *ITokenHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenHub.Contract.contract.Transact(opts, method, params...)
}

// GetBep2SymbolByContractAddr is a free data retrieval call binding the contract method 0xbd466461.
//
// Solidity: function getBep2SymbolByContractAddr(address contractAddr) view returns(bytes32)
func (_ITokenHub *ITokenHubCaller) GetBep2SymbolByContractAddr(opts *bind.CallOpts, contractAddr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ITokenHub.contract.Call(opts, &out, "getBep2SymbolByContractAddr", contractAddr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBep2SymbolByContractAddr is a free data retrieval call binding the contract method 0xbd466461.
//
// Solidity: function getBep2SymbolByContractAddr(address contractAddr) view returns(bytes32)
func (_ITokenHub *ITokenHubSession) GetBep2SymbolByContractAddr(contractAddr common.Address) ([32]byte, error) {
	return _ITokenHub.Contract.GetBep2SymbolByContractAddr(&_ITokenHub.CallOpts, contractAddr)
}

// GetBep2SymbolByContractAddr is a free data retrieval call binding the contract method 0xbd466461.
//
// Solidity: function getBep2SymbolByContractAddr(address contractAddr) view returns(bytes32)
func (_ITokenHub *ITokenHubCallerSession) GetBep2SymbolByContractAddr(contractAddr common.Address) ([32]byte, error) {
	return _ITokenHub.Contract.GetBep2SymbolByContractAddr(&_ITokenHub.CallOpts, contractAddr)
}

// GetContractAddrByBEP2Symbol is a free data retrieval call binding the contract method 0x59b92789.
//
// Solidity: function getContractAddrByBEP2Symbol(bytes32 bep2Symbol) view returns(address)
func (_ITokenHub *ITokenHubCaller) GetContractAddrByBEP2Symbol(opts *bind.CallOpts, bep2Symbol [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ITokenHub.contract.Call(opts, &out, "getContractAddrByBEP2Symbol", bep2Symbol)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractAddrByBEP2Symbol is a free data retrieval call binding the contract method 0x59b92789.
//
// Solidity: function getContractAddrByBEP2Symbol(bytes32 bep2Symbol) view returns(address)
func (_ITokenHub *ITokenHubSession) GetContractAddrByBEP2Symbol(bep2Symbol [32]byte) (common.Address, error) {
	return _ITokenHub.Contract.GetContractAddrByBEP2Symbol(&_ITokenHub.CallOpts, bep2Symbol)
}

// GetContractAddrByBEP2Symbol is a free data retrieval call binding the contract method 0x59b92789.
//
// Solidity: function getContractAddrByBEP2Symbol(bytes32 bep2Symbol) view returns(address)
func (_ITokenHub *ITokenHubCallerSession) GetContractAddrByBEP2Symbol(bep2Symbol [32]byte) (common.Address, error) {
	return _ITokenHub.Contract.GetContractAddrByBEP2Symbol(&_ITokenHub.CallOpts, bep2Symbol)
}

// GetMiniRelayFee is a free data retrieval call binding the contract method 0x149d14d9.
//
// Solidity: function getMiniRelayFee() view returns(uint256)
func (_ITokenHub *ITokenHubCaller) GetMiniRelayFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ITokenHub.contract.Call(opts, &out, "getMiniRelayFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMiniRelayFee is a free data retrieval call binding the contract method 0x149d14d9.
//
// Solidity: function getMiniRelayFee() view returns(uint256)
func (_ITokenHub *ITokenHubSession) GetMiniRelayFee() (*big.Int, error) {
	return _ITokenHub.Contract.GetMiniRelayFee(&_ITokenHub.CallOpts)
}

// GetMiniRelayFee is a free data retrieval call binding the contract method 0x149d14d9.
//
// Solidity: function getMiniRelayFee() view returns(uint256)
func (_ITokenHub *ITokenHubCallerSession) GetMiniRelayFee() (*big.Int, error) {
	return _ITokenHub.Contract.GetMiniRelayFee(&_ITokenHub.CallOpts)
}

// BatchTransferOutBNB is a paid mutator transaction binding the contract method 0x6e056520.
//
// Solidity: function batchTransferOutBNB(address[] recipientAddrs, uint256[] amounts, address[] refundAddrs, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubTransactor) BatchTransferOutBNB(opts *bind.TransactOpts, recipientAddrs []common.Address, amounts []*big.Int, refundAddrs []common.Address, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.contract.Transact(opts, "batchTransferOutBNB", recipientAddrs, amounts, refundAddrs, expireTime)
}

// BatchTransferOutBNB is a paid mutator transaction binding the contract method 0x6e056520.
//
// Solidity: function batchTransferOutBNB(address[] recipientAddrs, uint256[] amounts, address[] refundAddrs, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubSession) BatchTransferOutBNB(recipientAddrs []common.Address, amounts []*big.Int, refundAddrs []common.Address, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.Contract.BatchTransferOutBNB(&_ITokenHub.TransactOpts, recipientAddrs, amounts, refundAddrs, expireTime)
}

// BatchTransferOutBNB is a paid mutator transaction binding the contract method 0x6e056520.
//
// Solidity: function batchTransferOutBNB(address[] recipientAddrs, uint256[] amounts, address[] refundAddrs, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubTransactorSession) BatchTransferOutBNB(recipientAddrs []common.Address, amounts []*big.Int, refundAddrs []common.Address, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.Contract.BatchTransferOutBNB(&_ITokenHub.TransactOpts, recipientAddrs, amounts, refundAddrs, expireTime)
}

// BindToken is a paid mutator transaction binding the contract method 0x8eff336c.
//
// Solidity: function bindToken(bytes32 bep2Symbol, address contractAddr, uint256 decimals) returns()
func (_ITokenHub *ITokenHubTransactor) BindToken(opts *bind.TransactOpts, bep2Symbol [32]byte, contractAddr common.Address, decimals *big.Int) (*types.Transaction, error) {
	return _ITokenHub.contract.Transact(opts, "bindToken", bep2Symbol, contractAddr, decimals)
}

// BindToken is a paid mutator transaction binding the contract method 0x8eff336c.
//
// Solidity: function bindToken(bytes32 bep2Symbol, address contractAddr, uint256 decimals) returns()
func (_ITokenHub *ITokenHubSession) BindToken(bep2Symbol [32]byte, contractAddr common.Address, decimals *big.Int) (*types.Transaction, error) {
	return _ITokenHub.Contract.BindToken(&_ITokenHub.TransactOpts, bep2Symbol, contractAddr, decimals)
}

// BindToken is a paid mutator transaction binding the contract method 0x8eff336c.
//
// Solidity: function bindToken(bytes32 bep2Symbol, address contractAddr, uint256 decimals) returns()
func (_ITokenHub *ITokenHubTransactorSession) BindToken(bep2Symbol [32]byte, contractAddr common.Address, decimals *big.Int) (*types.Transaction, error) {
	return _ITokenHub.Contract.BindToken(&_ITokenHub.TransactOpts, bep2Symbol, contractAddr, decimals)
}

// TransferOut is a paid mutator transaction binding the contract method 0xaa7415f5.
//
// Solidity: function transferOut(address contractAddr, address recipient, uint256 amount, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubTransactor) TransferOut(opts *bind.TransactOpts, contractAddr common.Address, recipient common.Address, amount *big.Int, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.contract.Transact(opts, "transferOut", contractAddr, recipient, amount, expireTime)
}

// TransferOut is a paid mutator transaction binding the contract method 0xaa7415f5.
//
// Solidity: function transferOut(address contractAddr, address recipient, uint256 amount, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubSession) TransferOut(contractAddr common.Address, recipient common.Address, amount *big.Int, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.Contract.TransferOut(&_ITokenHub.TransactOpts, contractAddr, recipient, amount, expireTime)
}

// TransferOut is a paid mutator transaction binding the contract method 0xaa7415f5.
//
// Solidity: function transferOut(address contractAddr, address recipient, uint256 amount, uint64 expireTime) payable returns(bool)
func (_ITokenHub *ITokenHubTransactorSession) TransferOut(contractAddr common.Address, recipient common.Address, amount *big.Int, expireTime uint64) (*types.Transaction, error) {
	return _ITokenHub.Contract.TransferOut(&_ITokenHub.TransactOpts, contractAddr, recipient, amount, expireTime)
}

// UnbindToken is a paid mutator transaction binding the contract method 0xb99328c5.
//
// Solidity: function unbindToken(bytes32 bep2Symbol, address contractAddr) returns()
func (_ITokenHub *ITokenHubTransactor) UnbindToken(opts *bind.TransactOpts, bep2Symbol [32]byte, contractAddr common.Address) (*types.Transaction, error) {
	return _ITokenHub.contract.Transact(opts, "unbindToken", bep2Symbol, contractAddr)
}

// UnbindToken is a paid mutator transaction binding the contract method 0xb99328c5.
//
// Solidity: function unbindToken(bytes32 bep2Symbol, address contractAddr) returns()
func (_ITokenHub *ITokenHubSession) UnbindToken(bep2Symbol [32]byte, contractAddr common.Address) (*types.Transaction, error) {
	return _ITokenHub.Contract.UnbindToken(&_ITokenHub.TransactOpts, bep2Symbol, contractAddr)
}

// UnbindToken is a paid mutator transaction binding the contract method 0xb99328c5.
//
// Solidity: function unbindToken(bytes32 bep2Symbol, address contractAddr) returns()
func (_ITokenHub *ITokenHubTransactorSession) UnbindToken(bep2Symbol [32]byte, contractAddr common.Address) (*types.Transaction, error) {
	return _ITokenHub.Contract.UnbindToken(&_ITokenHub.TransactOpts, bep2Symbol, contractAddr)
}

// MemoryABI is the input ABI used to generate the binding from.
const MemoryABI = "[]"

// MemoryBin is the compiled bytecode used for deploying new contracts.
var MemoryBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209323a48b104f52ad6049da30c049f2d948fdd51593f63649bed67841938c956664736f6c63430006040033"

// DeployMemory deploys a new Ethereum contract, binding an instance of Memory to it.
func DeployMemory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Memory, error) {
	parsed, err := abi.JSON(strings.NewReader(MemoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MemoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Memory{MemoryCaller: MemoryCaller{contract: contract}, MemoryTransactor: MemoryTransactor{contract: contract}, MemoryFilterer: MemoryFilterer{contract: contract}}, nil
}

// Memory is an auto generated Go binding around an Ethereum contract.
type Memory struct {
	MemoryCaller     // Read-only binding to the contract
	MemoryTransactor // Write-only binding to the contract
	MemoryFilterer   // Log filterer for contract events
}

// MemoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemorySession struct {
	Contract     *Memory           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemoryCallerSession struct {
	Contract *MemoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MemoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemoryTransactorSession struct {
	Contract     *MemoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemoryRaw struct {
	Contract *Memory // Generic contract binding to access the raw methods on
}

// MemoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemoryCallerRaw struct {
	Contract *MemoryCaller // Generic read-only contract binding to access the raw methods on
}

// MemoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemoryTransactorRaw struct {
	Contract *MemoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemory creates a new instance of Memory, bound to a specific deployed contract.
func NewMemory(address common.Address, backend bind.ContractBackend) (*Memory, error) {
	contract, err := bindMemory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Memory{MemoryCaller: MemoryCaller{contract: contract}, MemoryTransactor: MemoryTransactor{contract: contract}, MemoryFilterer: MemoryFilterer{contract: contract}}, nil
}

// NewMemoryCaller creates a new read-only instance of Memory, bound to a specific deployed contract.
func NewMemoryCaller(address common.Address, caller bind.ContractCaller) (*MemoryCaller, error) {
	contract, err := bindMemory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryCaller{contract: contract}, nil
}

// NewMemoryTransactor creates a new write-only instance of Memory, bound to a specific deployed contract.
func NewMemoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MemoryTransactor, error) {
	contract, err := bindMemory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemoryTransactor{contract: contract}, nil
}

// NewMemoryFilterer creates a new log filterer instance of Memory, bound to a specific deployed contract.
func NewMemoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MemoryFilterer, error) {
	contract, err := bindMemory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemoryFilterer{contract: contract}, nil
}

// bindMemory binds a generic wrapper to an already deployed contract.
func bindMemory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MemoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memory *MemoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memory.Contract.MemoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memory *MemoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memory.Contract.MemoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memory *MemoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memory.Contract.MemoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Memory *MemoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Memory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Memory *MemoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Memory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Memory *MemoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Memory.Contract.contract.Transact(opts, method, params...)
}

// RLPDecodeABI is the input ABI used to generate the binding from.
const RLPDecodeABI = "[]"

// RLPDecodeBin is the compiled bytecode used for deploying new contracts.
var RLPDecodeBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220181baf58d2f79353375cbba9ae73c1d7d112f8842f92bc48cb128f13700b978964736f6c63430006040033"

// DeployRLPDecode deploys a new Ethereum contract, binding an instance of RLPDecode to it.
func DeployRLPDecode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPDecode, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPDecodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPDecodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPDecode{RLPDecodeCaller: RLPDecodeCaller{contract: contract}, RLPDecodeTransactor: RLPDecodeTransactor{contract: contract}, RLPDecodeFilterer: RLPDecodeFilterer{contract: contract}}, nil
}

// RLPDecode is an auto generated Go binding around an Ethereum contract.
type RLPDecode struct {
	RLPDecodeCaller     // Read-only binding to the contract
	RLPDecodeTransactor // Write-only binding to the contract
	RLPDecodeFilterer   // Log filterer for contract events
}

// RLPDecodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPDecodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPDecodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPDecodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPDecodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPDecodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPDecodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPDecodeSession struct {
	Contract     *RLPDecode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPDecodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPDecodeCallerSession struct {
	Contract *RLPDecodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPDecodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPDecodeTransactorSession struct {
	Contract     *RLPDecodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPDecodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPDecodeRaw struct {
	Contract *RLPDecode // Generic contract binding to access the raw methods on
}

// RLPDecodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPDecodeCallerRaw struct {
	Contract *RLPDecodeCaller // Generic read-only contract binding to access the raw methods on
}

// RLPDecodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPDecodeTransactorRaw struct {
	Contract *RLPDecodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPDecode creates a new instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecode(address common.Address, backend bind.ContractBackend) (*RLPDecode, error) {
	contract, err := bindRLPDecode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPDecode{RLPDecodeCaller: RLPDecodeCaller{contract: contract}, RLPDecodeTransactor: RLPDecodeTransactor{contract: contract}, RLPDecodeFilterer: RLPDecodeFilterer{contract: contract}}, nil
}

// NewRLPDecodeCaller creates a new read-only instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecodeCaller(address common.Address, caller bind.ContractCaller) (*RLPDecodeCaller, error) {
	contract, err := bindRLPDecode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPDecodeCaller{contract: contract}, nil
}

// NewRLPDecodeTransactor creates a new write-only instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPDecodeTransactor, error) {
	contract, err := bindRLPDecode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPDecodeTransactor{contract: contract}, nil
}

// NewRLPDecodeFilterer creates a new log filterer instance of RLPDecode, bound to a specific deployed contract.
func NewRLPDecodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPDecodeFilterer, error) {
	contract, err := bindRLPDecode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPDecodeFilterer{contract: contract}, nil
}

// bindRLPDecode binds a generic wrapper to an already deployed contract.
func bindRLPDecode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPDecodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPDecode *RLPDecodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPDecode.Contract.RLPDecodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPDecode *RLPDecodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPDecode.Contract.RLPDecodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPDecode *RLPDecodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPDecode.Contract.RLPDecodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPDecode *RLPDecodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPDecode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPDecode *RLPDecodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPDecode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPDecode *RLPDecodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPDecode.Contract.contract.Transact(opts, method, params...)
}

// RLPEncodeABI is the input ABI used to generate the binding from.
const RLPEncodeABI = "[]"

// RLPEncodeBin is the compiled bytecode used for deploying new contracts.
var RLPEncodeBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a7f8a9dea63ff28a573c4d26717079867be48faf1c6611bb94144eba7330cd6664736f6c63430006040033"

// DeployRLPEncode deploys a new Ethereum contract, binding an instance of RLPEncode to it.
func DeployRLPEncode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPEncode, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPEncodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// RLPEncode is an auto generated Go binding around an Ethereum contract.
type RLPEncode struct {
	RLPEncodeCaller     // Read-only binding to the contract
	RLPEncodeTransactor // Write-only binding to the contract
	RLPEncodeFilterer   // Log filterer for contract events
}

// RLPEncodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPEncodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPEncodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPEncodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPEncodeSession struct {
	Contract     *RLPEncode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPEncodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPEncodeCallerSession struct {
	Contract *RLPEncodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPEncodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPEncodeTransactorSession struct {
	Contract     *RLPEncodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPEncodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPEncodeRaw struct {
	Contract *RLPEncode // Generic contract binding to access the raw methods on
}

// RLPEncodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPEncodeCallerRaw struct {
	Contract *RLPEncodeCaller // Generic read-only contract binding to access the raw methods on
}

// RLPEncodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPEncodeTransactorRaw struct {
	Contract *RLPEncodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPEncode creates a new instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncode(address common.Address, backend bind.ContractBackend) (*RLPEncode, error) {
	contract, err := bindRLPEncode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// NewRLPEncodeCaller creates a new read-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeCaller(address common.Address, caller bind.ContractCaller) (*RLPEncodeCaller, error) {
	contract, err := bindRLPEncode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeCaller{contract: contract}, nil
}

// NewRLPEncodeTransactor creates a new write-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPEncodeTransactor, error) {
	contract, err := bindRLPEncode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeTransactor{contract: contract}, nil
}

// NewRLPEncodeFilterer creates a new log filterer instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPEncodeFilterer, error) {
	contract, err := bindRLPEncode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeFilterer{contract: contract}, nil
}

// bindRLPEncode binds a generic wrapper to an already deployed contract.
func bindRLPEncode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.RLPEncodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220330623cf29bc2b24ef9154ec313ab44fb4077fbaf305c00404422c7fd32ab7bf64736f6c63430006040033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SystemABI is the input ABI used to generate the binding from.
const SystemABI = "[{\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MANAGER_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SystemFuncSigs maps the 4-byte function signature to its string representation.
var SystemFuncSigs = map[string]string{
	"3dffc387": "BIND_CHANNELID()",
	"ab51bb96": "CODE_OK()",
	"51e80672": "CROSS_CHAIN_CONTRACT_ADDR()",
	"0bee7a67": "ERROR_FAIL_DECODE()",
	"96713da9": "GOV_CHANNELID()",
	"9dc09262": "GOV_HUB_ADDR()",
	"6e47b482": "INCENTIVIZE_ADDR()",
	"dc927faf": "LIGHT_CLIENT_ADDR()",
	"a1a11bf5": "RELAYERHUB_CONTRACT_ADDR()",
	"7942fd05": "SLASH_CHANNELID()",
	"43756e5c": "SLASH_CONTRACT_ADDR()",
	"4bf6c882": "STAKING_CHANNELID()",
	"c81b1662": "SYSTEM_REWARD_ADDR()",
	"fd6a6879": "TOKEN_HUB_ADDR()",
	"75d47a0a": "TOKEN_MANAGER_ADDR()",
	"70fd5bad": "TRANSFER_IN_CHANNELID()",
	"fc3e5908": "TRANSFER_OUT_CHANNELID()",
	"f9a2bbc7": "VALIDATOR_CONTRACT_ADDR()",
	"a78abc16": "alreadyInit()",
	"493279b1": "bscChainID()",
}

// SystemBin is the compiled bytecode used for deploying new contracts.
var SystemBin = "0x608060405234801561001057600080fd5b506102ef806100206000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806396713da9116100ad578063c81b166211610071578063c81b16621461021f578063dc927faf14610227578063f9a2bbc71461022f578063fc3e590814610237578063fd6a68791461023f5761012c565b806396713da9146101e35780639dc09262146101eb578063a1a11bf5146101f3578063a78abc16146101fb578063ab51bb96146102175761012c565b806351e80672116100f457806351e80672146101bb5780636e47b482146101c357806370fd5bad146101cb57806375d47a0a146101d35780637942fd05146101db5761012c565b80630bee7a67146101315780633dffc3871461015257806343756e5c14610170578063493279b1146101945780634bf6c882146101b3575b600080fd5b610139610247565b6040805163ffffffff9092168252519081900360200190f35b61015a61024c565b6040805160ff9092168252519081900360200190f35b610178610251565b604080516001600160a01b039092168252519081900360200190f35b61019c610257565b6040805161ffff9092168252519081900360200190f35b61015a61025c565b610178610261565b610178610267565b61015a61026d565b610178610272565b61015a610278565b61015a61027d565b610178610282565b610178610288565b61020361028e565b604080519115158252519081900360200190f35b610139610297565b61017861029c565b6101786102a2565b6101786102a8565b61015a6102ae565b6101786102b3565b606481565b600181565b61100181565b606081565b600881565b61200081565b61100581565b600281565b61100881565b600b81565b600981565b61100781565b61100681565b60005460ff1681565b600081565b61100281565b61100381565b61100081565b600381565b6110048156fea2646970667358221220f6bd4ecf625b4bcfbb66bc7843198370f23b5b7c20b542107de9b5c9ba2acf9464736f6c63430006040033"

// DeploySystem deploys a new Ethereum contract, binding an instance of System to it.
func DeploySystem(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *System, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SystemBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &System{SystemCaller: SystemCaller{contract: contract}, SystemTransactor: SystemTransactor{contract: contract}, SystemFilterer: SystemFilterer{contract: contract}}, nil
}

// System is an auto generated Go binding around an Ethereum contract.
type System struct {
	SystemCaller     // Read-only binding to the contract
	SystemTransactor // Write-only binding to the contract
	SystemFilterer   // Log filterer for contract events
}

// SystemCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemSession struct {
	Contract     *System           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemCallerSession struct {
	Contract *SystemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SystemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemTransactorSession struct {
	Contract     *SystemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRaw struct {
	Contract *System // Generic contract binding to access the raw methods on
}

// SystemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemCallerRaw struct {
	Contract *SystemCaller // Generic read-only contract binding to access the raw methods on
}

// SystemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemTransactorRaw struct {
	Contract *SystemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystem creates a new instance of System, bound to a specific deployed contract.
func NewSystem(address common.Address, backend bind.ContractBackend) (*System, error) {
	contract, err := bindSystem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &System{SystemCaller: SystemCaller{contract: contract}, SystemTransactor: SystemTransactor{contract: contract}, SystemFilterer: SystemFilterer{contract: contract}}, nil
}

// NewSystemCaller creates a new read-only instance of System, bound to a specific deployed contract.
func NewSystemCaller(address common.Address, caller bind.ContractCaller) (*SystemCaller, error) {
	contract, err := bindSystem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemCaller{contract: contract}, nil
}

// NewSystemTransactor creates a new write-only instance of System, bound to a specific deployed contract.
func NewSystemTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemTransactor, error) {
	contract, err := bindSystem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemTransactor{contract: contract}, nil
}

// NewSystemFilterer creates a new log filterer instance of System, bound to a specific deployed contract.
func NewSystemFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemFilterer, error) {
	contract, err := bindSystem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemFilterer{contract: contract}, nil
}

// bindSystem binds a generic wrapper to an already deployed contract.
func bindSystem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_System *SystemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _System.Contract.SystemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_System *SystemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.Contract.SystemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_System *SystemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _System.Contract.SystemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_System *SystemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _System.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_System *SystemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _System.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_System *SystemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _System.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_System *SystemCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "BIND_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_System *SystemSession) BINDCHANNELID() (uint8, error) {
	return _System.Contract.BINDCHANNELID(&_System.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) BINDCHANNELID() (uint8, error) {
	return _System.Contract.BINDCHANNELID(&_System.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_System *SystemCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_System *SystemSession) CODEOK() (uint32, error) {
	return _System.Contract.CODEOK(&_System.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_System *SystemCallerSession) CODEOK() (uint32, error) {
	return _System.Contract.CODEOK(&_System.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_System *SystemCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_System *SystemSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _System.Contract.CROSSCHAINCONTRACTADDR(&_System.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_System *SystemCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _System.Contract.CROSSCHAINCONTRACTADDR(&_System.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_System *SystemCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_System *SystemSession) ERRORFAILDECODE() (uint32, error) {
	return _System.Contract.ERRORFAILDECODE(&_System.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_System *SystemCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _System.Contract.ERRORFAILDECODE(&_System.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_System *SystemCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_System *SystemSession) GOVCHANNELID() (uint8, error) {
	return _System.Contract.GOVCHANNELID(&_System.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) GOVCHANNELID() (uint8, error) {
	return _System.Contract.GOVCHANNELID(&_System.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_System *SystemCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "GOV_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_System *SystemSession) GOVHUBADDR() (common.Address, error) {
	return _System.Contract.GOVHUBADDR(&_System.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_System *SystemCallerSession) GOVHUBADDR() (common.Address, error) {
	return _System.Contract.GOVHUBADDR(&_System.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_System *SystemCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "INCENTIVIZE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_System *SystemSession) INCENTIVIZEADDR() (common.Address, error) {
	return _System.Contract.INCENTIVIZEADDR(&_System.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_System *SystemCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _System.Contract.INCENTIVIZEADDR(&_System.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_System *SystemCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_System *SystemSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _System.Contract.LIGHTCLIENTADDR(&_System.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_System *SystemCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _System.Contract.LIGHTCLIENTADDR(&_System.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_System *SystemCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_System *SystemSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _System.Contract.RELAYERHUBCONTRACTADDR(&_System.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_System *SystemCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _System.Contract.RELAYERHUBCONTRACTADDR(&_System.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_System *SystemCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "SLASH_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_System *SystemSession) SLASHCHANNELID() (uint8, error) {
	return _System.Contract.SLASHCHANNELID(&_System.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) SLASHCHANNELID() (uint8, error) {
	return _System.Contract.SLASHCHANNELID(&_System.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_System *SystemCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "SLASH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_System *SystemSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _System.Contract.SLASHCONTRACTADDR(&_System.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_System *SystemCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _System.Contract.SLASHCONTRACTADDR(&_System.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_System *SystemCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_System *SystemSession) STAKINGCHANNELID() (uint8, error) {
	return _System.Contract.STAKINGCHANNELID(&_System.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _System.Contract.STAKINGCHANNELID(&_System.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_System *SystemCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_System *SystemSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _System.Contract.SYSTEMREWARDADDR(&_System.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_System *SystemCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _System.Contract.SYSTEMREWARDADDR(&_System.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_System *SystemCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_System *SystemSession) TOKENHUBADDR() (common.Address, error) {
	return _System.Contract.TOKENHUBADDR(&_System.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_System *SystemCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _System.Contract.TOKENHUBADDR(&_System.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_System *SystemCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TOKEN_MANAGER_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_System *SystemSession) TOKENMANAGERADDR() (common.Address, error) {
	return _System.Contract.TOKENMANAGERADDR(&_System.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_System *SystemCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _System.Contract.TOKENMANAGERADDR(&_System.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_System *SystemCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_System *SystemSession) TRANSFERINCHANNELID() (uint8, error) {
	return _System.Contract.TRANSFERINCHANNELID(&_System.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _System.Contract.TRANSFERINCHANNELID(&_System.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_System *SystemCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_System *SystemSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _System.Contract.TRANSFEROUTCHANNELID(&_System.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_System *SystemCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _System.Contract.TRANSFEROUTCHANNELID(&_System.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_System *SystemCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_System *SystemSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _System.Contract.VALIDATORCONTRACTADDR(&_System.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_System *SystemCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _System.Contract.VALIDATORCONTRACTADDR(&_System.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_System *SystemCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_System *SystemSession) AlreadyInit() (bool, error) {
	return _System.Contract.AlreadyInit(&_System.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_System *SystemCallerSession) AlreadyInit() (bool, error) {
	return _System.Contract.AlreadyInit(&_System.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_System *SystemCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _System.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_System *SystemSession) BscChainID() (uint16, error) {
	return _System.Contract.BscChainID(&_System.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_System *SystemCallerSession) BscChainID() (uint16, error) {
	return _System.Contract.BscChainID(&_System.CallOpts)
}

// TransferHelperABI is the input ABI used to generate the binding from.
const TransferHelperABI = "[]"

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
var TransferHelperBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220384675e16fcf4874ef3db109f5b14245017bbe24e4fbd824294bd89b9eefe0d464736f6c63430006040033"

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferHelper, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TransferHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct {
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
}

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct {
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct {
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct {
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct {
	Contract *TransferHelper // Generic contract binding to access the raw methods on
}

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct {
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct {
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) {
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) {
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperCaller{contract: contract}, nil
}

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) {
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperTransactor{contract: contract}, nil
}

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) {
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperFilterer{contract: contract}, nil
}

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
}
