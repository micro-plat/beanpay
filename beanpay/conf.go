package beanpay



var DBName = "db"
var QueueName = "queue"

//Config　配置数据库，消息队列的配置名称
func Config(db string, mq string) {
	DBName = db
	QueueName = mq
}