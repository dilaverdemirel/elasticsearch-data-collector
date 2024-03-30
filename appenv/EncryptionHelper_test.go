package appenv

import (
	"fmt"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {

	clear_text := "root"
	fmt.Println("Clear text : ", clear_text)
	got := Encrypt(clear_text)

	want := "root"

	fmt.Println("Encrypted text : ", got)
	got = Decrypt(got)
	fmt.Println("Decrypted text : ", got)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
