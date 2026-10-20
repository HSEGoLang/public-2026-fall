# testequal

Реализуйте четыре test helper по аналогии с assert и require:

```go
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{})
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{})
```

`T` — интерфейс из `t.go`. Его реализует `*testing.T`.
`AssertEqual` возвращает результат сравнения, `AssertNotEqual` — его отрицание.
При провале они сообщают об ошибке через `Errorf`, но выполнение теста
продолжается. `Require*` при провале также вызывают `FailNow`, завершая текущий
тест. При успехе ни `Errorf`, ни `FailNow` не вызываются.

Пометьте вспомогательные функции через `Helper`, чтобы сообщение указывало
на строку вызова в тесте, а не на внутренности вашей библиотеки. Вызывать
`Require*` следует из горутины самого теста: так устроен `testing.T.FailNow`.

## Равенство в этой задаче

Это ограниченный учебный API. Поддерживаются следующие динамические типы:

- `int`, `uint`, `int8`, `uint8`, `int16`, `uint16`, `int32`, `uint32`, `int64`, `uint64`;
- `string`;
- `map[string]string`;
- `[]int` и `[]byte`.

Числа и строки сравниваются по значению только при совпадении динамических
типов: `int(1)` и `int64(1)` не равны. `byte` и `rune` — псевдонимы
`uint8` и `int32`, поэтому поддерживаются. Отдельный именованный тип вроде
`type Age int` не входит в список.

Срезы равны при одинаковой длине и одинаковых элементах в том же порядке;
capacity не важна. У map должны совпадать все пары ключ/значение, порядок
обхода не важен. Входные значения изменять нельзя.

Два типизированных `nil` одного поддерживаемого типа равны:
`[]int(nil)` равен `[]int(nil)`. Nil и пустая коллекция различаются:
`[]int(nil)` не равен `[]int{}`; то же правило действует для `[]byte`
и `map[string]string`. Коллекции разных типов не равны.

Все остальные случаи считаются неравными, без panic: в частности,
`nil` и `nil`, `bool`, `float64`, `uintptr`, указатели, структуры и
неподдерживаемые коллекции. Поэтому `AssertNotEqual(t, true, true)`
успешен. Это намеренная граница учебного API; она отличается от testify.

## Сообщения

При провале сообщение должно содержать причину (`not equal` или `equal`),
значения `expected` и `actual` и пользовательское сообщение. Форматировать
значения нужно так, чтобы корректно отображались все поддерживаемые типы,
включая строки и map. Точный отступ не проверяется.

`msgAndArgs` задаёт пользовательское сообщение:

- нет аргументов — пустое сообщение;
- один аргумент — строка как есть, другое значение как `fmt.Sprintf("%+v", value)`;
- несколько аргументов — первый является строкой формата, остальные передаются
  в `fmt.Sprintf`. Вызывающий код гарантирует строковый первый аргумент.

Например, `AssertEqual(t, "cat", "dog", "%s differs", "animal")`
возвращает false, сообщает оба значения и animal differs, затем тест
продолжается.

Из корня репозитория:

```shell
go test ./06-Testing/testequal/...
go test -race ./06-Testing/testequal/...
```

Полезные справочные страницы: [testing.T](https://pkg.go.dev/testing#T),
[type assertions](https://go.dev/doc/effective_go#interface_conversions).
