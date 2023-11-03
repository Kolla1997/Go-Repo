package sMaps

import (
	"fmt"
)

func MapFunc() {
	fmt.Println(" \n Maps\n")
	grades := map[string]int{
		"Math":   98,
		"Sci":    99,
		"Social": 100,
	}
	fmt.Println(grades) // To print entire Map
	grades["Phy"] = 95  //To Add to map
	grades["Che"] = 91
	for key, val := range grades {
		fmt.Println(key, val) // To print key and value in Map
	}
	fmt.Println()
	delete(grades, "Che") //To delete from map

	fmt.Println("After Deleting from MAP\n")

	for key, val := range grades {
		fmt.Println(key, val) // To print key and value in Map
	}

	// IF - ELSE Loop
	if val, ok := grades["Che"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("NOT PRESENT")
	}

	SliceFunc()
	fmt.Println()
	fmt.Println()
}

func SliceFunc() {
	fmt.Println(" \n Slice \n")
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = arr1[0:]
	fmt.Println(s)

	s = append(s[:2], s[3:]...) // Delete from slice
	fmt.Println("After deleting : ", s)

	s = s[1:]
	fmt.Println(s)

	names := make([]string, 0)
	fmt.Println(names)
	names = append(names, "Hey", "Hi", "Hello")
	fmt.Println(names)
	names = names[1:]
	fmt.Println(names)

}
