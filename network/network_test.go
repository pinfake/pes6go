package network;

import "testing"

func TestRead(t *testing.T) {
    t.Run("Should return an error on short buffer", func(t *testing.T) {
        b := [] byte {};
        _, err := Read(b);
        if err == nil {
            t.Error( "No error on short byte array");
        }
    });
    t.Run("Should return mutated header", func(t *testing.T) {
        t.Error("No go!");
    });
}
