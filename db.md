# 数据字典

### 1. 账户信息[beanpay_account_info]

| 字段名       | 类型         | 默认值  | 为空  |             约束              | 描述                    |
| ------------ | ------------ | :-----: | :---: | :---------------------------: | :---------------------- |
| account_id   | number(20)   |  86000  |  否   |            PK,SEQ             | 帐户编号                |
| account_name | varchar2(32) |         |  否   |                               | 帐户名称                |
| ident        | varchar2(32) |         |  否   | UNQ(beanpay_account_info_eid) | 系统标识                |
| groups       | varchar2(32) |         |  否   | UNQ(beanpay_account_info_eid) | 用户分组                |
| eid          | varchar2(32) |         |  否   | UNQ(beanpay_account_info_eid) | 外部用户账户编号        |
| balance      | number(20,5) |    0    |  否   |                               | 帐户余额，单位：元      |
| credit       | number(20,5) |    0    |  否   |                               | 信用余额，单位：元      |
| status       | number(1)    |    0    |  否   |                               | 账户状态 0：正常 1:锁定 |
| create_time  | date         | sysdate |  否   |                               | 创建时间                |

### 2. 账户余额变动信息[beanpay_account_record]

| 字段名      | 类型           | 默认值  | 为空  |                  约束                  | 描述                                    |
| ----------- | -------------- | :-----: | :---: | :------------------------------------: | :-------------------------------------- |
| record_id   | number(20)     | 100000  |  否   |                 PK,SEQ                 | 变动编号                                |
| account_id  | number(20)     |         |  否   | UNQ(beanpay_account_record_account_id) | 帐户编号                                |
| trade_no    | varchar2(32)   |         |  否   | UNQ(beanpay_account_record_account_id) | 交易编号                                |
| ext_no      | varchar2(32)   |    0    |  是   |                                        | 拓展编号                                |
| trade_type  | number(1)      |    1    |  否   | UNQ(beanpay_account_record_account_id) | 交易类型 1:交易 2：手续费 3:佣金 4:红冲 5:平账   |
| change_type | number(1)      |         |  否   | UNQ(beanpay_account_record_account_id) | 变动类型 1:加款 2:提款 3：扣款 4：退款 5: 交易平账 6: 余额平账|
| amount      | number(20,5)   |         |  否   |                                        | 变动金额 单位：元                       |
| balance     | number(20,5)   |         |  否   |                                        | 帐户余额 单位：元                       |
| create_time | date           | sysdate |  否   |                                        | 创建时间                                |
| memo        | varchar2(1024) |         |  是   |                                        | 交易说明                                |
| ext         | varchar2(1024) |         |  是   |                                        | 扩展字段                                |

### 3. 服务包信息[beanpay_package_info]

| 字段名         | 类型         | 默认值  | 为空  |                 约束                 | 描述           |
| -------------- | ------------ | :-----: | :---: | :----------------------------------: | :------------- |
| pkg_id         | number(20)   | 620000  |  否   |                PK,SEQ                | 服务包编号     |
| account_id     | number(20)   |         |  否   | UNQ(beanpay_package_info_account_id) | 帐户编号       |
| spkg_id        | varchar2(32) |         |  否   | UNQ(beanpay_package_info_account_id) | 外部服务包编号 |
| pkg_name       | varchar2(32) |         |  否   |                                      | 服务包名称     |
| total_capacity | number(20)   |         |  否   |                                      | 总共可用数量   |
| total_remain   | number(20)   |         |  否   |                                      | 总共剩余数量   |
| capacity_daily | number(20)   |    0    |  否   |                                      | 日限制总数量   |
| deduct_today   | number(20)   |    0    |  否   |                                      | 今日扣减数量   |
| expires        | date         |         |  否   |                                      | 过期日期       |
| book_time      | date         | sysdate |  否   |                                      | 预订时间       |
| last_update    | date         |         |  否   |                                      | 上次变更时间   |

### 4. 服务包数量变动[beanpay_package_record]

| 字段名      | 类型           | 默认值  | 为空  |                  约束                  | 描述                                 |
| ----------- | -------------- | :-----: | :---: | :------------------------------------: | :----------------------------------- |
| record_id   | number(20)     | 600000  |  否   |                 PK,SEQ                 | 变动编号                             |
| pkg_id      | number(20)     |         |  否   |                                        | 服务包编号                           |
| account_id  | number(20)     |         |  否   | UNQ(beanpay_package_record_account_id) | 帐户编号                             |
| trade_no    | varchar2(32)   |         |  否   | UNQ(beanpay_package_record_account_id) | 外部交易编号                         |
| change_type | number(1)      |         |  否   | UNQ(beanpay_package_record_account_id) | 变动类型 1:添加,2:减少 3:扣除 4:退回 |
| create_time | date           | sysdate |  否   |                                        | 创建时间                             |
| num         | number(20)     |         |  否   |                                        | 变动数量                             |
| remain      | number(20)     |         |  否   |                                        | 剩余数量                             |
| ext         | varchar2(1024) |         |  是   |                                        | 扩展字段                             |