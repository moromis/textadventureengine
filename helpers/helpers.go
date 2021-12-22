package helpers

import (
	"bytes"
	"math/rand"
	"reflect"
	"strings"
)

func IndexInto2dArray(arr []int, col int, row int, width int) int {
	return arr[col+row*width]
}

func PickStringRandomly(arr []string) string {
	randomIndex := rand.Intn(len(arr))
	return arr[randomIndex]
}

// FROM: https://www.jeremymorgan.com/tutorials/go/learn-golang-casing/
func TitleCase(input string) string {
	input = strings.ReplaceAll(input, "_", " ")
	return strings.Title(strings.ToLower(input))
}

func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

func ListContains(list interface{}, element interface{}) (ok, found bool) {

	listValue := reflect.ValueOf(list)
	listKind := reflect.TypeOf(list).Kind()
	defer func() {
		if e := recover(); e != nil {
			ok = false
			found = false
		}
	}()

	if listKind == reflect.String {
		elementValue := reflect.ValueOf(element)
		return true, strings.Contains(listValue.String(), elementValue.String())
	}

	if listKind == reflect.Map {
		mapKeys := listValue.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if ObjectsAreEqual(mapKeys[i].Interface(), element) {
				return true, true
			}
		}
		return true, false
	}

	for i := 0; i < listValue.Len(); i++ {
		if ObjectsAreEqual(listValue.Index(i).Interface(), element) {
			return true, true
		}
	}
	return true, false

}
