//low stands for 25$
//middle stands for 50$
//high stands for 100$
//total is total :|
type Pocket struct {
	low int
	middle int
	high int
	total int
}

func Tickets(peopleInLine []int) string {
	var vasyasPocket Pocket
  
	for i := 0; i < len(peopleInLine); i++ {
		switch(peopleInLine[i]) {
			case 25:
				vasyasPocket.low++
			case 50:
				if vasyasPocket.low != 0 {
					vasyasPocket.low--
  			 		vasyasPocket.middle++
        			} else {
          				return "NO"
        			}
      			case 100:
        			if vasyasPocket.middle != 0 && vasyasPocket.low != 0{
          				vasyasPocket.low--
          				vasyasPocket.middle--
          				vasyasPocket.high++
        			} else if vasyasPocket.low >= 3 {
          				vasyasPocket.low -= 3
          				vasyasPocket.high++
        			} else {
          				return "NO"
        			}
    		}
  	}

  	return "YES"
}