package utilitas

import "errors"

type Riwayat struct {
	A        float64
	B        float64
	Operator string
	Hasil    float64
}

func CekGanjilGenap(angka int) string {
	if angka%2 == 0 {
		return "Genap"
	}
	return "Ganjil"
}

func Hitung(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("tidak bisa dibagi nol")
		}
		return a / b, nil
	default:
		return 0, errors.New("operator tidak valid")
	}
}

func DeretAngka(n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("n harus lebih besar dari 0")
	}
	hasil := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		hasil = append(hasil, i)
	}
	return hasil, nil
}
