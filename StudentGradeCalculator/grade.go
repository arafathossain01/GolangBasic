package main

func grade(avg float32) string {
	if avg > 100 || avg < 0 {
		return "Input Invalid."
	} else if avg >= 80 {
		return "You got A+"
	} else if avg >= 70 {
		return "You got A"
	} else if avg >= 60 {
		return "You got A-"
	} else if avg >= 50 {
		return "You got B"
	} else if avg >= 40 {
		return "You got C"
	} else if avg >= 33 {
		return "You got D"
	} else {
		return "You got F"
	}
}
