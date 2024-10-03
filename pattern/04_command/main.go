package main

import "fmt"

/*
 * Реализовать паттерн «КОМАНДА».
 *
 * Поведенческий паттерн.
 *
 * Объяснить применимость паттерна, его плюсы и минусы,
 * а также реальные примеры использования данного примера на практике.
 *
 * https://en.wikipedia.org/wiki/Command_pattern
 *
 * Превращает запросы в объекты.
 *
 * Используется при:
 * а) необходимости в операции отмены,
 * б) необходимости в параметризации объектов выполняемым действием (операцией),
 * в) необходимости в очередях, выполнении по расписанию, отмене.
 *
 * + отсутствие прямой зависимости между объектами между теми, кто вызывает
 *		метод, и теми, кто его выполняет
 * + принцип открытости/закрытости
 * + возможность собирать сложные команды из простых
 * - дополнительный слой абстракции
 */

type Command interface {
	Execute() bool
	Undo()
}

type BaseCommand struct {
	app    *Application
	editor *Editor
	backup string
}

func (c *BaseCommand) SaveBackup() {
	c.backup = c.editor.text
}

func (c *BaseCommand) Undo() {
	c.editor.text = c.backup
}

type CopyCommand struct {
	*BaseCommand
}

func NewCopyCommand(app *Application, editor *Editor) *CopyCommand {
	return &CopyCommand{
		&BaseCommand{app: app, editor: editor},
	}
}

func (c *CopyCommand) Execute() bool {
	c.app.clipboard = c.editor.GetText()
	return false
}

type PasteCommand struct {
	*BaseCommand
}

func NewPasteCommand(app *Application, editor *Editor) *PasteCommand {
	return &PasteCommand{
		&BaseCommand{app: app, editor: editor},
	}
}

func (c *PasteCommand) Execute() bool {
	c.SaveBackup()
	c.editor.ReplaceText(c.app.clipboard)
	return true
}

type CommandHistory struct {
	history []Command
}

func (h *CommandHistory) Push(c Command) {
	h.history = append(h.history, c)
}

func (h *CommandHistory) Pop() Command {
	if len(h.history) == 0 {
		return nil
	}
	last := h.history[len(h.history)-1]
	h.history = h.history[:len(h.history)-1]
	return last
}

type Editor struct {
	text string
}

func (e *Editor) GetText() string {
	return e.text
}

func (e *Editor) DeleteText() {
	e.text = ""
}

func (e *Editor) ReplaceText(text string) {
	e.text = text
}

type Application struct {
	clipboard    string
	editors      []*Editor
	activeEditor *Editor
	history      *CommandHistory
}

func NewApplication() *Application {
	return &Application{
		history: &CommandHistory{},
	}
}

func (a *Application) ExecuteCommand(command Command) {
	if command.Execute() {
		a.history.Push(command)
	}
}

func (a *Application) Undo() {
	command := a.history.Pop()
	if command != nil {
		command.Undo()
	}
}

func main() {
	app := NewApplication()
	editor := &Editor{text: "Hello, World!"}
	app.activeEditor = editor

	copyCmd := NewCopyCommand(app, editor)
	pasteCmd := NewPasteCommand(app, editor)

	fmt.Println("Initial text:", editor.text)

	app.ExecuteCommand(copyCmd)
	fmt.Println("Clipboard after copy:", app.clipboard)

	app.ExecuteCommand(pasteCmd)
	fmt.Println("Text after paste:", editor.text)
}
