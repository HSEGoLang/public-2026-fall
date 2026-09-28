// Code written by AI

package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	n    = 1000000
	step = n / 5
)

type MemMonitor struct {
	name string
}

func NewMemMonitor(name string) *MemMonitor {
	return &MemMonitor{name: name}
}

func (m *MemMonitor) measureMemory(fn func(n int)) {
	// Принудительная сборка мусора перед измерением
	runtime.GC()
	runtime.GC() // дважды для надежности
	time.Sleep(10 * time.Millisecond)

	var before runtime.MemStats
	runtime.ReadMemStats(&before)

	fmt.Printf("\n=== %s ===\n", m.name)
	fmt.Printf("Память до: %d KB\n", before.Alloc/1024)

	// Выполняем функцию
	start := time.Now()
	fn(n)
	duration := time.Since(start)

	var after runtime.MemStats
	runtime.ReadMemStats(&after)

	fmt.Printf("Память после: %d KB\n", after.Alloc/1024)
	fmt.Printf("Разница: +%d KB\n", int64(after.Alloc-before.Alloc)/1024)
	fmt.Printf("Время выполнения: %v\n", duration)
	fmt.Printf("Объектов в куче: %d\n", after.HeapObjects)
	fmt.Printf("GC запусков: %d\n", after.NumGC-before.NumGC)

	// Принудительная очистка памяти
	runtime.GC()
	time.Sleep(10 * time.Millisecond)

	var afterGC runtime.MemStats
	runtime.ReadMemStats(&afterGC)
	fmt.Printf("Память после GC: %d KB\n", afterGC.Alloc/1024)
	fmt.Printf("Освобождено GC: %d KB\n", int64(after.Alloc-afterGC.Alloc)/1024)
}

func f_usual(n int) {
	val := make(map[string]struct{})
	defer func() { fmt.Printf("usual: использовано %d адресов\n", len(val)) }()

	for i := 0; i < n; i++ {
		if i%(step) == 0 {
			fmt.Printf("defer: обработано %d элементов\n", i)
		}
		func() { val[fmt.Sprintf("%v", &i)] = struct{}{} }()
	}
}

func f_defer(n int) {
	val := make(map[string]struct{})
	defer func() { fmt.Printf("defer: использовано %d адресов\n", len(val)) }()

	for i := 0; i < n; i++ {
		if i%(step) == 0 {
			fmt.Printf("defer: обработано %d элементов\n", i)
		}
		defer func() { val[fmt.Sprintf("%v", &i)] = struct{}{} }()
	}
}

func main() {
	fmt.Printf("Сравнение потребления памяти (n = %d)\n", n)
	fmt.Printf("========================================\n")

	// Тест 1: обычные функции
	monitor1 := NewMemMonitor("f_usual")
	monitor1.measureMemory(f_usual)

	time.Sleep(100 * time.Millisecond)

	// Тест 2: defer с правильной передачей параметров
	monitor2 := NewMemMonitor("f_defer")
	monitor2.measureMemory(f_defer)
}
