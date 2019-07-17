// +build !oci

package sql

//CreateAccount 创建帐户信息
const CreateAccount = `INSERT INTO beanpay_account_info(account_name,uaid,balance,credit,status,create_time)values(
	@account_name,@uaid,0,0,0,now())`

//GetAccountByUaid 根据uaid查询帐户编号
const GetAccountByUaid = `select t.account_id,t.balance from beanpay_account_info t where t.uaid=@uaid`

//ChangeAmount 帐户加款
const ChangeAmount = `update beanpay_account_info t set t.balance=t.balance + @amount where t.account_id=@account_id
and t.balance + @amount >= 0`

//ExistsBalanceRecord 查询交易变动记录是否已存在
const ExistsBalanceRecord = `select count(0) from beanpay_account_record t where t.trade_no=@trade_no and t.account_id=@account_id`

//AddBalanceRecord 添加资金变动
const AddBalanceRecord = `insert into beanpay_account_record(account_id,trade_no,change_type,amount,balance,create_time)
select @account_id,@trade_no,@tp,@amount,t.balance,now() from beanpay_account_info t where t.account_id=@account_id`

//QueryBalanceRecord 查询余额资金变动信息
const QueryBalanceRecord = `select t.* from beanpay_account_record t 
where t.account_id=@account_id 
and t.create_time>to_date(@start,'yyyymmddhh24miss')
limit #pi,#ps`
