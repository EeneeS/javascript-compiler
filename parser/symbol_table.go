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
	return &SymbolTable{
		table:  make(map[string]Symbol),
		parent: st,
	}
}

func (st *SymbolTable) Define(name string, t SymbolType, value interface{}) {
}

func (st *SymbolTable) Resolve(name string) (*Symbol, bool) {
	if symbol, exists := st.table[name]; exists {
		return &symbol, true
	}
	if st.parent != nil {
		return st.parent.Resolve(name)
	}
	return nil, false
}
