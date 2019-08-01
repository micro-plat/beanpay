# beanpay接口开发规范 v1.0

## 一、文档说明

本文档提供beanpay 独立服务`apiserver`对外提供的`api`,`rpc`服务开发规范。


### 1.1 协议格式

- 编码： utf-8 编码，包括: 传输、加密、解密、数据转换等
- 协议： http服务(支持`GET`，`POST`)、rpc服务
- 请求: `键值对`或`json`
- 响应: `json`格式

### 1.2 参数签名

`apiserver`不提供签名，服务仅供内部系统使用


### 1.3 响应结果判断

不同的错误通过http status返回不同的值

| 状态码 | 说明                 |
| :----: | :------------------- |
|  200   | 请求、操作成功       |
|  201   | 请求成功，重复的请求 |
|  901   | 操作失败             |
|  903   | 金额错误             |
|  904   | 数量或余额不足       |
|  908   | 帐户、服务包不存在   |


## 二、钱包接口规范

### 2.1 创建钱包帐户

传入用户编号、名称创建钱包帐户。返回创建好的帐户信息
  
#### 2.1.1 请求参数

| 参数 |  类型  | 可空  | 示例  | 说明         |
| :--- | :----: | :---: | :---: | :----------- |
| eid  | string |  否   | 62356 | 外部用户编号 |
| name | string |  否   | 62356 | 用户名称     |


* 请求示例:

```sh
~/work/bin$ curl "http://192.168.4.121:9090/account/create?eid=86001&name=colin"

```

#### 2.1.2 响应参数

| 参数         |  类型  | 可空  | 示例  |   说明   |
| :----------- | :----: | :---: | :---: | :------: |
| account_id   | number |  否   | 86000 | 帐户编号 |
| account_name | string |  否   | colin | 帐户名称 |
| balance      | number |  否   |  100  | 帐户余额 |
| credit       | number |  否   |   0   | 授信金额 |

* 响应示例:
  
```sh
http.status:200
```

```json
{"account_id":86000,"account_name":"colin","balance":200,"credit":0}
```

> 请求支持幂等，重复调用返回的内容相同，但`http status`为`201`


### 2.2 查询钱包信息

传入用户编号查询钱包信息。返回帐户信息
  
#### 2.2.1 请求参数

| 参数 |  类型  | 可空  | 示例  | 说明         |
| :--- | :----: | :---: | :---: | :----------- |
| eid  | string |  否   | 62356 | 外部用户编号 |


* 请求示例:

```sh
~/work/bin$ curl "http://192.168.4.121:9090/account/query?eid=86001"
```

#### 2.2.2 响应参数

| 参数         |  类型  | 可空  | 示例  |   说明   |
| :----------- | :----: | :---: | :---: | :------: |
| account_id   | number |  否   | 86000 | 帐户编号 |
| account_name | string |  否   | colin | 帐户名称 |
| balance      | number |  否   |  100  | 帐户余额 |
| credit       | number |  否   |   0   | 授信金额 |

* 响应示例:
  
```sh
http.status:200
```

```json
{
    "account_id":86000,
    "account_name":"colin",
    "balance":200,
    "credit":0
}
```

> 请求支持幂等，重复调用返回的内容相同，但`http status`为`201`



### 2.3 钱包加款

传入用户编号、外部交易编号，加款金额进行钱包加款。返回加款记录信息
  
#### 2.3.1 请求参数

| 参数     |  类型  | 可空  |   示例   | 说明             |
| :------- | :----: | :---: | :------: | :--------------- |
| eid      | string |  否   |  62356   | 外部用户编号     |
| trade_no | string |  否   | 86009981 | 外部加款交易编号 |
| amount   | number |  否   |  10000   | 加款金额,单位分  |


* 请求示例:

```sh
~/work/bin$ curl "http://192.168.4.121:9090/account/balance/add?eid=86001&trade_no=8970876
&amount=200"
```

#### 2.3.2 响应参数

| 参数        |  类型  | 可空  |      示例      |       说明       |
| :---------- | :----: | :---: | :------------: | :--------------: |
| record_id   | number |  否   |       0        |   变动记录编号   |
| trade_no    | string |  否   |    8970876     | 外部加款交易编号 |
| account_id  | number |  否   |     86000      |     帐户编号     |
| amount      | number |  否   |      100       |     加款金额     |
| balance     | number |  否   |      100       |   加款后的余额   |
| change_type | number |  否   |       1        | 变动类型(加款:1) |
| create_time | string |  否   | 20190731172225 |     变动时间     |


* 响应示例:
  
```sh
http.status:200
```

```json
{
    "account_id":"86000",
    "amount":"200",
    "balance":"200",
    "change_type":"1",
    "create_time":"20190731172225",
    "record_id":"100000",
    "trade_no":"8970876"
}
```

> 请求支持幂等，重复调用返回的内容相同，但`http status`为`201`





### 2.4 钱包扣款

传入用户编号、外部交易编号，扣款金额进行钱包扣款。返回扣款记录信息
  
#### 2.4.1 请求参数

| 参数     |  类型  | 可空  |   示例   | 说明             |
| :------- | :----: | :---: | :------: | :--------------- |
| eid      | string |  否   |  62356   | 用户编号         |
| trade_no | string |  否   | 86009981 | 外部扣款交易编号 |
| amount   | number |  否   |  10000   | 扣款金额,单位分  |


* 请求示例:

```sh
~/work/bin$ curl "http://192.168.4.121:9090/account/balance/deduct?eid=86001&trade_no=8970876
&amount=200"
```

#### 2.4.2 响应参数

| 参数        |  类型  | 可空  |      示例      |       说明       |
| :---------- | :----: | :---: | :------------: | :--------------: |
| record_id   | number |  否   |       0        |   变动记录编号   |
| trade_no    | string |  否   |    8970876     | 外部扣款交易编号 |
| account_id  | number |  否   |     86000      |     帐户编号     |
| amount      | number |  否   |      100       |     扣款金额     |
| balance     | number |  否   |      100       |   扣款后的余额   |
| change_type | number |  否   |       2        | 变动类型(扣款:2) |
| create_time | string |  否   | 20190731172225 |     变动时间     |


* 响应示例:
  
```sh
http.status:200
```

```json
{
    "account_id":"86000",
    "amount":"-200",
    "balance":"0",
    "change_type":"2",
    "create_time":"20190731172225",
    "record_id":"100001",
    "trade_no":"8970876"
}
```

> 请求支持幂等，重复调用返回的内容相同，但`http status`为`201`



### 2.5 钱包退款

传入用户编号、外部扣款交易编号，退款金额进行退款。返回退款记录信息
  
#### 2.5.1 请求参数

| 参数     |  类型  | 可空  |   示例   | 说明             |
| :------- | :----: | :---: | :------: | :--------------- |
| eid      | string |  否   |  62356   | 用户编号         |
| trade_no | string |  否   | 86009981 | 外部扣款交易编号 |
| amount   | number |  否   |  10000   | 退款金额,单位分  |


* 请求示例:

```sh
~/work/bin$ curl "http://192.168.4.121:9090/account/balance/refund?eid=86001&trade_no=8970876
&amount=200"
```

#### 2.5.2 响应参数

| 参数        |  类型  | 可空  |      示例      |       说明       |
| :---------- | :----: | :---: | :------------: | :--------------: |
| record_id   | number |  否   |       0        |   变动记录编号   |
| trade_no    | string |  否   |    8970876     | 外部扣款交易编号 |
| account_id  | number |  否   |     86000      |     帐户编号     |
| amount      | number |  否   |      100       |     退款金额     |
| balance     | number |  否   |      100       |   退款后的余额   |
| change_type | number |  否   |       3        | 变动类型(退款:3) |
| create_time | string |  否   | 20190731172225 |     变动时间     |


* 响应示例:
  
```sh
http.status:200
```

```json
{
    "account_id":"86000",
    "amount":"200",
    "balance":"200",
    "change_type":"3",
    "create_time":"20190731172225",
    "record_id":"100002",
    "trade_no":"8970876"
}
```

> 请求支持幂等，重复调用返回的内容相同，但`http status`为`201`



### 2.6 变动记录查询

传入用户编号、开始时间、结束时间等查询资金变动记录
  
#### 2.6.1 请求参数

| 参数       |  类型  | 可空  |      示例      | 说明                        |
| :--------- | :----: | :---: | :------------: | :-------------------------- |
| eid        | string |  否   |     62356      | 用户编号                    |
| start_time | string |  否   | 20190731172225 | 开始时间                    |
| end_time   | string |  否   | 20190731172225 | 结束时间                    |
| pi         | number |  是   |       0        | 第几页，从0开始,未指定默认0 |
| ps         | number |  是   |       10       | 返回的数据行数,未指定默认10 |

* 请求示例:

```sh
curl "http://192.168.4.121:9090/account/record/query?eid=86001&start_time=20190731
&end_time=20190731&pi=0&ps=10"
```

#### 2.6.2 响应参数

| 参数        |  类型  | 可空  |      示例      |              说明              |
| :---------- | :----: | :---: | :------------: | :----------------------------: |
| record_id   | number |  否   |       0        |          变动记录编号          |
| trade_no    | string |  否   |    8970876     |        外部扣款交易编号        |
| account_id  | number |  否   |     86000      |            帐户编号            |
| amount      | number |  否   |      100       |            退款金额            |
| balance     | number |  否   |      100       |          退款后的余额          |
| change_type | number |  否   |       3        | 变动类型(加款:1,扣款:2,退款:3) |
| create_time | string |  否   | 20190731172225 |            变动时间            |


* 响应示例:
  
```sh
http.status:200
```

```json
[
    {
        "account_id":"86000",
        "amount":"200",
        "balance":"200",
        "change_type":"3",
        "create_time":"20190731172225",
        "record_id":"100002",
        "trade_no":"8970876"
    },
    {
        "account_id":"86000",
        "amount":"-200",
        "balance":"0",
        "change_type":"2",
        "create_time":"20190731172225",
        "record_id":"100001",
        "trade_no":"8970876"
    },
    {
        "account_id":"86000",
        "amount":"200",
        "balance":"200",
        "change_type":"1",
        "create_time":"20190731172225",
        "record_id":"100000",
        "trade_no":"8970876"
    }
]
```


## 三、服务包接口规范

### 3.1 创建服务包帐户

传入用户编号、服务包名称、服务包可用数量创建服务包帐户。返回创建好的服务包信息
  
#### 3.1.1 请求参数

| 参数    |  类型  | 可空  |   示例   | 说明                           |
| :------ | :----: | :---: | :------: | :----------------------------- |
| eid     | string |  否   |  62356   | 外部用户编号                   |
| sid     | string |  否   |  589766  | 外部服务包编号                 |
| name    | string |  否   |  62356   | 服务包名称                     |
| total   | number |  否   |   1000   | 可用总数                       |
| daily   | number |  是   |   1000   | 日限制数量，未指定时不限制     |
| expires | string |  是   | 20991231 | 过期时间，未指定时默认20991231 |


* 请求示例:

```sh
 curl "http://192.168.4.121:9090/package/create?eid=86001&sid=1000&name=colin
 &total=1000"
```

#### 3.1.2 响应参数

| 参数           |  类型  | 可空  |      示例      |      说明      |
| :------------- | :----: | :---: | :------------: | :------------: |
| pkg_id         | number |  否   |     620000     |   服务包编号   |
| account_id     | number |  否   |     86000      |    帐户编号    |
| spkg_id        | string |  否   |     colin      | 外部服务包编号 |
| pkg_name       | string |  否   |     colin      |   服务包名称   |
| total_capacity | number |  否   |      100       |   服务包总数   |
| total_remain   | number |  否   |       0        |  剩余可用数量  |
| capacity_daily | number |  否   |      100       |  日限制可用数  |
| deduct_today   | number |  否   |       0        |  今日扣减数量  |
| expires        | string |  否   | 20991231000000 |    过期日期    |
| book_time      | string |  否   | 20190731172225 |    订购日期    |
| last_update    | string |  否   | 20190731172225 |  上次修改日期  |


* 响应示例:
  
```sh
http.status:200
```

```json
{
    "pkg_id":620000,
    "account_id":86000,
    "spkg_id":"1000",
    "pkg_name":"colin",
    "total_capacity":1200,
    "total_remain":1000,
    "capacity_daily":1200,
    "deduct_today":200,
    "expires":"20991231000000",
    "book_time":"20190731172225",
    "last_update":"20190731172225"
}
```

> 请求支持幂等，重复调用返回的内容相同，但`http status`为`201`


### 3.2 查询服务包信息

传入用户编号、服务包编号查询服务包信息。返回帐户信息
  
#### 3.2.1 请求参数

| 参数 |  类型  | 可空  |  示例  | 说明           |
| :--- | :----: | :---: | :----: | :------------- |
| eid  | string |  否   | 62356  | 外部用户编号   |
| sid  | string |  否   | 589766 | 外部服务包编号 |

* 请求示例:

```sh
~/work/bin$ curl "http://192.168.4.121:9090/package/query?eid=86001&sid=1000"

```

#### 3.2.2 响应参数

| 参数           |  类型  | 可空  |      示例      |      说明      |
| :------------- | :----: | :---: | :------------: | :------------: |
| pkg_id         | number |  否   |     620000     |   服务包编号   |
| account_id     | number |  否   |     86000      |    帐户编号    |
| spkg_id        | string |  否   |     colin      | 外部服务包编号 |
| pkg_name       | string |  否   |     colin      |   服务包名称   |
| total_capacity | number |  否   |      100       |   服务包总数   |
| total_remain   | number |  否   |       0        |  剩余可用数量  |
| capacity_daily | number |  否   |      100       |  日限制可用数  |
| deduct_today   | number |  否   |       0        |  今日扣减数量  |
| expires        | string |  否   | 20991231000000 |    过期日期    |
| book_time      | string |  否   | 20190731172225 |    订购日期    |
| last_update    | string |  否   | 20190731172225 |  上次修改日期  |


* 响应示例:
  
```sh
http.status:200
```

```json
{
    "pkg_id":620000,
    "account_id":86000,
    "spkg_id":"1000",
    "pkg_name":"colin",
    "total_capacity":1200,
    "total_remain":1000,
    "capacity_daily":1200,
    "deduct_today":200,
    "expires":"20991231000000",
    "book_time":"20190731172225",
    "last_update":"20190731172225"
}
```


### 3.3 服务包数量增加

传入用户编号、外部交易编号，增加数量进行服务包增加。返回增加记录信息
  
#### 3.3.1 请求参数

| 参数     |  类型  | 可空  |   示例   | 说明             |
| :------- | :----: | :---: | :------: | :--------------- |
| eid      | string |  否   |  62356   | 外部用户编号     |
| sid      | string |  否   |  589766  | 外部服务包编号   |
| trade_no | string |  否   | 86009981 | 外部增加交易编号 |
| num      | number |  否   |  10000   | 增加数量         |


* 请求示例:

```sh
~/work/bin$ curl "http://192.168.4.121:9090/package/capacity/add?eid=86001&sid=1000&trade_no=8970876
&num=200"
```

#### 3.3.2 响应参数

| 参数        |  类型  | 可空  |      示例      |       说明       |
| :---------- | :----: | :---: | :------------: | :--------------: |
| record_id   | number |  否   |     600000     |   变动记录编号   |
| pkg_id      | number |  否   |     620000     |    服务包编号    |
| account_id  | number |  否   |     86000      |     帐户编号     |
| trade_no    | string |  否   |    8970876     | 外部增加交易编号 |
| num         | number |  否   |      100       |     增加数量     |
| remain      | number |  否   |      100       |     剩余数量     |
| change_type | number |  否   |       1        | 变动类型(增加:1) |
| create_time | string |  否   | 20190731172225 |     变动时间     |


* 响应示例:
  
```sh
http.status:200
```

```json
{
    "record_id":"600000",
    "pkg_id":"620000",
    "account_id":"86000",
    "trade_no":"8970876",
    "change_type":"1",
    "create_time":"20190731172225",
    "num":"200",
    "remain":"1200"
}
```

> 请求支持幂等，重复调用返回的内容相同，但`http status`为`201`





### 3.4 服务包数量扣减

传入用户编号、外部交易编号，扣减数量进行服务包扣减。返回扣减记录信息
  
#### 3.4.1 请求参数

| 参数     |  类型  | 可空  |   示例   | 说明             |
| :------- | :----: | :---: | :------: | :--------------- |
| eid      | string |  否   |  62356   | 外部用户编号     |
| sid      | string |  否   |  589766  | 外部服务包编号   |
| trade_no | string |  否   | 86009981 | 外部增加交易编号 |
| num      | number |  否   |  10000   | 扣减数量         |


* 请求示例:

```sh
~/work/bin$ curl "http://192.168.4.121:9090/package/capacity/deduct?eid=86001&sid=1000&trade_no=8970876&num=200"
```

#### 3.4.2 响应参数

| 参数        |  类型  | 可空  |      示例      |       说明       |
| :---------- | :----: | :---: | :------------: | :--------------: |
| record_id   | number |  否   |     600000     |   变动记录编号   |
| pkg_id      | number |  否   |     620000     |    服务包编号    |
| account_id  | number |  否   |     86000      |     帐户编号     |
| trade_no    | string |  否   |    8970876     | 外部扣减交易编号 |
| num         | number |  否   |      100       |     扣减数量     |
| remain      | number |  否   |      100       |     剩余数量     |
| change_type | number |  否   |       1        | 变动类型(扣减:2) |
| create_time | string |  否   | 20190731172225 |     变动时间     |


* 响应示例:
  
```sh
http.status:200
```

```json
{
    "account_id":"86000",
    "change_type":"2",
    "create_time":"20190731172225",
    "num":"-200",
    "pkg_id":"620000",
    "record_id":"600001",
    "remain":"1000",
    "trade_no":"8970876"
}
```



### 3.5 服务包数量退回

传入用户编号、外部扣减交易编号，退回数量进行退回。返回退回记录信息
  
#### 3.5.1 请求参数

| 参数     |  类型  | 可空  |   示例   | 说明             |
| :------- | :----: | :---: | :------: | :--------------- |
| eid      | string |  否   |  62356   | 用户编号         |
| sid      | string |  否   |  589766  | 外部服务包编号   |
| trade_no | string |  否   | 86009981 | 外部扣款交易编号 |
| num      | number |  否   |  10000   | 退回金额,单位分  |


* 请求示例:

```sh
~/work/bin$ curl "http://192.168.4.121:9090/package/capacity/refund?eid=86001&sid=1000&trade_no=8970876&num=200
```

#### 3.5.2 响应参数

| 参数        |  类型  | 可空  |      示例      |       说明       |
| :---------- | :----: | :---: | :------------: | :--------------: |
| record_id   | number |  否   |     600000     |   变动记录编号   |
| pkg_id      | number |  否   |     620000     |    服务包编号    |
| account_id  | number |  否   |     86000      |     帐户编号     |
| trade_no    | string |  否   |    8970876     | 外部扣减交易编号 |
| num         | number |  否   |      100       |     退回数量     |
| remain      | number |  否   |      100       |     剩余数量     |
| change_type | number |  否   |       1        | 变动类型(退回:3) |
| create_time | string |  否   | 20190731172225 |     变动时间     |


* 响应示例:
  
```sh
http.status:200
```

```json
{
    "account_id":"86000",
    "change_type":"3",
    "create_time":"20190731172225",
    "num":"200",
    "pkg_id":"620000",
    "record_id":"600002",
    "remain":"1200",
    "trade_no":"8970876"
}
```



### 3.6 变动记录查询

传入用户编号、开始时间、结束时间等查询资金变动记录
  
#### 3.6.1 请求参数

| 参数       |  类型  | 可空  |      示例      | 说明                        |
| :--------- | :----: | :---: | :------------: | :-------------------------- |
| eid        | string |  否   |     62356      | 用户编号                    |
| sid        | string |  是   |     589766     | 外部服务包编号              |
| start_time | string |  否   | 20190731172225 | 开始时间                    |
| end_time   | string |  否   | 20190731172225 | 结束时间                    |
| pi         | number |  是   |       0        | 第几页，从0开始,未指定默认0 |
| ps         | number |  是   |       10       | 返回的数据行数,未指定默认10 |

* 请求示例:

```sh
curl "http://192.168.4.121:9090/package/record/query?eid=86001&sid=1000&start_time=20190731&end_time=20190731&pi=0&ps=2"
```

#### 3.6.2 响应参数

| 参数        |  类型  | 可空  |      示例      |              说明              |
| :---------- | :----: | :---: | :------------: | :----------------------------: |
| record_id   | number |  否   |       0        |          变动记录编号          |
| pkg_id      | number |  否   |     620000     |           服务包编号           |
| trade_no    | string |  否   |    8970876     |        外部扣款交易编号        |
| num         | number |  否   |      100       |            退回金额            |
| remain      | number |  否   |      100       |          退回后的余额          |
| change_type | number |  否   |       3        | 变动类型(增加:1,扣减:2,退回:3) |
| create_time | string |  否   | 20190731172225 |            变动时间            |


* 响应示例:
  
```sh
http.status:200
```

```json
[
    {
        "change_type":"3",
        "create_time":"20190731172225",
        "num":"200",
        "pkg_id":"620000",
        "record_id":"600002",
        "remain":"1200",
        "trade_no":"8970876"
    },
    {
        "change_type":"2",
        "create_time":"20190731172225",
        "num":"-200",
        "pkg_id":"620000",
        "record_id":"600001",
        "remain":"1000",
        "trade_no":"8970876"
    }
]
```
