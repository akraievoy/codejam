package algo

import (
	"math/big"
	"reflect"
	"testing"
)

func TestRatToContFrac1(t *testing.T) {
	a := big.NewRat(61, 14)
	actual := RatToContFrac(a)
	expected := ContFracNew(4,2,1,4)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RatToContFrac(big.NewRat(61, 14)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestRatToContFrac2(t *testing.T) {
	a := big.NewRat(105, 39)
	actual := RatToContFrac(a)
	expected := ContFracNew(2,1,2,4)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RatToContFrac(big.NewRat(105, 39)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestRatToContFrac3(t *testing.T) {
	a := big.NewRat(0, 100)
	actual := RatToContFrac(a)
	expected := ContFracNew(0)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RatToContFrac(big.NewRat(0, 100)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestRatToContFrac4(t *testing.T) {
	a := big.NewRat(100, 100)
	actual := RatToContFrac(a)
	expected := ContFracNew(1)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("RatToContFrac(big.NewRat(100, 100)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestContFracToRat1(t *testing.T) {
	cf := ContFracNew(4,2,1,4)
	actual := ContFracToRat(cf)
	expected := big.NewRat(61, 14)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ContFracToRat(cf) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestContFracToRat2(t *testing.T) {
	cf := ContFracNew(2,1,2,4)
	actual := ContFracToRat(cf)
	expected := big.NewRat(105, 39)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ContFracToRat(cf) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestTrailingOne1(t *testing.T) {
	cf := ContFracNew(3, 7, 15, 2, 7, 1, 4, 2)
	actual := ContFracToRat(TrailingOne(cf))
	expected := ContFracToRat(cf)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ContFracToRat(TrailingOne(cf)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestTrailingOne2(t *testing.T) {
	cf := ContFracNew(1)
	actual := ContFracToRat(TrailingOne(cf))
	expected := ContFracToRat(cf)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ContFracToRat(TrailingOne(cf)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestTrailingOne3(t *testing.T) {
	cf := ContFracNew(10)
	actual := ContFracToRat(TrailingOne(cf))
	expected := ContFracToRat(cf)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ContFracToRat(TrailingOne(cf)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestBestApproximate(t *testing.T) {
	actual := BestApproximate(big.NewRat(314155, 100000), big.NewRat(314165, 100000))
	expected := big.NewRat(355, 113)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("BestApproximate(big.NewRat(314155, 100000), big.NewRat(314165, 100000)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestBestApproximate2(t *testing.T) {
	actual := BestApproximate(big.NewRat(0, 1), big.NewRat(1, 1))
	expected := big.NewRat(1, 2)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("BestApproximate(big.NewRat(0, 1), big.NewRat(1, 1)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestBestApproximate3(t *testing.T) {
	actual := BestApproximate(big.NewRat(1, 1), big.NewRat(2, 1))
	expected := big.NewRat(3, 2)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("BestApproximate(big.NewRat(1, 1), big.NewRat(2, 1)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestBestApproximate4(t *testing.T) {
	actual := BestApproximate(big.NewRat(3, 44), big.NewRat(28, 56))
	expected := big.NewRat(1, 3)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("BestApproximate(big.NewRat(3, 44), big.NewRat(28, 56)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}

func TestBestApproximate5(t *testing.T) {
	actual := BestApproximate(big.NewRat(0, 1), big.NewRat(-1, -3))
	expected := big.NewRat(1, 4)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("BestApproximate(big.NewRat(0, 1), big.NewRat(-1, -3)) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
	}
}