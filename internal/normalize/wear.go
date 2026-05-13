package normalize

import "strings"

var wearMap = map[string]string{
	"FN": "Factory New",
	"MW": "Minimal Wear",
	"FT": "Field-Tested",
	"WW": "Well-Worn",
	"BS": "Battle-Scarred",
}

func WearShortcut(name string) string {
	upper := strings.ToUpper(name)

	for abbr, full := range wearMap {
		wrapped := "(" + abbr + ")"
		if strings.Contains(upper, wrapped) {
			idx := strings.Index(upper, wrapped)
			return name[:idx] + "(" + full + ")" + name[idx+len(wrapped):]
		}
	}

	words := strings.Fields(name)
	if len(words) == 0 {
		return name
	}
	last := strings.ToUpper(words[len(words)-1])
	if full, ok := wearMap[last]; ok {
		sep := ""
		if len(words) > 1 {
			sep = " "
		}
		prefix := strings.Join(words[:len(words)-1], " ")
		return prefix + sep + "(" + full + ")"
	}

	return name
}

func WearShortcuts(names []string) []string {
	out := make([]string, len(names))
	for i, n := range names {
		out[i] = WearShortcut(n)
	}
	return out
}
