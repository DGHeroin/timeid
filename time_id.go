package timeid

import (
    "sync"
    "time"
)

var (
    customEpoch int64 = 1577836800000 // 1274257488896 in 2010-05-19 08:24:48
    stepBits    uint8 = 22
)

type (
    Node struct {
        mu        sync.Mutex
        time      int64
        epoch     time.Time
        step      int64
        timeShift uint8
        stepMask  int64
    }
)

func NewNode() *Node {
    n := &Node{}
    n.timeShift = stepBits
    n.stepMask = -1 ^ (-1 << stepBits)

    var curTime = time.Now()
    // add time.Duration to curTime to make sure we use the monotonic clock if available
    n.epoch = curTime.Add(time.Unix(customEpoch/1000, (customEpoch%1000)*1000000).Sub(curTime))
    return n
}
func (n *Node) Generate() int64 {
    n.mu.Lock()
    defer n.mu.Unlock()
    now := time.Since(n.epoch).Nanoseconds() / 1000000
    if now == n.time {
        n.step = (n.step + 1) & n.stepMask

        if n.step == 0 {
            for now <= n.time {
                now = time.Since(n.epoch).Nanoseconds() / 1000000
            }
        }
    } else {
        n.step = 0
    }
    n.time = now

    r := (now)<<n.timeShift |
        (n.step)

    return r
}

func (n *Node) Reverse(id int64) (time.Time, int64) {
    ms := id>>stepBits + customEpoch
    t := time.Unix(0, int64(time.Millisecond)*ms)

    step := id & n.stepMask
    return t, step
}
