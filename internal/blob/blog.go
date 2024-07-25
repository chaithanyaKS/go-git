package blob

type Blob struct {
	Data []byte
}

func New(data []byte) Blob {
	return Blob{Data: data}
}
