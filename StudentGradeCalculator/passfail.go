package main

func passFail(bangla float32, english float32, math float32, physics float32) string {
	if bangla < 40 || english < 40 || math < 40 || physics < 40 {
		return "Fail"
	} else {
		return "Pass"
	}
}
