package qubits

// Qubit is structure for qubits
type Qubit struct{ 
	Qubit float64
}

type Qubits []*Qubit

func Exconst(num float64) (r *Qubit) {
	r = new(Qubit)
	r.Qubit = num
	return r
}
