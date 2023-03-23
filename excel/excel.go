package excel

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type WorkOrder struct {
	name     string
	contents string
}

func CreateFile(workOrders map[string]map[string]string) {
	now := time.Now().Unix()
	filename := fmt.Sprintf("work_orders_%v.xlsx", now)

	file := excelize.NewFile()

	for workOrderSetName, workOrderSet := range workOrders {
		sheetNum := file.NewSheet(workOrderSetName)

		workOrderSetSlice := []WorkOrder{}
		for name, contents := range workOrderSet {
			workOrderSetSlice = append(workOrderSetSlice, WorkOrder{
				name:     name,
				contents: contents,
			})
		}

		sort.Slice(workOrderSetSlice, func(a, b int) bool {
			return workOrderSetSlice[a].name < workOrderSetSlice[b].name
		})

		row := 1
		for workOrderName, contents := range workOrderSet {
			a := fmt.Sprintf("A%v", row)
			b := fmt.Sprintf("B%v", row)
			file.SetCellValue(workOrderSetName, a, workOrderName)
			file.SetCellValue(workOrderSetName, b, contents)
			row++
		}

		file.SetActiveSheet(sheetNum)
	}

	err := file.SaveAs(filename)
	if err != nil {
		log.Fatal(err)
	}
}
