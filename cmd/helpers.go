package cmd

import "fmt"

func stringInt(i int) string {
	return fmt.Sprintf("%d", i)
}

func boolYesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
