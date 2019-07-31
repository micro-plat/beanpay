// +build !oci

package sql

//CreateAccount 创建帐户信息
const CreateAccount = `INSERT INTO beanpay_account_info(account_name,eid,balance,credit,status,create_time)values(
	@name,@eid,0,0,0,now())`

//GetAccountByeid 根据eid查询帐户编号
const GetAccountByeid = `select t.account_id,t.account_name,t.eid,t.balance,t.credit from beanpay_account_info t where t.eid=@eid`

//ChangeAmount 帐户加款
const ChangeAmount = `update beanpay_account_info t set t.balance=t.balance + @amount where t.account_id=@account_id
and t.balance + @amount >= 0`

//ExistsBalanceRecord 查询交易变动记录是否已存在
const ExistsBalanceRecord = `select count(0) from beanpay_account_record t 
where t.trade_no=@trade_no and t.account_id=@account_id
and t.change_type=@tp
and abs(t.amount) >= @max_amount`

//GetBalanceRecord 查询交易变动记录是否已存在
const GetBalanceRecord = `select  t.record_id,t.account_id,
t.trade_no,t.change_type,t.amount,t.balance,DATE_FORMAT(t.create_time, '%Y%m%d%H%i%s') create_time from beanpay_account_record t 
where t.record_id=@record_id`

//GetBalanceRecordByTradeNo 查询交易变动记录是否已存在
const GetBalanceRecordByTradeNo = `select t.record_id,t.account_id,
t.trade_no,t.change_type,t.amount,t.balance,DATE_FORMAT(t.create_time, '%Y%m%d%H%i%s') create_time
 from beanpay_account_record t 
where t.trade_no=@trade_no and t.account_id=@account_id
and t.change_type=@tp`

//AddBalanceRecord 添加资金变动
const AddBalanceRecord = `insert into beanpay_account_record
(account_id,trade_no,change_type,amount,balance,create_time)
select @account_id,@trade_no,@tp,@amount,t.balance,now()
 from beanpay_account_info t where t.account_id=@account_id`

//QueryBalanceRecord 查询余额资金变动信息
const QueryBalanceRecord = `select t.record_id,t.account_id,
t.trade_no,t.change_type,t.amount,t.balance,DATE_FORMAT(t.create_time, '%Y%m%d%H%i%s') create_time
from beanpay_account_record t 
where t.account_id = @account_id 
and t.create_time >= STR_TO_DATE(@start,'%Y%m%d')
and t.create_time < DATE_ADD(STR_TO_DATE(@end,'%Y%m%d'),interval 1 day)
order by t.record_id desc
limit #pf,#ps`
