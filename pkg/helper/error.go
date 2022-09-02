package helper

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
