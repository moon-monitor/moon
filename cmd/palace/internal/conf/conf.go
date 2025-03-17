package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"buf.build/go/protoyaml"
	"github.com/joho/godotenv"

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
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.GetUser(), c.GetPassword(), c.GetHost(), c.GetPort(), dbName)
		}
	}
	return dsn
}

func (c *Data_Database) IsGroup() bool {
	return !validate.TextIsNull(c.GetDsn())
}

func Load(cfgPath string) (*Bootstrap, error) {
	_ = godotenv.Load()

	var cfg Bootstrap
	err := filepath.Walk(cfgPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml") {
			yamlBytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			err = protoyaml.UnmarshalOptions{
				Path: cfgPath,
			}.Unmarshal(resolveEnv(yamlBytes), &cfg)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func resolveEnv(content []byte) []byte {
	regex := regexp.MustCompile(`\$\{(\w+)(?::([^}]+))?}`)

	return regex.ReplaceAllFunc(content, func(match []byte) []byte {
		matches := regex.FindSubmatch(match)
		envKey := string(matches[1])
		var defaultValue string

		if len(matches) > 2 {
			defaultValue = string(matches[2])
		}

		if value, exists := os.LookupEnv(envKey); exists {
			return []byte(value)
		}
		return []byte(defaultValue)
	})
}
