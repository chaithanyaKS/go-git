package blob

type Blob struct {
	Oid  string
	Data []byte
}

func New(data []byte) *Blob {
	return &Blob{Data: data}
}

func (b *Blob) AssignOid(oid string) {
	b.Oid = oid
}
func (b *Blob) GetData() (string, error) {
	return string(b.Data), nil
}
