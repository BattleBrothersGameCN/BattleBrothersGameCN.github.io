package main

import (
	"fmt"
	"log"
	"strings"
	"path/filepath"
	"os"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func newCommand() *cobra.Command {
	var (
		src = ""
		precentage = 0.7
		format = imaging.JPEG
	)

	cmd := &cobra.Command{
		Use: "resizing",
		RunE: func(cmd *cobra.Command, args []string) error {
			img, err := imaging.Open(src)
			if err != nil {
				return err
			}
			srcName := filepath.Base(src)
			srcDir := filepath.Dir(src)
			srcExt := filepath.Ext(srcName)
			srcName = strings.TrimSuffix(srcName, srcExt)
			destPath := fmt.Sprintf("%s/%s.%s", srcDir, srcName, strings.ToLower(format.String()))

			bound := img.Bounds()
			size := bound.Size()
			dest := imaging.Resize(img, int(float64(size.X) * precentage), 0, imaging.Lanczos)
			return imaging.Save(dest, destPath)
		},
	}

	cmd.Flags().StringVar(&src, "src", "", "src file to resize")
	_ = cmd.MarkFlagRequired("src")
	return cmd
}

func main() {
	log.SetOutput(os.Stderr)
	cmd := newCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}