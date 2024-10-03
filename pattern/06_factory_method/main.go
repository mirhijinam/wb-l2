package main

import "fmt"

/*
 * Реализовать паттерн «ФАБРИЧНЫЙ МЕТОД».
 *
 * Порождающий паттерн.
 *
 * Объяснить применимость паттерна, его плюсы и минусы,
 * а также реальные примеры использования данного примера на практике.
 *
 * https://en.wikipedia.org/wiki/Factory_method_pattern
 *
 * Определение общего интерфейса для создания объектов главного класса,
 * позволяя подклассам изменять тип создаваемых объектов.
 *
 * Используется при:
 * а) неопределенности типов и зависимостей,
 * б) желании предоставить возможность расширения фреймворка.
 *
 * + нет привязки к конкретным классам продуктам
 * + код производства продуктов в одном месте
 * + упрощение добавления новых типов продуктов
 * + принцип открытости/закрытости
 * - параллельные иерархии классов
 */

type Button interface {
	Render()
}

type WindowsButton struct{}

func (b *WindowsButton) Render() {
	fmt.Println("Rendering Windows Button")
}

type HTMLButton struct{}

func (b *HTMLButton) Render() {
	fmt.Println("Rendering HTML Button")
}

type Dialog interface {
	CreateButton() Button
	Render()
}

type WindowsDialog struct{}

func (d *WindowsDialog) CreateButton() Button {
	return &WindowsButton{}
}

func (d *WindowsDialog) Render() {
	button := d.CreateButton()
	button.Render()
}

type WebDialog struct{}

func (d *WebDialog) CreateButton() Button {
	return &HTMLButton{}
}

func (d *WebDialog) Render() {
	button := d.CreateButton()
	button.Render()
}

func main() {
	configOS := "Windows"

	var dialog Dialog
	if configOS == "Windows" {
		dialog = &WindowsDialog{}
	} else if configOS == "Web" {
		dialog = &WebDialog{}
	} else {
		fmt.Println("Error! Unknown operating system.")
		return
	}

	dialog.Render()
}
