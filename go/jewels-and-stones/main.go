package main

func main(jewels string, stones string) int {
	var jewelsCounter int

	for i, iv := range jewels {
		for j, jv := range stones {
			if iv == jv {
				jewelsCounter += 1
			}
		}
	}

	return jewelsCounter
}
