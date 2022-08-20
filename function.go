package function

import (
	"net/http"
	"strconv"
	"context"
	"fmt"
	"os"
	"github.com/go-rod/rod"
	"github.com/jomei/notionapi"
)

func Function(w http.ResponseWriter, r *http.Request) {
	// Get the pomodoro num from Indify
	page := rod.New().MustConnect().MustPage(os.Getenv("INDIFY_URL"))
	numel := page.MustElementX("/html/body/div[1]/div[1]/div/div[1]/p")
	// fmt.Println(numel.MustText())
	pomnum, err := strconv.ParseFloat(numel.MustText(), 64)
	if err != nil {
		fmt.Printf("unable to convert the num of Pomodoro: %v", err)
		return
	}

	// Get the latest date
	nc := notionapi.NewClient(notionapi.Token(os.Getenv("NOTION_INTEGRATION_TOKEN")))
	query := &notionapi.DatabaseQueryRequest{
		Sorts: []notionapi.SortObject{
			{
				Property: "Date",
				Direction: notionapi.SortOrderDESC,
			},
		},
	}
	qr, err := nc.Database.Query(context.Background(), notionapi.DatabaseID(os.Getenv("NOTION_DATABASE_ID")), query)
	if err != nil {
		fmt.Printf("unable to get dates from Notion: %v", err)
		return
	}
	dates := qr.Results

	// Update the study hour in the latest date
	pur := &notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"Total Study Hour": notionapi.NumberProperty{
				Number: pomnum / 2,
			},
		},
	}
	up, err := nc.Page.Update(context.Background(), notionapi.PageID(dates[0].ID), pur)

	if err != nil {
		fmt.Printf("unable to update the date props: %v", err)
	}
	fmt.Println(up.URL)

	// Reset the pomodoro num
	resetel := page.MustElementX("/html/body/div[1]/div[1]/div/div[2]/button")
	resetel.MustClick()
}