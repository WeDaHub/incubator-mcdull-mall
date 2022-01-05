package env

import (
	"fmt"

	"github.com/spf13/viper"
)

type Conf struct {
	Domain string
	Mini   WxMiniApp
	Db     MySQL
}

type Server struct {
	Domain string
}

type WxMiniApp struct {
	AppId  string // 小程序AppId
	Secret string // 小程序Secret
}

type MySQL struct {
	Addr   string // 数据库地址
	User   string // 数据库用户
	Pass   string // 数据库密码
	DbName string // DB名称
}

// LoadConf 读取配置
func LoadConf() *Conf {
	v := viper.New()
	v.SetConfigName("application")
	v.AddConfigPath("conf/")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	conf := &Conf{
		v.GetString("server.domain"),
		WxMiniApp{
			AppId:  v.GetString("wxmini.appid"),
			Secret: v.GetString("wxmini.appsecret"),
		},
		MySQL{
			Addr:   v.GetString("database.addr"),
			User:   v.GetString("database.user"),
			Pass:   v.GetString("database.pass"),
			DbName: v.GetString("database.dbname"),
		},
	}
	return conf
}
