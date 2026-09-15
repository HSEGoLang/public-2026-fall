# expr

Реализуйте функцию `expr.Eval`, вычисляющую значение арифметического дерева.

Узел `Expression` либо хранит число (`OperationValue`), либо задаёт бинарную операцию над двумя поддеревьями:

- `OperationAdd` — сложение;
- `OperationSubtract` — вычитание;
- `OperationMultiply` — умножение;
- `OperationDivide` — целочисленное деление.

Пример для `(7 + 5) * 3`:

```go
expression := &expr.Expression{
	Operation: expr.OperationMultiply,
	Left: &expr.Expression{
		Operation: expr.OperationAdd,
		Left:      &expr.Expression{Value: 7},
		Right:     &expr.Expression{Value: 5},
	},
	Right: &expr.Expression{Value: 3},
}
```

Zero value поля `Operation` равен `OperationValue`, поэтому лист можно записать коротко: `&expr.Expression{Value: 7}`. Поля `Left` и `Right` у листа игнорируются.

`Eval` должна:

- вернуть ошибку, совместимую с `ErrInvalidExpression`, для `nil`, неизвестной операции или бинарного узла без одного из операндов;
- вернуть ошибку, совместимую с `ErrDivisionByZero`, при делении на ноль;
- сохранять исходную ошибку в цепочке так, чтобы её можно было найти через `errors.Is`;
- не изменять дерево.

Тип операции текущего узла проверяется до вычисления поддеревьев. Поэтому узел
с неизвестной `Operation` возвращает `ErrInvalidExpression`, даже если одно из
его поддеревьев тоже содержит ошибку.

Переполнение `int64` можно не проверять: арифметика должна вести себя как обычные операции Go.

## Проверка

Из корня репозитория:

```shell
go test ./02-Language/expr/...
```
