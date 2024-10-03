package main

import "fmt"

/*
 * Реализовать паттерн «ЦЕПОЧКА ОБЯЗАННОСТЕЙ».
 *
 * Поведенческий паттерн.
 *
 * Объяснить применимость паттерна, его плюсы и минусы,
 * а также реальные примеры использования данного примера на практике.
 *
 * https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
 *
 * Последовательная обработка запросов.
 *
 * Используется при:
 *  а) обработке различных запросов различными способами,
 *  б) строгом порядке обработки запросов,
 *  в) потенциальном расширении набора обработчиков.
 *
 * + уменьшение зависимости между клиентами и обработчиков
 * + принцип открытости/закрытости
 * + принцип единственной обязанности
 * - возможность отсутствия необходимого обработчика
 */

type Support interface {
	SetNext(next Support)
	HandleRequest(level int)
}

type BaseSupport struct {
	next Support
}

func (b *BaseSupport) SetNext(next Support) {
	b.next = next
}

type JuniorSupport struct {
	BaseSupport
}

func (j *JuniorSupport) HandleRequest(level int) {
	if level == 1 {
		fmt.Println("Junior Support: Обработал простой запрос.")
	} else if j.next != nil {
		j.next.HandleRequest(level)
	}
}

type MiddleSupport struct {
	BaseSupport
}

func (m *MiddleSupport) HandleRequest(level int) {
	if level == 2 {
		fmt.Println("Middle Support: Обработал средний запрос.")
	} else if m.next != nil {
		m.next.HandleRequest(level)
	}
}

type SeniorSupport struct {
	BaseSupport
}

func (s *SeniorSupport) HandleRequest(level int) {
	if level == 3 {
		fmt.Println("Senior Support: Обработал сложный запрос.")
	} else if s.next != nil {
		s.next.HandleRequest(level)
	}
}

func main() {
	junior := &JuniorSupport{}
	middle := &MiddleSupport{}
	senior := &SeniorSupport{}

	junior.SetNext(middle)
	middle.SetNext(senior)

	fmt.Println("Запрос уровня 1:")
	junior.HandleRequest(1)

	fmt.Println("\nЗапрос уровня 2:")
	junior.HandleRequest(2)

	fmt.Println("\nЗапрос уровня 3:")
	junior.HandleRequest(3)
}
