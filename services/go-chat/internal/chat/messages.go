package chat

import "math/rand"

var Messages = []string{
	"Hello!",
	"Hey there 👋",
	"Ping from Go",
	"How are you?",
	"Go says hi 🚀",
}

func PickMessage() string {
	return Messages[rand.Intn(len(Messages))]
}
