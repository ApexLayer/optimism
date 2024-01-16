// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const CrossL2InboxStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/interop/CrossL2Inbox.sol:CrossL2Inbox\",\"label\":\"roots\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_bytes32,t_bool))\"},{\"astId\":1001,\"contract\":\"src/interop/CrossL2Inbox.sol:CrossL2Inbox\",\"label\":\"unconsumedMessages\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1002,\"contract\":\"src/interop/CrossL2Inbox.sol:CrossL2Inbox\",\"label\":\"chainState\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_struct(ChainState)1005_storage)\"},{\"astId\":1003,\"contract\":\"src/interop/CrossL2Inbox.sol:CrossL2Inbox\",\"label\":\"crossL2Sender\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/interop/CrossL2Inbox.sol:CrossL2Inbox\",\"label\":\"messageSourceChain\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_bytes32\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes32,t_mapping(t_bytes32,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(bytes32 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_bytes32,t_bool)\"},\"t_mapping(t_bytes32,t_struct(ChainState)1005_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e struct ChainState)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_struct(ChainState)1005_storage\"},\"t_struct(ChainState)1005_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ChainState\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var CrossL2InboxStorageLayout = new(solc.StorageLayout)

var CrossL2InboxDeployedBin = "0x6080604052600436106100965760003560e01c8063c220426b11610069578063e5507be31161004e578063e5507be31461022d578063e591b2821461024f578063e63c390c1461027757600080fd5b8063c220426b146101b5578063db10b9a9146101f557600080fd5b80632cfea9fe1461009b5780633d6d0dd4146100e957806354fd4d501461013b57806360a687aa14610191575b600080fd5b3480156100a757600080fd5b506100cf6100b6366004610b2e565b6002602052600090815260409020805460019091015482565b604080519283526020830191909152015b60405180910390f35b3480156100f557600080fd5b506003546101169073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e0565b34801561014757600080fd5b506101846040518060400160405280600581526020017f302e302e3100000000000000000000000000000000000000000000000000000081525081565b6040516100e09190610b47565b34801561019d57600080fd5b506101a760045481565b6040519081526020016100e0565b3480156101c157600080fd5b506101e56101d0366004610b2e565b60016020526000908152604090205460ff1681565b60405190151581526020016100e0565b34801561020157600080fd5b506101e5610210366004610bba565b600060208181529281526040808220909352908152205460ff1681565b34801561023957600080fd5b5061024d610248366004610d09565b61028a565b005b34801561025b57600080fd5b5061011673deaddeaddeaddeaddeaddeaddeaddeaddead000281565b61024d610285366004610dcb565b61056e565b60408101514614610322576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f43726f73734c32496e626f783a2074617267657420636861696e20646f65732060448201527f6e6f74206d61746368000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b600061032d8261091e565b60008181526001602052604090205490915060ff166103a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f43726f73734c32496e626f783a20756e6b6e6f776e206d6573736167650000006044820152606401610319565b60035473ffffffffffffffffffffffffffffffffffffffff1661dead14610451576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f43726f73734c32496e626f783a2063616e206f6e6c792074726967676572206f60448201527f6e652063616c6c20706572206d657373616765000000000000000000000000006064820152608401610319565b6060820151600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055602080830151600455600082815260019091526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055608083015160c084015160a085015160e08601516104f793929190610ab2565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055600060045560405190915082907f608b51d991a28926c87c94dae8c72df6a62c5f22b359bb418bf204355b39fa7d9061056190841515815260200190565b60405180910390a2505050565b3373deaddeaddeaddeaddeaddeaddeaddeaddead000214610611576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f43726f73734c32496e626f783a206f6e6c7920706f737469652063616e20646560448201527f6c69766572206d61696c000000000000000000000000000000000000000000006064820152608401610319565b60005b81811015610919576002600084848481811061063257610632610e40565b90506020028101906106449190610e6f565b6000013581526020019081526020016000206001015483838381811061066c5761066c610e40565b905060200281019061067e9190610e6f565b604001351161070f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f43726f73734c32496e626f783a20626c6f636b4e756d626572206d757374206260448201527f6520696e6372656173696e6700000000000000000000000000000000000000006064820152608401610319565b604051806040016040528084848481811061072c5761072c610e40565b905060200281019061073e9190610e6f565b60200135815260200184848481811061075957610759610e40565b905060200281019061076b9190610e6f565b6040013590526002600085858581811061078757610787610e40565b90506020028101906107999190610e6f565b3581526020808201929092526040016000908120835181559290910151600192830155808585858181106107cf576107cf610e40565b90506020028101906107e19190610e6f565b600001358152602001908152602001600020600085858581811061080757610807610e40565b90506020028101906108199190610e6f565b60200135815260200190815260200160002060006101000a81548160ff02191690831515021790555060005b83838381811061085757610857610e40565b90506020028101906108699190610e6f565b610877906060810190610ead565b905081101561090657600180600086868681811061089757610897610e40565b90506020028101906108a99190610e6f565b6108b7906060810190610ead565b858181106108c7576108c7610e40565b90506020020135815260200190815260200160002060006101000a81548160ff02191690831515021790555080806108fe90610f1c565b915050610845565b508061091181610f1c565b915050610614565b505050565b60e0810151805160209182018190206040805193840182905283019190915260009182906060016040516020818303038152906040528051906020012090506000846060015185608001518660a001518760c001516040516020016109b7949392919073ffffffffffffffffffffffffffffffffffffffff94851681529290931660208301526040820152606081019190915260800190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828252805160209182012081840152828201949094528051808303820181526060830182528051908501208785015197820151608084019890985260a0808401989098528151808403909801885260c08301825287519785019790972060e0830152610100808301979097528051808303909701875261012090910190525083519301929092207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001792915050565b6000806000610ac2866000610b10565b905080610af8576308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b600080855160208701888b5af1979650505050505050565b600080603f83619c4001026040850201603f5a021015949350505050565b600060208284031215610b4057600080fd5b5035919050565b600060208083528351808285015260005b81811015610b7457858101830151858201604001528201610b58565b81811115610b86576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008060408385031215610bcd57600080fd5b50508035926020909101359150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff81118282101715610c2f57610c2f610bdc565b60405290565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c5957600080fd5b919050565b600082601f830112610c6f57600080fd5b813567ffffffffffffffff80821115610c8a57610c8a610bdc565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610cd057610cd0610bdc565b81604052838152866020858801011115610ce957600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215610d1b57600080fd5b813567ffffffffffffffff80821115610d3357600080fd5b908301906101008286031215610d4857600080fd5b610d50610c0b565b823581526020830135602082015260408301356040820152610d7460608401610c35565b6060820152610d8560808401610c35565b608082015260a083013560a082015260c083013560c082015260e083013582811115610db057600080fd5b610dbc87828601610c5e565b60e08301525095945050505050565b60008060208385031215610dde57600080fd5b823567ffffffffffffffff80821115610df657600080fd5b818501915085601f830112610e0a57600080fd5b813581811115610e1957600080fd5b8660208260051b8501011115610e2e57600080fd5b60209290920196919550909350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112610ea357600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610ee257600080fd5b83018035915067ffffffffffffffff821115610efd57600080fd5b6020019150600581901b3603821315610f1557600080fd5b9250929050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f74577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(CrossL2InboxStorageLayoutJSON), CrossL2InboxStorageLayout); err != nil {
		panic(err)
	}

	layouts["CrossL2Inbox"] = CrossL2InboxStorageLayout
	deployedBytecodes["CrossL2Inbox"] = CrossL2InboxDeployedBin
	immutableReferences["CrossL2Inbox"] = false
}