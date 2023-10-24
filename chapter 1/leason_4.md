# Flow control
## condition
### if else
if else adalah syntax golang yang digunakan untuk menenrukan kondisi a atau b dalam sebuah program
```
package main

import "fmt"

func main() {
    var left = 10
    var right = 11

    if left == right {
        fmt.Println("left is equal to right")
    } else {
        fmt.Println("left is not equal to right")
    }
}
```
### swicth
swicth memiliki konsep yang sama dengan if else tapi untuk swicth cuma yang membedakan jika tidak mendapat kondisi yang sesuai, dengan switch dapat memberikan prosess default
```
package main

import "fmt"

func main() {
    distric := "jogja"

    switch distric {
        case "jogja":
            fmt.Println("welcome to jogja")
        case "semarang": 
            fmt.Println("welcome to semarang")
        default:
            fmt.Println("where you go?")
    }
}
```

## looping
### for
for diguakan untuk mengulang poses dalam suatu program, contohnya for dapat digunakan untuk menguraikan semua semua value dalam `array`. dalam golang sendiri looping hanya terdiri dari `for` namun implementasi terbagi menjadi 3.
```
package main

import "fmt"

func main() {
    var name = "hilda"

	for i := 0; i < len(name); i++ {
		fmt.Println(string(name[i]))
	}

	i := 0
	for i < len(name) {
		fmt.Println(string(name[i]))
		i++
	}

	for i, val := range name {
		fmt.Println(i, string(val))
	}
}
```
### funcition
```
```

### scope
```
```