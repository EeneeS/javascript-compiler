package parser

type SymbolType string

const (
	Int    SymbolType = "Integer"
	String SymbolType = "String"
	Float  SymbolType = "Float"
	Object SymbolType = "Object"
)

type Symbol struct {
	Name  string
	Type  SymbolType
	Value interface{}
}

type SymbolTable struct {
	table  map[string]Symbol
	parent *SymbolTable
}

func NewSymbolTable(st *SymbolTable) *SymbolTable {
	return &SymbolTable{table: make(map[string]Symbol), parent: st}
}
