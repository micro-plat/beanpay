package beanpay

import "github.com/micro-plat/beanpay/beanpay/const/confs"

// Config 配置数据库，消息队列的配置名称
func Config(db string, mq string) {
	confs.DBName = db
	confs.QueueName = mq
}

var localMode = true

//SetRemoteMode beanpay为远程服务模式
func SetRemoteMode() {
	localMode = false
}
