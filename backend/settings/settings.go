package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	DB           string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func Init() (err error) {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 指定配置文件类型 专用于从远程配置信息时指定配置文件类型
	viper.AddConfigPath(".")      // 还可以在工作目录中查找配置
	err = viper.ReadInConfig()    // 查找并读取配置文件
	if err != nil {
		// 处理读取配置文件的错误
		return
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
	}
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal(Conf) failed,err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal(Conf) failed,err:%v\n", err)
		}
	})
	fmt.Printf("%#v\n")
	return
}
