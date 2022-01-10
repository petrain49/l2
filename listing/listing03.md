Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

// false т.к. тип err == PathError
// интерфейс содержит информацию о типе переменной, которая была присвоена интерфейсу, и ссылку на значение
// в инофрмации о типе хранится сам тип, размер, хэш и т.д
// пустой интерфейс не содержит информацию о базовом типе и не имеет методов, т.е. переменная любого типа соответствует пустому интерфейсу

```