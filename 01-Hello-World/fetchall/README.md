## fetchall

В этой задаче нужно написать консольную утилиту,
которая принимает на вход произвольное количество http URL'ов и скачивает их содержимое **конкурентно**.

Программа не должна останавливаться на невалидном URL'e.
Текст ответов можно игнорировать.
Вместо этого можно залогировать прогресс в произвольном формате.

За основу можно взять [пример `fetchall` из первой лекции](../../lectures/01-intro/fetchall/fetchall.go). В `01-Hello-World/fetchall/main.go` нужны импорты и **обе** функции: `main` и `fetch`. Сохраните строку `//go:build !solution` и пустую строку после неё. `go fetch(...)` запускает обработку очередного URL, а чтение из `ch` ждёт результаты; подробное устройство горутин и каналов разберём позже. Для первого ДЗ достаточно собрать и запустить лекционный пример, затем применить [форматтеры и проверки](../sum/README.md).

Все команды ниже выполняются из корня личного репозитория. На Windows используйте **Git Bash**.

Пример:
```
$ fetchall https://gopl.io golang.org http://golang.org
Get golang.org: unsupported protocol scheme ""
1.05s    11071  http://golang.org
2.18s     4154  https://gopl.io
2.18s elapsed
```

В примере логируются времена обработки индивидуальных запросов, размеры ответов и общее время работы программы.
Можно видеть, что общее время работы равно максимуму, а не сумме времён индивидуальных запросов.

### Проверка решения

Для запуска тестов нужно выполнить следующую команду:

```
go test -v ./01-Hello-World/fetchall/...
```

### Запуск программы

```
go run ./01-Hello-World/fetchall https://go.dev https://go.dev/doc/
```

### Компиляция

```
go install ./01-Hello-World/fetchall
```

После выполнения в `$GOPATH/bin` появится исполняемый файл с именем `fetchall`.

### Линтер

Установите [golangci-lint](https://github.com/golangci/golangci-lint), если вы ещё этого не сделали, и проверьте решение перед отправкой:
```
golangci-lint -v run ./01-Hello-World/fetchall/...
```

### Ссылки

1. Чтение аргументов командной строки: https://gobyexample.com/command-line-arguments
2. HTTP запрос: https://golang.org/pkg/net/http/
3. Запуск горутин: https://gobyexample.com/goroutines
4. Ожидание завершения горутин: https://gobyexample.com/channels
4. Замер времени: https://golang.org/pkg/time/#Since
