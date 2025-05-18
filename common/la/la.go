package la

import "math"

type Vector struct {
	Data []float32
}

func (vector *Vector) L1() float32 {
	var sum float64

	for _, i := range vector.Data {
		sum += math.Abs(float64(i))
	}

	return float32(sum)
}

func (vector *Vector) L2() float32 {
	var sum float64

	for _, i := range vector.Data {
		sum += math.Pow(float64(i), 2)
	}

	return float32(math.Sqrt(sum))
}

func (v *Vector) Dot(o *Vector) float32 {
	var sum float32
	for i := range v.Data {
		sum += v.Data[i] * o.Data[i]
	}
	return sum
}

func (v *Vector) CosineSimiliarity(o *Vector) float32 {
	return v.Dot(o) / (v.L2() * o.L2())
}

func (v *Vector) Equal(o *Vector) bool {
	if len(v.Data) != len(o.Data) {
		return false
	}

	for i := range v.Data {
		if v.Data[i] != o.Data[i] {
			return false
		}
	}

	return true
}
