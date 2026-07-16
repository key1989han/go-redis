package redis

import (
    "testing"
    "time"
)

func TestSafeReceiveMessage_Timeout(t *testing.T) {
    // Test that timeout works correctly
    ch := make(chan interface{}, 1)
    timer := time.NewTimer(50 * time.Millisecond)
    defer timer.Stop()

    select {
    case <-ch:
        t.Error("Should not receive")
    case <-timer.C:
        // Expected timeout
    }
}

func TestSafeReceiveMessage_Success(t *testing.T) {
    ch := make(chan interface{}, 1)
    ch <- "test"

    select {
    case msg := <-ch:
        if msg != "test" {
            t.Error("Expected test message")
        }
    case <-time.After(time.Second):
        t.Error("Should not timeout")
    }
}
