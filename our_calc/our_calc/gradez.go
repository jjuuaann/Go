/*

Group:
Juan Monsalve
Emily Howard
Chris Hobbs
Ed Shea

-------------------------------------------------------------

A Grading calculator created in Go

Program asks user/student for the wight of
Tests	QUizes	Homeworks
asks for the number of assignments, and assignment grades

will ouput your test, quiz and homework averages
and overall average.

*/

package main

import (
	"fmt"
	//"log"
	//"net/http"
	//"strings"
)

func grading() {

	//variables needed for arrays
	var testGrades []float32
	var nTests int

	var quizGrades []float32
	var nQuizzes int

	var homeworkGrades []float32
	var nHomeworks int

	//weights
	var testsWeight float32
	var quizzesWeight float32
	var homeworksWeight float32

	fmt.Println("Please input the value of each assignment as a percentage.")

	fmt.Print("Tests: ")
	fmt.Scanln(&testsWeight)

	fmt.Print("Quiz: ")
	fmt.Scanln(&quizzesWeight)

	fmt.Print("Homework: ")
	fmt.Scanln(&homeworksWeight)

	fmt.Println("--------------------------------------------------------------")
	fmt.Println(" ")

	fmt.Println("How many tests have you had? ")
	testOptions(&testGrades, nTests)

	fmt.Println("--------------------------------------------------------------")
	fmt.Println(" ")

	fmt.Println("How many Quizes have you had? ")
	quizOptions(&quizGrades, nQuizzes)

	fmt.Println("--------------------------------------------------------------")
	fmt.Println(" ")

	fmt.Println("How many Homeworks have you had? ")
	homeworkOptions(&homeworkGrades, nHomeworks)

	calculations(testGrades, quizGrades, homeworkGrades, testsWeight, quizzesWeight, homeworksWeight)

}

/*
loops through adding grades to testGrades
Indicators and Pointers are a MUST

*/

func testOptions(testGrades *[]float32, nTests int) {

	var counter int
	fmt.Scanln(&nTests)

	for i := 0; i < nTests; i++ {
		counter++
		fmt.Print("What was your grade on Test ")
		fmt.Print(counter)
		fmt.Println("?")
		var scoreInput float32
		fmt.Scanln(&scoreInput)
		*testGrades = append(*testGrades, scoreInput)
	}

}

func quizOptions(quizGrades *[]float32, nQuizzes int) {

	var counter int
	fmt.Scanln(&nQuizzes)

	for i := 0; i < nQuizzes; i++ {
		counter++
		fmt.Print("What was your grade on Quiz ")
		fmt.Print(counter)
		fmt.Println("?")
		var scoreInput float32
		fmt.Scanln(&scoreInput)
		*quizGrades = append(*quizGrades, scoreInput)
	}

}

func homeworkOptions(homeworkGrades *[]float32, nHomeworks int) {

	var counter int
	fmt.Scanln(&nHomeworks)

	for i := 0; i < nHomeworks; i++ {
		counter++
		fmt.Print("What was your grade on Homework ")
		fmt.Print(counter)
		fmt.Println("?")
		var scoreInput float32
		fmt.Scanln(&scoreInput)
		*homeworkGrades = append(*homeworkGrades, scoreInput)
	}

}

/*
Calculations
*/

func calculations(testGrades, quizGrades, homeworkGrades []float32, testsWeight, quizzesWeight, homeworksWeight float32) {

	/*

		for loops that reads through array maps
		 _ is used to avoid a set decleration value, (float, int, String...)
	*/
	var tAverage float32 = 0
	for _, value := range testGrades {
		tAverage += value
	}

	var qAverage float32 = 0
	for _, value := range quizGrades {
		qAverage += value
	}

	var hAverage float32 = 0
	for _, value := range homeworkGrades {
		hAverage += value
	}

	var tPercent float32
	var qPercent float32
	var hPercent float32
	var sumPercent float32
	var grade float32

	fmt.Print("Tests:")
	fmt.Println(testGrades)
	fmt.Print("Quizess:")
	fmt.Println(quizGrades)
	fmt.Print("Homeworks:")
	fmt.Println(homeworkGrades)

	fmt.Println("--------------------------------------------------------------")

	/*
		we also print like this
	*/

	/*
		fmt.Print("Tests Average:")
		fmt.Println(tAverage / float32(len(testGrades)))
		fmt.Print("Quizes average:")
		fmt.Println(qAverage / float32(len(testGrades)))
		fmt.Print("homework Average:")
		fmt.Println(hAverage / float32(len(homeworkGrades)))

	*/

	tAverage = tAverage / float32(len(testGrades))
	fmt.Print("Your Test Average: ")
	fmt.Println(tAverage)
	qAverage = qAverage / float32(len(quizGrades))
	fmt.Print("Your Quiz Average: ")
	fmt.Println(qAverage)
	hAverage = hAverage / float32(len(homeworkGrades))
	fmt.Print("Your Test Average: ")
	fmt.Println(hAverage)

	fmt.Println("--------------------------------------------------------------")

	tPercent = tAverage * (testsWeight / 100)
	qPercent = qAverage * (quizzesWeight / 100)
	hPercent = hAverage * (homeworksWeight / 100)

	sumPercent = tPercent + qPercent + hPercent

	grade = sumPercent / ((testsWeight / 100) + (quizzesWeight / 100) + (homeworksWeight / 100))

	fmt.Print("Your Grade: ")
	fmt.Println(grade)

}
