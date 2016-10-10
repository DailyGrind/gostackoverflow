package main

import "fmt"
import "os"
import "time"
import "html"
import "strings"
import "github.com/philgebhardt/mdcat"
import "github.com/philgebhardt/Stack-on-Go/stackongo"

func main() {
	// Or just add them to New()
	// cyan := color.New(color.Cyan, color.Bold)

	session := stackongo.NewSession("stackoverflow")

	questions, err := session.Search(os.Args[1], map[string]string{"sort": "relevance"})
	// questions, err := session.Search(os.Args[1], map[string]string{"sort": "votes"})
	if err != nil {
		fmt.Printf(err.Error())
	}

	question_items := questions.Items
	question_ids := []int{}

	for i := range [3]int{0, 1, 2} {
		// printQuestion(question)
		question_ids = append(question_ids, question_items[i].Question_id)
    }

    question_details, err := session.GetQuestions(question_ids, map[string]string{"sort": "activity", "filter": "!9YdnSJ*_T"})

	for _, question := range question_details.Items {
		
		// cyan := color.New(color.FgCyan, color.Bold)
		// color.Cyan(fmt.Sprintf("%v\n", html.UnescapeString(question.Title)))

		markdown_string := fmt.Sprintf("# [%s](%s) \n*%s* %s\n\n%s\n", 
			html.UnescapeString(question.Title),
			question.Link,
			html.UnescapeString(question.Owner.Display_name),
			time.Unix(question.Creation_date, 0),
			html.UnescapeString(question.Body_markdown))

		mdcat.Print(strings.NewReader(markdown_string))
	}
	// fmt.Printf("Questions: %v\n", questions)
}

// func printQuestion(question Question) {
// 	fmt.Printf("%v\n", html.UnescapeString(question.Title))
// 	fmt.Printf("%v %v\n", question.Owner.Display_name, time.Unix(question.Creation_date, 0))
// 	fmt.Printf("%v\n\n", question.Link)
// 	fmt.Printf("%v\n\n\n", question)
// }