package evolution

import (
	"math/rand"
)

// define a chromosome
type Chromosome struct {
	Fitness int
	Gene    []float64
}

func NewChromosome(len int) *Chromosome {
	return &Chromosome{
		Fitness: 0,
		Gene: func() []float64 {
			g := make([]float64, len)
			for i, _ := range g {
				g[i] = rand.Float64()
			}
			return g
		}(),
	}
}

// define a subpopulation
type Subpopulation struct {
	SubpSize    int           // size of subpopulation
	ChromSize   int           // size of chromosome
	Chromosomes []*Chromosome // gene pool for neurons
}

func NewSubpopulation(s, l int) *Subpopulation {
	return &Subpopulation{
		SubpSize:  s,
		ChromSize: l,
		Chromosomes: func() []*Chromosome {
			c := make([]*Chromosome, s)
			for i, _ := range c {
				c[i] = NewChromosome(l)
			}
			return c
		}(),
	}
}

// one point crossover
func (s *Subpopulation) Crossover1P(c1, c2 int) {
	cut := rand.Intn(s.SubpSize)
	for i := cut; i < s.Size; i++ {
		g1 := s.Chromosomes[c1].Gene[i]
		g2 := s.Chromosomes[c2].Gene[i]

		s.Chromosomes[c1].Gene[i] = g2
		s.Chromosomes[c2].Gene[i] = g1
	}
}
