# Методы и интерфейсы

В домашнем задании шесть задач, каждая по 100 баллов.

| Задача | Что практикуем |
| --- | --- |
| [lrucache](lrucache/README.md) | Свой тип с состоянием, методы и реализация интерфейса Cache |
| [linefilter](linefilter/README.md) | Функция, принимающая Reader, Writer и callback |
| [otp](otp/README.md) | Собственные реализации Reader и Writer, обёртки над потоками |
| [retryread](retryread/README.md) | Ошибки, ограниченные повторы, errors.Is/As и Unwrap |
| [yamlembed](yamlembed/README.md) | Embedding и методы, которые библиотека вызывает через интерфейсы |
| [ciletters](ciletters/README.md) | Генерация текста через text/template и встраивание шаблона с go:embed |

Можно идти в порядке таблицы. Для `otp` в условии разобраны XOR и нужные
правила Read/Write; для `yamlembed` — YAML, теги и интерфейсы библиотеки.
Перед `retryread` пригодятся бонусные слайды об ошибках и `errors.As`.
Для `ciletters` понадобятся условные блоки и циклы в шаблоне; ссылки есть в условии.

Команды из условий запускаются из корня репозитория. Например:

```shell
go test ./04-Interfaces/lrucache/...
```
