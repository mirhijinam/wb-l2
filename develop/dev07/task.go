package main

import (
	"sync"
)

/* Or-channel.
 *
 * Реализовать функцию, которая будет объединять один или более done-каналов
 * в single-канал, если один из его составляющих каналов закроется.
 * Одним из вариантов было бы очевидно написать выражение при помощи select,
 * которое бы реализовывало эту связь, однако иногда неизестно общее число
 * done-каналов, с которыми вы работаете в рантайме.
 * В этом случае удобнее использовать вызов единственной функции, которая,
 * приняв на вход один или более or-каналов, реализует весь функционал.
 *
 * Определение функции:
 * var or func(channels ...<-chan interface{}) <-chan interface{}
 *
 * Пример использования функции:
 * sig := func(after time.Duration) <- chan interface{} {
 * 	c := make(chan interface{})
 * 	go func() {
 * 		defer close(c)
 * 		time.Sleep(after)
 * }()
 * return c
 * }
 *
 * start := time.Now()
 * <-or (
 * 	sig(2*time.Hour),
 * 	sig(5*time.Minute),
 * 	sig(1*time.Second),
 * 	sig(1*time.Hour),
 * 	sig(1*time.Minute),
 * )
 *
 * fmt.Printf(“fone after %v”, time.Since(start))
 */

func or(channels ...<-chan interface{}) <-chan interface{} {
	orDone := make(chan interface{})
	defer close(orDone)

	var wg sync.WaitGroup
	wg.Add(1)

	for _, c := range channels {
		go func(channel <-chan interface{}) {
			defer wg.Done()

			for v := range channel {
				orDone <- v
			}
		}(c)
	}

	wg.Wait()

	return orDone
}
