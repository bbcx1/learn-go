package models

type BillItem struct {
	Name  string
	Price float64
	Qty   uint
}

type Bill struct {
	Items  []BillItem
	People uint
}
