package main

import (
	"fmt"
	"math"
)

const precision = 0.0001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// fmtパッケージは変数を文字列で出力する際に
	// errorインタフェースを確認しているので、fmt.Sprint(e)を呼び出すと
	// ErrNegativeSqrt.Error()が呼ばれ無限ループになる
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	mathZ := math.Sqrt(x)
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z * z - x) / (2 * z)
		if math.Abs(z - mathZ) < precision {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
