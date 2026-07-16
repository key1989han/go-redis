package redis

import "time"

// SafeReceiveMessage wraps ReceiveMessage with timeout protection
func SafeReceiveMessage(ps *PubSub, timeout time.Duration) (*Message, error) {
    ch := ps.Channel()
    timer := time.NewTimer(timeout)
    defer timer.Stop()

    select {
    case msg, ok := <-ch:
        if !ok {
            return nil, ErrClosed
        }
        return msg, nil
    case <-timer.C:
        return nil, ErrTimeout
    }
}
