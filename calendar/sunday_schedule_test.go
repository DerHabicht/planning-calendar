package calendar

import (
	"testing"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/stretchr/testify/assert"
)

func TestSundaySchoolSchedule(t *testing.T) {
	cal := &cal.Calendar{}
	cal.AddHoliday(SundaySchool...)

	testDate := time.Date(2023, 11, 5, 0, 0, 0, 0, time.Local)

	act, _, _ := cal.IsHoliday(testDate)

	assert.True(t, act)
}
