package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(App)

type App struct {
	*Neo4j `mapstructure:"neo4j"`
}

type Neo4j struct {
	Uri      string `mapstructure:"uri"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func Init() error {
	viper.SetConfigFile("./conf/conf.yaml") // 指定配置文件
	//viper.AddConfigPath("../conf/")         // 指定查找配置文件的路径
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {             // 读取配置信息失败
		//panic(fmt.Errorf("Fatal error dao file: %s \n", err))
		return err
	}
	//将配置文件序列化成app
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper Unmarshal failed,err: %s", err.Error())
	}
	// 监控配置文件变化
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config update")
		// 添加通知事件
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper Unmarshal failed,err: %s", err.Error())
		}
	})
	return nil
}
