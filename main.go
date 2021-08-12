package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must enter an exercise name from 'A Tour of Go' Tutorial.")
		return
	}
	switch exercise := os.Args[1]; exercise {
	case "loops-and-functions":
		runLoopsAndFunctionsExercise()
	case "slices":
		runSlicesExercise()
	case "maps":
		runMapsExercise()
	case "fibonacci":
		runFibonacciExercise()
	case "stringer":
		runStringerExercise()
	case "errors":
		runErrorsExercise()
	case "reader":
		runReaderExercise()
	case "rot-reader":
		runRot13ReaderExercise()
	case "images":
		runImagesExercise()
	case "equivalent-binary-trees":
		runEquivalentBinaryTreesExercise()
	case "web-crawler":
		runWebCrawlerExercise()
	default:
		fmt.Printf("Invalid exercise: %v\n", exercise)
	}
}
