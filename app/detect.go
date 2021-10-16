package app

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"golang.org/x/sync/errgroup"
)

func NewDetectMode(app *App) error {
	return newDetectMode(app)
}

func newDetectMode(app *App) error {
	input := app.params.Input
	p := filepath.Join(input.ArchiveDir, "data", "tweet.js")
	tw, err := ReadTweetsJSON(p)
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(context.Background())
	queue := make(chan Tweet)

	g.Go(func() error {
		defer close(queue)

		for _, t := range *tw {
			select {
			case queue <- t.Tweet:
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		return nil
	})

	ret := make(chan string)
	for i := 0; i < 10; i++ {
		g.Go(func() error {
			for t := range queue {
				r := fmt.Sprintln("ID:", t.ID, "CreatedAt:", t.CreatedAt.Format(time.RFC3339))

				select {
				case ret <- r:
				case <-ctx.Done():
					return ctx.Err()
				}
			}

			return nil
		})
	}

	go func() {
		g.Wait()
		close(ret)
	}()

	for t := range ret {
		fmt.Println(t)
	}

	return nil
}
