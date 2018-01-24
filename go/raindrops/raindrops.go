package raindrops

import "fmt"

func Convert(input int) string {

	var response string = ""

	if input % 3 == 0 {
		response += "Pling"
	}
	if (input % 5 == 0) {
		response += "Plang"
	}
	if (input % 7 == 0) {
		response += "Plong"
	}

	if len(response) == 0 {
		// None of the above conditions were met
		return fmt.Sprintf("%v", input)
	}

	return response
}