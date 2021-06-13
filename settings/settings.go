package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	Port      int    `mapstructure:"port"`
	MachineID int64  `mapstructure:"machine_id"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
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
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	//方式一：直接指定配置文件路径（相对路径或者绝对路径）
	//相对路径：相对执行的可执行文件的相对路径
	//绝对路径：系统中实际的文件路径

	//方式二：指定配置文件名和配置问价的位置，viper自行查找可用的配置文件
	//配置文件名不需要带后缀
	//配置文件位置可配置多个
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")

	//viper.SetConfigType("yaml") // 从远程比如etcd读取配置使用

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println("viper.Read In config() failed, err:", err)
		return err
	}

	// 将配置反序列化到conf中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper unmarshal conf,", err)
	}

	// 监视配置文件改变
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println("viper unmarshal conf,", err)
		}
	})
	return nil
}
