package main

import "fmt"

/*
 * Реализовать паттерн «СТРОИТЕЛЬ».
 *
 * Порождающий паттерн.
 *
 * Объяснить применимость паттерна, его плюсы и минусы,
 * а также реальные примеры использования данного примера на практике.
 *
 * https://en.wikipedia.org/wiki/Builder_pattern
 *
 * Выносит конструирование объекта за пределы класса, поручив это другим
 * объектам, называемым строителями.
 *
 * Используется при:
 * a)  избавлении от "телескопического" конструктора с множеством
 * 		опциональных параметров,
 * б) создании нескольких представителей объектов одинаковых по структуре,
 * 		но с различными деталями. Например, спорт-кар и внедорожник,
 * в) создании сложных составных объектов
 *
 * + пошаговое создание продукта
 * + повторное использование кода
 * + изоляция сложностей сборки
 * - усложнение введением абстракций
 * - привязка к определенным строителям 
 * 	(как в нашем случае, клиент должен знать результат какого билдера
 * 		ему необходим, и Директор не может возвращать объект напрямую)
 */

type Car struct {
	seats        int
	engine       string
	tripComputer bool
	gps          bool
}

type Manual struct {
	seats        int
	engine       string
	tripComputer bool
	gps          bool
}

type Builder interface {
	Reset()
	SetSeats(int)
	SetEngine(string)
	SetTripComputer(bool)
	SetGPS(bool)
}

type CarBuilder struct {
	car *Car
}

func (b *CarBuilder) Reset() {
	b.car = &Car{}
}

func (b *CarBuilder) SetSeats(seats int) {
	b.car.seats = seats
}

func (b *CarBuilder) SetEngine(engine string) {
	b.car.engine = engine
}

func (b *CarBuilder) SetTripComputer(hasTripComputer bool) {
	b.car.tripComputer = hasTripComputer
}

func (b *CarBuilder) SetGPS(hasGPS bool) {
	b.car.gps = hasGPS
}

func (b *CarBuilder) GetResult() *Car {
	return b.car
}

type CarManualBuilder struct {
	manual *Manual
}

func (b *CarManualBuilder) Reset() {
	b.manual = &Manual{}
}

func (b *CarManualBuilder) SetSeats(seats int) {
	b.manual.seats = seats
}

func (b *CarManualBuilder) SetEngine(engine string) {
	b.manual.engine = engine
}

func (b *CarManualBuilder) SetTripComputer(hasTripComputer bool) {
	b.manual.tripComputer = hasTripComputer
}

func (b *CarManualBuilder) SetGPS(hasGPS bool) {
	b.manual.gps = hasGPS
}

func (b *CarManualBuilder) GetResult() *Manual {
	return b.manual
}

type Director struct{}

func (d *Director) ConstructSportsCar(builder Builder) {
	builder.Reset()
	builder.SetSeats(2)
	builder.SetEngine("SportEngine")
	builder.SetTripComputer(true)
	builder.SetGPS(true)
}

func main() {
	// В данном паттерне клиентом может использоваться сущность Директор, 
	// который определяет порядок вызова конкретных шагов по построению
	// необходимого объекта, тем самым абстрагируя клиента от конструирования
	//
	// Наша реализация Директора не привязывается к конкретному типу
	// создаваемого объекта. Все реализуется через метод
	director := &Director{}

	// Можно заметить, что Директор не знает, какой именно объект он создает

	// В данном блоке создается сама машина
	carBuilder := &CarBuilder{}
	director.ConstructSportsCar(carBuilder)
	car := carBuilder.GetResult()
	fmt.Printf("Car: %+v\n", car)

	// В данном блоке создается мануал по эксплуатации машины
	manualBuilder := &CarManualBuilder{}
	director.ConstructSportsCar(manualBuilder)
	manual := manualBuilder.GetResult()
	fmt.Printf("Manual: %+v\n", manual)
}
