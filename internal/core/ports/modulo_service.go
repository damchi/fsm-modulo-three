package ports

// ModuloService defines a port to compute modulo.
type ModuloService interface {
	Compute(binary string, mod int) (int, error)
}
