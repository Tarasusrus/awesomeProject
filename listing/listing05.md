Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
напечатает "ok".

функция test() возвращает *customError, который также удовлетворяет интерфейсу error, 
поскольку *customError реализует метод Error(). Функция test() возвращает nil,
```
