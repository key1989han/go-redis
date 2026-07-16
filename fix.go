package redis

import "sync"

type SafePubSub struct {
    mu      sync.Mutex
    channel chan *Message
    closed  bool
}

func (ps *SafePubSub) Receive() (*Message, error) {
    ps.mu.Lock()
    if ps.closed {
        ps.mu.Unlock()
        return nil, ErrClosed
    }
    ps.mu.Unlock()
    msg, ok := <-ps.channel
    if !ok {
        return nil, ErrClosed
    }
    return msg, nil
}

func (ps *SafePubSub) Close() error {
    ps.mu.Lock()
    defer ps.mu.Unlock()
    ps.closed = true
    close(ps.channel)
    return nil
}
