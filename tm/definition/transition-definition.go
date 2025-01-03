package definition

type TransitionFunction struct {
	currentState     string
	readCharacter    string
	writeCharacter   string
	destinationState string
	move             MoveDirection
}

type MoveDirection int

const (
	LEFT MoveDirection = iota
	RIGHT
)

func (t TransitionFunction) Print() {
	print("\u03B4( ")
	print(t.currentState)
	print(", ")
	print(t.readCharacter)
	print(" ) = { ")
	print(t.writeCharacter)
	print(", ")
	print(t.destinationState)
	print(", ")
	if t.move == LEFT {
		print("L")
	} else {
		print("R")
	}
	print(" }")
}
