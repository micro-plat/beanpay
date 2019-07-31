drop table beanpay_account_info;
create table beanpay_account_info
(
	account_id number(20) not null ,
	account_name varchar2(32) not null ,
	eid varchar2(32) not null ,
	balance number(20) default 0 not null ,
	credit number(20) default 0 not null ,
	status number(1) default 0 not null ,
	create_time date default sysdate not null
);


comment on table beanpay_account_info is '账户信息';
	comment on column beanpay_account_info.account_id is '帐户编号';	
	comment on column beanpay_account_info.account_name is '帐户名称';	
	comment on column beanpay_account_info.eid is '外部用户账户编号';	
	comment on column beanpay_account_info.balance is '帐户余额，单位：分';	
	comment on column beanpay_account_info.credit is '信用余额，单位：分';	
	comment on column beanpay_account_info.status is '账户状态 0：正常 1:锁定';	
	comment on column beanpay_account_info.create_time is '创建时间';



alter table beanpay_account_info
	add constraint pk_account_info primary key(account_id);
alter table beanpay_account_info
	add constraint beanpay_account_info_eid unique(eid);



drop sequence seq_account_info_id;
create sequence seq_account_info_id
	minvalue 86000
	maxvalue 99999999999
	start with 86000
	cache 20;
	