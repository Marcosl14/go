package main

func main() {

}

func add(a, b interface{}) interface{} { // the same as any
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		return aString + bString
	}

	return nil
}

func addWithGenerics[T int | float64 | string](a, b T) T {
	return a + b
}

func convert[T float32 | float64, K int | string](input T) K {
	var output K
	return output
}

// go run app.go
