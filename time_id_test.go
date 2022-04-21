package timeid

import (
    "testing"
    "time"
)

func TestNewNode(t *testing.T) {
    ts, _ := time.Parse("2006-01-02 15:04:05", "2022-01-01 00:00:00.000")
    t.Log(ts, ts.UnixMilli())
    t.Log("max step in one millisecond:",-1^(-1<<(StepBits)))
    t.Log("max date support:", time.Unix(0, (-1^(-1<<(64-StepBits)) + ts.UnixMilli())*int64(time.Millisecond)))
}

func TestNode_Generate(t *testing.T) {
    node := NewNode()
    a := node.Generate()
    b := node.Generate()
    c := node.Generate()

    t.Log(a)
    t.Log(b)
    t.Log(c)
}

func TestNode_Reverse(t *testing.T) {
    node := NewNode()
    a := node.Generate()
    b := node.Generate()
    c := node.Generate()

    at, as := IdReverse(a)
    bt, bs := IdReverse(b)
    ct, cs := IdReverse(c)

    t.Log(a, at, as)
    t.Log(b, bt, bs)
    t.Log(c, ct, cs)
}
