
 drop table beanpay_package_record;

	create table beanpay_package_record(
		record_id bigint  not null PRIMARY KEY AUTO_INCREMENT  comment '变动编号' ,
		pkg_id bigint  not null    comment '服务包编号' ,
		account_id bigint  not null    comment '帐户编号' ,
		trade_no varchar(32)  not null    comment '外部交易编号' ,
		change_type tinyint(1)  not null    comment '变动类型 1:添加 2：扣除 3：退回' ,
		num bigint  not null    comment '变动数量' ,
		remain bigint  not null    comment '剩余数量' ,
		create_time datetime default current_timestamp not null    comment '创建时间' 
				
  )COMMENT='服务包数量变动';

 




	drop index beanpay_package_record_account_id ON beanpay_package_record;
 create unique index beanpay_package_record_account_id ON beanpay_package_record(account_id,trade_no);
 
