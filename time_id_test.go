package timeid

import (
    "testing"
    "time"
)

func TestNewNode(t *testing.T) {
    ts, _ := time.Parse("2006-01-02 15:04:05", "2010-05-19 08:24:48.8964")
    t.Log(ts, ts.UnixMilli())
    t.Log("max step in one millisecond:",-1^(-1<<(22)))
    t.Log("max date support:", time.Unix(0, (-1^(-1<<(64-22)) + ts.UnixMilli())*int64(time.Millisecond)))
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

    at, as := node.Reverse(a)
    bt, bs := node.Reverse(b)
    ct, cs := node.Reverse(c)

    t.Log(a, at, as)
    t.Log(b, bt, bs)
    t.Log(c, ct, cs)
}
