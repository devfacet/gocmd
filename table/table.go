/*
 * gocmd
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

// Package table provides functions for handling tables in terminal
package table

import (
	"errors"
	"fmt"
)

// Options represents the options that can be set when creating a new table
type Options struct {
	Data [][]string
}

// New returns a table by the given options
func New(o Options) (*Table, error) {
	// Init vars
	t := Table{
		data: o.Data,
	}
	return &t, nil
}

// Table represent a table
type Table struct {
	data     [][]string
	colSizes map[int]int
}

// Data returns the data of the table
func (t *Table) Data() [][]string {
	return t.data
}

// SetData sets the table data by the given row, column and value
func (t *Table) SetData(row, col int, val string) error {
	// Check row and column
	if row < 1 {
		return errors.New("invalid row")
	} else if col < 1 {
		return errors.New("invalid column")
	}

	// Increase the row capacity if it's necessary
	if row > len(t.data) {
		nt := make([][]string, row)
		copy(nt, t.data)
		t.data = nt
	}

	// Increase the column capacity if it's necessary
	if col > len(t.data[row-1]) {
		nr := make([]string, col)
		copy(nr, t.data[row-1])
		t.data[row-1] = nr
	}

	// Set the value
	t.data[row-1][col-1] = val

	// Set the column size for alignment
	if t.colSizes == nil {
		t.colSizes = make(map[int]int)
	}

	if len(val) > t.colSizes[col-1] {
		t.colSizes[col-1] = len(val)
	}

	return nil
}

// SetRow sets a row by the given row number and column values
func (t *Table) SetRow(row int, cols ...string) error {
	// Iterate columns and set data
	for i, v := range cols {
		if err := t.SetData(row, i+1, v); err != nil {
			return err
		}
	}
	return nil
}

// AddRow adds a row at the end of the table by the given column values
func (t *Table) AddRow(cols ...string) error {
	row := len(t.data) + 1
	// Iterate columns and set data
	for i, v := range cols {
		if err := t.SetData(row, i+1, v); err != nil {
			return err
		}
	}
	return nil
}

// FormattedData returns the formatted table data
func (t *Table) FormattedData() string {
	// Iterate over the rows and prepare result
	result := ""
	rowVal := ""
	colSize := ""
	for _, row := range t.data {
		rowVal = ""
		for i, c := range row {
			colSize = fmt.Sprintf("%d", t.colSizes[i])
			rowVal += fmt.Sprintf("%-"+colSize+"s\t", c)
		}
		result += fmt.Sprintf("%s\n", rowVal)
	}

	return result
}
