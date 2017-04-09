package network;

import "testing"

func TestDecodeShouldReturnErrorOnShortByteSlice(t *testing.T) {
    b := [] byte {};
    _, err := Decode(b);
    if err == nil {
        t.Error( "No error on short byte array");
    }
}
