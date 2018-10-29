package UF

import (
	"math"
	"problem"
	"solution"
	"util"
)

type UF6 struct {
	name                string
	numberOfVariables   int
	numberOfObjectives  int
	numberOfConstraints int
	upper               []float64
	lower               []float64
	rand                *util.Random
}

func (this *UF6) NewUF6(nd, no, nc int) {
	this.name = "UF6"
	this.numberOfVariables = nd
	this.numberOfObjectives = no
	this.numberOfConstraints = nc
	this.upper = make([]float64, nd, nd)
	this.lower = make([]float64, nd, nd)

	this.lower[0] = 0
	this.upper[0] = 1
	for i := 1; i < nd; i++ {
		this.lower[i] = -1
		this.upper[i] = 1
	}

	this.rand = new(util.Random)
	this.rand.NewRand()
}

func (this *UF6) GetName() string {
	return this.name
}

func (this *UF6) GetNumberOfObjectives() int {
	return this.numberOfObjectives
}

func (this *UF6) GetNumberOfVariables() int {
	return this.numberOfVariables
}

func (this *UF6) GetNumberOfConstraints() int {
	return this.numberOfConstraints
}

func (this *UF6) CreateSolution() *solution.Solution {
	solution := new(solution.Solution)
	solution.NewSolution(this.numberOfVariables, this.numberOfObjectives, this.numberOfConstraints, this.lower, this.upper)
	for i := 0; i < this.numberOfVariables; i++ {
		value := this.rand.Float64() * (this.upper[i] - this.lower[i])
		solution.SetVariableValue(i, value)
	}

	return solution
}

func (this *UF6) Evaluate(solution *solution.Solution) {
	x := problem.GetReal(solution)
	f := []float64{0, 0}

	nx := this.numberOfVariables

	count1 := 0
	count2 := 0
	sum1 := 0.0
	sum2 := 0.0
	prod1 := 1.0
	prod2 := 1.0
	var yj, hj, pj float64
	N := 2.0
	E := 0.1

	for j := 2; j <= nx; j++ {
		yj = x[j-1] - math.Sin(6.0*PI*x[0]+float64(j)*PI/float64(nx))
		pj = math.Cos(20.0 * yj * PI / math.Sqrt(float64(j)))

		if j%2 == 0 {
			sum2 += yj * yj
			prod2 *= pj
			count2++
		} else {
			sum1 += yj * yj
			prod1 *= pj
			count1++
		}
	}

	hj = 2.0 * (0.5/N + E) * math.Sin(2.0*N*PI*x[0])

	if hj < 0.0 {
		hj = 0.0
	}

	f[0] = x[0] + hj + 2.0*(4.0*sum1-2.0*prod1+2.0)/float64(count1)
	f[1] = 1.0 - x[0] + hj + 2.0*(4.0*sum2-2.0*prod2+2.0)/float64(count2)

	solution.SetObjective(0, f[0])
	solution.SetObjective(1, f[1])
}
