package codegen

import "strconv"

// return an int in binary format
func decToBin(val string) string {
	number, err := strconv.Atoi(val)
	if err != nil {
		panic("strconv.Atoi err: %s " + err.Error() + "\n")
	}
	return strconv.FormatInt(int64(number), 2)
}
