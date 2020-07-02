package backend

type ProcessResult interface {
	OK() bool
	Errors() []string
	Warnings() []string
	Messages() []string
}

type Message interface {
}

type Processor interface {
	Process(m Message) ProcessResult
}
