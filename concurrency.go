package main 

import "fmt"

func f(n int) {

	for index :=0; index < 10; index ++ {
		fmt.Println(n, ":", index)
	}
}

func main() {

for index :=0; index < 10; index ++ {
		go f(index)
	}
	

	var input string
	//fmt.Println("index")
	fmt.Scanln(&input)

}