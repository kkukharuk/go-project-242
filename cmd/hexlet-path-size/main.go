package main

import (
	"context"
	"fmt"
	"os"

	pathsize "code"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "<path>",
		Action: func(ctx context.Context, c *cli.Command) error {
			p := c.Args().First()
			if p == "" {
				return fmt.Errorf("path is required")
			}
			size, err := pathsize.GetSize(p)
			if err != nil {
				return err
			}
			fmt.Printf("%s\t%s\n", pathsize.FormatSize(size), p)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
