package main

import "fmt"

/*
 * Реализовать паттерн «ПОСЕТИТЕЛЬ».
 *
 * Поведенческий паттерн.
 *
 * Объяснить применимость паттерна, его плюсы и минусы,
 * а также реальные примеры использования данного примера на практике.
 *
 * https://en.wikipedia.org/wiki/Visitor_pattern
 *
 * Выполняет какую-то операцию над всеми элементами сложной структуры 
 * объектов разных классов.
 *
 * Используется при:
 * a) невозможности добавить в классы новый функционал,
 * б) ситуации, в которой функционал необходим не всем объектам в структуре.
 *
 * + упрощенное добавление функционала
 * + объединение родственных операций в одном классе
 * + возможность накапливать состояние при обходе
 * - неоправданность при частой схеме иерархий
 * - может привести к нарушению инкапсуляции
 */

type Shape interface {
	Move(int, int)
	Draw()
	Accept(Visitor)
}

type Dot struct {
	id   int
	x, y int
}

func (d *Dot) Move(x, y int) {
	d.x += x
	d.y += y
}

func (d *Dot) Draw() {
	fmt.Printf("Drawing Dot at (%d, %d)\n", d.x, d.y)
}

func (d *Dot) Accept(v Visitor) {
	v.VisitDot(d)
}

type Circle struct {
	id int
	x  int
	y  int
	r  int
}

func (c *Circle) Move(x, y int) {
	c.x += x
	c.y += y
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing Circle at (%d, %d) with radius %d\n", c.x, c.y, c.r)
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

type Rectangle struct {
	ID     int
	left   int
	top    int
	width  int
	height int
}

func (r *Rectangle) Move(x, y int) {
	r.left += x
	r.top += y
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing Rectangle at (%d, %d) with width %d and height %d\n", r.left, r.top, r.width, r.height)
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}

type Visitor interface {
	VisitDot(d *Dot)
	VisitCircle(c *Circle)
	VisitRectangle(r *Rectangle)
}

type XMLExportVisitor struct{}

func (x *XMLExportVisitor) VisitDot(d *Dot) {
	fmt.Printf("<Dot id='%d' x='%d' y='%d' />\n", d.id, d.x, d.y)
}

func (x *XMLExportVisitor) VisitCircle(c *Circle) {
	fmt.Printf("<Circle id='%d' x='%d' y='%d' radius='%d' />\n", c.id, c.x, c.y, c.r)
}

func (x *XMLExportVisitor) VisitRectangle(r *Rectangle) {
	fmt.Printf("<Rectangle id='%d' left='%d' top='%d' width='%d' height='%d' />\n", r.ID, r.left, r.top, r.width, r.height)
}

type Application struct {
	AllShapes []Shape
}

func (app *Application) Export() {
}

func main() {
	shapes := []Shape{
		&Dot{id: 1, x: 10, y: 20},
		&Circle{id: 2, x: 15, y: 25, r: 5},
		&Rectangle{ID: 3, left: 30, top: 40, width: 10, height: 20},
	}

	exportVisitor := &XMLExportVisitor{}
	for _, shape := range shapes {
		shape.Accept(exportVisitor)
	}
}
