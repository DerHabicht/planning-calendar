package cap

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

func TestCapFiscalYear_FyWeek(t *testing.T) {
	fy2024 := NewCapFiscalYear(2024)

	d := date.New(2024, time.March, 14)

	wk, err := fy2024.FyWeek(d)
	assert.NoError(t, err)
	assert.Equal(t, wk, 24)

	d = date.New(2024, time.December, 25)
	_, err = fy2024.FyWeek(d)
	assert.Error(t, err)
}
