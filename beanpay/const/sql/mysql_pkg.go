// +build !oracle

package sql

//CreatePackage 创建服务包信息
const CreatePackage = `INSERT INTO beanpay_package_info(account_id,spkg_id,pkg_name,total_capacity,total_remain,capacity_daily,
	deduct_today,expires,book_time,last_update)values(@account_id,@spkg_id,@name,@total,@total,@daily,0,STR_TO_DATE(@expires,'%Y%m%d'),now(),now())`

//GetPackageBySPKG 根据spkg_id查询服务包编号
const GetPackageBySPKG = `select t.pkg_id,t.account_id,t.spkg_id,t.pkg_name,t.total_capacity,
t.total_remain,t.capacity_daily,t.deduct_today,
DATE_FORMAT(t.expires, '%Y%m%d%H%i%s') expires,
DATE_FORMAT(t.book_time, '%Y%m%d%H%i%s') book_time,
DATE_FORMAT(t.last_update, '%Y%m%d%H%i%s') last_update
from beanpay_package_info t 
where t.spkg_id=@spkg_id and t.account_id=@account_id`

//ChangePackage 服务数量变更
const ChangePackage = `update beanpay_package_info t 
set t.total_remain=t.total_remain + @capacity ,
t.total_capacity=t.total_capacity+@total,
t.deduct_today=t.deduct_today + @capacity
where t.pkg_id=@pkg_id
and t.total_remain + @capacity >= 0 
and (t.capacity_daily = 0 or t.capacity_daily - t.deduct_today + @capacity>0)`

//ExistsPackageRecord 查询交易变动记录是否已存在
const ExistsPackageRecord = `select count(0) from beanpay_package_record t 
where t.trade_no=@trade_no and t.pkg_id=@pkg_id
and t.change_type=@change_type
and abs(t.num) >= @max_num`

//GetPackageRecordByTradeNo 查询交易变动记录是否已存在
const GetPackageRecordByTradeNo = `select t.record_id,t.pkg_id,t.account_id,
t.trade_no,t.change_type,t.num,t.remain,DATE_FORMAT(t.create_time, '%Y%m%d%H%i%s') create_time
 from beanpay_package_record t 
where t.trade_no=@trade_no and pkg_id=@pkg_id
and t.change_type=@change_type`

//AddPackageRecord 添加服务包变动记录
const AddPackageRecord string = `insert into beanpay_package_record
(pkg_id,account_id,trade_no,change_type,num,remain,create_time,ext)
select t.pkg_id,t.account_id,@trade_no,@change_type,@capacity,t.total_remain,now(),@ext from beanpay_package_info t
 where t.pkg_id=@pkg_id`

//QueryPackageRecord 查询务包变动记录
const QueryPackageRecord string = `select t.record_id,t.pkg_id,t.trade_no,t.change_type,t.num,t.remain,
DATE_FORMAT(t.create_time, '%Y%m%d%H%i%s') create_time
from beanpay_package_record t 
where t.account_id=@account_id
and (t.pkg_id=@pkg_id or @pkg_id=0)
and t.create_time >= STR_TO_DATE(@start,'%Y%m%d')
and t.create_time < DATE_ADD(STR_TO_DATE(@end,'%Y%m%d'),interval 1 day)
order by t.record_id desc
limit #pf,#ps`
