package leetcode

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
)

var (
	URL string = "https://leetcode.com"
)

type Problems struct {
	Problem []Stat_status_pairs `json:"stat_status_pairs"`
}

type Stat_status_pairs struct {
	Stat       stat       `json:"stat"`
	Difficulty difficulty `json:"difficulty"`
	Piad_only  bool       `json:"paid_only"`
}

type stat struct {
	Question_id         int    `json:"question_id"`
	Question_title      string `json:"question__title"`
	Question_title_slug string `json:"question__title_slug"`
	Total_acs           int    `json:"total_acs"`
	Total_submitted     int    `json:"total_submitted"`
}

type difficulty struct {
	Level int `json:"level"`
	Alias string
}

func decodeJSON(data []byte) []Stat_status_pairs {
	p := Problems{}

	if err := json.Unmarshal(data, &p); err != nil {
		log.Fatal(err)
	}

	return p.Problem
}

func PickAProblem(checked_questions []int) Stat_status_pairs {
	problems := decodeJSON(getAllProblems())

	sort.Slice(problems[:], func(i, j int) bool {
		return problems[i].Difficulty.Level*problems[i].Stat.Total_submitted > problems[j].Difficulty.Level*problems[j].Stat.Total_submitted
	})

	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)
	//i := r1.Intn(len(problems))
	i := 0
	for j := 0; j < len(checked_questions); j++ {
		if checked_questions[j] == problems[i].Stat.Question_id {
			//i = r1.Intn(len(problems))
			i++
			j = 0
		}
	}

	switch problems[i].Difficulty.Level {
	case 1:
		problems[i].Difficulty.Alias = "Easy"
	case 2:
		problems[i].Difficulty.Alias = "Medium"
	case 3:
		problems[i].Difficulty.Alias = "Hard"
	}

	return problems[i]
}

func getAllProblems() []byte {
	resp, err := http.Get(URL + "/api/problems/all")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
