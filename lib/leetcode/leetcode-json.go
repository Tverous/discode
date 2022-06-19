package leetcode

var (
	URL    string = "https://leetcode.com"
	ListId        = map[string]string{
		"LeetCode Curated Algo 170": "\\\"552y65ke\"\\",
		"LeetCode Curated SQL 70":   "\\\"5htp6xyg\\\"",
		"Top 100 Liked Questions":   "\\\"79h8rn6\\\"",
		"Top Amazon Questions":      "\\\"7p5x763\\\"",
		"Top Facebook Questions":    "\\\"7p59281\\\"",
		"Top Google Questions":      "\\\"7p55wqm\\\"",
		"Top Interview Questions":   "\\\"wpwgkgt\\\"",
		"Top Microsoft Questions":   "\\\"55vr69d7\\\"",
	}
)

type GraphQLDataObj struct {
	Data Data `json:"data"`
}

type Data struct {
	ProblemsetQuestionList ProblemsetQuestionList `json:"problemsetQuestionList"`
}

type ProblemsetQuestionList struct {
	Total     int        `json:"total"`
	Questions []Question `json:"questions"`
}

type Question struct {
	AcRate             float32    `json:"acRate"`
	Title              string     `json:"title"`
	TitleSlug          string     `json:"titleSlug"`
	Difficulty         string     `json:"difficulty"`
	FrontendQuestionId string     `json:"frontendQuestionId"`
	TopicTags          []TopicTag `json:"topicTags"`
}

type TopicTag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
