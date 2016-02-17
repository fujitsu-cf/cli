package terminal

import (
	"time"

	. "github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/cf/trace"
)

type DebugPrinter struct{}

func (DebugPrinter) Print(title, dump string) {
	trace.Logger.Printf("\n%s [%s]\n%s\n", HeaderColor(T(title)), time.Now().Format(time.RFC3339), trace.Sanitize(dump))
}
