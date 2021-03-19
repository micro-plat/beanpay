package mysql
 
//beanpay_package_record 服务包数量变动
const beanpay_package_record=`
	DROP TABLE IF EXISTS beanpay_package_record;
	CREATE TABLE IF NOT EXISTS beanpay_package_record (
		record_id bigint  not null auto_increment comment '变动编号' ,
		pkg_id bigint  not null  comment '服务包编号' ,
		account_id bigint  not null  comment '帐户编号' ,
		trade_no varchar(32)  not null  comment '外部交易编号' ,
		change_type tinyint  not null  comment '变动类型 1:添加,2:减少 3:扣除 4:退回' ,
		create_time datetime default current_timestamp not null  comment '创建时间' ,
		num bigint  not null  comment '变动数量' ,
		remain bigint  not null  comment '剩余数量' ,
		ext varchar(1024)    comment '扩展字段' 
		,primary key (record_id)
		,unique index beanpay_package_record_account_id(account_id,trade_no,change_type)
	) ENGINE=InnoDB auto_increment = 600000 DEFAULT CHARSET=utf8mb4 COMMENT='服务包数量变动'`