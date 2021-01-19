// +build !oracle

package sql

import _ "github.com/micro-plat/beanpay/beanpay/const/sql/mysql"

//CreateAccount 创建帐户信息
const CreateAccount = `INSERT INTO beanpay_account_info(
	account_name,ident,groups,eid,balance,credit,status,create_time)values(
	@name,@ident,@groups,@eid,0,0,0,now())`

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
const GetAccountByeid = `select t.account_id,t.account_name,
t.eid,ifnull(t.balance,0) balance,ifnull(t.credit,0) credit from beanpay_account_info t where 
t.ident=@ident and t.groups=@groups and t.eid=@eid`

//ChangeAmount 帐户加款
const ChangeAmount = `update beanpay_account_info t set t.balance=t.balance + @amount where t.account_id=@account_id
and t.balance + t.credit + @amount >= 0 `

//ExistsBalanceRecord 查询交易变动记录是否已存在
const ExistsBalanceRecord = `select count(0) from beanpay_account_record t 
where t.trade_no=@trade_no and t.account_id=@account_id
and t.change_type=@change_type and t.trade_type=@trade_type`

//GetBalanceRecord 查询交易变动记录是否已存在
const GetBalanceRecord = `select  t.record_id,t.account_id,t.trade_type,t.memo,
t.trade_no,t.change_type,t.amount,t.balance,DATE_FORMAT(t.create_time, '%Y%m%d%H%i%s') create_time from beanpay_account_record t 
where t.record_id=@record_id`

//GetBalanceRecordByTradeNo 查询交易变动记录是否已存在
const GetBalanceRecordByTradeNo = `select t.record_id,t.account_id,t.trade_type,t.memo,
t.trade_no,t.change_type,t.amount,t.balance balance,DATE_FORMAT(t.create_time, '%Y%m%d%H%i%s') create_time
 from beanpay_account_record t 
where t.trade_no=@trade_no 
and t.account_id=@account_id
and t.change_type=@change_type
and t.trade_type=@trade_type`

//AddBalanceRecord 添加资金变动
const AddBalanceRecord = `insert into beanpay_account_record
(account_id,trade_no,ext_no,change_type,amount,balance,create_time,trade_type,ext,memo)
select @account_id,@trade_no,@ext_no,@change_type,@amount,t.balance,now(),@trade_type,@ext,@memo
 from beanpay_account_info t where t.account_id=@account_id`

//QueryBalanceRecordCount 查询余额资金变动信息
const QueryBalanceRecordCount = `select count(1)
from beanpay_account_record t 
INNER JOIN beanpay_account_info a ON a.account_id = t.account_id
where t.create_time >= DATE_FORMAT(@start,'%Y%m%d')
and t.create_time < DATE_ADD(DATE_FORMAT(@end,'%Y%m%d'),interval 1 day)
and a.groups like CONCAT('',@types,'%')
and a.account_name like concat('%',@account_name,'%')
&t.account_id &t.change_type &t.trade_type &a.groups &a.ident
`

//QueryBalanceRecord 查询余额资金变动信息
const QueryBalanceRecord = `select t.record_id,t.account_id,t.memo,t.trade_type,
t.trade_no,t.change_type,t.amount,t.balance,t.create_time,a.account_name,a.eid,a.groups
from beanpay_account_record t 
INNER JOIN beanpay_account_info a ON a.account_id = t.account_id
where  t.create_time >= DATE_FORMAT(@start,'%Y%m%d')
and t.create_time < DATE_ADD(DATE_FORMAT(@end,'%Y%m%d'),interval 1 day)
and a.groups like CONCAT('',@types,'%')
and a.account_name like concat('%',@account_name,'%')
&t.account_id &t.change_type &t.trade_type &a.groups
order by t.record_id desc
limit #pageSize offset #currentPage
`

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
and t.trade_type=@trade_type
and t.account_id=@account_id
and t.change_type=@change_type
for update
`

// QueryTradedAmount 查询已交易金额
const QueryTradedAmount = `
select 
sum(t.amount)
from beanpay_account_record t 
where t.account_id=@account_id
and t.trade_type=@trade_type
and t.change_type=@change_type
and t.ext_no=@ext_no
`

// CheckRefundAmount 检查退款金额
const CheckRefundAmount = `
SELECT 
  IF(ABS(@deduct_amount) - ABS(IFNULL(SUM(t.amount),0)) - ABS(@amount) >=0,TRUE,FALSE) can_refund,
  IFNULL(SUM(t.amount),0) refund_amount 
FROM
  beanpay_account_record t 
WHERE t.account_id = @account_id 
  AND t.trade_type = @trade_type 
  AND t.change_type = @change_type 
  AND t.ext_no = @ext_no 
`

//QueryAccountListCount 获取账户信息列表条数
const QueryAccountListCount = `
select count(1)
from beanpay_account_info t
where  t.groups like CONCAT('',@types,'%')
and if(isnull(@account_name)||@account_name='',1=1,t.account_name like concat('%',@account_name,'%'))
and if(isnull(@eid)||@eid='',1=1,t.eid=@eid)
and if(isnull(@groups)||@groups='',1=1,t.groups=@groups)
and if(isnull(@ident)||@ident='',1=1,t.ident=@ident)
and if(isnull(@status)||@status='',1=1,t.status=@status)`

//QueryAccountList 查询账户信息列表数据
const QueryAccountList = `
select
	t.account_id,
	t.account_name,
	t.ident,
	t.groups,
    t.eid,
    ifnull(t.balance,0) balance,
    ifnull(t.credit,0) credit,
    t.create_time,
    t.status
    from beanpay_account_info t
where t.groups like CONCAT('',@types,'%')
and if(isnull(@eid)||@eid='',1=1,t.eid=@eid)
and if(isnull(@account_name)||@account_name='',1=1,t.account_name like concat('%',@account_name,'%'))
and if(isnull(@groups)||@groups='',1=1,t.groups=@groups)
and if(isnull(@ident)||@ident='',1=1,t.ident=@ident)
and if(isnull(@status)||@status='',1=1,t.status=@status)
order by t.account_id desc
limit #pageSize offset #currentPage
`
