package helpers

func PanicChckErr(e error) {
	if e != nil {
		panic(e)
	}
}
