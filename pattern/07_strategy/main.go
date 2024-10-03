package main

import "fmt"

/*
 * Реализовать паттерн «СТРАТЕГИЯ».
 *
 * Поведенческий паттерн.
 *
 * Объяснить применимость паттерна, его плюсы и минусы,
 * а также реальные примеры использования данного примера на практике.
 *
 * https://en.wikipedia.org/wiki/Strategy_pattern
 *
 * Определяет семейство схожих алгоритмов и помещает каждый из них в 
 * собственный класс, после чего возможно предоставить взаимозаменяемость 
 * алгоритмов прямо во время исполнения.
 *
 * Используется при:
 * а) использовании разных вариаций какого-то алгоритма внутри объекта,
 * б) существовании множества похожих классов, немного отличающихся поведением,
 * в) сокрытии деталей реализации алгоритмов от других классов,
 * г) реализации вариаций алгоритмов с помощью крупного условного блока.
 *
 * + замена алгоритмов налету
 * + изоляция кода и данных алгоритмов от других классов
 * + уход к делегированию от наследования
 * + принцип открытости/закрытости
 * - усложнение доп. классами
 * - необходимость знания о разнице между алгоритмами на стороне клиента
 */

type Strategy interface {
	Execute(n1, n2 int) int
}

type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) Execute(n1, n2 int) int {
	return n1 + n2
}

type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) Execute(n1, n2 int) int {
	return n1 - n2
}

type ConcreteStrategyMultiply struct{}

func (s *ConcreteStrategyMultiply) Execute(n1, n2 int) int {
	return n1 * n2
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(n1, n2 int) int {
	return c.strategy.Execute(n1, n2)
}

func main() {
	var action string
	var n1, n2 int

	fmt.Println("Введите первую цифру:")
	fmt.Scan(&n1)
	fmt.Println("Введите вторую цифру:")
	fmt.Scan(&n2)
	fmt.Println("Введите действие (add, sub, mul):")
	fmt.Scan(&action)

	context := &Context{}

	switch action {
	case "add":
		context.SetStrategy(&ConcreteStrategyAdd{})
	case "sub":
		context.SetStrategy(&ConcreteStrategySubtract{})
	case "mul":
		context.SetStrategy(&ConcreteStrategyMultiply{})
	default:
		fmt.Println("Неизвестное действие!")
		return
	}

	result := context.ExecuteStrategy(n1, n2)

	fmt.Printf("Результат: %d\n", result)
}
