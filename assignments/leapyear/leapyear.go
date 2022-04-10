package leapyear

//LeapYear is used to test whether given year is leap year or not
func LeapYear(year int) string {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return "LeapYear"
	}
	return "NotaLeapYear"
}
