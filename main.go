package wechat_service

import "wechat-service/conf"

var path = "./conf/app.json"

func main() {
	err := conf.InitConfig(path)
	if err != nil {
		panic(err)
	}
}
