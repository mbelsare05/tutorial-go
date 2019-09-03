package tutorial

// FuzzerEntrypoint is the method Fuzzbuzz will repeatedly
// run with new tests to try and find bugs in BrokenMethod
func FuzzerEntrypoint(Data []byte) int {
	Fuzz(Data)
	return 0
}
