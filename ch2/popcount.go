package ch2

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1)
	}
}

func Popcount(x uint64) int {
	return int(pc[byte(x>>0)] +
		pc[byte(x>>8)] +
		pc[byte(x>>16)] +
		pc[byte(x>>24)] +
		pc[byte(x>>32)] +
		pc[byte(x>>40)] +
		pc[byte(x>>48)] +
		pc[byte(x>>56)])
}
