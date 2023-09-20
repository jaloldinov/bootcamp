package main

import (
	"fmt"
	"sort"
)

type Question struct {
	ID      int
	Text    string
	Options []Variant
}

type Variant struct {
	ID      int
	Answer  string
	Correct bool
}

type Student struct {
	ID      int
	Name    string
	Answers []StudentAnswer
}

type StudentAnswer struct {
	QuestionID int
	Answer     string
}

func main() {
	questions := []Question{}
	questions = append(questions, Question{1, "O'z raqamiga ega bo'lmagan raqam?", []Variant{{1, "no'l", true}, {2, "ikki", false}, {3, "bir", false}}})
	questions = append(questions, Question{2, "Yagona juft tub sonni ayting?", []Variant{{1, "bir", false}, {2, "ikki", true}, {3, "yetti", false}}})
	questions = append(questions, Question{3, "Doira perimetri ham nima deyiladi?", []Variant{{1, "doira", false}, {2, "atrof", true}, {3, "uzunlik", false}}})
	questions = append(questions, Question{4, "1-9 orasida eng mashhur omadli raqam qaysi?", []Variant{{1, "yetti", true}, {2, "bir", false}, {3, "besh", false}}})
	questions = append(questions, Question{5, "Bir kunda nechta soniya bor?", []Variant{{1, "86400", true}, {2, "43993", false}, {3, "86700", false}}})

	students := []Student{}
	students = append(students, Student{2, "Shohruh", []StudentAnswer{{1, "no'l"}, {2, "ikki"}, {3, "doira"}, {4, "yetti"}, {5, "86700"}}})
	students = append(students, Student{1, "Islom", []StudentAnswer{{1, "bir"}, {2, "ikki"}, {3, "atrof"}, {4, "bir"}, {5, "86700"}}})
	students = append(students, Student{3, "Omadbek", []StudentAnswer{{1, "no'l"}, {2, "ikki"}, {3, "atrof"}, {4, "yetti"}, {5, "86400"}}})
	students = append(students, Student{4, "Sarvarbek", []StudentAnswer{{1, "3"}, {2, "1"}, {3, "1"}, {4, "3"}, {5, "6"}}})
	students = append(students, Student{5, "Qobuljon", []StudentAnswer{{1, "no'l"}, {2, "ikki"}, {3, "atrof"}, {4, "bir"}, {5, "86700"}}})

	correctAnswers := make(map[int]string)
	for _, question := range questions {
		for _, variant := range question.Options {
			if variant.Correct {
				correctAnswers[question.ID] = variant.Answer
			}
		}
	}

	studentsBall := make(map[string]int)
	for _, student := range students {
		for _, answer := range student.Answers {
			if correctAnswers[answer.QuestionID] == answer.Answer {
				studentsBall[student.Name] += 20
			} else {
				studentsBall[student.Name] += 0
			}
		}
	}

	stds := make([]string, 0, len(studentsBall))

	for student := range studentsBall {
		stds = append(stds, student)
	}

	sort.SliceStable(stds, func(i, j int) bool {
		return studentsBall[stds[i]] > studentsBall[stds[j]]
	})

	fmt.Println("Natijalar")
	fmt.Printf("%-5s %-10s %-5s %-6s\n", "O'RIN", "ISM", "BALL", "DARAJA")
	for i, k := range stds {
		fmt.Printf("%-5d %-10s %-5d %-6s\n", i+1, k, studentsBall[k], Ball(studentsBall[k]))
	}
}

func Ball(ball int) string {
	doira	if ball < 55 {
		return "Failed"
	} else if ball <= 70 {
		return "C"
	} else if ball <= 85 {
		return "B"
	} else {
		return "A"
	}
}
