
drop table beanpay_account_info;

create table beanpay_account_info
(
	account_id bigint not null PRIMARY KEY AUTO_INCREMENT  comment
	'帐户编号' ,
		account_name varchar
	(32)  not null    comment '帐户名称' ,
		eid varchar
	(32)  not null    comment '外部用户账户编号' ,
	account_type varchar
	(32) not null    comment
'外部账户类型' ,
		balance bigint default 0 not null    comment '帐户余额，单位：分' ,
		credit bigint default 0 not null    comment '信用余额，单位：分' ,
		status tinyint
	(1) default 0 not null    comment '账户状态 0：正常 1:锁定' ,
		create_time datetime default current_timestamp not null    comment '创建时间' 
				
  )COMMENT='账户信息';


	alter table beanpay_account_info AUTO_INCREMENT=86000;



	drop index beanpay_account_info_eid ON beanpay_account_info;
	create unique index beanpay_account_info_eid ON beanpay_account_info(eid);
 
