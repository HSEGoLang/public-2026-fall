//go:build !solution

package speller

// Spell returns the American English spelling of n.
func Spell(n int64) string {
	_ = scales

	return spellUnderThousand(n)
}

func spellUnderThousand(n int64) string {
	_, _, _ = n, smallNumbers, tens

	return ""
}
