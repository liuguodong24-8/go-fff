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
	MongoDb struct {
		Addr       string
		MongodbUri string
		Timeout    int
		Source     string
		Database   string
		Username   string
		Password   string
		PoolLimit  int
	}
	Chain struct {
		Address   string
		WsAddress string
	}
}
