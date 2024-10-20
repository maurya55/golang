package main

func main() {

	m := make(map[string]string)

	m["name"] = "wap"
	m["dle"] = "ins"

	delete(m, "dle")

	println(m["nam"])

	// val, ok := m["nae"]

	// println(val, ok)

}
