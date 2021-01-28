package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
	"strconv"
)

const (
	NumOfProfessor = 11
)

type Survey struct {
	Id string // Student Name
	Scores []int // Scores of each Teacher
	Shingaku bool // whether want to go to graduate school or not
}

func NewSurvey(id string, scores []int, shingaku bool) *Survey {
	return &Survey{Id: id, Scores: scores, Shingaku: shingaku}
}

func NewSurveyFromLine(str string) *Survey {

	slice := strings.Split(str, ",")
	if len(slice) != int(NumOfProfessor+2) {
		return &Survey{}
	}

	var id string
	var scores []int
	var shingaku bool

	id = slice[0]

	strScores := slice[1:len(slice)-1]
	for _, str := range strScores {
		value, err := strconv.Atoi(str)
		if err != nil {
			println(err)
		}
		scores = append(scores, value)
	}

	shingaku, err := strconv.ParseBool(slice[len(slice)-1])
	if err != nil {
		println(err)
	}

	return NewSurvey(id, scores, shingaku)
}

func NewSurveyOnlyId(id string) *Survey {
	return &Survey{Id: id}
}

func main(){
	output := flag.String("output", "notfound.txt", "output file name")
	csv := flag.String("result", "./results.csv", "The survey result file")
	list := flag.String("input", "./input_data.txt", "The list of student")
	flag.Parse()

	results := loadSurveyFromFile(*csv, NewSurveyFromLine)
	userList := loadSurveyFromFile(*list, NewSurveyOnlyId)

	exists := make(map[string]Survey)
	var notfound []string

	for _, v := range results {
		exists[v.Id] = v
	}
	for _, v := range userList {
		if _, ok := exists[v.Id]; !ok {
			notfound = append(notfound, v.Id)
		}
	}

	saveToFile(*output, notfound)
	println("Number of students that not found in results:"+strconv.Itoa(len(notfound)))
}

func loadSurveyFromFile(csv string, function func (string) *Survey) (results []Survey) {
	fp, err := os.Open(csv)
	if err != nil {
		return nil
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		survey := function(scanner.Text())
		results = append(results, *survey)
	}

	return results
}

func saveToFile(output string, notfound []string) {
	fp, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	for _, user := range notfound {
		_, err := fp.WriteString(user+"\n")
		if err != nil {
		panic(err)
		}
	}
}