package main

import (
	"fmt"
	"os"
	"path"

	"github.com/martinlindhe/go-subber/common"
	"github.com/martinlindhe/go-subber/download"
	"github.com/martinlindhe/go-subber/srt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file        = kingpin.Arg("file", "A .srt (to clean) or video file (to fetch subs).").Required().File()
	verbose     = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	keepAds     = kingpin.Flag("keep-ads", "Do not strip advertisement captions.").Bool()
	skipBackups = kingpin.Flag("skip-backups", "Do not make backup (.srt.org) of original .srt").Bool()
	language    = kingpin.Flag("language", "Language.").Default("en").String()
	filter      = kingpin.Flag("filter", "Filter (none, capslock).").Default("none").String()
)

func main() {
	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	inFileName := (*file).Name()

	if len(inFileName) < 1 {
		fmt.Printf("File name required\n")
		os.Exit(0)
	}

	action(inFileName)
}

func action(inFileName string) {

	ext := path.Ext(inFileName)
	if ext == ".srt" {
		srt.CleanupSrt(inFileName, *filter, *skipBackups, *keepAds)
		return
	}

	subFileName := inFileName[0:len(inFileName)-len(ext)] + ".srt"

	if common.Exists(subFileName) {
		fmt.Println("Subs found locally, not downloading ...")
		srt.CleanupSrt(subFileName, *filter, *skipBackups, *keepAds)
		return
	}

	fmt.Printf("Downloading subs for input file ...\n")

	captions, err := download.FindSubs(inFileName, *language, *keepAds)
	common.Check(err)

	srt.WriteSrt(captions, subFileName)
}
