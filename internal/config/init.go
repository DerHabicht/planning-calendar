package config

import (
	"os"
	"path/filepath"

	"github.com/derhabicht/planning-calendar/internal/logging"
)

func InitConfig() {
	SetDefault("home_location.lat", 0.0)
	SetDefault("home_location.long", 0.0)
	SetDefault("home_location.tz", "UTC")
	SetDefault("cover_logo", "")
	SetDefault("lunar_calibration_date", "2010-01-15") // Date of a New Moon for calibrating the LaTeX code that paints the moon in the calendar.

	cwd, _ := os.Getwd()
	logging.Trace().Str("cwd", cwd).Msg("")

	cfgDir, err := ConfigDir()
	if err != nil {
		logging.Error().Err(err).Msg("could not find user config dir")
	} else {
		logging.Debug().Str("usrCfgDir", cfgDir).Msg("user config dir found")
	}

	SetConfigType("yaml")
	SetConfigName("config")
	AddConfigPath(filepath.Join(cfgDir, "cfg"))
	err = ReadInConfig()
	if err != nil {
		path := filepath.Join(cfgDir, "cfg", "config.yaml")
		logging.Warn().Err(err).Str("path", path).Msg("config file not found, creating a default config")
		err = WriteConfigAs(path)
		if err != nil {
			logging.Error().Err(err).Msg("failed to create default config file")
		}
	}
}
