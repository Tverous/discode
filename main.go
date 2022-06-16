package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	// "encoding/json"
	"fmt"

)

// "fmt"
// "leetcode-bot/lib/leetcode"
// "log"
// "os"
// "strconv"
// "strings"

// "github.com/bwmarrin/discordgo"
// "github.com/joho/godotenv"

func main() {

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal(err)
	// }

	// dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	// if err != nil {
	// 	log.Fatal("error creating Discord session,", err)
	// 	return
	// }

	// var messages, tmp_msg []*discordgo.Message
	// var beforeID string
	// for i := 0; i < 4; i++ {
	// 	tmp_msg, err = dg.ChannelMessages(os.Getenv("CHANNEL_ID"), 100, beforeID, "", "")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if len(tmp_msg) == 0 {
	// 		break
	// 	}

	// 	messages = append(messages, tmp_msg...)
	// }

	// checked_question := make([]int, len(messages))
	// for _, val := range messages {
	// 	author := val.Author
	// 	if author.ID == os.Getenv("DISCORD_BOT_ID") {
	// 		qid, err := strconv.Atoi(strings.Split(strings.Split(val.Content, "\n")[2], ".")[0])
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		checked_question = append(checked_question, qid)
	// 	}
	// }

	// p := leetcode.PickAProblem(checked_question)
	// fmt.Println(p)
	// msg := fmt.Sprintf("Hi @everyone \n"+
	// 	"#A Leetcode A Day, %v \n"+
	// 	"%v. %v\n"+
	// 	"%v/problems/%v",
	// 	p.Difficulty.Alias, p.Stat.Question_id, p.Stat.Question_title, leetcode.URL, p.Stat.Question_title_slug)

	// if _, err := dg.ChannelMessageSend(os.Getenv("CHANNEL_ID"), msg); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := dg.Close(); err != nil {
	// 	log.Fatal(err)
	// }
}
