package mysql
 
//beanpay_package_info 服务包信息
const beanpay_package_info=`
	DROP TABLE IF EXISTS beanpay_package_info;
	CREATE TABLE IF NOT EXISTS beanpay_package_info (
		pkg_id bigint  not null auto_increment comment '服务包编号' ,
		account_id bigint  not null  comment '帐户编号' ,
		spkg_id varchar(32)  not null  comment '外部服务包编号' ,
		pkg_name varchar(32)  not null  comment '服务包名称' ,
		total_capacity bigint  not null  comment '总共可用数量' ,
		total_remain bigint  not null  comment '总共剩余数量' ,
		capacity_daily bigint default 0 not null  comment '日限制总数量' ,
		deduct_today bigint default 0 not null  comment '今日扣减数量' ,
		expires datetime  not null  comment '过期日期' ,
		book_time datetime default current_timestamp not null  comment '预订时间' ,
		last_update datetime  not null  comment '上次变更时间' 
		,primary key (pkg_id)
		,unique index beanpay_package_info_account_id(account_id,spkg_id)
	) ENGINE=InnoDB auto_increment = 620000 DEFAULT CHARSET=utf8mb4 COMMENT='服务包信息'`