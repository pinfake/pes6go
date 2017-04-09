package network;

import "testing"

func TestDecode(t *testing.T) {
    t.Run("Should return an error on short buffer", func(t *testing.T) {
        b := [] byte {};
        _, err := Decode(b);
        if err == nil {
            t.Error( "No error on short byte array");
        }
    });
}
