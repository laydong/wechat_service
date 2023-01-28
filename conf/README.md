# 项目整体配置

## 项目名称
app_name = "wechat-server"
## 监听端口
http_listen = "0.0.0.0:10080"
## 是否开启代理
"is_true": true,
## 代理到服务域名
"wx_url": "http://127.0.0.1",


[kafka]
## kafka brokers
brokers = ["192.168.3.40:9092", "192.168.3.41:9092", "192.168.3.42:9092"]
cert_file = ""
key_file = ""
ca_file = ""
verify_ssl = false

[[wx_app]]
## 填写唯一得微信标识
"wx_id": 10000,
## 是否给前端获取到微信编号 1获取 0不获取
"type": 0,
## 填写唯一得appname
app_name = "weixin1"
## 微信app_id
app_id = "#######"
## 微信app_secret
app_secret = "#######"
## 消息加解密Token
token = "laya"
## 消息加解密秘钥
aes_key = "weixin1"
## 微信支付账号-商户id
mch_id = "weixin1"
## 微信支付秘钥-商户密钥
mch_key = "weixin1"
## 服务器IP地址(微信支付)
ip_address = "weixin1"
## 允许调用的接口列表，NULL表示不限制
calls = ""

[[wx_app]]
# app3配置
## 填写唯一得appname
app_name = "weixin3"
## 微信app_id
app_id = "weixin3"
## 微信app_secret
app_secret = "weixin3"
## 消息加解密Token
token = "weixin3"
## 消息加解密秘钥
aes_key = "weixin3"
## 微信支付账号-商户id
mch_id = "weixin3"
## 微信支付秘钥-商户密钥
mch_key = "weixin3"
## 服务器IP地址(微信支付)
ip_address = "weixin3"
## 允许调用的接口列表，NULL表示不限制
calls = "/api|/api/new"
