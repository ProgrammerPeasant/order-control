package utils

import (
	"fmt"
	"log"
	_ "log"
	"os"
	_ "os"
	"time"

	"github.com/xuri/excelize/v2"
)

// CreateExcelReport — пример генерации Excel-файла
func CreateExcelReport(data [][]string) (string, error) {
	// Создаём новую книгу Excel
	f := excelize.NewFile()
	sheetName := "Sheet1"
	f.NewSheet(sheetName)

	// Заполняем ячейки
	for rowIndex, row := range data {
		for colIndex, cellValue := range row {
			cellName, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
			if err := f.SetCellValue(sheetName, cellName, cellValue); err != nil {
				return "", err
			}
		}
	}

	// Генерируем уникальное имя файла
	fileName := fmt.Sprintf("report_%d.xlsx", time.Now().UnixNano())
	// Сохраняем файл во временную директорию
	if err := f.SaveAs("/tmp/" + fileName); err != nil {
		return "", err
	}

	return fileName, nil
}

// DeleteFile — удаляем файл после скачивания/использования
func DeleteFile(filePath string) {
	if err := os.Remove(filePath); err != nil {
		log.Printf("Ошибка при удалении файла %s: %v", filePath, err)
	}
}
