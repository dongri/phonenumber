package phonenumber

import "testing"

func BenchmarkParseWithLandLine(b *testing.B) {
	benchmarks := []struct {
		number  string
		country string
	}{
		{"+371 (67) 881-727", "LV"},
		{"00371 (67) 881-727", "LV"},
		{"003726823000", "EE"},
		{"+3726823000", "EE"},
		{"3726823000", "EE"},
		{"7499 709 88 33", "RU"},
		{"22 (483) 53-34", "PL"},
		{"48224835334", "PL"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.number, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ParseWithLandLine(bm.number, bm.country)
			}
		})
	}
}
