package integer

func MAX_INT() int {
	max := int32(0)
	for i := int32(0); i >= 0; i++ {
		max = i
	}
	return int(max)
}
func MIN_INT() int {
	min := int32(0)
	for i := int32(0); i <= 0; i-- {
		min = i
	}
	return int(min)
}
