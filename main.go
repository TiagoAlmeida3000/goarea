package goarea

import "math"

// Pi é uma proporção numerica
const Pi = 3.1416

//Circ é responsavel por clacular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect pe responsavel por calcular a area
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
