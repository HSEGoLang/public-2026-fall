## urlfetch

В этой задаче нужно написать консольную утилиту,
которая принимает на вход произвольное количество http URL'ов и последовательно скачивает их содержимое.

При обработке невалидного URL'а программа должна завершаться с ненулевым exit кодом.

За основу можно взять [пример `urlfetch` из первой лекции](../../lectures/01-intro/urlfetch/urlfetch.go): он уже показывает весь алгоритм. Перенесите импорты и тело `main` в `01-Hello-World/urlfetch/main.go`, сохранив строку `//go:build !solution` и пустую строку после неё. Затем запустите форматтеры, тесты и линтер по [инструкции sum](../sum/README.md). На первом ДЗ можно пользоваться кодом лекции.

Все команды ниже выполняются из корня личного репозитория. На Windows используйте **Git Bash**.

### Примеры

Успешный запуск:
```
$ urlfetch https://golang.org https://go.dev
<!DOCTYPE html>
<html lang="en">
<meta charset="utf-8">
<meta name="description" content="Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.">
...
```

Неуспешный запуск:
```
$ urlfetch golang.org       
fetch: Get golang.org: unsupported protocol scheme ""
```

### Проверка решения

Для запуска тестов нужно выполнить следующую команду:

```
go test -v ./01-Hello-World/urlfetch/...
```

### Запуск программы

```
go run ./01-Hello-World/urlfetch https://go.dev
```

### Компиляция

```
go install ./01-Hello-World/urlfetch
```

После выполнения в `$GOPATH/bin` появится исполняемый файл с именем `urlfetch`.

### Линтер

Установите [golangci-lint](https://github.com/golangci/golangci-lint), если вы ещё этого не сделали, и проверьте решение перед отправкой:
```
golangci-lint -v run ./01-Hello-World/urlfetch/...
```

### Walkthrough

1. Чтение аргументов командной строки: https://gobyexample.com/command-line-arguments
2. HTTP запрос: https://golang.org/pkg/net/http/
