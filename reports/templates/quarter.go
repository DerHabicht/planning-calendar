package templates

// OKRHeaderTemplate is the header row for the OKR tables on the back side of the Quarter page. The placeholders are the
// ISO week numbers for the weeks of that quarter.
const OKRHeaderTemplate = `Objectives & Key Results & +W01 & +W02 & +W03 & +W04 & +W05 & +W06 & +W07 & +W08 & +W09 & +W10 & +W11 & +W12 & +W13`

// QuarterTemplate is the LaTeX code for the quarter pages of the calendar.
// Template placeholders are:
//
//	+CYQ		the calendar year quarter of this page
//	+FYQ		the fiscal year quarter of this page
//	+M1..3CMD	placeholders for the minimonth macros that make up the content of the front of the quarter page
//	+OKR_HDR	placeholder for the processed OKRHeaderTemplate
const QuarterTemplate = `\chapter*{\textbf{+CYQ/+FYQ}}

\begin{center}
+M1CMD+M2CMD+M3CMD
\end{center}

\pagebreak

\section*{OKRs}
\subsection*{Home}
\begin{tabularx}{\textwidth}{|p{1.5in}|p{1.5in}||X|X|X|X|X|X|X|X|X|X|X|X|X|}
    \hline
    +OKR_HDR \\ \hline\hline
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \hline
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \hline
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \hline
\end{tabularx}

\subsection*{Work}
\begin{tabularx}{\textwidth}{|p{1.5in}|p{1.5in}||X|X|X|X|X|X|X|X|X|X|X|X|X|}
    \hline
    +OKR_HDR \\ \hline\hline
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \hline
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \hline
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \hline
\end{tabularx}

\subsection*{CAP}
\begin{tabularx}{\textwidth}{|p{1.5in}|p{1.5in}||X|X|X|X|X|X|X|X|X|X|X|X|X|}
    \hline
    +OKR_HDR \\ \hline\hline
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \hline
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \hline
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \cline{2-15}
    &            &    &    &    &    &    &    &    &    &    &    &    &    &    \\ \hline
\end{tabularx}
`
