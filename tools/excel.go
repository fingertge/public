// ***********************************************************************************************
// ***                                     G O L A N D                                         ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2023/10/19 13:18:32
// * Proj: work
// * Pack: tools
// * File: excel.go
// *----------------------------------------------------------------------------------------------
// * Overviews:
// *----------------------------------------------------------------------------------------------
// * Functions:
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
package tools

import (
	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
)

type Excel struct {
	f *excelize.File
}

func (e *Excel) Open(file string) error {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return err
	}
	e.f = f
	return nil
}

func (e *Excel) Close() {
	e.f.Close()
	e.f = nil
}

func (e *Excel) GetCellValue(sheetName, pos string) (string, error) {
	return e.f.GetCellValue(sheetName, pos)
}

// 按列获取
func (e *Excel) GetColsValue(sheetName, startPos, endPos string) ([]string, error) {
	cols, err := e.f.GetCols(sheetName)
	if err != nil {
		return []string{}, err
	}
	colStart, rowStart, err := excelize.CellNameToCoordinates(startPos)
	if err != nil {
		return []string{}, err
	}
	colEnd, rowEnd, err := excelize.CellNameToCoordinates(endPos)
	if err != nil {
		return []string{}, err
	}
	if len(cols) < colStart-1 {
		return []string{}, errors.Errorf("start col position out of range")
	}
	if colStart != colEnd {
		return []string{}, errors.New("startPos and endPos is not same col")
	}
	if rowEnd < rowStart {
		return []string{}, errors.Errorf("end row position must bigger than start row position")
	}

	rows := cols[colStart-1]
	if len(rows) < rowEnd {
		return []string{}, errors.Errorf("end row position out of range")
	}
	colTmp := make([]string, rowEnd-rowStart+1)
	copy(colTmp, rows[rowStart-1:rowEnd])
	return colTmp, nil
}

// 按行获取
func (e *Excel) GetRowsValue(sheetName, startPos, endPos string) ([]string, error) {
	rows, err := e.f.GetRows(sheetName)
	if err != nil {
		return []string{}, err
	}
	colStart, rowStart, err := excelize.CellNameToCoordinates(startPos)
	if err != nil {
		return []string{}, err
	}
	colEnd, rowEnd, err := excelize.CellNameToCoordinates(endPos)
	if err != nil {
		return []string{}, err
	}
	if len(rows) < rowStart-1 {
		return []string{}, errors.Errorf("start row position out of range")
	}
	if rowStart != rowEnd {
		return []string{}, errors.New("startPos and endPos is not same row")
	}
	if colEnd < colStart {
		return []string{}, errors.Errorf("end col position must bigger than start col position")
	}
	cols := rows[rowStart-1]
	if len(cols) < colEnd {
		return []string{}, errors.Errorf("end col position out of range")
	}
	rowTmp := make([]string, colEnd-colStart+1)
	copy(rowTmp, cols[colStart-1:colEnd])
	return rowTmp, nil
}

// 列做一维
func (e *Excel) GetColMatrix(sheetName, startPos, endPos string) ([][]string, error) {
	cols, err := e.f.GetCols(sheetName)
	if err != nil {
		return [][]string{}, err
	}
	colStart, rowStart, err := excelize.CellNameToCoordinates(startPos)
	if err != nil {
		return [][]string{}, err
	}
	colEnd, rowEnd, err := excelize.CellNameToCoordinates(endPos)
	if err != nil {
		return [][]string{}, err
	}
	if colEnd < colStart {
		return [][]string{}, errors.New("end col position must bigger than start col position")
	}
	if len(cols) < colEnd {
		return [][]string{}, errors.Errorf("end col position out of range")
	}
	if rowEnd < rowStart {
		return [][]string{}, errors.Errorf("end row position must bigger than start row position")
	}
	if len(cols[colStart-1]) < rowEnd {
		return [][]string{}, errors.Errorf("end row position out of range")
	}
	var matrix [][]string
	for i := colStart - 1; i < colEnd; i++ {
		row := make([]string, rowEnd-rowStart+1)
		copy(row, cols[i][rowStart-1:rowEnd])
		matrix = append(matrix, row)
	}
	return matrix, nil
}

// 行做一维
func (e *Excel) GetRowMatrix(sheetName, startPos, endPos string) ([][]string, error) {
	rows, err := e.f.GetRows(sheetName)
	if err != nil {
		return [][]string{}, err
	}
	colStart, rowStart, err := excelize.CellNameToCoordinates(startPos)
	if err != nil {
		return [][]string{}, err
	}
	colEnd, rowEnd, err := excelize.CellNameToCoordinates(endPos)
	if err != nil {
		return [][]string{}, err
	}
	if rowEnd < rowStart {
		return [][]string{}, errors.Errorf("end row position must bigger than start row position")
	}
	if len(rows) < rowEnd {
		return [][]string{}, errors.Errorf("end row position out of range")
	}
	if colEnd < colStart {
		return [][]string{}, errors.New("end col position must bigger than start col position")
	}
	if len(rows[rowStart-1]) < colEnd {
		return [][]string{}, errors.Errorf("end col position out of range")
	}
	var matrix [][]string
	for i := rowStart - 1; i < rowEnd; i++ {
		col := make([]string, colEnd-colStart+1)
		copy(col, rows[i][colStart-1:colEnd])
		matrix = append(matrix, col)
	}
	return matrix, nil
}
