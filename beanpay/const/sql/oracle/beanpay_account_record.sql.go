package oracle

const beanpay_account_record = `create table beanpay_account_record(
		record_id number(20) default 100000 not null ,
		account_id number(20)  not null ,
		trade_no varchar2(32)  not null ,
		ext_no varchar2(32) default 0  ,
		trade_type number(1) default 1 not null ,
		change_type number(1)  not null ,
		amount number(20,5)  not null ,
		balance number(20,5)  not null ,
		create_time date default sysdate not null ,
		memo varchar2(1024)   ,
		ext varchar2(1024)   
		);
	

	comment on table beanpay_account_record is '账户余额变动信息';
	comment on column beanpay_account_record.record_id is '变动编号';	
	comment on column beanpay_account_record.account_id is '帐户编号';	
	comment on column beanpay_account_record.trade_no is '交易编号';	
	comment on column beanpay_account_record.ext_no is '拓展编号';	
	comment on column beanpay_account_record.trade_type is '交易类型 1:交易 2：手续费 3:佣金';	
	comment on column beanpay_account_record.change_type is '变动类型 1:加款 2：提款 3：扣款 4：退款';	
	comment on column beanpay_account_record.amount is '变动金额 单位：元';	
	comment on column beanpay_account_record.balance is '帐户余额 单位：元';	
	comment on column beanpay_account_record.create_time is '创建时间';	
	comment on column beanpay_account_record.memo is '交易说明';	
	comment on column beanpay_account_record.ext is '扩展字段';	
	

 
	alter table beanpay_account_record
	add constraint pk_beanpay_account_record primary key(record_id);
	alter table beanpay_account_record
	add constraint beanpay_record_account_id unique(account_id,change_type,trade_no,trade_type);
	
	create sequence seq_account_record_id
	increment by 1
	minvalue 1
	maxvalue 99999999999
	start with 1
	cache 20;`
