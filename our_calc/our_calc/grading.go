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
	"os"
)

func grading() {

	var testGrade1, testGrade2, testGrade3, testGrade4 float32
	var quizGrade1, quizGrade2, quizGrade3, quizGrade4 float32
	var homeworkGrade1, homeworkGrade2, homeworkGrade3, homeworkGrade4 float32

	var testsWeight float32
	var quizzesWeight float32
	var homeworksWeight float32

	var howManyTest int
	var howManyQuiz int
	var howManyHomework int

	fmt.Println("")

	fmt.Println("Please input the value of each assignment as a percentage.")

	fmt.Print("Tests: ")
	fmt.Scanln(&testsWeight)

	fmt.Print("Quiz: ")
	fmt.Scanln(&quizzesWeight)

	fmt.Print("Homework: ")
	fmt.Scanln(&homeworksWeight)

	fmt.Println("-------------------------------------------------")

	fmt.Println("If you have not done any of the assignments asked please put 0")

	fmt.Println("How many tests have you had? ")
	tOptions(&testGrade1, &testGrade2, &testGrade3, &testGrade4, &howManyTest)

	fmt.Println("How many Quizes have you had? ")
	qOptions(&quizGrade1, &quizGrade2, &quizGrade3, &quizGrade4, &howManyQuiz)

	fmt.Println("How many homework assignments have you had? ")
	hOptions(&homeworkGrade1, &homeworkGrade2, &homeworkGrade3, &homeworkGrade4, &howManyHomework)

	calculations(testsWeight, quizzesWeight, homeworksWeight, testGrade1, testGrade2, testGrade3, testGrade4,
		quizGrade1, quizGrade2, quizGrade3, quizGrade4, homeworkGrade1, homeworkGrade2, homeworkGrade3, homeworkGrade4,
		howManyTest, howManyQuiz, howManyHomework)

}

func tOptions(testGrade1, testGrade2, testGrade3, testGrade4 *float32, howManyTest *int) {

	var input int

	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
		fmt.Println("invalid input")
		return
	}
	switch input {
	case 1:
		fmt.Println("What was your grade on this assignment?")
		fmt.Scanln(testGrade1)
		*howManyTest = 1

	case 2:
		fmt.Println("What was your grade on the first assignment?")
		fmt.Scanln(testGrade1)
		fmt.Println("What was your grade on the second assignment?")
		fmt.Scanln(testGrade2)
		*howManyTest = 2

	case 3:
		fmt.Println("What was your grade on the first assignment?")
		fmt.Scanln(testGrade1)
		fmt.Println("What was your grade on the second assignment?")
		fmt.Scanln(testGrade2)
		fmt.Println("What was your grade on the third assignment?")
		fmt.Scanln(testGrade3)
		*howManyTest = 3

	case 4:
		fmt.Println("What was your grade on the first assignment?")
		fmt.Scanln(testGrade1)
		fmt.Println("What was your grade on the second assignment?")
		fmt.Scanln(testGrade2)
		fmt.Println("What was your grade on the third assignment?")
		fmt.Scanln(testGrade3)
		fmt.Println("What was your grade on the fourth assignment?")
		fmt.Scanln(testGrade4)
		*howManyTest = 4

	default:
		fmt.Println("I can only accept up to 4 assignments, please try again")
		os.Exit(2)

	}
	return
}

func qOptions(quizGrade1, quizGrade2, quizGrade3, quizGrade4 *float32, howManyQuiz *int) {

	var input int

	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
		fmt.Println("invalid input")
		return
	}
	switch input {
	case 1:
		fmt.Println("What was your grade on this assignment?")
		fmt.Scanln(quizGrade1)
		*howManyQuiz = 1

	case 2:
		fmt.Println("What was your grade on the first assignment?")
		fmt.Scanln(quizGrade1)
		fmt.Println("What was your grade on the second assignment?")
		fmt.Scanln(quizGrade2)
		*howManyQuiz = 2
	//	os.Exit(2)

	case 3:
		fmt.Println("What was your grade on the first assignment?")
		fmt.Scanln(quizGrade1)
		fmt.Println("What was your grade on the second assignment?")
		fmt.Scanln(quizGrade2)
		fmt.Println("What was your grade on the third assignment?")
		fmt.Scanln(quizGrade3)
		*howManyQuiz = 3

	case 4:
		fmt.Println("What was your grade on the first assignment?")
		fmt.Scanln(quizGrade1)
		fmt.Println("What was your grade on the second assignment?")
		fmt.Scanln(quizGrade2)
		fmt.Println("What was your grade on the third assignment?")
		fmt.Scanln(quizGrade3)
		fmt.Println("What was your grade on the fourth assignment?")
		fmt.Scanln(quizGrade4)
		*howManyQuiz = 4

	default:
		fmt.Println("I can only accept up to 4 assignments, please try again")
		os.Exit(2)

	}
}

func hOptions(homeworkGrade1, homeworkGrade2, homeworkGrade3, homeworkGrade4 *float32, howManyHomework *int) {

	var input int

	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
		fmt.Println("invalid input")
		return
	}
	switch input {
	case 1:
		fmt.Println("What was your grade on this assignment?")
		fmt.Scanln(homeworkGrade1)
		*howManyHomework = 1

	case 2:
		fmt.Println("What was your grade on the first assignment?")
		fmt.Scanln(homeworkGrade1)
		fmt.Println("What was your grade on the second assignment?")
		fmt.Scanln(homeworkGrade2)
		*howManyHomework = 2

	case 3:
		fmt.Println("What was your grade on the first assignment?")
		fmt.Scanln(homeworkGrade1)
		fmt.Println("What was your grade on the second assignment?")
		fmt.Scanln(&homeworkGrade2)
		fmt.Println("What was your grade on the third assignment?")
		fmt.Scanln(homeworkGrade3)
		*howManyHomework = 3

	case 4:
		fmt.Println("What was your grade on the first assignment?")
		fmt.Scanln(homeworkGrade1)
		fmt.Println("What was your grade on the second assignment?")
		fmt.Scanln(homeworkGrade2)
		fmt.Println("What was your grade on the third assignment?")
		fmt.Scanln(homeworkGrade3)
		fmt.Println("What was your grade on the fourth assignment?")
		fmt.Scanln(homeworkGrade4)
		*howManyHomework = 4

	default:
		fmt.Println("I can only accept up to 4 assignments, please try again")
		os.Exit(2)

	}
}

func calculations(testsWeight float32, quizzesWeight float32, homeworksWeight float32, testGrade1 float32, testGrade2 float32, testGrade3 float32, testGrade4 float32, quizGrade1 float32, quizGrade2 float32, quizGrade3 float32, quizGrade4 float32, homeworkGrade1 float32, homeworkGrade2 float32, homeworkGrade3 float32, homeworkGrade4 float32, howManyTest int, howManyQuiz int, howManyHomework int) {

	fmt.Println(testsWeight)
	fmt.Println(quizzesWeight)
	fmt.Println(homeworksWeight)
	fmt.Println(testGrade1)
	fmt.Println(testGrade2)
	fmt.Println(testGrade3)
	fmt.Println(testGrade4)
	fmt.Println(quizGrade1)
	fmt.Println(quizGrade2)
	fmt.Println(quizGrade3)
	fmt.Println(quizGrade4)
	fmt.Println(homeworkGrade1)
	fmt.Println(homeworkGrade2)
	fmt.Println(homeworkGrade3)
	fmt.Println(homeworkGrade4)

	var tAverage float32
	var qAverage float32
	var hAverage float32
	var tPercent float32
	var qPercent float32
	var hPercent float32
	var sumPercent float32
	var grade float32

	tAverage = (testGrade1 + testGrade2 + testGrade3 + testGrade4) / 4
	qAverage = (quizGrade1 + quizGrade2 + quizGrade3 + quizGrade4) / 4
	hAverage = (homeworkGrade1 + homeworkGrade2 + homeworkGrade3 + homeworkGrade4) / 4

	tPercent = tAverage * (testsWeight / 100)
	qPercent = qAverage * (quizzesWeight / 100)
	hPercent = hAverage * (homeworksWeight / 100)

	sumPercent = tPercent + qPercent + hPercent

	grade = sumPercent / ((testsWeight / 100) + (quizzesWeight / 100) + (homeworksWeight / 100))

	fmt.Println("-------------------------------------------------")

	fmt.Print("Your Test Average: ")
	fmt.Println(tAverage)
	fmt.Print("Your Quiz Average: ")
	fmt.Println(qAverage)
	fmt.Print("Your Homework Average: ")
	fmt.Println(hAverage)
	fmt.Println("-------------------------------------------------")

	fmt.Print("your grade: ")
	fmt.Println(grade)
}
