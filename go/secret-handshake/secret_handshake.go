package secret

const testVersion = 1

func Handshake(code uint) []string {
	codeParts := []string{}

	for code != 0 {
		if code&1 != 0 {
			codeParts = append(codeParts, "wink")
		}

		if code&2 != 0 {
			codeParts = append(codeParts, "double blink")
		}

		if code&4 != 0 {
			codeParts = append(codeParts, "close your eyes")
		}

		if code&8 != 0 {
			codeParts = append(codeParts, "jump")
		}

		if code&16 != 0 {
			oldCodeParts := codeParts
			codeParts = []string{}
			for i := len(oldCodeParts) - 1; i >= 0; i-- {
				codeParts = append(codeParts, oldCodeParts[i])
			}

		}
		code = code >> 6
	}
	return codeParts
}
