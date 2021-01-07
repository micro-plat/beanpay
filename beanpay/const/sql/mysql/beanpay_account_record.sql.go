package mysql

const beanpay_account_record = `
DROP TABLE IF EXISTS beanpay_account_record;
	CREATE TABLE IF NOT EXISTS beanpay_account_record (
		record_id BIGINT(20)  not null AUTO_INCREMENT comment '变动编号' ,
		account_id BIGINT(20)  not null  comment '帐户编号' ,
		trade_no VARCHAR(32)  not null  comment '交易编号' ,
		ext_no VARCHAR(32) default 0   comment '拓展编号' ,
		trade_type TINYINT(1) default 1 not null  comment '交易类型 1:交易 2：手续费 3:佣金' ,
		change_type TINYINT(1)  not null  comment '变动类型 1:加款 2：提款 3：扣款 4：退款' ,
		amount decimal(20,5)  not null  comment '变动金额 单位：元' ,
		balance decimal(20,5)  not null  comment '帐户余额 单位：元' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		memo VARCHAR(1024)    comment '交易说明' ,
		ext VARCHAR(1024)    comment '扩展字段' ,
		PRIMARY KEY (record_id),
		UNIQUE KEY beanpay_account_record_account_id (account_id,change_type,trade_no,trade_type)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='账户余额变动信息';`
