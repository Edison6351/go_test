package for_example

import (
	"fmt"
)

func for_example() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			fmt.Println("求余2大于1")
		}
		fmt.Println(n)
	}

}
