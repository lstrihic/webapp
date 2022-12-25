package main

import (
	"github.com/lstrihic/webapp/pkg/config"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"os"
	"strings"
	"time"
)

var (
	// service version
	version = "develop"

	// configuration file location
	configFile string

	// service configuration
	cfg config.Config

	// service logger
	logger zerolog.Logger
)

var rootCMD = cobra.Command{
	Use: "service",
	RunE: func(cmd *cobra.Command, args []string) error {
		fxApp := fx.New(
			fx.Supply(&logger),
			fx.Supply(&cfg),
			fx.Invoke(func(log *zerolog.Logger) {
				// TODO: ...
				log.Info().Msg("TODO: implement")
			}),
		)

		if err := fxApp.Start(cmd.Context()); err != nil {
			return err
		}
		<-fxApp.Done()
		return fxApp.Stop(cmd.Context())
	},
}

func main() {
	// start service
	if err := rootCMD.Execute(); err != nil {
		logger.Error().Err(err).Msg("Failed to start service")
	}
}

func init() {
	// init logger
	logger = log.
		Level(zerolog.DebugLevel).
		Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}).
		With().
		Caller().
		Str("version", version).
		Str("service_name", "service").
		Logger()

	// initialize cobra
	cobra.OnInitialize(initConfig)
	rootCMD.PersistentFlags().StringVar(&configFile, "config", "data/config.yaml", "config file (default is data/config.yaml)")
}

func initConfig() {
	// configure viper
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	// bind env
	_ = viper.BindEnv("port", "PORT")

	// try to read config
	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(&cfg, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "config"
		decoderConfig.IgnoreUntaggedFields = true
	})
}
