package templates

// MinimonthWeekHeaderTemplate is the LaTeX code to print a minimonth header. The placeholders are to allow
// highlighting of the corresponding doomsday.
const MinimonthWeekHeaderTemplate = `W & +M & +T & +W & +H & +F & +S & +U \\`

// MinimonthWeekTemplate is the LaTeX code for each individual week row of a minimonth. The placeholders should be
// replaced with the day of the month represented by the cell.
const MinimonthWeekTemplate = `+W & +D1 & +D2 & +D3 & +D4 & +D5 & +D6 & +D7 \\`

// MinimonthTemplate is the LaTeX macro for generating a miniature month within the calendar. It is produced as a macro
// because each month could be used in several places within the calendar's contents.
//
// The placeholders are:
//
//	+COMMAND		the LaTeX macro command being defined
//	+MONTH			the full name of the month and year represented by this minimonth
//	+WEEK_HEADER	placeholder for the processed MinimonthWeekHeaderTemplate
//	+W1..6			placeholder for the processed MinimonthWeekTemplate
const MinimonthTemplate = `
\newcommand{+COMMAND}{\fbox{\begin{minipage}{0.24\textwidth}
          \centering
          {\Large\textbf{+MONTH}}\vspace{\baselineskip}
          \begin{tabularx}{\textwidth}{r|rrrrrrr}
              \toprule
 			  +WEEK_HEADER
              \midrule
              +W1
              +W2
              +W3
              +W4
              +W5
              +W6
          \end{tabularx}
\end{minipage}}}
`
