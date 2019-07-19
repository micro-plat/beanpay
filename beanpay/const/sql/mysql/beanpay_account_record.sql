
 drop table beanpay_account_record;

	create table beanpay_account_record(
		record_id bigint  not null PRIMARY KEY AUTO_INCREMENT  comment '变动编号' ,
		account_id bigint  not null    comment '帐户编号' ,
		trade_no varchar(32)  not null    comment '交易编号' ,
		change_type tinyint(1)  not null    comment '变动类型 1:加款 2：扣款 3：退款' ,
		amount bigint  not null    comment '变动金额 单位：分' ,
		balance bigint  not null    comment '帐户余额 单位：分' ,
		create_time datetime default current_timestamp not null    comment '创建时间' 
				
  )COMMENT='账户余额变动信息';

 




	drop index beanpay_account_record_account_id ON beanpay_account_record;
 create unique index beanpay_account_record_account_id ON beanpay_account_record(account_id,trade_no);
 
