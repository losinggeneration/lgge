package timer

import (
	"testing"
	"time"
)

func TestReset(te *testing.T) {
	t := NewTimer()
	t.Reset()

	if !t.start.Equal(t.last) {
		te.Error("Expected start time and last update time to be equal after reset")
	}

	last := *t

	time.Sleep(5 * time.Microsecond)

	t.Reset()
	if last.start.Equal(t.start) {
		te.Error("Did not expect start times to be equal after Reset")
	}

	if last.last.Equal(t.last) {
		te.Error("Did not expect last update times to be equal after Reset")
	}
}

func TestUpdate(te *testing.T) {
	t := NewTimer()
	t.Reset()
	start := t.start
	last := t.last

	time.Sleep(5 * time.Millisecond)
	t.Update()

	if !start.Equal(t.start) {
		te.Error("Did not expect start time to change after call to Update")
	}

	if last.Equal(t.last) {
		te.Error("Expected last update time to change after call to Update")
	}

	if start.Equal(t.last) {
		te.Error("Did not expect start & last update time to be equal after call to Update")
	}

	if t.frames != 1 {
		te.Error("Did not expect frames to be 0 after call to Update")
	}

	last = t.last

	time.Sleep(5 * time.Millisecond)
	t.Update()

	if !start.Equal(t.start) {
		te.Error("Did not expect start time to change after call to Update")
	}

	if last.Equal(t.last) {
		te.Error("Expected last update time to change after call to Update")
	}

	if start.Equal(t.last) {
		te.Error("Did not expect start & last update time to be equal after call to Update")
	}

	if t.frames != 2 {
		te.Error("Did not expect frames to be 0 after call to Update")
	}
}

func TestTime(t *testing.T) {
	timer := NewTimer()
	timer.Reset()

	time.Sleep(5 * time.Millisecond)
	if now := timer.Time(); now <= .001 {
		t.Error("Not eneough time has passed when Time was called")
	}

	time.Sleep(5 * time.Millisecond)
	if now := timer.Time(); now <= .001 {
		t.Error("Not eneough time has passed when Time was called")
	}
}

func TestDelta(t *testing.T) {
	timer := NewTimer()
	timer.Reset()

	time.Sleep(5 * time.Millisecond)
	if now := timer.Delta(); now <= .001 {
		t.Error("Not eneough time has passed when Delta was called")
	}

	time.Sleep(5 * time.Millisecond)
	if now := timer.Delta(); now <= .001 {
		t.Error("Not eneough time has passed when Delta was called")
	}
}

func TestFPS(te *testing.T) {
	timer := NewTimer()
	timer.Reset()
	expected := 10

	if f := timer.FPS(); f != 0 {
		te.Errorf("Expected FPS to return 0 but got %v", f)
	}

	for i := 0; i < expected; i++ {
		timer.Update()
	}

	time.Sleep(1015 * time.Millisecond)

	if f := timer.FPS(); f != expected {
		te.Errorf("Expected FPS to return %v but got %v", expected, f)
	}
}
