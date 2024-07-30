package entry

type Entry struct {
	Name string
	Oid  string
}

func New(name string, oid string) Entry {
	return Entry{Name: name, Oid: oid}
}

type Entries []Entry

func (e Entries) Len() int {
	return len(e)
}
