
drop table beanpay_account_info;

CREATE TABLE `beanpay_account_info` (
  `account_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '帐户编号',
  `account_name` varchar(32) NOT NULL COMMENT '帐户名称',
  `ident` varchar(32) NOT NULL COMMENT '系统标识',
  `groups` varchar(32) NOT NULL COMMENT '用户分组',
  `eid` varchar(32) NOT NULL COMMENT '外部用户账户编号',
  `balance` bigint(20) NOT NULL DEFAULT '0' COMMENT '帐户余额，单位：分',
  `credit` bigint(20) NOT NULL DEFAULT '0' COMMENT '信用余额，单位：分',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '账户状态 0：正常 1:锁定',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`account_id`),
  UNIQUE KEY `beanpay_account_info_eid` (`ident`,`groups`,`eid`)
) ENGINE=InnoDB AUTO_INCREMENT=86000 DEFAULT CHARSET=utf8mb4 COMMENT='账户信息';
 
