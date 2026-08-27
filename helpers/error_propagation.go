package helpers

func ChckErr(e error) {
	if e != nil {
		panic(e)
	}
}
