import "math"

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
  
	isNegative := true
	maxSum, currentSum := math.MinInt32, 0
  
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 { isNegative = false }
    
		currentSum += numbers[i]
		if currentSum < 0 {
			currentSum = 0
		} else if maxSum < currentSum {
			maxSum = currentSum
		}
	}
  
	if isNegative == true { return 0 }
  
	return maxSum
}