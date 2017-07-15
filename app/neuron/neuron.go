package neuron

import (
	"math/rand"
	"time"
)

var Ne *Neuron

func init() {
	rand.Seed(time.Now().UnixNano())
	M = rand.Float64() * 2
	B = rand.Float64() * 2
	Ne = NewNeuron(2)
	train(Ne, 100000, 0.2)
}

type Neuron struct {
	weights []float64
	bias    float64
}

func NewNeuron(n int32) *Neuron {
	var i int32
	w := make([]float64, n, n)
	for i = 0; i < n; i++ {
		w[i] = ((rand.Float64() * 2) - 1)
	}
	return &Neuron{
		weights: w,
		bias:    ((rand.Float64() * 2) - 1),
	}
}

func (p *Neuron) Process(inputs []float64) float64 {
	sum := p.bias
	for i, input := range inputs {
		sum += input * p.weights[i]
	}
	return heaviside(sum)
}
func heaviside(f float64) float64 {
	if f < 0 {
		return 0
	}
	return 1
}

// func sigmoid(x float64) float64 {
// 	return 1 / (1 + math.Exp(-x))
// }

func verify(p *Neuron) float64 {
	var correctAnswers float64
	for i := 0; i < 100; i++ {
		point := []float64{
			rand.Float64()*201 - 101,
			rand.Float64()*201 - 101,
		}
		result := p.Process(point)
		if result == isAboveLine(point, f) {
			correctAnswers++
		}
	}

	return correctAnswers
}

var M, B float64

func f(x float64) float64 {
	return M*x + B
}
func isAboveLine(point []float64, f func(float64) float64) float64 {
	x := point[0]
	y := point[1]
	if y > f(x) {
		// is above line
		return 1
	}
	// is below line
	return 0
}

func (p *Neuron) Adjust(inputs []float64, delta float64, learningRate float64) {
	for i, input := range inputs {
		p.weights[i] += input * delta * learningRate
	}
	p.bias += delta * learningRate
}

func train(p *Neuron, iters int, rate float64) {
	for i := 0; i < iters; i++ {
		point := []float64{
			rand.Float64()*201 - 101,
			rand.Float64()*201 - 101,
		}
		actual := p.Process(point)
		expected := isAboveLine(point, f)
		delta := expected - actual
		p.Adjust(point, delta, rate)
	}
}
