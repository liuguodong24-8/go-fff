package config

// Setting 配置对象
var Setting setting

// setting 配置
type setting struct {
	App struct {
		Port int64
	}
	Log struct {
		Channel string
		Level   int
		Output  string
		Stack   bool
	}
	Chain struct {
		Address            string
		WsAddress          string
		NftContractOwnKey  string
		NftContractAddress string
		ToAddress          string
	}
	Java struct {
		MintUrl string
		Type    string
	}
}
