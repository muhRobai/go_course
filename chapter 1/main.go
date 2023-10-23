package main

import "fmt"

func main() {
	var number = 10
	{
		{
			name := "joko"
			number += 10
			fmt.Println(number, name)
		}
		fmt.Println(number)
	}

	{

		var i = 10
		for i > 0 {
			fmt.Println(i)
			i--
		}
	}

	{
		var name = "lala"
		for i := 0; i < len(name); i++ {
			fmt.Println(string(name[i]))
		}
	}

	{
		var name = "asas"
		for key, value := range name {
			fmt.Println(key, string(value))
		}
	}

}
