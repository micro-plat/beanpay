package mysql

const beanpay_package_record = `
DROP TABLE IF EXISTS beanpay_package_record;
	CREATE TABLE IF NOT EXISTS beanpay_package_record (
		record_id BIGINT(20)  not null AUTO_INCREMENT comment '变动编号' ,
		pkg_id BIGINT(20)  not null  comment '服务包编号' ,
		account_id BIGINT(20)  not null  comment '帐户编号' ,
		trade_no VARCHAR(32)  not null  comment '外部交易编号' ,
		change_type TINYINT(1)  not null  comment '变动类型 1:添加,2:减少 3:扣除 4:退回' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		num BIGINT(20)  not null  comment '变动数量' ,
		remain BIGINT(20)  not null  comment '剩余数量' ,
		PRIMARY KEY (record_id),
		UNIQUE KEY beanpay_package_record_account_id (account_id,change_type,trade_no)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='服务包数量变动';`
