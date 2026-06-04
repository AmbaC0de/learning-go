package main

func reformat(message string, formatter func(string) string) string {
	firstResult := formatter(message)
	secondResult := formatter(firstResult)
	thirdResult := formatter(secondResult)

	return "TEXTIO: " + thirdResult
}
