package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var sum int
	for _, v := range s.score {
		sum += v
	}
	resultFloat := float64(sum) / float64(len(s.score))
	return resultFloat
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for i := 0; i < len(s.score); i++ {
		if s.score[i] > max {
			max = s.score[i]
			name = s.name[i]
		}
	}
	return max, name
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for i := 0; i < len(s.score); i++ {
		if s.score[i] < min {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return min, name
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Printf("Input %d Student's name: ", i+1)
		fmt.Scanln(&name)
		a.name = append(a.name, name)

		var score int
		fmt.Printf("Input %s Score : ", name)
		fmt.Scanln(&score)
		a.score = append(a.score, score)
	}

	fmt.Printf("\nAverage Score Srtudent is %.f", a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Printf("\nMax score students is : %s ( %d )", nameMax, scoreMax)

	scoreMin, nameMin := a.Min()
	fmt.Printf("\nMax score students is : %s ( %d )", nameMin, scoreMin)
}
