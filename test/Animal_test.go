package test

import (
	"github.com/6tail/tyme4go/tyme"
	"testing"
)

func TestAnimal1(t *testing.T) {
	animal := tyme.Animal{}.FromIndex(0)
	excepted := "蛟"
	got := animal.String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestAnimal2(t *testing.T) {
	animal, err := tyme.Animal{}.FromName("蛟")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	excepted := "龙"
	got := animal.Next(1).String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
