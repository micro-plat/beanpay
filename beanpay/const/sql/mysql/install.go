package mysql

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

func init() {

	//注册服务包
	hydra.DBCli.OnStarting(func(c hydra.ICli) error {
		hydra.Installer.DB.AddSQL(
			beanpay_account_info,
			beanpay_account_record)
		if c.IsSet("pkg") && types.GetBool(c.String("pkg")) {
			hydra.Installer.DB.AddSQL(
				beanpay_package_info,
				beanpay_package_record,
			)
		}
		return nil
	})

}
