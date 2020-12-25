package oracle

const beanpay_package_info = `create table beanpay_package_info(
		pkg_id number(20) default 620000 not null ,
		account_id number(20)  not null ,
		spkg_id varchar2(32)  not null ,
		pkg_name varchar2(32)  not null ,
		total_capacity number(20)  not null ,
		total_remain number(20)  not null ,
		capacity_daily number(20) default 0 not null ,
		deduct_today number(20) default 0 not null ,
		expires date  not null ,
		book_time date default sysdate not null ,
		last_update date  not null 
		);
	

	comment on table beanpay_package_info is '服务包信息';
	comment on column beanpay_package_info.pkg_id is '服务包编号';	
	comment on column beanpay_package_info.account_id is '帐户编号';	
	comment on column beanpay_package_info.spkg_id is '外部服务包编号';	
	comment on column beanpay_package_info.pkg_name is '服务包名称';	
	comment on column beanpay_package_info.total_capacity is '总共可用数量';	
	comment on column beanpay_package_info.total_remain is '总共剩余数量';	
	comment on column beanpay_package_info.capacity_daily is '日限制总数量';	
	comment on column beanpay_package_info.deduct_today is '今日扣减数量';	
	comment on column beanpay_package_info.expires is '过期日期';	
	comment on column beanpay_package_info.book_time is '预订时间';	
	comment on column beanpay_package_info.last_update is '上次变更时间';	
	

 
	alter table beanpay_package_info
	add constraint pk_beanpay_package_info primary key(pkg_id);
	alter table beanpay_package_info
	add constraint beanpay_package_info_account_id unique(account_id,spkg_id);
	
	create sequence seq_package_info_pkg_id
	increment by 1
	minvalue 1
	maxvalue 99999999999
	start with 1
	cache 20;`
