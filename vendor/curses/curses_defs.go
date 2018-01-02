package curses

// #define _Bool int
// #include <curses.h>
import "C"

const (
	ERR           = C.ERR
	OK            = C.OK
	A_ALTCHARSET  = C.A_ALTCHARSET
	A_BLINK       = C.A_BLINK
	A_BOLD        = C.A_BOLD
	A_DIM         = C.A_DIM
	A_INVIS       = C.A_INVIS
	A_PROTECT     = C.A_PROTECT
	A_REVERSE     = C.A_REVERSE
	A_STANDOUT    = C.A_STANDOUT
	A_UNDERLINE   = C.A_UNDERLINE
	A_ATTRIBUTES  = C.A_ATTRIBUTES
	A_CHARTEXT    = C.A_CHARTEXT
	A_COLOR       = C.A_COLOR
	WA_ALTCHARSET = C.WA_ALTCHARSET
	WA_BLINK      = C.WA_BLINK
	WA_BOLD       = C.WA_BOLD
	WA_DIM        = C.WA_DIM
	WA_INVIS      = C.WA_INVIS
	WA_LEFT       = C.WA_LEFT
	WA_PROTECT    = C.WA_PROTECT
	WA_REVERSE    = C.WA_REVERSE
	WA_RIGHT      = C.WA_RIGHT
	WA_STANDOUT   = C.WA_STANDOUT
	WA_UNDERLINE  = C.WA_UNDERLINE
	COLOR_BLACK   = C.COLOR_BLACK
	COLOR_BLUE    = C.COLOR_BLUE
	COLOR_GREEN   = C.COLOR_GREEN
	COLOR_CYAN    = C.COLOR_CYAN
	COLOR_RED     = C.COLOR_RED
	COLOR_MAGENTA = C.COLOR_MAGENTA
	COLOR_YELLOW  = C.COLOR_YELLOW
	COLOR_WHITE   = C.COLOR_WHITE
	KEY_BREAK     = C.KEY_BREAK
	KEY_DOWN      = C.KEY_DOWN
	KEY_UP        = C.KEY_UP
	KEY_LEFT      = C.KEY_LEFT
	KEY_RIGHT     = C.KEY_RIGHT
	KEY_HOME      = C.KEY_HOME
	KEY_BACKSPACE = C.KEY_BACKSPACE
	KEY_F0        = C.KEY_F0
	KEY_DL        = C.KEY_DL
	KEY_IL        = C.KEY_IL
	KEY_DC        = C.KEY_DC
	KEY_IC        = C.KEY_IC
	KEY_EIC       = C.KEY_EIC
	KEY_CLEAR     = C.KEY_CLEAR
	KEY_EOS       = C.KEY_EOS
	KEY_EOL       = C.KEY_EOL
	KEY_SF        = C.KEY_SF
	KEY_SR        = C.KEY_SR
	KEY_NPAGE     = C.KEY_NPAGE
	KEY_PPAGE     = C.KEY_PPAGE
	KEY_STAB      = C.KEY_STAB
	KEY_CTAB      = C.KEY_CTAB
	KEY_CATAB     = C.KEY_CATAB
	KEY_ENTER     = C.KEY_ENTER
	KEY_SRESET    = C.KEY_SRESET
	KEY_RESET     = C.KEY_RESET
	KEY_PRINT     = C.KEY_PRINT
	KEY_LL        = C.KEY_LL
	KEY_A1        = C.KEY_A1
	KEY_A3        = C.KEY_A3
	KEY_B2        = C.KEY_B2
	KEY_C1        = C.KEY_C1
	KEY_C3        = C.KEY_C3
	KEY_BTAB      = C.KEY_BTAB
	KEY_BEG       = C.KEY_BEG
	KEY_CANCEL    = C.KEY_CANCEL
	KEY_CLOSE     = C.KEY_CLOSE
	KEY_COMMAND   = C.KEY_COMMAND
	KEY_COPY      = C.KEY_COPY
	KEY_CREATE    = C.KEY_CREATE
	KEY_END       = C.KEY_END
	KEY_EXIT      = C.KEY_EXIT
	KEY_FIND      = C.KEY_FIND
	KEY_HELP      = C.KEY_HELP
	KEY_MARK      = C.KEY_MARK
	KEY_MESSAGE   = C.KEY_MESSAGE
	KEY_MOVE      = C.KEY_MOVE
	KEY_NEXT      = C.KEY_NEXT
	KEY_OPEN      = C.KEY_OPEN
	KEY_OPTIONS   = C.KEY_OPTIONS
	KEY_PREVIOUS  = C.KEY_PREVIOUS
	KEY_REDO      = C.KEY_REDO
	KEY_REFERENCE = C.KEY_REFERENCE
	KEY_REFRESH   = C.KEY_REFRESH
	KEY_REPLACE   = C.KEY_REPLACE
	KEY_RESTART   = C.KEY_RESTART
	KEY_RESUME    = C.KEY_RESUME
	KEY_SAVE      = C.KEY_SAVE
	KEY_SBEG      = C.KEY_SBEG
	KEY_SCANCEL   = C.KEY_SCANCEL
	KEY_SCOMMAND  = C.KEY_SCOMMAND
	KEY_SCOPY     = C.KEY_SCOPY
	KEY_SCREATE   = C.KEY_SCREATE
	KEY_SDC       = C.KEY_SDC
	KEY_SDL       = C.KEY_SDL
	KEY_SELECT    = C.KEY_SELECT
	KEY_SEND      = C.KEY_SEND
	KEY_SEOL      = C.KEY_SEOL
	KEY_SEXIT     = C.KEY_SEXIT
	KEY_SFIND     = C.KEY_SFIND
	KEY_SHELP     = C.KEY_SHELP
	KEY_SHOME     = C.KEY_SHOME
	KEY_SIC       = C.KEY_SIC
	KEY_SLEFT     = C.KEY_SLEFT
	KEY_SMESSAGE  = C.KEY_SMESSAGE
	KEY_SMOVE     = C.KEY_SMOVE
	KEY_SNEXT     = C.KEY_SNEXT
	KEY_SOPTIONS  = C.KEY_SOPTIONS
	KEY_SPREVIOUS = C.KEY_SPREVIOUS
	KEY_SPRINT    = C.KEY_SPRINT
	KEY_SREDO     = C.KEY_SREDO
	KEY_SREPLACE  = C.KEY_SREPLACE
	KEY_SRIGHT    = C.KEY_SRIGHT
	KEY_SRSUME    = C.KEY_SRSUME
	KEY_SSAVE     = C.KEY_SSAVE
	KEY_SSUSPEND  = C.KEY_SSUSPEND
	KEY_SUNDO     = C.KEY_SUNDO
	KEY_SUSPEND   = C.KEY_SUSPEND
	KEY_UNDO      = C.KEY_UNDO
)
