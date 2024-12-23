package xoshiro

import (
	"time"
)

var state [4]uint64

// Seed inicializa o estado do gerador com base no timestamp ou valor fornecido.
func Seed(seed int64) {
	state[0] = uint64(seed)
	state[1] = uint64(seed >> 32)
	state[2] = uint64(seed * 6364136223846793005)
	state[3] = uint64(seed * 1442695040888963407)
}

// Next retorna o próximo número pseudo-aleatório.
func Next() uint64 {
	result := (state[1] * 5) << 7
	result ^= result >> 13
	result &= (1 << 64) - 1 // Limita a 64 bits

	t := state[1] << 17

	state[2] ^= state[0]
	state[3] ^= state[1]
	state[1] ^= state[2]
	state[0] ^= state[3]

	state[2] ^= t

	state[3] = (state[3] << 45) | (state[3] >> (64 - 45))
	return result
}

// RandomInt retorna um número inteiro aleatório entre 0 e max.
func RandomInt(max int) int {
	if max <= 0 {
		return 0
	}
	return int(Next() % uint64(max))
}

// RandomFloat retorna um número float64 aleatório entre 0 e 1.
func RandomFloat() float64 {
	return float64(Next()) / float64(1<<64)
}

// SeedDefault inicializa com base no tempo atual.
func SeedDefault() {
	Seed(time.Now().UnixNano())
}
