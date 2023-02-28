package conf

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io/ioutil"
)

var c *Config

type Config struct {
	App
	WxApps []WxApps `json:"wx_app"`
	Applet []Applet `json:"applet"`
	Redis  Redis    `json:"redis"`
}

type (
	App struct {
		AppName      string `json:"app_name"`
		Port         string `json:"port"`
		HttpListen   string `json:"http_listen"`
		RunMode      string `json:"run_mode"`
		Version      string `json:"version"`
		Logger       string `json:"logger"`
		WxAuthDomain string `json:"wx_auth_domain"`
		WxUrl        string `json:"wx_url"`
		IsTrue       bool   `json:"is_true"`
	}
	WxApps struct {
		WxId      int    `json:"wx_id"`
		Type      int    `json:"type"`
		Name      string `json:"name"`
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
		Token     string `json:"token"`
		AesKey    string `json:"aes_key"`
	}
	Applet struct {
		Name      string `json:"name"`
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}
	Redis struct {
		Auth        string `json:"auth"`
		DB          int    `json:"db"`
		Host        string `json:"host"`
		PoolSize    int    `json:"pool_size"`
		MinIdleConn int    `json:"min_idle_conn"`
	}
)

func InitConfig(confPath string) error {
	viper.SetConfigFile(confPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.WatchConfig()
	InitConf(confPath)
	viper.OnConfigChange(func(_ fsnotify.Event) {
		InitConf(confPath)
	})
	return nil
}

func InitConf(confPath string) *Config {
	c = &Config{
		App: App{
			AppName:      viper.GetString("app_name"),
			RunMode:      viper.GetString("run_mode"),
			Port:         viper.GetString("port"),
			HttpListen:   viper.GetString("http_listen"),
			Logger:       viper.GetString("logger"),
			Version:      viper.GetString("version"),
			WxAuthDomain: viper.GetString("wx_auth_domain"),
			WxUrl:        viper.GetString("wx_url"),
			IsTrue:       viper.GetBool("is_true"),
		},
		Redis: Redis{
			Host:        viper.GetString("redis.host"),
			Auth:        viper.GetString("redis.auth"),
			PoolSize:    viper.GetInt("redis.pool_size"),
			MinIdleConn: viper.GetInt("redis.min_idle_conn"),
			DB:          viper.GetInt("redis.db"),
		},
	}
	data, _ := ioutil.ReadFile(confPath)
	json.Unmarshal(data, &c)
	return c
}

func GetBaseConf() App        { return c.App }
func GetWxApps() []WxApps     { return c.WxApps }
func GetApplet() []Applet     { return c.Applet }
func GetRdbConf() Redis       { return c.Redis }
func GetRunMode() string      { return c.App.RunMode }
func GetAppName() string      { return c.App.AppName }
func GetWxAuthDomain() string { return c.App.WxAuthDomain }
