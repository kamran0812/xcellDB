package readconfig

type Config struct {
	Columns    []string `mapstructure:"DB_COLUMNS"`
	Connection string   `mapstructure:"DB_CONNECTION"`
}
