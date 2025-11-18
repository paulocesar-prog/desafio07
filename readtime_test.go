package readtime

import "testing"

const sample = "Simples como voar. Chega ser lacrimejante. Giropops Strigus Girus"

func TestCount(t *testing.T) {
	if CountWords(sample) != 9 {
		t.Fatal("esperado 9 palavras")
	}
}

func TestReadingTime(t *testing.T) {
	min, _ := ReadingTime(sample, 200)
	if min != 0.05 {
		t.Fatalf("esperado 0.05, obtido %.2f", min)
	}
}

func TestErrorPPM(t *testing.T) {
	if _, err := ReadingTime(sample, 50); err == nil {
		t.Fatal("deveria falhar para PPM < 80")
	}
}