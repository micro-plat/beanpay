module github.com/micro-plat/beanpay

go 1.14

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/micro-plat/hydra v0.0.0-00010101000000-000000000000
	github.com/micro-plat/lib4go v1.0.9
	github.com/zkfy/go-oci8 v0.0.0-20180327092318-ad9f59dedff0
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	google.golang.org/appengine v1.6.7 // indirect
)

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra

replace github.com/micro-plat/lib4go => ../../../github.com/micro-plat/lib4go
