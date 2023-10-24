# Frist Go Program
```
package main

import "fmt"

func main() {
    fmt.Println("Hello World!)
}
```

## variable
### var 
`var` adalah tipe variabel yang dapat berubah. jika sebuah variabel di deklarasi dengan menggunakan var maka isi valur dari variabel tersebut dapat berubah. deklarasi variabel `var` juga dapat diganti menjadi `=:` 
```
    package main

    import "fmt"

    func main() {
        var number = 0
        fmt.Println("value number: " + number)

        number = 1
        fmt.Println("value number: " + number)
    }
```
### constant
`const` adalah tipe variabel yang tidak dapat berubah valuenya atau variabel tetap. jika sebuah variabel dideklarasikan dengan menggunakan `const` maka variable tersebut bersifat absolut dan tidak berubah ubah. jika tipe variable `const` diubah maka akan muncul error
```
    package main

    import "fmt"

    func main() {
        const pi = 3.14
        var r = 12
        
        result := pi * r * r

        fmt.Println("luas lingkaran: ", result)
    }
```

## basic data type
### string
tipe data `string` string adalah tipe data yang digunakan untuk menyimpan teks. dalam golang `string` sendiri terdiri dari slice dan byte.
```
    package main

    import "fmt"

    func main() {
        var name = "Lukas"
        fmt.Println("Hello ", name)
    }
```
### rune
tipe data `rune` dapat dikatakan mirip dengan `int32` valuenya sama. tapi `rune` sendiri merepresentasikan unique code. `rune` memungkinkan golang melakukan input dalam bentuk code yang tidak umum seperti emoji, bahasa jepang, china dll
```
    package main

    import "fmt"

    func main() {
        smile := '😀'

        fmt.Printf("Data type of %c is %T and the rune value is %U \n", smile, smile, smile);    
    }
``` 
### float
tipe data `float` terdiri dari 2, `float32` dan `float32`. tipe data `float` biasanya digunakan untuk menyimpan nilai decimal suatu angka. perbedaan dari `float32` dan `float64` pada seberapa besar angka yang dapat disimpan `float32` dapat menampung 32 bit sedangkan `float64` dapat menampung 64 bit
```
    package main

    import "fmt"

    func main() {
        var decimal float32
        decimal = 0.3322211330

        fmt.Println(decimal)
    }
```
### uint
tipe data `uint` adalah tipe data yang biasa digunakan untuk menyimpan number dengan value positif. tipe data `uint` senditi terdiri dari `uint8`, `uint16`, `uint32`, `uint64`.
```
    package main

    import "fmt"

    func main() {
        var numb uint8 = 0
        fmt.Println(numb)
    }
```
### int
tipe data `int` adalah tipe data yang paling sering digunakan untuk menyimpan number dengan value negative maupun positif. tipe data `int` memiliki terdiri dari `int8`, `int16`, `int32`, `int64`.
```
    package main

    func main() {
        var numb int8 = 0
        fmt.Println(numb)
    }
```
### bool
tipe data `bool` adalah tipe data yang digunakan untuk menyatakan kebenaran. tipe `bool` disebut juga sebagai tipe data kondisi logika.
```
    package main

    func main() {
        var isHot bool = true
        fmt.Println(isHot)
    }
``` 

## Operators and Expressions
### aritmatic operators,
operator aritmatika adalah operator yang digunakan untuk proses perhitungan matematika seperti `-`, `+`, `*`, `/`, dan `%`
```
package main

import "fmt"

func main() {
    var length = 10
    var width = 5

    fmt.Println("rectagle area = ", length * width)
    
    // modulus
    fmt.Println("modulus of 2", 2%2)
}
```
### logic operators,
operator logika adalah operator yang digunakan untuk proses logika komputer seperti `&&`, `||`, `!`
```
package main

import "fmt"

func main() {
    var eat = true
	var happy = true

	fmt.Println("if you eat you will happy", eat && happy)
}
```
### comparation operators,
operator perbandingan adalah operator yang digunakan untuk membandingan 2 variabel seperti `==`, `!=`, `>=`, `<=`, `>`, `<`.
```
package main

import "fmt"

func main() {
    var right = 10
    var left = 11

    fmt.Println("rigth is greather than left? ", right > left)
}
```

