package main

import "fmt"

/*
 * Реализовать паттерн «ФАСАД».
 *
 * Структурный паттерн.
 *
 * Объяснить применимость паттерна, его плюсы и минусы,
 * а также реальные примеры использования данного примера на практике.
 *
 * https://en.wikipedia.org/wiki/Facade_pattern
 *
 * Разделяет сложную подсистему зависимостей и библиотек.
 *
 * Используется при:
 * а) задействовании только части возможностей подсистемы,
 * б) распределении ее возможности по отдельным слоям.
 *
 * + Изоляция клиента от сложностей.
 * - Возможность стать слишком толстым и привязанным ко всем компонентам.
 */

type VideoFile struct {
	filename string
}

func NewVideoFile(filename string) *VideoFile {
	return &VideoFile{filename: filename}
}

type OggCompressionCodec struct{}

type MPEG4CompressionCodec struct{}

type CodecFactory struct{}

func (c *CodecFactory) Extract(file *VideoFile) interface{} {
	fmt.Println("Extracting codec from file:", file.filename)
	return nil
}

type BitrateReader struct{}

func (br *BitrateReader) Read(filename string, codec interface{}) string {
	fmt.Println("Reading file:", filename, "with codec")
	return "buffer"
}

func (br *BitrateReader) Convert(buffer string, codec interface{}) string {
	fmt.Println("Converting buffer with codec")
	return "converted_file"
}

type AudioMixer struct{}

func (am *AudioMixer) Fix(file string) string {
	fmt.Println("Fixing audio in file:", file)
	return "fixed_" + file
}

// Фасад
type VideoConverter struct{}

func (vc *VideoConverter) Convert(filename, format string) string {
	file := NewVideoFile(filename)
	sourceCodec := new(CodecFactory).Extract(file)

	var destinationCodec interface{}
	if format == "mp4" {
		destinationCodec = new(MPEG4CompressionCodec)
	} else {
		destinationCodec = new(OggCompressionCodec)
	}

	buffer := new(BitrateReader).Read(filename, sourceCodec)
	result := new(BitrateReader).Convert(buffer, destinationCodec)
	result = new(AudioMixer).Fix(result)
	return result
}

// Приложение
func main() {
	converter := &VideoConverter{}
	mp4 := converter.Convert("funny-cats-video.ogg", "mp4")
	fmt.Println("Converted file saved:", mp4)
}
