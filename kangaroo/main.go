package main

// ! CHALLENGE URL: https://www.hackerrank.com/challenges/kangaroo

func main() {
	result := kangaroo(0, 3, 4, 2)
	println(result)
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v1 <= v2 {
		return "NO"
	}
	if (x2-x1)%(v1-v2) == 0 {
		return "YES"
	}
	return "NO"
}
