package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func GetProblems(listId string, tags []string) Data {
	req, err := http.NewRequest("POST", URL+"/graphql/", bytes.NewBuffer(MakeGraphQLQuery(listId, tags)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	p := GraphQLDataObj{}
	if err := json.Unmarshal(data, &p); err != nil {
		log.Fatal(err)
	}

	return p.Data
}

func PickOneProblem(difficulty, listId string, tags []string, solvedQuestions mapset.Set[string]) Question {
	problems := GetProblems(listId, tags)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var idx int = r1.Intn(len(problems.ProblemsetQuestionList.Questions))
	if difficulty != "" {
		for ; difficulty != problems.ProblemsetQuestionList.Questions[idx].Difficulty || solvedQuestions.Contains(problems.ProblemsetQuestionList.Questions[idx].FrontendQuestionId); idx++ {
			if idx == len(problems.ProblemsetQuestionList.Questions)-1 {
				idx = 0
			}
		}
	}

	var elems []string
	for i := 0; i < len(problems.ProblemsetQuestionList.Questions[idx].TopicTags); i++ {
		elems = append(elems, problems.ProblemsetQuestionList.Questions[idx].TopicTags[i].Slug)
	}
	problems.ProblemsetQuestionList.Questions[idx].TopicsStr = strings.Join(elems, ",")

	return problems.ProblemsetQuestionList.Questions[idx]
}

type Filter struct {
	listId string
	tags   []string
}

func MakeGraphQLQuery(listId string, tags []string) []byte {
	f := &Filter{
		listId: listId,
		tags:   tags,
	}
	var queryString = fmt.Sprintf("{\"query\":\"query { problemsetQuestionList: questionList(categorySlug: \\\"\\\" filters: %+v) { total: totalNum questions: data { acRate, difficulty freqBar frontendQuestionId: questionFrontendId isFavor paidOnly: isPaidOnly status title titleSlug topicTags { name id slug } hasSolution hasVideoSolution } } }\",\"variables\":{}}", *f)
	var query = []byte(queryString)

	return query
}
