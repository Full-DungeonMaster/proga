package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Сортировка пузырьком
func bubbleSort(numb []int) {
	n := len(numb)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numb[j] > numb[j+1] {
				numb[j], numb[j+1] = numb[j+1], numb[j]
			}
		}
	}
}

// Чтение из файла
func readInput(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	var numbers []int
	for _, s := range strings.Fields(line) {
		var num int
		_, err := fmt.Sscanf(s, "%d", &num)
		if err != nil {
			return nil, fmt.Errorf("Неверный формат числа: %v", err)
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

// Запись в файл
func writeOutput(filename string, numbers []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%v", numbers)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	numbers, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	bubbleSort(numbers)

	err = writeOutput("output.txt", numbers)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Println("Результат записан в output.txt")
}
