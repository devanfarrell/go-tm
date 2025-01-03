package definition

type Definition struct {
	description         string
	states              []string
	inputAlphabet       []string
	tapeAlphabet        []string
	transitionFunctions []TransitionFunction
	initialState        string
	blackCharacter      string
	finalStates         []string
}

func (def Definition) Print() {
	print(def.description)
	def.printStates()
	def.printInputAlphabet()
	def.printTapeAlphabet()
}

func (d Definition) printStates() {
	print("Q\t = { ")
	for i, state := range d.states {
		if i != 0 {
			print(", ")
		}
		print(state)

	}
	println(" }")
}

func (d Definition) printInputAlphabet() {
	print("\u03A3\t = { ")
	for i, letter := range d.inputAlphabet {
		if i != 0 {
			print(", ")
		}
		print(letter)
	}
	println(" }")
}

func (d Definition) printTapeAlphabet() {
	print("\u0393\t = { ")
	for i, letter := range d.tapeAlphabet {
		if i != 0 {
			print(", ")
		}
		print(letter)
	}
	println(" }")
}
