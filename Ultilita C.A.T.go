package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при открытии файла %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при чтении файла %s: %v\n", filename, err)
	}
}

func main() {

	var help bool
	flag.BoolVar(&help, "h", false, "Показать справку")
	flag.Parse()

	if help {
		fmt.Println("Использование: cat [опции] [файлы]")
		fmt.Println("Опции:")
		fmt.Println("  -h    Показать эту справку")
		fmt.Println("\nВывод содержимого файлов или объединение нескольких файлов.")
		return
	}

	files := flag.Args()

	if len(files) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при чтении из stdin: %v\n", err)
		}
		return
	}

	for _, file := range files {
		readFile(file)
	}
}
