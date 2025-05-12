package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected 1 agrument: time_between_reqs")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("Error scraping feeds: %v\n", err)
		}
	}

	return nil
}

func scrapeFeeds(s *state) error {
    feed, err := s.db.GetNextFeedToFetch(context.Background())
    if err != nil {
        return err
    }
    
    feedContent, err := fetchFeed(context.Background(), feed.Url)
    if err != nil {
        return err
    }

    err = s.db.MarkFeedFetched(context.Background(), feed.ID)
    if err != nil {
        return err
    }
    
    for _, item := range feedContent.Channel.Item {
        fmt.Printf("- %s\n", item.Title)
    }
    
    return nil
}