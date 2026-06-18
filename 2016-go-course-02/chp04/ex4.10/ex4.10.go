package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aliwatters/go-course-02/chp04/github"
)

func age(t time.Time) int {
	return int(time.Since(t).Hours() / 24) // days
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var newItems []*github.Issue
	var oldItems []*github.Issue
	var ancientItems []*github.Issue

	for _, item := range result.Items {
		switch {
		case age(item.CreatedAt) < 30:
			newItems = append(newItems, item)
			break
		case age(item.CreatedAt) < 366:
			oldItems = append(oldItems, item)
			break
		default:
			ancientItems = append(ancientItems, item)
		}
	}

	fmt.Println("\nIssues less than 30 days old")
	for _, item := range newItems {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("\nIssues greater than 30 days old and less than 365 days old")
	for _, item := range oldItems {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("\nIssues older than 365 days")
	for _, item := range ancientItems {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println()
}
