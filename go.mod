module github.com/micro-plat/beanpay

go 1.12

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/go-sql-driver/mysql v1.4.1
	github.com/micro-plat/hydra v0.12.2
	github.com/micro-plat/lib4go v0.2.1
	github.com/micro-plat/zkcli v0.0.0-20190522060924-e37c30ff0771
	github.com/urfave/cli v1.22.1
	github.com/zkfy/go-oci8 v0.0.0-20180327092318-ad9f59dedff0
)

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra

replace github.com/micro-plat/lib4go => ../../../github.com/micro-plat/lib4go
