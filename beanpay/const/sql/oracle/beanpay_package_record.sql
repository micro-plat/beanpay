
drop table beanpay_package_record;

create table beanpay_package_record
(
	record_id number(20) not null ,
	pkg_id number(20) not null ,
	account_id number(20) not null ,
	trade_no varchar2(32) not null ,
	change_type number(1) not null ,
	num number(20) not null ,
	remain number(20) not null ,
	create_time date default sysdate not null
);


comment on table beanpay_package_record is '服务包数量变动';
	comment on column beanpay_package_record.record_id is '变动编号';	
	comment on column beanpay_package_record.pkg_id is '服务包编号';	
	comment on column beanpay_package_record.account_id is '帐户编号';	
	comment on column beanpay_package_record.trade_no is '外部交易编号';	
	comment on column beanpay_package_record.change_type is '变动类型 1:添加 2：扣除 3：退回';	
	comment on column beanpay_package_record.num is '变动数量';	
	comment on column beanpay_package_record.remain is '剩余数量';	
	comment on column beanpay_package_record.create_time is '创建时间';



alter table beanpay_package_record
	add constraint pk_package_record primary key(record_id);
alter table beanpay_package_record
	add constraint beanpay_pkg_record_account_id unique(account_id,trade_no);

drop sequence seq_package_record_id;

create sequence seq_package_record_id
	minvalue 600000
	maxvalue 99999999999
	start with 600000
	cache 20;
	