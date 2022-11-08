package main

import (
	"discode/lib/leetcode"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatal("error creating Discord session,", err)
		return
	}

	var messages, tmp_msg []*discordgo.Message
	var beforeID string = ""
	for i := 0; i < 10; i++ {
		tmp_msg, err = dg.ChannelMessages(os.Getenv("CHANNEL_ID"), 100, beforeID, "", "")
		if err != nil {
			log.Fatal(err)
		}
		if len(tmp_msg) == 0 {
			break
		}

		messages = append(messages, tmp_msg...)
		beforeID = tmp_msg[len(tmp_msg)-1].ID
	}

	solvedQuestions := mapset.NewSet[string]()
	pastTopics := mapset.NewSet[string]()
	for _, val := range messages {
		author := val.Author
		if author.Discriminator == os.Getenv("DISCORD_BOT_ID") {
			if strings.Split(strings.Split(val.Content, "\n")[1], ",")[0] == "#A Leetcode A Day" {
				qid := strings.Split(strings.Split(val.Content, "\n")[2], ".")[0]

				solvedQuestions.Add(qid)
				if len(solvedQuestions.ToSlice()) >= 90 {
					break
				}
			}
			if strings.Split(strings.Split(val.Content, "\n")[1], "#")[0] == "本週 Topic " {
				tid := strings.Split(strings.Split(val.Content, "\n")[1], "#")[1]
				pastTopics.Add(tid)
			}
		}
	}

	topInterviewQuestions := leetcode.GetProblems(leetcode.ListId["Top Interview Questions"], []string{"\\\"\\\""})
	topInterviewTags := mapset.NewSet[string]()
	for i := 0; i < len(topInterviewQuestions.ProblemsetQuestionList.Questions); i++ {
		for j := 0; j < len(topInterviewQuestions.ProblemsetQuestionList.Questions[i].TopicTags); j++ {
			topInterviewTags.Add(topInterviewQuestions.ProblemsetQuestionList.Questions[i].TopicTags[j].Slug)
		}
	}

	var tag string
	var prefix_msg string
	if int(time.Now().Weekday()) == 1 {
		tagSlice := topInterviewTags.ToSlice()
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		idx := r1.Intn(len(tagSlice))
		for ; pastTopics.Contains(tagSlice[idx]); idx++ {
			if idx == len(tagSlice)-1 {
				idx = 0
			}
		}

		tag = "\\\"" + tagSlice[idx] + "\\\""
		prefix_msg = fmt.Sprintf("Hi @everyone \n"+
					 "This week's topic: #%v\n", tagSlice[idx])
	} else {
		for i := 0; i < len(messages); i++ {
			author := messages[i].Author
			if author.Discriminator == os.Getenv("DISCORD_BOT_ID") {
				if strings.Split(strings.Split(messages[i].Content, "\n")[1], "#")[0] == "本週 Topic " {
					tag = "\\\"" + strings.Split(strings.Split(messages[i].Content, "\n")[1], "#")[1] + "\\\""
					break
				}
			}
		}

		prefix_msg = fmt.Sprintf("Hi @everyone\n")
	}

	fmt.Println(tag)
	var p leetcode.Question
	allProblems := leetcode.GetProblems("\\\"\\\"", []string{tag})
	numOfEasy := 0
	numOfMedium := 0
	numOfHard := 0
	for i := 0; i < len(allProblems.ProblemsetQuestionList.Questions); i++ {
		if !solvedQuestions.Contains(allProblems.ProblemsetQuestionList.Questions[i].FrontendQuestionId) {
			if "Easy" == allProblems.ProblemsetQuestionList.Questions[i].Difficulty {
				numOfEasy++
			} else if "Medium" == allProblems.ProblemsetQuestionList.Questions[i].Difficulty {
				numOfMedium++
			} else if "Hard" == allProblems.ProblemsetQuestionList.Questions[i].Difficulty {
				numOfHard++
			}
		}
	}
	fmt.Println(numOfEasy, numOfMedium, numOfHard)

	if numOfHard >= numOfMedium && numOfHard >= numOfEasy {
		p = leetcode.PickOneProblem("Hard", "\\\"\\\"", []string{tag}, solvedQuestions)
	} else if numOfMedium >= numOfHard && numOfMedium >= numOfEasy {
		p = leetcode.PickOneProblem("Medium", "\\\"\\\"", []string{tag}, solvedQuestions)
	} else {
		p = leetcode.PickOneProblem("Easy", "\\\"\\\"", []string{tag}, solvedQuestions)
	}
	fmt.Println(p)

	msg := fmt.Sprintf(prefix_msg+
		"#A Leetcode A Day, %v \n"+
		"%v. %v\n",
		p.Difficulty, p.FrontendQuestionId, p.Title)

	var embedMsg discordgo.MessageEmbed
	embedMsg.URL = leetcode.URL + "/problems/" + p.TitleSlug
	embedMsg.Type = discordgo.EmbedTypeRich
	embedMsg.Title = p.FrontendQuestionId + ". " + p.Title
	embedMsg.Fields = append(embedMsg.Fields, &discordgo.MessageEmbedField{
		"Difficulty",
		p.Difficulty,
		true,
	})
	embedMsg.Fields = append(embedMsg.Fields, &discordgo.MessageEmbedField{
		"Acceptance Rate",
		fmt.Sprintf("%.2f%%", p.AcRate),
		true,
	})
	embedMsg.Fields = append(embedMsg.Fields, &discordgo.MessageEmbedField{
		"Has Solution/Is Paid",
		fmt.Sprintf("%v/%v", p.HasSolution, p.IsPaid),
		true,
	})
	embedMsg.Fields = append(embedMsg.Fields, &discordgo.MessageEmbedField{
		"Topics",
		p.TopicsStr,
		false,
	})

	if _, err := dg.ChannelMessageSendComplex(os.Getenv("CHANNEL_ID"), &discordgo.MessageSend{
		Content: msg,
		Embed:   &embedMsg,
	}); err != nil {
		log.Fatal(err)
	}

	if err := dg.Close(); err != nil {
		log.Fatal(err)
	}
}
