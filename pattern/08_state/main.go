package main

import "fmt"

/*
 * Реализовать паттерн «СОСТОЯНИЕ».
 *
 * Поведенческий паттерн.
 *
 * Объяснить применимость паттерна, его плюсы и минусы, 
 * а также реальные примеры использования данного примера на практике.
 *
 * https://en.wikipedia.org/wiki/State_pattern
 *
 * Позволяет объектам менять поведение в зависимости от состояния.
 *
 * Используется при:
 * а) необходимости кардинально менять поведение объекта,
 * б) множестве больших условных операторов внутри класса.
 *
 * + концентрация кода, связанного с определенным состоянием
 * + упрощение контекста
 * - усложнение кода при небольшом количестве редко изменяемых состояний
 */

type State interface {
	clickLock()
	clickPlay()
	clickNext()
	clickPrevious()
}

type AudioPlayer struct {
	state State
}

func NewAudioPlayer() *AudioPlayer {
	player := &AudioPlayer{}
	player.changeState(&ReadyState{player: player})
	return player
}

func (p *AudioPlayer) changeState(state State) {
	p.state = state
}

func (p *AudioPlayer) clickLock() {
	p.state.clickLock()
}

func (p *AudioPlayer) clickPlay() {
	p.state.clickPlay()
}

func (p *AudioPlayer) clickNext() {
	p.state.clickNext()
}

func (p *AudioPlayer) clickPrevious() {
	p.state.clickPrevious()
}

type ReadyState struct {
	player *AudioPlayer
}

func (s *ReadyState) clickLock() {
	fmt.Println("Locking player.")
	s.player.changeState(&LockedState{player: s.player})
}

func (s *ReadyState) clickPlay() {
	fmt.Println("Playing music.")
	s.player.changeState(&PlayingState{player: s.player})
}

func (s *ReadyState) clickNext() {
	fmt.Println("Skipping to next song.")
}

func (s *ReadyState) clickPrevious() {
	fmt.Println("Skipping to previous song.")
}

type PlayingState struct {
	player *AudioPlayer
}

func (s *PlayingState) clickLock() {
	fmt.Println("Locking player.")
	s.player.changeState(&LockedState{player: s.player})
}

func (s *PlayingState) clickPlay() {
	fmt.Println("Stopping playback.")
	s.player.changeState(&ReadyState{player: s.player})
}

func (s *PlayingState) clickNext() {
	fmt.Println("Fast forwarding.")
}

func (s *PlayingState) clickPrevious() {
	fmt.Println("Rewinding.")
}

type LockedState struct {
	player *AudioPlayer
}

func (s *LockedState) clickLock() {
	fmt.Println("Unlocking player.")
	s.player.changeState(&ReadyState{player: s.player})
}

func (s *LockedState) clickPlay() {
	fmt.Println("Player is locked.")
}

func (s *LockedState) clickNext() {
	fmt.Println("Player is locked.")
}

func (s *LockedState) clickPrevious() {
	fmt.Println("Player is locked.")
}

func main() {
	player := NewAudioPlayer()

	player.clickPlay()
	player.clickNext()
	player.clickLock()
	player.clickPlay()
	player.clickLock()
	player.clickPrevious()
}
