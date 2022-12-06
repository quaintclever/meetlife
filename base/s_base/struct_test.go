package sbase

import (
	"fmt"
	"testing"
)

type Test struct {
	tg  TestGroup
	msg string
}

type TestGroup struct {
	ts *Test
}

func TestStructLoop(t *testing.T) {
	ts := &Test{
		msg: "testMsg",
	}
	ts.tg.ts = ts

	fmt.Printf("ts:%p, ts:%+v\n", ts, ts)
	fmt.Println(ts == ts.tg.ts)
	// update ts msg
	ts.msg = "testMsg2"
	fmt.Println(ts.msg)
	fmt.Println(ts.tg.ts.msg)
}
