package mysql

const beanpay_package_info = `
DROP TABLE IF EXISTS beanpay_package_info;
	CREATE TABLE IF NOT EXISTS beanpay_package_info (
		pkg_id BIGINT(20)  not null AUTO_INCREMENT comment '服务包编号' ,
		account_id BIGINT(20)  not null  comment '帐户编号' ,
		spkg_id VARCHAR(32)  not null  comment '外部服务包编号' ,
		pkg_name VARCHAR(32)  not null  comment '服务包名称' ,
		total_capacity BIGINT(20)  not null  comment '总共可用数量' ,
		total_remain BIGINT(20)  not null  comment '总共剩余数量' ,
		capacity_daily BIGINT(20) default 0 not null  comment '日限制总数量' ,
		deduct_today BIGINT(20) default 0 not null  comment '今日扣减数量' ,
		expires DATETIME  not null  comment '过期日期' ,
		book_time DATETIME default current_timestamp not null  comment '预订时间' ,
		last_update DATETIME  not null  comment '上次变更时间' ,
		PRIMARY KEY (pkg_id),
		UNIQUE KEY beanpay_package_info_account_id (account_id,spkg_id)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='服务包信息';`
