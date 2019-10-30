drop table beanpay_account_record;
create table beanpay_account_record
(
	record_id number(20) not null ,
	account_id number(20) not null ,
	trade_no varchar2(32) not null ,
	reduct_no varchar(32) default 0 not null ,
	change_type number(1) not null ,
	trade_type number(1) default 1 not null,
	amount number(20) not null ,
	balance number(20) not null ,
	create_time date default sysdate not null
);


comment on table beanpay_account_record is '账户余额变动信息';
	comment on column beanpay_account_record.record_id is '变动编号';	
	comment on column beanpay_account_record.account_id is '帐户编号';	
	comment on column beanpay_account_record.trade_no is '交易编号';
	comment on column beanpay_account_record.reduct_no is '扣款编号，退款时检查用';	
	comment on column beanpay_account_record.change_type is '变动类型 1:加款 2:提款 3：扣款 4：退款';	
    comment on column beanpay_account_record.trade_type is '交易类型 1:订单交易 2:手续费';	
	comment on column beanpay_account_record.amount is '变动金额 单位：分';	
	comment on column beanpay_account_record.balance is '帐户余额 单位：分';	
	comment on column beanpay_account_record.create_time is '创建时间';



alter table beanpay_account_record
	add constraint pk_account_record primary key(record_id);
alter table beanpay_account_record
	add constraint beanpay_acct_record_account_id unique(account_id,trade_no,change_type);

drop sequence seq_account_record_id;

create sequence seq_account_record_id
	minvalue 100000
	maxvalue 99999999999
	start with 100000
	cache 20;
	