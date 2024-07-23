package plancal

type Calendar struct {
	fiscalYear int
}

func NewCalendar(fiscalYear int) Calendar {
	return Calendar{
		fiscalYear: fiscalYear,
	}
}
