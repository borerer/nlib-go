package utils

func LogError(err error) {
	if err != nil {
		println(err.Error())
	}
}
