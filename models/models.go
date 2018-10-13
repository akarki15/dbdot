package models

type ForeignKey struct {
	From, To string
}

type Column struct {
	Name     string
	DataType string
}

type TableAndColumns struct {
	Name    string
	Columns []Column
}
