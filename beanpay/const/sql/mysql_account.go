// +build !oracle

package sql

//CreateAccount 创建帐户信息
const CreateAccount = `INSERT INTO beanpay_account_info(
	account_name,ident,group,eid,,balance,credit,status,create_time)values(
	@name,@ident,@group,@eid,0,0,0,now())`

//GetAccountByeid 根据eid查询帐户编号
const GetAccountByeid = `select t.account_id,t.account_name,
t.eid,t.balance,t.credit from beanpay_account_info t where 
t.ident=@ident and t.group=@group and t.eid=@eid`

//ChangeAmount 帐户加款
const ChangeAmount = `update beanpay_account_info t set t.balance=t.balance + @amount where t.account_id=@account_id
and t.balance + @amount >= 0`

//ExistsBalanceRecord 查询交易变动记录是否已存在
const ExistsBalanceRecord = `select count(0) from beanpay_account_record t 
where t.trade_no=@trade_no and t.account_id=@account_id
and t.change_type=@tp and t.trade_type=@trade_type`

//GetBalanceRecord 查询交易变动记录是否已存在
const GetBalanceRecord = `select  t.record_id,t.account_id,t.trade_type,
t.trade_no,t.change_type,t.amount,t.balance,DATE_FORMAT(t.create_time, '%Y%m%d%H%i%s') create_time from beanpay_account_record t 
where t.record_id=@record_id`

//GetBalanceRecordByTradeNo 查询交易变动记录是否已存在
const GetBalanceRecordByTradeNo = `select t.record_id,t.account_id,t.trade_type,
t.trade_no,t.change_type,t.amount,t.balance,DATE_FORMAT(t.create_time, '%Y%m%d%H%i%s') create_time
 from beanpay_account_record t 
where t.trade_no=@trade_no 
and t.account_id=@account_id
and t.change_type=@tp
and t.trade_type=@trade_type`

//AddBalanceRecord 添加资金变动
const AddBalanceRecord = `insert into beanpay_account_record
(account_id,trade_no,deduct_no,change_type,amount,balance,create_time,trade_type)
select @account_id,@trade_no,@deduct_no,@tp,@amount,t.balance,now(),@trade_type
 from beanpay_account_info t where t.account_id=@account_id`

//QueryBalanceRecord 查询余额资金变动信息
const QueryBalanceRecord = `select t.record_id,t.account_id,t.trade_type,
t.trade_no,t.change_type,t.amount,t.balance,DATE_FORMAT(t.create_time, '%Y%m%d%H%i%s') create_time
from beanpay_account_record t 
where t.account_id = @account_id 
and t.create_time >= STR_TO_DATE(@start,'%Y%m%d')
and t.create_time < DATE_ADD(STR_TO_DATE(@end,'%Y%m%d'),interval 1 day)
order by t.record_id desc
limit #pf,#ps`

//LockDuductRecord 锁扣款记录
const LockDuductRecord = `
select 
(-1*t.amount) amount
from beanpay_account_record t 
where t.trade_no=@trade_no 
and t.trade_type=@trade_type
and t.account_id=@account_id
and t.change_type=@tp
for update
`

// QueryRefundAmount 查询已退款金额
const QueryRefundAmount = `
select 
sum(t.amount)
from beanpay_account_record t 
where t.account_id=@account_id
and t.trade_type=@trade_type
and t.change_type=@tp
and t.deduct_no=@deduct_no
`
