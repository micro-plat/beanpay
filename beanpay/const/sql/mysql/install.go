package mysql

import "github.com/micro-plat/hydra"

func init() {
	hydra.Installer.DB.AddSQL(
		beanpay_account_info,
		beanpay_account_record,
		beanpay_package_info,
		beanpay_package_record,
	)
}
