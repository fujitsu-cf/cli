package formatters

import (
	. "github.com/fujitsu-cf/cli/cf/i18n"
)

func Allowed(allowed bool) string {
	if allowed {
		return T("allowed")
	}
	return T("disallowed")
}
