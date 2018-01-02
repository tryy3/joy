package crypto

type RandomSource interface {
	GetRandomValues(array []byte) (b []byte)
}
