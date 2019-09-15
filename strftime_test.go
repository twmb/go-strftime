package strftime

import (
	"testing"

	"time"
)

// The tests below are written according to package time's date;
// thus we use the same timezone and unix second as package time.
var tz, _ = time.LoadLocation("MST")
var refTime = time.Unix(1136239445, 0).In(tz)

func TestAppendFormat(t *testing.T) {
	for i, test := range []struct {
		format string
		exp    string
	}{
		{"", ``},

		// non escapes
		{"%", `%`},
		{"|%|", `|%|`},
		{"|% ", `|% `},
		{"%| x", `%| x`},
		{"|x|%| x", `|x|%| x`},

		{"%a", `Mon`},
		{"%A", `Monday`},
		{"%b", `Jan`},
		{"%h", `Jan`},
		{"%B", `January`},
		{"%c", `Mon Jan  2 15:04:05 2006`}, // locale dependent but we hardcode
		{"%C", `20`},
		{"%d", `02`},
		{"%D", `01/02/06`},
		{"%e", ` 2`},
		{"%F", `2006-01-02`},
		{"%G", `2006`},
		{"%g", `06`},
		{"%H", `15`},
		{"%I", `03`},
		{"%j", `002`},
		{"%k", `15`},
		{"%l", ` 3`},
		{"%m", `01`},
		{"%M", `04`},
		{"%n", "\n"},
		{"%p", `PM`},
		{"%P", `pm`}, // locale dependent but we hardcode
		{"%r", `03:04:05 PM`},
		{"%R", `15:04`},
		{"%s", `1136239445`},
		{"%S", `05`},
		{"%t", "\t"},
		{"%T", `15:04:05`},
		{"%u", `1`},
		{"%U", `01`},
		{"%V", `01`},
		{"%w", `1`},
		{"%W", `01`},
		{"%x", `02/01/06`},
		{"%X", `15:04:05`},
		{"%y", `06`},
		{"%Y", `2006`},
		{"%z", `-0700`},
		{"%Z", `MST`},
		{"%%", `%`},
		{"%+", `Mon Jan  2 15:04:05 MST 2006`}, // date(1), but we hardcode

		// unsupported modifiers; stripped
		{"%E", ``},
		{"%EF", ``},
		{"%O", ``},
		{"%OF", ``},
		{"%O|", `|`}, // only %O since non-ascii follows

		// all in one
		{
			benchFmt,
			`Mon|Monday|Jan|January|Mon Jan  2 15:04:05 2006|20|02|01/02/06| 2|||2006-01-02|2006|06|Jan|15|03|002|15| 3|01|04|||PM|pm|03:04:05 PM|15:04|1136239445|05|	|15:04:05|1|01|01|1|01|02/01/06|15:04:05|06|2006|-0700|MST|%|%*|%`,
		},
	} {
		got := AppendFormat(nil, test.format, refTime)
		if string(got) != test.exp {
			t.Errorf("%d [%s]: got != exp;\n%s\n%s\n", i, test.format, got, test.exp)
		}
	}
}

const benchFmt = "%a|%A|%b|%B|%c|%C|%d|%D|%e|%E|%EF|%F|%G|%g|%h|%H|%I|%j|%k|%l|%m|%M|%O|%OF|%p|%P|%r|%R|%s|%S|%t|%T|%u|%U|%V|%w|%W|%x|%X|%y|%Y|%z|%Z|%%|%*|%"

var benchTime = time.Date(2014, time.July, 2, 11, 57, 42, 234098432, time.UTC)

func BenchmarkAppendFormat(b *testing.B) {
	var dst []byte
	for i := 0; i < b.N; i++ {
		dst = AppendFormat(dst[:0], benchFmt, benchTime)
	}
}
