# beanpay

提供钱包帐户的创建、查询，余额加款、扣款、退款、查询功能

提供服务包创建、查询，数量增加、减少、撤销、查询功能

√ 　支持钱包余额的加、扣、退

√ 　支持服务包数量的增、减、撤

√ 　独立服务运行，调用 api,rpc 服务集成

√ 　集成服务运行，使用源代码直接集成

√ 　高并发支持

√ 　基于 hydra 构建

[接口开发规范](https://github.com/micro-plat/beanpay/blob/master/api.md)

## 一、独立服务

独立启动`apiserver`对外提供`api`,`rpc`服务

#### 1、启动`apiserver`

- 编译`apiserver`服务

```sh
 ~/work/bin$ go install -tags "prod" github.com/micro-plat/beanpay/apiserver #mysql

 或

 ~/work/bin$ go install -tags "prod oci" github.com/micro-plat/beanpay/apiserver #oracle
```

- 安装服务, 根据向导设置数据库连接串, 生成数据库表结构

```sh
apiserver install -r zk://192.168.106.18 -c t
```

- 启动服务

```sh
apiserver start -r zk://192.168.106.18 -c t
```

#### 2、测试服务

- 传入用户编号、名称创建钱包帐户。返回创建好的帐户信息

```sh
~/work/bin$ curl "http://192.168.4.121:9090/account/create?eid=86001&name=colin"

{"account_id":86000,"account_name":"colin","balance":200,"credit":0}
```

- 传入用户编号或商户编号,外部交易编号(幂等判断)，加款金额进行钱包加款。返回加款记录信息

```sh
~/work/bin$ curl "http://192.168.4.121:9090/account/balance/add?eid=86001&trade_no=8970876&amount=200"
{"account_id":"86000","amount":"200","balance":"200","change_type":"1","create_time":"20190731172225","record_id":"100000","trade_no":"8970876"}
```

- 传入用户编号或商户编号,外部交易编号(幂等判断)，加款金额进行钱包扣款。返回扣款记录信息

```sh
~/work/bin$ curl "http://192.168.4.121:9090/account/balance/deduct?eid=86001&trade_no=8970876&amount=200"
{"account_id":"86000","amount":"-200","balance":"0","change_type":"2","create_time":"20190731172225","record_id":"100001","trade_no":"8970876"}
```

- 传入用户编号、外部扣款交易编号(幂等判断)，退款金额进行钱包退款，退款金额不能大于扣款金额，同一笔扣款只允许一次退款操作。返回退款记录信息

```sh
~/work/bin$ curl "http://192.168.4.121:9090/account/balance/refund?eid=86001&trade_no=8970876&amount=200"
{"account_id":"86000","amount":"200","balance":"200","change_type":"3","create_time":"20190731172225","record_id":"100002","trade_no":"8970876"}
```

[其它接口，参考开发规范](https://github.com/micro-plat/beanpay/blob/master/api.md)

## 二、代码集成

将数据表创建到业务系统，或独立库，业务系统使用本地代码，直接进行帐户、服务包操作

#### 1、创建数据库

使用 beanpay 创建数据库，支持 mysql, oracle

- 编译 beanpay

```sh
 ~/work/bin$ go install github.com/micro-plat/beanpay #mysql

 或

 ~/work/bin$ go install -tags "oci" github.com/micro-plat/beanpay # oracle

```

- 创建数据表

`beanpay [注册中心地址] [平台名称]` 即可将 `beanpay` 需要的表创建到`/平台/var/db/db` 配置的数据库中。

```sh
~/work/bin$ beanpay zk://192.168.0.109 mall #读取/mall/var/db/db 创建数据库

或

~/work/bin$ beanpay zk://192.168.0.109 mall mdb #读取 /mall/var/db/mdb 创建数据库

```

#### 2、编码

- 创建帐户

```go
account, err := beanpay.CreateAccount(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("name"))
if err != nil {
    return err
}

return account
```

- 帐户加款

```go
record, err := beanpay.AddAmount(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("amount"))
if err != nil {
    return err
}
return record

```

- 帐户扣款

```go
record, err := beanpay.DeductAmount(ctx,
		ctx.Request.GetString("eid"),
		ctx.Request.GetString("trade_no"),
		ctx.Request.GetInt("amount"))
if err != nil {
    return err
}
return record

```

> 其它操作请查看[http://github.com/micro-plat/beanpay/beanpay/service.go](https://github.com/micro-plat/beanpay/blob/master/beanpay/service.go)
