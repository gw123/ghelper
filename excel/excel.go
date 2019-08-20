package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
)

//xlsx文件解析
func ExcelParse(fileName string) error {
	filePath := "./upload/" + fileName
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return err
	}

	//获取行数
	//length := len(xlFile.Sheets[0].Rows)
	//开辟除表头外的行数的数组内存
	//resourceArr := make([]string, length-1)
	//遍历sheet
	for _, sheet := range xlFile.Sheets {
		//遍历每一行
		for _, row := range sheet.Rows {
			//跳过第一行表头信息
			//if rowIndex == 0 {
			//	for _, cell := range row.Cells {
			//		text := cell.String()
			//		fmt.Printf("%s\n", text)
			//	}
			//	continue
			//}
			//遍历每一个单元
			//fmt.Println("\t", rowIndex)
			for _, cell := range row.Cells {
				text := cell.String()
				if text != "" {
					fmt.Printf("%s \t", text)
				}
			}
			fmt.Println()
		}
	}
	return nil
}

func main() {
	err := ExcelParse("shou.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
