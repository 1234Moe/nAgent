package cliReg

import (
	"errors"
	"github.com/urfave/cli/v2"
	"nhentaiAgnet/function"
	"nhentaiAgnet/structType"
	"os"
	"strconv"
)

func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "download",
			Aliases:     []string{"d"},
			Usage:       "download comicId",
			Description: "",
			Action: func(context *cli.Context) error {
				if context.NArg() < 1 {
					return errors.New("comic id is required")
				}
				//get comic id
				comicId, err := strconv.Atoi(context.Args().Get(0))
				if err != nil {
					return errors.New("comic id is not number")
				}
				//get output dir
				outputDir := context.String("output")
				//check output dir is exist
				if _, err := os.Stat(outputDir); outputDir != "" && os.IsNotExist(err) {
					return errors.New("output dir is not exist")
				}
				//get thread count
				threadCount := context.Int("thread")
				if threadCount < 1 {
					threadCount = 1
				}
				//get comic download struct
				comicDownload := structType.ComicDownload{
					ComicId:           comicId,
					OutputDir:         outputDir,
					ProxyUrl:          context.String("proxy"),
					ThreadCount:       threadCount,
					Retry:             !context.Bool("noRetry"),
					DirNameUseComicId: context.Bool("idDir"),
					Zip:               context.Bool("zip"),
				}
				err = function.Download(comicDownload)
				return err
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "output",
					Aliases: []string{"o"},
					Usage:   "Set output directory",
				},
				&cli.IntFlag{
					Name:    "thread",
					Aliases: []string{"t"},
					Value:   5,
					Usage:   "Set download thread number",
				},
				&cli.BoolFlag{
					Name:    "noRetry",
					Aliases: []string{"nr"},
					Value:   false,
					Usage:   "Disable retry when download pictures failed",
				},
				&cli.BoolFlag{
					Name:    "idDir",
					Aliases: []string{"id"},
					Value:   false,
					Usage:   "Use comic id as directory name",
				},
				&cli.BoolFlag{
					Name:    "zip",
					Aliases: []string{"z"},
					Value:   false,
					Usage:   "Create a zip file and delete origin dir",
				},
			},
		},
	}
}
func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "stdout",
			Aliases: []string{"so"},
			Value:   "text",
			Usage:   "Set result information print to stdout type json or text (Not Finish)",
		},
		&cli.StringFlag{
			Name:    "proxy",
			Aliases: []string{"p"},
			Usage:   "Set proxy server",
		},
	}
}
