package types

type BotCommand struct {
	Command string
	Value   string
	Params  map[string]string
}
