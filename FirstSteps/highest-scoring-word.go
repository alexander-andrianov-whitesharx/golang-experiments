import "strings"

func High(s string) (exitWord string) {
	words := strings.Fields(s)
  
  	// I have used 'difference' variable to get clear values of single characters
	var difference, totalSum, tempSum byte = 96, 0, 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			tempSum += words[i][j] - difference
		}
		if tempSum > totalSum {
			totalSum = tempSum
			exitWord = words[i]
	  }
    
	  tempSum = 0
	}
  
	return exitWord
}