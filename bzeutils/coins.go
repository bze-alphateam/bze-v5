package bzeutils

import "strings"

func IsIBCDenom(denom string) bool {
	return strings.HasPrefix(denom, "ibc/")
}
