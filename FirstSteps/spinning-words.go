import "strings"

func SpinWords(str string) string {
	wordsArray := strings.Fields(str)
	for i := 0; i < len(wordsArray); i++ {
		if len(wordsArray[i]) >= 5 {
			singleWord := []rune(wordsArray[i])
			for i, j := 0, len(singleWord) - 1; i < j; i, j = i + 1, j - 1 {
				singleWord[i], singleWord[j] = singleWord[j], singleWord[i]
			}
      
			wordsArray[i] = string(singleWord)
		}
	}
	
	return strings.Join(wordsArray, " ")
}