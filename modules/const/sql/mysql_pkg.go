// +build !oci

package sql

//CreatePackage 创建帐户信息
const CreatePackage = `INSERT INTO beanpay_package_info(account_id,spkg_id,pkg_name,total_capacity,total_remain,capacity_daily,
	deduct_today,expires,book_time,last_update)values(@account_id,@spkg_id,@name,@total,@total,@daily,0,@expires,now(),now())`

//GetPackageBySPKG 根据spkg_id查询帐户编号
const GetPackageBySPKG = `select t.pkg_id,t.total_remain from beanpay_package_info t 
where t.spkg_id=@spkg_id and t.account_id=@account_id`

//ChangePackage 服务数量变更
const ChangePackage = `update beanpay_package_info t 
set t.total_remain=t.total_remain + @capacity 
where t.pkg_id=@pkg_id
and t.total_remain + @capacity >= 0 
and t.capacity_daily - t.deduct_today + @capacity>0`

//ExistsPackageRecord 查询交易变动记录是否已存在
const ExistsPackageRecord = `select count(0) from beanpay_package_record t where t.trade_no=@trade_no and
 t.pkg_id=@pkg_id`

//AddPackageRecord 添加资金变动
const AddPackageRecord = `insert into beanpay_package_record(pkg_id,account_id,trade_no,change_type,num,remain,create_time)
select t.pkg_id,t.account_id,@trade_no,@tp,@num,t.total_remain,now() from beanpay_package_info t where t.pkg_id=@pkg_id`

//QueryPackageRecord 查询余额资金变动信息
const QueryPackageRecord = `select t.* from beanpay_package_record t 
where t.pkg_id=@pkg_id 
and t.create_time>to_date(@start,'yyyymmddhh24miss')
limit #pi,#ps`
