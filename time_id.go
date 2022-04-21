package timeid

import (
    "sync"
    "time"
)

var (
    EpochMS  int64 = 1640995200000 // 2022-01-01 00:00:00.000 - max date 2056-11-04 03:53:47.775
    StepBits uint8 = 20 // 1048575 step in a millisecond
    StepMask int64 = -1 ^ (-1 << StepBits)
)

type (
    Node struct {
        mu    sync.Mutex
        time  int64
        epoch time.Time
        step  int64
    }
)

func NewNode() *Node {
    n := &Node{}
    var curTime = GetTime()
    n.epoch = curTime.Add(time.Unix(EpochMS/1000, (EpochMS%1000)*1000000).Sub(curTime))
    return n
}
func (n *Node) Generate() int64 {
    n.mu.Lock()
    defer n.mu.Unlock()
    var now = time.Since(n.epoch).Nanoseconds() / 1000000
    if now == n.time {
        n.step = (n.step + 1) & StepMask

        if n.step == 0 {
            for now <= n.time {
                now = time.Since(n.epoch).Nanoseconds() / 1000000
            }
        }
    } else {
        n.step = 0
    }
    n.time = now

    r := (now)<<StepBits |
        (n.step)

    return r
}

func IdReverse(id int64) (time.Time, int64) {
    ms := id>>StepBits + EpochMS
    t := time.Unix(0, int64(time.Millisecond)*ms)
    step := id & StepMask
    return t, step
}
func IdReverseMs(id int64) int64 {
    t, _ := IdReverse(id)
    return t.UnixNano() / 1000000
}

func TimeMsToId(ms int64, step int64) int64 {
    return (ms-EpochMS)<<StepBits | (step)
}
func TimeParseMs(layout, value string) (int64, error) {
    t, err := time.Parse(layout, value)
    if err != nil {
        return 0, err
    }
    return t.UnixNano() / int64(time.Millisecond), nil
}
func TimeMsToTime(ms int64) time.Time {
    return time.Unix(0, int64(time.Millisecond)*ms).UTC()
}
func timelyMs(layout string, ms int64) int64 {
    value, _ := time.Parse(layout, TimeMsToTime(ms).Format(layout))
    return value.UTC().UnixNano() / 1000000
}
func TimeMsToMinutely(ms int64) int64 {
    return timelyMs("2006-01-02 15:04", ms)
}
func TimeMsToHourly(ms int64) int64 {
    return timelyMs("2006-01-02 15", ms)
}
func TimeMsToDaily(ms int64) int64 {
    return timelyMs("2006-01-02", ms)
}
func TimeMsToMonthly(ms int64) int64 {
    return timelyMs("2006-01", ms)
}
func TimeMsToYearly(ms int64) int64 {
    return timelyMs("2006", ms)
}
func GetTime() time.Time {
    return time.Now().UTC()
}
