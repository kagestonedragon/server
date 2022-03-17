package configs

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const ServiceName = "todo"

var options = []option{
	{"config", "string", "", "config file"},

	{"server.http.port", "int", 8091, "server http port"},
	{"server.http.timeout_sec", "int", 86400, "server http connection timeout"},
	{"server.grpc.port", "int", 9091, "server grpc port"},
	{"server.grpc.timeout_sec", "int", 86400, "server grpc connection timeout"},

	{"logger.level", "string", "emerg", "Level of logging. A string that correspond to the following levels: emerg, alert, crit, err, warning, notice, info, debug"},
	{"logger.time_format", "string", "2006-01-02T15:04:05.999999999", "Date format in logs"},

	{"sentry.enabled", "bool", false, "Enables or disables sentry"},
	{"sentry.dsn", "string", "https://829c0fb5737e4fc19997a076d355ece5@sentry.dev.kubedev.ru/4", "Data source name. Sentry addr"},
	{"sentry.environment", "string", "local", "The environment to be sent with events."},

	{"tracer.enabled", "bool", false, "Enables or disables tracing"},
	{"tracer.host", "string", "127.0.0.1", "The tracer host"},
	{"tracer.port", "int", 5775, "The tracer port"},
	{"tracer.name", "string", "todo", "The tracer name"},

	{"metrics.enabled", "bool", false, "Enables or disables metrics"},
	{"metrics.port", "int", 9153, "server http port"},

	{"limiter.enabled", "bool", false, "Enables or disables limiter"},
	{"limiter.limit", "float64", 10000.0, "Limit tokens per second"},

	{"postgres.master.host", "string", "", "postgres master host"},
	{"postgres.master.port", "int", 0000, "postgres master port"},
	{"postgres.master.user", "string", "", "postgres master user"},
	{"postgres.master.password", "string", "", "postgres master password"},
	{"postgres.master.database_name", "string", "", "postgres master database name"},
	{"postgres.master.secure", "string", "disable", "postgres master SSL support"},
	{"postgres.master.max_conns_pool", "int", 150, "max number of connections pool postgres"},

	{"postgres.replica.host", "string", "", "postgres master host"},
	{"postgres.replica.port", "int", 0000, "postgres master port"},
	{"postgres.replica.user", "string", "", "postgres master user"},
	{"postgres.replica.password", "string", "", "postgres master password"},
	{"postgres.replica.database_name", "string", "", "postgres master database name"},
	{"postgres.replica.secure", "string", "disable", "postgres master SSL support"},
	{"postgres.replica.max_conns_pool", "int", 150, "max number of connections pool postgres"},

	{"cache.lifetime", "int", 60, "lifetime of repository cache (in seconds)"},
}

type Config struct {
	Server struct {
		GRPC struct {
			Port       int
			TimeoutSec int `mapstructure:"timeout_sec"`
		}
		HTTP struct {
			Port       int
			TimeoutSec int `mapstructure:"timeout_sec"`
		}
	}
	Logger struct {
		Level      string
		TimeFormat string
	}
	Sentry struct {
		Enabled     bool
		Dsn         string
		Environment string
	}
	Tracer struct {
		Enabled bool
		Host    string
		Port    int
		Name    string
	}
	Metrics struct {
		Enabled bool
		Port    int
	}
	Limiter struct {
		Enabled bool
		Limit   float64
	}
	Postgres struct {
		Master  Database
		Replica Database
	}
	Cache struct {
		Lifetime int
	}
}

type option struct {
	name        string
	typing      string
	value       interface{}
	description string
}

type Database struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string `mapstructure:"database_name"`
	Secure       string
	MaxConnsPool int `mapstructure:"max_conns_pool"`
}

// NewConfig returns and prints struct with config parameters
func NewConfig() *Config {
	return &Config{}
}

// read gets parameters from environment variables, flags or file.
func (c *Config) Read() error {
	viper.SetEnvPrefix(ServiceName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, o := range options {
		switch o.typing {
		case "string":
			pflag.String(o.name, o.value.(string), o.description)
		case "int":
			pflag.Int(o.name, o.value.(int), o.description)
		case "bool":
			pflag.Bool(o.name, o.value.(bool), o.description)
		case "float64":
			pflag.Float64(o.name, o.value.(float64), o.description)
		default:
			viper.SetDefault(o.name, o.value)
		}
	}

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

	if fileName := viper.GetString("config"); fileName != "" {
		viper.SetConfigName(fileName)
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			return errors.Wrap(err, "failed to read from file")
		}
	}

	if err := viper.Unmarshal(c); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
	}
	return nil
}

func (c *Config) Print() error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, string(b))
	return nil
}
