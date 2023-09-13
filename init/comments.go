package init

import (
	"context"
	"demo/prisma/db"
	"log"
)

func InitComments() error {
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
