package templates

// MonthWeekdayHeaderTemplate is the LaTeX code that generates the header row of the month. The placeholders are the
// days of the week, allowing the processor to highlight the doomsday of the year.
const MonthWeekdayHeaderTemplate = `KFM & & & +MON & +TUE & +WED & +THU & +FRI & +SAT & +SUN \\`

// MonthDayTemplate is the LaTeX code for each individual day in the month.
// The placeholders are:
//
//	+DY					the day of the month represented by this cell; the abbreviated month should be included if
//						crossing month boundaries
//	+HD					the abbreviation of any holiday that falls on this day
//	+MP					the full date---in ISO 8601 format---for the TikZ moon phase calculation
//	+YD					the ordinal day number
//	+SR					the sunrise time for the configured location
//	+MJD				the Modified Julian Day number
//	+SS					the sunset time for the configured location
//
// const MonthDayTemplate = `+DY\\+HD\moon{+FD}\\\vspace{1em}\hspace{1em}+YD\hfill{}+SR\\+MJD\hfill{}+SS`
const MonthDayTemplate = `+DY\\+HD+MP\\\vspace{1em}\hspace{1em}+YD\hfill{}+SR\\+MJD\hfill{}+SS`

// MonthTemplate is the LaTeX code for generating a month page.
// Template placeholders are:
//
//	+M					the full name and calendar year of the current month
//	+PREV_CMD			the command which invokes the processed MinimonthTemplate for the previous month
//	+NEXT_CMD			the command which invokes the processed MinimonthTemplate for the next month
//	+WEEKDAYS			the macro for the processed MonthWeekdayHeaderTemplate
//	+FT1..6				the fiscal year trimester that this week (row) belongs to
//	+FQ1..6				the fiscal year quarter that this week (row) belongs to
//	+FW1..6				this week's (row's) fiscal year week number
//	+AQ1..6				the calendar year quarter that this week (row) belongs to
//	+AS1..6				the sprint that this week (row) belongs to
//	+IW1..6				this week's (row's) playing card and ISO week number
//	+D01..42			the day data for the cells in the calendar
const MonthTemplate = `\chapter*{\textbf{+M}\hfill{\footnotesize+PREV_CMD+NEXT_CMD}}
\begin{center}
{\ttfamily%
\begin{tabular}{|c|lr|c|c|c|c|c|c|c|}
\hline
    +WEEKDAYS
\hline
    \multirow{5}{0.5in}{} & FY    & +FT1/+FQ1 & \multirow[t]{5}{1in}{+D01} & \multirow[t]{5}{1in}{+D02} & \multirow[t]{5}{1in}{+D03} & \multirow[t]{5}{1in}{+D04} & \multirow[t]{5}{1in}{+D05} & \multirow[t]{5}{1in}{+D06} & \multirow[t]{5}{1in}{+D07} \\
                        &       &      +FW1 &                            &                            &                            &                            &                            &                            &                            \\
                        & AG7IF &      +AQ1 &                            &                            &                            &                            &                            &                            &                            \\
                        &       &      +AS1 &                            &                            &                            &                            &                            &                            &                            \\
                        & ISO   &      +IW1 &                            &                            &                            &                            &                            &                            &                            \\
\hline
    \multirow{5}{0.5in}{} & FY    & +FT2/+FQ2 & \multirow[t]{5}{1in}{+D08} & \multirow[t]{5}{1in}{+D09} & \multirow[t]{5}{1in}{+D10} & \multirow[t]{5}{1in}{+D11} & \multirow[t]{5}{1in}{+D12} & \multirow[t]{5}{1in}{+D13} & \multirow[t]{5}{1in}{+D14} \\
                        &       &      +FW2 &                            &                            &                            &                            &                            &                            &                            \\
                        & AG7IF &      +AQ2 &                            &                            &                            &                            &                            &                            &                            \\
                        &       &      +AS2 &                            &                            &                            &                            &                            &                            &                            \\
                        & ISO   &      +IW2 &                            &                            &                            &                            &                            &                            &                            \\
\hline
    \multirow{5}{0.5in}{} & FY    & +FT3/+FQ3 & \multirow[t]{5}{1in}{+D15} & \multirow[t]{5}{1in}{+D16} & \multirow[t]{5}{1in}{+D17} & \multirow[t]{5}{1in}{+D18} & \multirow[t]{5}{1in}{+D19} & \multirow[t]{5}{1in}{+D20} & \multirow[t]{5}{1in}{+D21} \\
                        &       &      +FW3 &                            &                            &                            &                            &                            &                            &                            \\
                        & AG7IF &      +AQ3 &                            &                            &                            &                            &                            &                            &                            \\
                        &       &      +AS3 &                            &                            &                            &                            &                            &                            &                            \\
                        & ISO   &      +IW3 &                            &                            &                            &                            &                            &                            &                            \\
\hline
    \multirow{5}{0.5in}{} & FY    & +FT4/+FQ4 & \multirow[t]{5}{1in}{+D22} & \multirow[t]{5}{1in}{+D23} & \multirow[t]{5}{1in}{+D24} & \multirow[t]{5}{1in}{+D25} & \multirow[t]{5}{1in}{+D26} & \multirow[t]{5}{1in}{+D27} & \multirow[t]{5}{1in}{+D28} \\
                        &       &      +FW4 &                            &                            &                            &                            &                            &                            &                            \\
                        & AG7IF &      +AQ4 &                            &                            &                            &                            &                            &                            &                            \\
                        &       &      +AS4 &                            &                            &                            &                            &                            &                            &                            \\
                        & ISO   &      +IW4 &                            &                            &                            &                            &                            &                            &                            \\
\hline
    \multirow{5}{0.5in}{} & FY    & +FT5/+FQ5 & \multirow[t]{5}{1in}{+D29} & \multirow[t]{5}{1in}{+D30} & \multirow[t]{5}{1in}{+D31} & \multirow[t]{5}{1in}{+D32} & \multirow[t]{5}{1in}{+D33} & \multirow[t]{5}{1in}{+D34} & \multirow[t]{5}{1in}{+D35} \\
                        &       &      +FW5 &                            &                            &                            &                            &                            &                            &                            \\
                        & AG7IF &      +AQ5 &                            &                            &                            &                            &                            &                            &                            \\
                        &       &      +AS5 &                            &                            &                            &                            &                            &                            &                            \\
                        & ISO   &      +IW5 &                            &                            &                            &                            &                            &                            &                            \\
\hline
    \multirow{5}{0.5in}{} & FY    & +FT6/+FQ6 & \multirow[t]{5}{1in}{+D36} & \multirow[t]{5}{1in}{+D37} & \multirow[t]{5}{1in}{+D38} & \multirow[t]{5}{1in}{+D39} & \multirow[t]{5}{1in}{+D40} & \multirow[t]{5}{1in}{+D41} & \multirow[t]{5}{1in}{+D42} \\
                        &       &      +FW6 &                            &                            &                            &                            &                            &                            &                            \\
                        & AG7IF &      +AQ6 &                            &                            &                            &                            &                            &                            &                            \\
                        &       &      +AS6 &                            &                            &                            &                            &                            &                            &                            \\
                        & ISO   &      +IW6 &                            &                            &                            &                            &                            &                            &                            \\
\hline
\end{tabular}
}
\end{center}
`
