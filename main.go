package main

import (
	"context"
	"demo/prisma/db"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	if err := initPosts(); err != nil {
		panic(err)
	}

	if err := initComments(); err != nil {
		panic(err)
	}
}

func initPosts() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	// create a post
	createdPost, err := client.Post.CreateOne(
		db.Post.Title.Set("Hi from Prisma!"),
		db.Post.Published.Set(true),
		db.Post.Desc.Set("Prisma is a database toolkit and makes databases easy."),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdPost, "", "  ")
	fmt.Printf("created post: %s\n", result)

	// find a single post
	post, err := client.Post.FindUnique(
		db.Post.ID.Equals(createdPost.ID),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ = json.MarshalIndent(post, "", "  ")
	fmt.Printf("post: %s\n", result)

	// for optional/nullable values, you need to check the function and create two return values
	// `desc` is a string, and `ok` is a bool whether the record is null or not. If it's null,
	// `ok` is false, and `desc` will default to Go's default values; in this case an empty string (""). Otherwise,
	// `ok` is true and `desc` will be "my description".
	desc, ok := post.Desc()
	if !ok {
		return fmt.Errorf("post's description is null")
	}

	fmt.Printf("The posts's description is: %s\n", desc)

	return nil
}

func initComments() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	post, err := client.Post.CreateOne(
		db.Post.Title.Set("My new post"),
		db.Post.Published.Set(true),
		db.Post.Desc.Set("Hi there."),
		db.Post.ID.Set("123"),
	).Exec(ctx)
	if err != nil {
		return err
	}

	log.Printf("post: %+v", post)

	// then create a comment
	comments, err := client.Comment.CreateOne(
		db.Comment.Content.Set("my description"),
		// link the post we created before
		db.Comment.Post.Link(
			db.Post.ID.Equals(post.ID),
		),
	).Exec(ctx)
	if err != nil {
		return err
	}

	log.Printf("post: %+v", comments)

	return nil
}
