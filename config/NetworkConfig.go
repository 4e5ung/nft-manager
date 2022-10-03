package config

type NetworkInfo struct {
	Name                 string
	ChainId              int64
	RpcURL               string
	SmartContractAddress string
}

var (
	EthTestnet = NetworkInfo{ChainId: 3, Name: "Ropsten", RpcURL: "https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161", SmartContractAddress: "0xB970B7eE2d752c164d86F5865e9bBb0fF6035D52"}
	BscTestnet = NetworkInfo{ChainId: 97, Name: "BSC testnet", RpcURL: "https://data-seed-prebsc-1-s1.binance.org:8545", SmartContractAddress: "0x9B448C561c9bCd2020694B9719D2480C8B10685F"}
	//EdgeTestnet = NetworkInfo{ChainId: 80001, Name: "Mumbai", RpcURL: "https://matic-mumbai.chainstacklabs.com", SmartContractAddress: "0xA69D5FC40165d34cD25EAf6512E0056A86DDECBe"}
	//PolygonTestnet = NetworkInfo{ChainId: 80001, Name: "Mumbai", RpcURL: "https://rpc-mumbai.matic.today", SmartContractAddress: "0xA69D5FC40165d34cD25EAf6512E0056A86DDECBe"}
	EdgeTestnet = NetworkInfo{ChainId: 100, Name: "Edge", RpcURL: "http://rpc-20220830074026599300000002-2117606179.ap-northeast-1.elb.amazonaws.com", SmartContractAddress: "0xFA852BEA5f22a3358D4b682Ad2a55fCF3916A91D"}
)

//var testNetworkInfo = []NetworkInfo{EdgeTestnet}

//func (m NetworkInfo) GetNetworkInfo(name string)(abc NetworkInfo) {
//
//
//	return testNetworkInfo
//}
