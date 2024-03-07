package simpletest

import "testing"

func TestDivide(t *testing.T) {
	testCases := []struct {
		name      string
		a, b      float64
		want      float64
		wantError bool
	}{
		{"Позитивний тест", 10.0, 2.0, 5.0, false},
		{"Негативний тест: Ділення на нуль", 10.0, 0, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Divide(tc.a, tc.b)
			if tc.wantError {
				if err == nil {
					t.Errorf("очікувалась помилка, але отримано nil")
				}
			} else {
				if err != nil {
					t.Errorf("не очікувалась помилка, але отримано: %v", err)
				}
				if got != tc.want {
					t.Errorf("отримано %v, очікувалося %v", got, tc.want)
				}
			}
		})
	}
}
