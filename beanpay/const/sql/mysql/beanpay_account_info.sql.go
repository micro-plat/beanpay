package mysql

const beanpay_account_info = `
	DROP TABLE IF EXISTS beanpay_account_info;
	CREATE TABLE IF NOT EXISTS beanpay_account_info (
		account_id BIGINT(20)  not null AUTO_INCREMENT comment '帐户编号' ,
		account_name VARCHAR(32)  not null  comment '帐户名称' ,
		ident VARCHAR(32)  not null  comment '系统标识' ,
		groups VARCHAR(32)  not null  comment '用户分组' ,
		eid VARCHAR(32)  not null  comment '外部用户账户编号' ,
		balance decimal(20,5) default 0 not null  comment '帐户余额，单位：元' ,
		credit decimal(20,5) default 0 not null  comment '信用余额，单位：元' ,
		status TINYINT(1) default 0 not null  comment '账户状态 0：正常 1:锁定' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (account_id),
		UNIQUE KEY beanpay_account_info_eid (eid,groups,ident)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='账户信息';
`
