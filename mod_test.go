package modtest

import "testing"

func TestMod(t *testing.T) {
	 want := "Hello, world."
    if got := Mod(); got != want {
        t.Errorf("Mod() = %q, want %q", got, want)
    }
}