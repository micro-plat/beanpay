drop table beanpay_account_info;

create table BEANPAY_ACCOUNT_INFO
(
  ACCOUNT_ID   NUMBER(20) not null,
  ACCOUNT_NAME VARCHAR2(32) not null,
  IDENT        VARCHAR2(32) not null,
  GROUPS       VARCHAR2(32) not null,
  EID          VARCHAR2(64) not null,
  BALANCE      NUMBER(20) default 0 not null,
  CREDIT       NUMBER(20) default 0 not null,
  STATUS       NUMBER(1) default 0 not null,
  CREATE_TIME  DATE default sysdate not null
);
-- Add comments to the table 
comment on table BEANPAY_ACCOUNT_INFO
  is '账户信息';
-- Add comments to the columns 
comment on column BEANPAY_ACCOUNT_INFO.ACCOUNT_ID
  is '帐户编号';
comment on column BEANPAY_ACCOUNT_INFO.ACCOUNT_NAME
  is '帐户名称';
comment on column BEANPAY_ACCOUNT_INFO.IDENT
  is '系统标识';
comment on column BEANPAY_ACCOUNT_INFO.GROUPS
  is '用户分组';
comment on column BEANPAY_ACCOUNT_INFO.EID
  is '外部用户账户编号';
comment on column BEANPAY_ACCOUNT_INFO.BALANCE
  is '帐户余额，单位：分';
comment on column BEANPAY_ACCOUNT_INFO.CREDIT
  is '信用余额，单位：分';
comment on column BEANPAY_ACCOUNT_INFO.STATUS
  is '账户状态 0：正常 1:锁定';
comment on column BEANPAY_ACCOUNT_INFO.CREATE_TIME
  is '创建时间';

alter table BEANPAY_ACCOUNT_INFO
  add constraint PK_ACCOUNT_INFO primary key (ACCOUNT_ID);
alter table BEANPAY_ACCOUNT_INFO
  add constraint BEANPAY_ACCOUNT_INFO_EID unique (IDENT, GROUPS, EID);



drop sequence seq_account_info_id;
create sequence seq_account_info_id
	minvalue 86000
	maxvalue 99999999999
	start with 86000
	cache 20;
	