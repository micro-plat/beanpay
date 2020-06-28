// +build oracle

package sql

//CreatePackage 创建服务包信息
const CreatePackage = `INSERT INTO beanpay_package_info
(pkg_id,account_id,spkg_id,pkg_name,total_capacity,total_remain,capacity_daily,deduct_today,expires,book_time,last_update)values
(seq_package_info_pkg_id.nextval,@account_id,@spkg_id,@name,@total,@total,@daily,0,to_date(@expires,'yyyymmdd'),sysdate,sysdate)`

//GetPackageBySPKG 根据spkg_id查询服务包编号
const GetPackageBySPKG = `select t.pkg_id,t.account_id,t.spkg_id,t.pkg_name,t.total_capacity,
t.total_remain,t.capacity_daily,t.deduct_today,
to_char(t.expires, 'yyyymmddhh24miss') expires,
to_char(t.book_time, 'yyyymmddhh24miss') book_time,
to_char(t.last_update, 'yyyymmddhh24miss') last_update
from beanpay_package_info t 
where t.spkg_id=@spkg_id and t.account_id=@account_id`

//ChangePackage 服务数量变更
const ChangePackage = `update beanpay_package_info t 
set t.total_remain=t.total_remain + @capacity ,
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
t.trade_no,t.change_type,t.num,t.remain,to_char(t.create_time, 'yyyymmddhh24miss') create_time
 from beanpay_package_record t 
where t.trade_no=@trade_no and pkg_id=@pkg_id
and t.change_type=@change_type`

//AddPackageRecord 添加服务包变动记录
const AddPackageRecord string = `insert into beanpay_package_record(record_id,pkg_id,account_id,trade_no,change_type,num,remain,create_time,ext)
select seq_package_record_id.nextval,t.pkg_id,t.account_id,@trade_no,@change_type,@capacity,t.total_remain,sysdate,@ext from beanpay_package_info t where t.pkg_id=@pkg_id`

//QueryPackageRecord 查询余额资金变动信息
const QueryPackageRecord = `select 
l2.record_id,l2.pkg_id,l2.trade_no,l2.change_type,l2.num,l2.remain, l2.create_time
from(select rownum rn,l1.* from(	
	select t.record_id,t.pkg_id,t.trade_no,t.change_type,t.num,t.remain,
to_char(t.create_time, 'yyyymmddhh24miss') create_time
from beanpay_package_record t 
where t.account_id=@account_id
and (t.pkg_id=@pkg_id or @pkg_id=0)
and t.create_time >= to_date(@start,'yyyymmdd')
and t.create_time < to_date(@end,'yyyymmdd')+1
order by t.record_id desc
) l1
where rownum <= (@pi+1) * @ps) l2 
where l2.rn > (@pi) * @ps`
