package main

import "fmt"
import "os"
import "time"
import "html"
import "strings"
import "github.com/philgebhardt/mdcat"
import "github.com/philgebhardt/Stack-on-Go/stackongo"

func main() {

	session := stackongo.NewSession("stackoverflow")

	questions, err := session.Search(os.Args[1], map[string]string{"sort": "relevance"})
	if err != nil {
		fmt.Printf(err.Error())
	}

	question_items := questions.Items
	question_ids := []int{}

	for i := range [3]int{0, 1, 2} {
		question_ids = append(question_ids, question_items[i].Question_id)
    }

	// filter='!9YdnSJ*_T' used to request body_markdown
    question_details, err := session.GetQuestions(question_ids, map[string]string{"sort": "activity", "filter": "!9YdnSJ*_T"})

	for _, question := range question_details.Items {
		
		markdown_string := fmt.Sprintf("# [%s](%s) \n*%s* %s\n\n%s\n", 
			html.UnescapeString(question.Title),
			question.Link,
			html.UnescapeString(question.Owner.Display_name),
			time.Unix(question.Creation_date, 0),
			html.UnescapeString(question.Body_markdown))

		mdcat.Print(strings.NewReader(markdown_string))
	}
}
