package buffer

type View []byte

func NewView(size int) View {
	return make(View, size)
}
