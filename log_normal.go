// TO_VIEW

//go:build !debug

package nostr

func debugLogf(str string, args ...any) {
	return
}
