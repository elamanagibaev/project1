package main

import (
	"bufio"
	"fmt"
	"os"
	"projectOne/logics"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("❌ Использование: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("❌ Ошибка при открытии входного файла:", err)
		return
	}
	defer in.Close()

	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("❌ Ошибка при создании выходного файла:", err)
		return
	}
	defer out.Close()

	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	writer.Flush()

	fmt.Println("✅ Файл успешно обработан.")

	for scanner.Scan() {
		line := scanner.Text()
		line = logics.ProcessUp(line)
		line = logics.ProcessLow(line)
		line = logics.ProcessCap(line)
		line = logics.ProcessHex(line)
		line = logics.ProcessBin(line)
		line = logics.ProcessUpN(line)
		line = logics.ProcessLowN(line)
		line = logics.ProcessCapN(line)

		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("❌ Ошибка записи:", err)
			return
		}
	}

	writer.Flush()
}
