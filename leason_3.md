# Frist Go Program
```
package main

import "fmt"

func main() {
    fmt.Println("Hello World!)
}
```

1. variable
    - var 
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
    - constant
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

2. basic data type
- string
- rune
- float
- uint
- int
- bool
- error
3. operators and expressions
- aritmatic operators
    - operator aritmatika adalah operator yang digunakan untuk proses perhitungan matematika seperti `-`, `+`, `*`, `/`, dan `%`
- logic operators
    - operator logika adalah operator yang digunakan untuk proses logika komputer seperti `&&`, `||`, `!`
- comparation operators
    - operator perbandingan adalah operator yang digunakan untuk membandingan 2 variabel seperti `==`, `!=`, `>=`, `<=`, `>`, `<`,

