package qsoft

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"qsoft_test/internal/qsoft/models"
	"strconv"
	"time"
)

const (
	month      = time.January
	day        = 1
	hour       = 0
	minute     = 0
	second     = 0
	nanosec    = 0
	hoursInDay = 24
)

type service struct {
	logger zerolog.Logger
}

func New(logger zerolog.Logger) *service {
	return &service{
		logger: logger,
	}
}

func (s *service) Days(ctx context.Context, year string) (string, error) {
	y, err := convertToNum(year)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to convert year")
		return "", models.ErrParamNotNum
	}

	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to load location")
		return "", err
	}
	date := time.Date(y, month, day, hour, minute, second, nanosec, loc)

	var days string
	if time.Now().After(date) {
		days = daysGone(date)
	} else {
		days = daysLeft(date)
	}

	return days, nil
}

func daysLeft(date time.Time) string {
	left := time.Until(date)
	days := left / hoursInDay

	return fmt.Sprintf("days left: %v", int(days.Hours()))
}

func daysGone(date time.Time) string {
	gone := time.Since(date)
	days := gone / hoursInDay

	return fmt.Sprintf("days gone: %v", int(days.Hours()))
}

func convertToNum(year string) (int, error) {
	y, err := strconv.Atoi(year)
	if err != nil {
		return 0, err
	}

	return y, nil
}
