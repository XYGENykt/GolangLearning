//как создать срез из массива

baseArray := [8]string{"Anna", "Max", "Eva", "Leo", "Nina", "Tom", "Sophie", "Chris"} // базовый массив

slice1 := baseArray[1:5] // со 2-го по 5-й элемент
slice2 := baseArray[:3]  // с 1-го по 3-й элемент
slice3 := baseArray[4:]  // с 5-го до конца

fmt.Println(slice1) // [Max Eva Leo Nina]
fmt.Println(slice2) // [Anna Max Eva]
fmt.Println(slice3) // [Nina Tom Sophie Chris]

//функция append

numbers := []int{10, 20, 30}
numbers = append(numbers, 40, 50)

fmt.Println(numbers) // [10 20 30 40 50]

//функция append для создания нового слайса из массива

NewSlice := append(Slice[1:2], Slice[3:4])