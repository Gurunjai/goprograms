package sample

import "testing"

func TestPow2(t *testing.T) {
	in := map[int]int {
		77: 64,
		64: 64,
		10: 8,
		70000: 65536,
	}

	for num, want := range(in) {
		if got := pow2(num); got != want {
			t.Fatalf("Mismatch for %v \n\t\t Got: %v, \t Want: %v\n", num, got, want)
		}
	}
}

func TestGetMod(t *testing.T) {
	in := map[int]int {
		77: 13,
		64: 0,
		10: 2,
		70000: 4464,
	}
		
	for num, want := range(in) {
		if got := GetMod(num); got != want {
			t.Fatalf("Mismatched for %v \n\t\t Got: %v, \t Want: %v\n", num, got, want)
		}
	}
}

func TestParity(t *testing.T) {
	in := map[int]int {
		64: 1,
		70: 1,
		66: 0,
		10: 0,
		90: 0,
		65534: 1,
	}

	for num, want := range(in) {
		if got := Parity(uint64(num)); got != want {
			t.Fatalf("Mismatch for %v \n\t\t Got: %v, \t Want: %v\n", num, got, want)
		}
	}
}

func TestIsPow2(t *testing.T) {
	in := map[int]bool {
		64: true,
		70: false,
		65536: true,
		70000: false,
		1024: true,
		2050: false,
	}

	for num, want := range(in) {
		if got := IsPow2(num); got != want {
			t.Fatalf("Mismatch for %v\n \t\t Got: %v, \t Want: %v\n", num, got, want)
		}
	}
}

func TestGetRemainder(t *testing.T) {
	in := [][]int {
		[]int { 70, 64, 6},
		[]int { 70000, 65536, 4464},
		[]int { 15, 2, 1},
		[]int { 32768, 32768, 0},
		[]int { 35, 26, -1},
	}

	for _, val := range(in) {
		num1, num2, want := val[0], val[1], val[2]
		if got := GetRemainder(num1, num2); got != want {
			t.Fatalf("Mismatch for num1: %v, num2: %v\n \t\t Got: %v, \t Want: %v\n",
				num1, num2, got, want)
		}
	}
}

func TestMsbSet(t *testing.T) {
	in := map[int]int {
		1: 0,
		65536: 16,
		70000: 16,
		25: 4,
		1029: 10,
	}

	for num, want := range(in) {
		if got := msbSet(num); got != want {
			t.Fatalf("Mismatch for %v\n \t\t Got: %v, \t Want: %v\n", num, got, want)
		}
	}
}

func TestStrReplace(t *testing.T) {
	in := map[string]string {
		"          oll2345     eH65    ": "Hello",
	}

	for str, want := range(in) {
		if got := strReverse(str); got != want {
			t.Fatalf("Mismatch for %q\n \t Got: %q\n \t Want: %q\n", str, got, want)
		}
	}
}

func TestHasher(t *testing.T) {
	t.Errorf("Out Value Index for 0xa5a5a5f0 & 0x80: %v\n", destHasher(0xa5a5a5f0, 0x80))
	t.Errorf("Out Value Index for 0xa8a8a5f0 & 0x80: %v\n", destHasher(0xa8a8a5f0, 0x80))
	t.Errorf("Out Value Index for 0xa0a8a5f0 & 0x80: %v\n", destHasher(0xa0a8a5f0, 0x80))
	t.Errorf("Out Value Index for 0x9890a5f0 & 0x80: %v\n", destHasher(0x9890a5f0, 0x80))
}