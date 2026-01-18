package test

import (
	"testing"

	"github.com/6tail/tyme4go/tyme"
)

func TestDogDay0(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2011, 7, 14)
	excepted := "初伏"
	got := d.GetDogDay().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "初伏"
	got = d.GetDogDay().GetDog().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "初伏第1天"
	got = d.GetDogDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestDogDay1(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2011, 7, 23)
	excepted := "初伏"
	got := d.GetDogDay().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "初伏"
	got = d.GetDogDay().GetDog().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "初伏第10天"
	got = d.GetDogDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestDogDay2(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2011, 7, 24)
	excepted := "中伏"
	got := d.GetDogDay().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "中伏"
	got = d.GetDogDay().GetDog().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "中伏第1天"
	got = d.GetDogDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestDogDay3(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2011, 8, 12)
	excepted := "中伏"
	got := d.GetDogDay().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "中伏"
	got = d.GetDogDay().GetDog().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "中伏第20天"
	got = d.GetDogDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestDogDay4(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2011, 8, 13)
	excepted := "末伏"
	got := d.GetDogDay().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "末伏"
	got = d.GetDogDay().GetDog().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "末伏第1天"
	got = d.GetDogDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}

func TestDogDay5(t *testing.T) {
	d, _ := tyme.SolarDay{}.FromYmd(2011, 8, 22)
	excepted := "末伏"
	got := d.GetDogDay().GetName()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "末伏"
	got = d.GetDogDay().GetDog().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}

	excepted = "末伏第10天"
	got = d.GetDogDay().String()
	if excepted != got {
		t.Errorf("excepted: %v, got: %v", excepted, got)
	}
}
