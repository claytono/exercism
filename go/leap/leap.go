package leap

const testVersion = 2

// It's good style to write a comment here documenting IsLeapYear.
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
        switch {
        case year % 400 == 0:
                return true
        case year % 100 == 0:
                return false
        case year % 4 == 0:
                return true
        }
        return false
}
