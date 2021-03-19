package mysql
 
//beanpay_account_info 账户信息
const beanpay_account_info=`
	DROP TABLE IF EXISTS beanpay_account_info;
	CREATE TABLE IF NOT EXISTS beanpay_account_info (
		account_id bigint  not null auto_increment comment '帐户编号' ,
		account_name varchar(32)  not null  comment '帐户名称' ,
		ident varchar(32)  not null  comment '系统标识' ,
		groupx varchar(32)  not null  comment '用户分组' ,
		eid varchar(32)  not null  comment '外部用户账户编号' ,
		balance bigint default 0 not null  comment '帐户余额，单位：分' ,
		credit bigint default 0 not null  comment '信用余额，单位：分' ,
		status tinyint default 0 not null  comment '账户状态 0：正常 1:锁定' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (account_id)
		,unique index beanpay_account_info_eid(ident,groupx,eid)
	) ENGINE=InnoDB auto_increment = 86000 DEFAULT CHARSET=utf8mb4 COMMENT='账户信息'`