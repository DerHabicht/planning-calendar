package calendar

import (
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/ag7if/go-latex"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-calendar/calendar"
	"github.com/derhabicht/planning-calendar/internal/config"
	"github.com/derhabicht/planning-calendar/internal/logging"
	"github.com/derhabicht/planning-calendar/reports"
)

func configureLaTeXCompiler(logger logging.Logger) (*latex.Compiler, error) {
	cfgDir, err := config.ConfigDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to find config directory")
	}

	cacheDir, err := config.CacheDir()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to find cache directory")
	}

	assetDir := filepath.Join(cfgDir, "assets")

	compiler := latex.NewCompiler(assetDir, cacheDir, logger.DefaultLogger())

	return &compiler, nil
}

func generateLaTeX(cal calendar.Calendar, compiler *latex.Compiler, outputFile files.File) error {
	planningCal, err := reports.NewCalendar(cal)
	if err != nil {
		return errors.WithStack(err)
	}

	assets := []string{config.GetString("cover_logo")}

	err = compiler.GenerateLaTeX(planningCal, outputFile, assets)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func BuildCalendar(year int, outputFile files.File, logger logging.Logger) error {

	cal := calendar.NewCalendar(year)

	compiler, err := configureLaTeXCompiler(logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = generateLaTeX(cal, compiler, outputFile)
	if err != nil {
		return errors.WithStack(err)
	}

	err = compiler.CompileLaTeX(outputFile)
	if err != nil {
		return err
	}

	return nil
}
