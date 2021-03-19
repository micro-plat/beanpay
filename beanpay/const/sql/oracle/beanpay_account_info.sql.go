package oracle

const beanpay_account_info = `create table beanpay_account_info(
		account_id number(20) default 86000 not null ,
		account_name varchar2(32)  not null ,
		ident varchar2(32)  not null ,
		groupx varchar2(32)  not null ,
		eid varchar2(32)  not null ,
		balance number(20,5) default 0 not null ,
		credit number(20,5) default 0 not null ,
		status number(1) default 0 not null ,
		create_time date default sysdate not null 
		);
	

	comment on table beanpay_account_info is '账户信息';
	comment on column beanpay_account_info.account_id is '帐户编号';	
	comment on column beanpay_account_info.account_name is '帐户名称';	
	comment on column beanpay_account_info.ident is '系统标识';	
	comment on column beanpay_account_info.groupx is '用户分组';	
	comment on column beanpay_account_info.eid is '外部用户账户编号';	
	comment on column beanpay_account_info.balance is '帐户余额，单位：元';	
	comment on column beanpay_account_info.credit is '信用余额，单位：元';	
	comment on column beanpay_account_info.status is '账户状态 0：正常 1:锁定';	
	comment on column beanpay_account_info.create_time is '创建时间';	
	

 
	alter table beanpay_account_info
	add constraint pk_beanpay_account_info primary key(account_id);
	alter table beanpay_account_info
	add constraint beanpay_account_info_eid unique(eid,groupx,ident);
	
	create sequence seq_account_info_id
	increment by 1
	minvalue 1
	maxvalue 99999999999
	start with 1
	cache 20;`
