package messages

type Message interface {
	fun(s string) string
}

type Package struct {
	m1 Message
	m2 Message
}
type RealMessage struct {
}

func (rm RealMessage) fun(s string) string {
	return s + " ..."
}

type MessageExtractor struct {
	RealMessage
}
