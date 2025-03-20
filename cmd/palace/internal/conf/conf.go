package conf

import (
	"fmt"

	"github.com/moon-monitor/moon/pkg/config"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func (c *Bootstrap) IsDev() bool {
	return c.GetEnvironment() == config.Environment_DEV
}

func (c *Data_Database) GenDsn(dbName string) string {
	dsn := c.GetDsn()
	switch c.GetDriver() {
	case config.Database_MYSQL:
		if validate.TextIsNull(dsn) {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.GetUser(), c.GetPassword(), c.GetHost(), c.GetPort(), dbName, c.GetParams())
		}
	}
	return dsn
}

func (c *Data_Database) IsGroup() bool {
	return !validate.TextIsNull(c.GetDsn())
}
