import "strings"

func toWeirdCase(str string) (exitString string) {
	runes := []rune(str)
	substringsArray, tempArray, tempString := strings.Split(string(runes), " "), []string{}, ""
  
	for i := 0; i < len(substringsArray); i++ {
		for j := 0; j < len(substringsArray[i]); j++ {
			if j % 2 == 0 {
				tempString += strings.ToUpper(string(substringsArray[i][j]))
			} else if j % 2 == 1 {
				tempString += strings.ToLower(string(substringsArray[i][j]))
			} else {
				tempString += string(substringsArray[i][j])
			}
		}
    
		tempArray = append(tempArray, tempString)
	  	tempString = ""
	}
  
	exitString = strings.Join(tempArray, " ")
	return exitString
}