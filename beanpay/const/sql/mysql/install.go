
package mysql

import (
	"github.com/micro-plat/hydra"
	_ "github.com/go-sql-driver/mysql"
)
		
func init() {
	//注册服务包
	hydra.OnReadying(func() error {
		hydra.Installer.DB.AddSQL(
		beanpay_account_info,
		beanpay_account_record,
		beanpay_package_info,
		beanpay_package_record,
		
		)
		return nil
	}) 
}
