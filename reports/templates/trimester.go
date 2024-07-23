package templates

// TrimesterTemplate is the LaTeX code for the trimester pages of the calendar.
// Template placeholders are:
//
//	+T			the full name of the trimester
//	+M1..4CMD	placeholders for the minimonth macros that make up the content of the trimester page
const TrimesterTemplate = `\chapter*{\textbf{+T}}

\vspace{1.5in}

+M1CMD+M2CMD+M3CMD+M4CMD
`
