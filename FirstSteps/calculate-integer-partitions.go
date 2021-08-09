func Partitions(n int) int {
	return CalculatePartitions(n, n)
}

func CalculatePartitions(currentSum int, maxValue int) int {
	if currentSum == maxValue {
		return 1 + CalculatePartitions(currentSum, maxValue - 1);
	}
	if maxValue == 0 || currentSum < 0 {
		return 0
	}
	if maxValue == 1 || currentSum == 0 {
		return 1
	}
  
	return CalculatePartitions(currentSum, maxValue - 1) + CalculatePartitions(currentSum - maxValue, maxValue)
}