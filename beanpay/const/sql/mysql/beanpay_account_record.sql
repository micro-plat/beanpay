
drop table beanpay_account_record;


CREATE TABLE `beanpay_account_record` (
  `record_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '变动编号',
  `account_id` bigint(20) NOT NULL COMMENT '帐户编号',
  `trade_no` varchar(32) NOT NULL COMMENT '交易编号',
  `ext_no` varchar(32) DEFAULT '0' COMMENT '拓展编号',
  `trade_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '交易类型 1:交易 2：手续费 3:佣金',
  `change_type` tinyint(1) NOT NULL COMMENT '变动类型 1:加款 2：提款 3：扣款 4：退款',
  `amount` bigint(20) NOT NULL COMMENT '变动金额 单位：分',
  `balance` bigint(20) NOT NULL COMMENT '帐户余额 单位：分',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ext` varchar(2024) COMMENT '扩展字段',
  PRIMARY KEY (`record_id`),
  UNIQUE KEY `beanpay_account_record_account_id` (`account_id`,`trade_no`,`change_type`,`trade_type`)
) ENGINE=InnoDB AUTO_INCREMENT=8001 DEFAULT CHARSET=utf8mb4 COMMENT='账户余额变动信息';