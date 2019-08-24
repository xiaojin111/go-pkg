package micro

import (
	"github.com/jinmukeji/go-pkg/mysqldb"
	"github.com/micro/go-micro/config"
)

// NewDbClientFromConfig 通过 Micro Config 的配置创建 DbClient
func NewDbClientFromConfig(cfgKey ...string) (*mysqldb.DbClient, error) {
	opts := mysqldb.NewOptions()
	if err := config.Get(cfgKey...).Scan(&opts); err != nil {
		return nil, err
	}

	return mysqldb.NewDbClientFromOptions(opts)
}
