package mysql
 
//beanpay_account_record 账户余额变动信息
const beanpay_account_record=`
	DROP TABLE IF EXISTS beanpay_account_record;
	CREATE TABLE IF NOT EXISTS beanpay_account_record (
		record_id bigint  not null auto_increment comment '变动编号' ,
		account_id bigint  not null  comment '帐户编号' ,
		trade_no varchar(32)  not null  comment '交易编号' ,
		ext_no varchar(32) default 0   comment '拓展编号' ,
		trade_type tinyint default 1 not null  comment '交易类型 1:交易 2：手续费 3:佣金' ,
		change_type tinyint  not null  comment '变动类型 1:加款 2：提款 3：扣款 4：退款 5: 交易平账 6: 余额平账' ,
		amount decimal(20,5)  not null  comment '变动金额 单位：元' ,
		balance decimal(20,5)  not null  comment '帐户余额 单位：元' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		memo varchar(1024)    comment '交易说明' ,
		ext varchar(1024)    comment '扩展字段' 
		,primary key (record_id)
		,unique index beanpay_account_record_account_id(account_id,trade_no,trade_type,change_type)
	) ENGINE=InnoDB auto_increment = 100000 DEFAULT CHARSET=utf8mb4 COMMENT='账户余额变动信息'`