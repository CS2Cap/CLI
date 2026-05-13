package normalize

import "testing"

func TestWearShortcut(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"AK-47 | Redline FT", "AK-47 | Redline (Field-Tested)"},
		{"AK-47 | Redline (FT)", "AK-47 | Redline (Field-Tested)"},
		{"AK-47 | Redline (Field-Tested)", "AK-47 | Redline (Field-Tested)"},
		{"★ Bayonet | Doppler FN", "★ Bayonet | Doppler (Factory New)"},
		{"M4A4 | Howl MW", "M4A4 | Howl (Minimal Wear)"},
		{"AWP | Dragon Lore BS", "AWP | Dragon Lore (Battle-Scarred)"},
		{"Desert Eagle | Blaze WW", "Desert Eagle | Blaze (Well-Worn)"},
		{"Sticker | Team Liquid (Holo)", "Sticker | Team Liquid (Holo)"},
		{"", ""},
		{"AK-47 | Redline", "AK-47 | Redline"},
	}
	for _, tt := range tests {
		got := WearShortcut(tt.input)
		if got != tt.want {
			t.Errorf("WearShortcut(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestWearShortcuts(t *testing.T) {
	got := WearShortcuts([]string{"AK-47 | Redline FT", "M4A4 | Howl MW"})
	want := []string{"AK-47 | Redline (Field-Tested)", "M4A4 | Howl (Minimal Wear)"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
