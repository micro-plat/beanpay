// +build oracle

package sql

//CreateAccount 创建帐户信息
const CreateAccount = `INSERT INTO beanpay_account_info(account_id,account_name,
	ident,groups,eid,balance,credit,status,create_time)values(seq_account_info_id.nextval,
	@name,@ident,@groups,@eid,0,0,0,sysdate)`

//UpdateAccount 修改帐户信息
const UpdateAccount = `
update beanpay_account_info t
set t.account_name = @name
where t.ident = @ident
and t.groups = @groups
and t.eid = @eid
`

//SetCreditAmount 设置授信金额
const SetCreditAmount = `UPDATE 
beanpay_account_info b 
SET
b.credit = @credit 
WHERE b.account_id = @account_id`

//GetAccountByeid 根据eid查询帐户编号
const GetAccountByeid = `select t.account_id,t.account_name,t.eid,t.balance,
t.credit from beanpay_account_info t where
t.ident=@ident and t.groups=@groups and t.eid=@eid`

//ChangeAmount 帐户加款
const ChangeAmount = `update beanpay_account_info t set t.balance=t.balance + @amount where t.account_id=@account_id
and t.balance + t.credit + @amount >= 0`

//ExistsBalanceRecord 查询交易变动记录是否已存在
const ExistsBalanceRecord = `select count(0) from beanpay_account_record t 
where t.trade_no=@trade_no 
and t.trade_type=@trade_type
and t.account_id=@account_id
and t.change_type=@change_type`

//GetBalanceRecord 查询交易变动记录是否已存在
const GetBalanceRecord = `select  t.record_id,t.account_id,t.trade_type,t.memo,
t.trade_no,t.change_type,t.amount,t.balance,to_char(t.create_time, 'yyyymmddhh24miss') create_time from beanpay_account_record t 
where t.record_id=@record_id`

//GetBalanceRecordByTradeNo 查询交易变动记录是否已存在
const GetBalanceRecordByTradeNo = `select t.record_id,t.account_id,t.trade_type,t.memo,
t.trade_no,t.change_type,t.amount,t.balance,to_char(t.create_time, 'yyyymmddhh24miss') create_time
 from beanpay_account_record t 
where t.trade_no=@trade_no and t.account_id=@account_id
and t.change_type=@change_type and t.trade_type=@trade_type`

//AddBalanceRecord 添加资金变动
const AddBalanceRecord = `insert into beanpay_account_record
(record_id,account_id,trade_no,ext_no,change_type,amount,balance,create_time,trade_type,ext,memo)
select seq_account_record_id.nextval,@account_id,@trade_no,@ext_no,@change_type,@amount,t.balance,sysdate,@trade_type,@ext,@memo
 from beanpay_account_info t where t.account_id=@account_id`

//QueryBalanceRecordCount 查询余额资金变动信息
const QueryBalanceRecordCount = `select 
l2.record_id,l2.account_id,l2.memo,
l2.trade_no,l2.change_type,l2.amount,l2.balance,l2.create_time
from(select rownum rn,l1.* from(select t.record_id,t.account_id,t.memo,
t.trade_no,t.change_type,t.amount,t.balance,t.trade_type,to_char(t.create_time, 'yyyymmddhh24miss') create_time from  
beanpay_account_record t where t.account_id = @account_id and 
t.create_time >= to_date(@start,'yyyymmdd')
and t.create_time < to_date(@end,'yyyymmdd')+1
`

//QueryBalanceRecord 查询余额资金变动信息
const QueryBalanceRecord = `select count(1) from  
beanpay_account_record t where t.account_id = @account_id and 
t.create_time >= to_date(@start,'yyyymmdd')
and t.create_time < to_date(@end,'yyyymmdd')+1
order by t.record_id desc) l1
where rownum <= (@pi+1) * @ps) l2 
where l2.rn > (@pi) * @ps`

// LockAccount 锁账户
const LockAccount = `
SELECT 
  a.account_id
FROM
  beanpay_account_info a 
WHERE a.account_id = @account_id 
FOR UPDATE
`

//LockTradeRecord 锁交易记录
const LockTradeRecord = `
select 
(-1*t.amount) amount 
from beanpay_account_record t 
where t.trade_no=@trade_no 
and t.account_id=@account_id
and t.change_type=@change_type
and t.trade_type=@trade_type
for update
`

// QueryTradedAmount 查询已交易金额
const QueryTradedAmount = `
select 
sum(t.amount)
from beanpay_account_record t 
where t.account_id=@account_id
and t.change_type=@change_type
and t.trade_type=@trade_type
and t.ext_no=@ext_no
`

//QueryAccountListCount 获取账户信息列表条数
const QueryAccountListCount = `
select count(1)
from beanpay_account_info t
where  t.groups like '%' || @types || '%'
 &t.account_name
 &t.eid
 &t.groups
 &t.ident
 &t.status`

//QueryAccountList 查询账户信息列表数据
const QueryAccountList = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.account_id,
                               t.account_name,
                               t.ident,
                               t.groups,
                               t.eid,
                               t.balance,
                               t.credit,
                               to_char(t.create_time,'yyyy-MM-dd HH24:mi:ss') create_time,
                               t.status
                          from beanpay_account_info t
                         where t.groups like '%' || @types || '%'
                         &t.account_name &t.eid 
                         &t.groups  &t.ident
                         &t.status
                         order by t.account_id desc) R
                 where rownum <= @size) L
         where L.rn > @currentPage) TAB1
`
