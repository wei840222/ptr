package goptr

import "testing"

func TestFloat32(t *testing.T) {
	var f float32 = 123.12

	if *Float32(f) != f {
		t.Error("did not get a pointer back")
	}
}

func TestFloat64(t *testing.T) {
	var f float64 = 123.12

	if *Float64(f) != f {
		t.Error("did not get a pointer back")
	}
}

func TestInt(t *testing.T) {
	var i int = 123

	if *Int(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestInt32(t *testing.T) {
	var i int32 = 123

	if *Int32(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestInt64(t *testing.T) {
	var i int64 = 123

	if *Int64(i) != i {
		t.Error("did not get a pointer back")
	}
}

func TestString(t *testing.T) {
	var i string = "What was the question to life, universe and everything?"

	if *String(i) != i {
		t.Error("did not get a pointer back")
	}
}