package readconfig

import (
	"github.com/kamran0812/xcelldb/persist"
	"github.com/spf13/viper"
)

func Load() map[string]int {
	viper.AddConfigPath("./readConfig")
	viper.SetConfigName("test") // Register config file name (no extension)
	viper.SetConfigType("yaml") // Look for specific type
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return persist.ColList(config.Columns)

}
