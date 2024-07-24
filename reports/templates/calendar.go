package templates

// DoomsdayTableRowTemplate is the LaTeX code for an individual row of the Doomsday reference table. The generated rows
// should cover a span of two years prior to this calendar's fiscal year, up to two years after. The current year's
// values should be wraped in a `\textbf{}` macro.
//
// The placeholders are:
//
//	+Y					the year represented by this table entry
//	+DD					the day of this years' Doomsday
const DoomsdayTableRowTemplate = `+Y & +DD \\
`

// DoomsdayTableTemplate is the LaTeX code for creating the table of Doomsdays for the adjacent years.
// The placeholder is the collection of DoomsdayTableRowTemplate rows.
const DoomsdayTableTemplate = `\begin{tabular}{rc}
\toprule
\textbf{Year} & \textbf{Doomsday} \\
\midrule
+DD_TABLE_ROWS
\bottomrule
\end{tabular}
`

// SolsticeTableTemplate is the LaTeX code for creating the table of solstices and equinoxes that occur during the span
// covered by this calendar.
// The placeholders are:
//
//	+FYS				the calendar year of the first winter solstice on this calendar (FY-1)
//	+WS1				the ACP 125 DTG of the first winter solstice on this calendar
//	+VE					the ACP 125 DTG of this calendar's vernal equinox
//	+SS					the ACP 125 DTG of this calendar's summer solstice
//	+AE					the ACP 125 DTG of this calendar's autumnal equinox
//	+FYE				the calendar year of the second winter solstice on this calendar (FY)
//	+WS1				the ACP 125 DTG of the second winter solstice on this calendar
const SolsticeTableTemplate = `\begin{tabular}{llr}
\toprule
\Capricorn  & Winter Solstice (+FYS)  & +WS1 \\
\Aries      & Vernal Equinox          & +VE  \\
\Cancer     & Summer Solstice         & +SS  \\
\Libra      & Autumnal Equinox        & +AE  \\
\Capricorn  & Winter Solstice (+FYE)  & +WS2 \\
\bottomrule
\end{tabular}
`

// HolidayAbbvRowTemplate is the LaTeX code for a row in the holiday abbreviations table.
// The placeholders are:
//
//	+ABBV				the holiday's abbreviation
//	+FULL_NAME			the holiday's full name
//	+CY1Act				the actual date of the holiday in the first calendar year (FY-1)
//	+CY1Obs				the observed date of the holiday (if different) in the first calendar year (FY-1)
//	+CY2Act				the actual date of the holiday in the second calendar year (FY)
//	+CY2Obs				the observed date of the holiday (if different) in the second calendar year (FY)
const HolidayAbbvRowTemplate = `+ABBV & +FULL_NAME & +CY1Act & +CY1Obs & +CY2Act & +CY2Obs \\
`

// CalendarTemplate is the LaTeX code for the master calendar file.
// All other templates are injected into this template before compiling.
// The placeholders for other templates in this package are:
//
//	+DOOMSDAYS			placeholder for the processed DoomsdayTableTemplate
//	+SOLSTICES			placeholder for the processed SolsticeTableTemplate
//	+HOLIDAYS			placeholder for all of the processed HolidayAbbvRowTemplate rows, sorted by occurence date
//	+ABBVS				placeholder for all of the processed HolidayAbbvRowTemplate rows, sorted by abbreviation
//	+T1..4				placeholders for the processed TrimesterTemplate
//	+Q1..5				placeholders for the processed QuarterTemplate
//	+M01..15			placeholders for the processed MonthTemplate
//	+MINIMONTH_CMDS		placeholder for all of the processed MinimonthTemplate
//
// The other placeholders in this template are:
//
//	+LCD				the Lunar Calibration Date that the Tikz uses to calculate and then draw the phase of the moon
//	+TITLE_COLOR		the color to use for the title box outline on the calendar's title page
//	+CAL_START			the full name and year of the first full month page in this calendar
//	+CAL_END			the full name and year of the last full month page in this calendar
//	+JP_START			the starting year of this calendar, expressed as the year of Julian Period A
//	+JP_END				the ending year of this calendar, expressed as the year of Julian Period A
//	+PIC				the picture to typeset on the title page of the calendar
//	+CY1				the first calendar year covered in this calendar (i.e. FY-1)
//	+CY2				the second calendar year covered in this calendar (i.e. FY)
const CalendarTemplate = `\documentclass[10pt]{book}

\usepackage{booktabs}
\usepackage{fancyhdr}
\usepackage[landscape,margin=0.5in]{geometry}
\usepackage{graphicx}
\usepackage{marvosym}
\usepackage{multirow}
\usepackage[autolanguage]{numprint}
\usepackage{tabularx}
\usepackage{textcomp}
\usepackage{tikz}
\usepackage{titlepic}
\usepackage{titlesec}
\usepackage{units}
\usepackage{xcolor}

\renewcommand{\multirowsetup}{\raggedleft}
\renewcommand{\headrulewidth}{0pt}

\fancypagestyle{plain}{\fancyhf{}}
\pagestyle{plain}

\titlespacing*{\chapter}{0pt}{0pt}{20pt}
\titleformat{\chapter}[display]{\normalfont\huge}{\chaptertitlename\ \thechapter}{20pt}{\Huge}

\usetikzlibrary{calendar,fpu}

\tikzset{
    moon colour/.style={
        moon fill/.style={
            fill=#1
        }
    },
    sky colour/.style={
        sky draw/.style={
            draw=#1
        },
        sky fill/.style={
            fill=#1
        }
    },
    southern hemisphere/.style={
        rotate=180
    }
}

\makeatletter
\pgfcalendardatetojulian{+LCD}{\c@pgf@counta}
\def\synodicmonth{29.530588853}
\newcommand{\moon}[2][]{%
    \edef\checkfordate{\noexpand\in@{-}{#2}}%
    \checkfordate%
    \ifin@%
    \pgfcalendardatetojulian{#2}{\c@pgf@countb}%
    \pgfkeys{/pgf/fpu=true,/pgf/fpu/output format=fixed}%
    \pgfmathsetmacro\dayssincenewmoon{\the\c@pgf@countb-\the\c@pgf@counta-(7/24+11/(24*60))}%
    \pgfmathsetmacro\lunarage{mod(\dayssincenewmoon,\synodicmonth)}
    \pgfkeys{/pgf/fpu=false}%%
    \else%
    \def\lunarage{#2}%
    \fi%
    \pgfmathsetmacro\leftside{ifthenelse(\lunarage<=\synodicmonth/2,cos(360*(\lunarage/\synodicmonth)),1)}%
    \pgfmathsetmacro\rightside{ifthenelse(\lunarage<=\synodicmonth/2,-1,-cos(360*(\lunarage/\synodicmonth))}%
    \tikz [moon colour=white,sky colour=black,#1]{
        \draw [moon fill, sky draw] (0,0) circle [radius=1ex];
        \draw [sky draw, sky fill] (0,1ex)
        arc (90:-90:\rightside ex and 1ex)
        arc (-90:90:\leftside ex and 1ex)
        -- cycle;
    }%
}
\newcommand{\newmoon}{\moon{0}}
\newcommand{\waxingcrescent}{\moon{\synodicmonth/8}}
\newcommand{\firstquartermoon}{\moon{2*\synodicmonth/8}}
\newcommand{\waxinggibbous}{\moon{3*\synodicmonth/8}}
\newcommand{\fullmoon}{\moon{4*\synodicmonth/8}}
\newcommand{\waninggibbous}{\moon{5*\synodicmonth/8}}
\newcommand{\thirdquartermoon}{\moon{6*\synodicmonth/8}}
\newcommand{\waningcrescent}{\moon{7*\synodicmonth/8}}
\makeatother

+MINIMONTH_CMDS

\begin{document}

\begin{titlepage}
    \setlength{\fboxrule}{10pt}
    \begin{center}
        \Huge%

        \vspace*{0.25in}

        \fcolorbox{+TITLE_COLOR}{white}{%
            \begin{minipage}{0.5\textwidth}
                \begin{center}
                    AG7IF Planning Calendar

                    +CAL_START to +CAL_END

                    Julian Period A +JP_START to +JP_END
                \end{center}
            \end{minipage}
        }

        \vfill

        \includegraphics[height=4.5in]{+PIC}

        \vspace*{0.25in}

    \end{center}
\end{titlepage}

\section*{+CY1/+CY2 Holidays}
\begin{tabularx}{\textwidth}{lXrrrr}
    \toprule
    Abbreviation & Holiday & +CY1 Occurrence & +CY1 Observed & +CY2 Occurrence & +CY2 Observed \\
    \midrule
    +HOLIDAYS
    \bottomrule
\end{tabularx}

\section*{Reference Tables}
\begin{center}
    \vfill
    \hfill
    \begin{tabular}{rc}
        \toprule
        \textbf{Century} & \textbf{Doomsday} \\
        \midrule
        1700 & U \\
        1800 & F \\
        1900 & W \\
        2000 & T \\
        \bottomrule
    \end{tabular}
    \hfill
    +DOOMSDAYS
    \hfill
    +SOLSTICES
    \vfill
    \hfill
    \begin{tabular}{rrr}
        \toprule
        \textbf{Unit}      & \textbf{Days}        & \textbf{Equivalent} \\
        \midrule
        Julian Period      &   \numprint{2914695} & $\unit{\numprint{7980}}{yr}$ \\
        megaday (Mdy)      &   \numprint{1000000} & $\unit{\numprint{2740}}{yr}$ \\
        100 kdy            &    \numprint{100000} & $\unit{274}{yr}$ \\
        10 kdy             &     \numprint{10000} & $\unit{27.4}{yr}$ \\
        kiloday (kdy)      &      \numprint{1000} & $\unit{2.74}{yr}$ \\
        hectaday (hdy)     &                  100 & $\unit{14.29}{wk}$ \\
        decaday (dady)     &                   10 & $\unit{1.429}{wk}$ \\
        day (dy)           &                    1 & $\unit{\numprint{86400}}{s}$ \\
        deciday (ddy)      &                  0.1 & $\unit{2.4}{h}$ \\
        centiday (cdy)     &                 0.01 & $\unit{14.4}{min}$ \\
        milliday (mdy)     &                0.001 & $\unit{1.44}{min}$ \\
        microday (\textmu{}dy) &               0.0001 & $\unit{8.64}{s}$ \\
        \bottomrule
    \end{tabular}
    \hfill
    \begin{tabular}{rr|rr|rr|rr}
        \toprule
        \textbf{cd} & \textbf{time} &\textbf{cd} & \textbf{time} &\textbf{cd} & \textbf{time} &\textbf{cd} & \textbf{time} \\
        \midrule
        0.00 & 0000 & 0.25 & 0600 & 0.50 & 1200 & 0.75 & 1800 \\
        0.01 & 0014 & 0.26 & 0614 & 0.51 & 1214 & 0.76 & 1814 \\
        0.02 & 0028 & 0.27 & 0628 & 0.52 & 1228 & 0.77 & 1828 \\
        0.03 & 0043 & 0.28 & 0643 & 0.53 & 1243 & 0.78 & 1843 \\
        0.04 & 0057 & 0.29 & 0657 & 0.54 & 1257 & 0.79 & 1857 \\
        0.05 & 0112 & 0.30 & 0712 & 0.55 & 1312 & 0.80 & 1912 \\
        0.06 & 0126 & 0.31 & 0726 & 0.56 & 1326 & 0.81 & 1926 \\
        0.07 & 0140 & 0.32 & 0740 & 0.57 & 1340 & 0.82 & 1940 \\
        0.08 & 0155 & 0.33 & 0755 & 0.58 & 1355 & 0.83 & 1955 \\
        0.09 & 0209 & 0.34 & 0809 & 0.59 & 1409 & 0.84 & 2009 \\
        0.10 & 0224 & 0.35 & 0824 & 0.60 & 1424 & 0.85 & 2024 \\
        0.11 & 0238 & 0.36 & 0838 & 0.61 & 1438 & 0.86 & 2038 \\
        0.12 & 0252 & 0.37 & 0852 & 0.62 & 1452 & 0.87 & 2052 \\
        0.13 & 0307 & 0.38 & 0907 & 0.63 & 1507 & 0.88 & 2107 \\
        0.14 & 0321 & 0.39 & 0921 & 0.64 & 1521 & 0.89 & 2121 \\
        0.15 & 0336 & 0.40 & 0936 & 0.65 & 1536 & 0.90 & 2136 \\
        0.16 & 0350 & 0.41 & 0950 & 0.66 & 1550 & 0.91 & 2150 \\
        0.17 & 0404 & 0.42 & 1004 & 0.67 & 1604 & 0.92 & 2204 \\
        0.18 & 0419 & 0.43 & 1019 & 0.68 & 1619 & 0.93 & 2219 \\
        0.19 & 0433 & 0.44 & 1033 & 0.69 & 1633 & 0.94 & 2233 \\
        0.20 & 0448 & 0.45 & 1048 & 0.70 & 1648 & 0.95 & 2248 \\
        0.21 & 0502 & 0.46 & 1102 & 0.71 & 1702 & 0.96 & 2302 \\
        0.22 & 0516 & 0.47 & 1116 & 0.72 & 1716 & 0.97 & 2316 \\
        0.23 & 0531 & 0.48 & 1131 & 0.73 & 1731 & 0.98 & 2331 \\
        0.24 & 0545 & 0.49 & 1145 & 0.74 & 1745 & 0.99 & 2345 \\
        \bottomrule
    \end{tabular}
    \hfill
    \begin{tabular}{cr|cr}
        \toprule
        \textbf{ACPTZ} & \textbf{UTC} & \textbf{ACPTZ} & \textbf{UTC} \\
        \midrule
        Z & $ +0$ &   &       \\
        A & $ +1$ & N &  $-1$ \\
        B & $ +2$ & O &  $-2$ \\
        C & $ +3$ & P &  $-3$ \\
        D & $ +4$ & Q &  $-4$ \\
        E & $ +5$ & R &  $-5$ \\
        F & $ +6$ & S &  $-6$ \\
        G & $ +7$ & T &  $-7$ \\
        H & $ +8$ & U &  $-8$ \\
        I & $ +9$ & V &  $-9$ \\
        J & $+10$ & W & $-10$ \\
        K & $+11$ & X & $-11$ \\
        M & $+12$ & Y & $-12$ \\
        \bottomrule
    \end{tabular}
    \hspace*{\fill}
    \vspace*{\fill}
\end{center}

\pagebreak

\noindent
\begin{tabularx}{\textwidth}{XXXX}
    \minifirstoctober      &
    \minifirstnovember     &
    \minifirstdecember     &
    \minifirstjanuary     \\

    \minifebruary          &
    \minimarch             &
    \miniapril             &
    \minimay              \\

    \minijune              &
    \minijuly              &
    \miniaugust            &
    \minisecondseptember  \\

    \minisecondoctober     &
    \minisecondnovember    &
    \miniseconddecember    &
    \minisecondjanuary    \\
\end{tabularx}

\pagebreak

\noindent
\begin{tabularx}{\textwidth}{XXXX}
    \minifirstextfebruary      &
    \minifirstextmarch         &
    \minifirstextapril         &
    \minifirstextmay          \\

    \miniextjune               &
    \miniextjuly               &
    \miniextaugust             &
    \miniextseptember         \\

    \miniextoctober            &
    \miniextnovember           &
    \miniextdecember           &
    \miniextjanuary           \\

    \minisecondextfebruary     &
    \minisecondextmarch        &
    \minisecondextapril        &
    \minisecondextmay         \\
\end{tabularx}

\pagebreak

\section*{Holiday Abbreviations}
\begin{tabularx}{\textwidth}{lXrrrr}
    \toprule
    Abbreviation & Holiday & +CY1 Occurrence & +CY1 Observed & +CY2 Occurrence & +CY2 Observed \\
    \midrule
    +ABBVS
    \bottomrule
\end{tabularx}

\newpage

+T1

+Q1

+M01

+M02

+M03

+Q2

+M04

+T2

+M05

+M06

+Q3

+M07

+M08

+T3

+M09

+Q4

+M10

+M11

+M12

+T4

+Q5

+M13

+M14

+M15

\end{document}
`
