package main

import (
	"github.com/ezeoleaf/GobotTweet/config"
	"github.com/urfave/cli/v2"
)

func getFlags(c *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "topic",
			Aliases:     []string{"t"},
			Value:       "",
			Usage:       "topic for searching repos",
			Destination: &cfg.Topic,
		},
		&cli.StringFlag{
			Name:        "lang",
			Aliases:     []string{"l"},
			Value:       "",
			Usage:       "language for searching repos",
			Destination: &cfg.Language,
		},
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Value:       "./config.json",
			Usage:       "path to config file",
			Destination: &cfg.ConfigFile,
		},
		&cli.IntFlag{
			Name:        "time",
			Aliases:     []string{"x"},
			Value:       15,
			Usage:       "periodicity of tweet in minutes",
			Destination: &cfg.Periodicity,
		},
		&cli.IntFlag{
			Name:        "cache",
			Aliases:     []string{"r"},
			Value:       50,
			Usage:       "size of cache for no repeating repositories",
			Destination: &cfg.CacheSize,
		},
		&cli.StringFlag{
			Name:        "hashtag",
			Aliases:     []string{"ht"},
			Value:       "",
			Usage:       "list of comma separated hashtags",
			Destination: &cfg.Hashtags,
		},
		&cli.BoolFlag{
			Name:        "tweet-language",
			Aliases:     []string{"tl"},
			Value:       false,
			Usage:       "bool for allowing twetting the language of the repo",
			Destination: &cfg.TweetLanguage,
		},
		&cli.BoolFlag{
			Name:        "safe-mode",
			Aliases:     []string{"sf"},
			Value:       false,
			Usage:       "bool for safe mode. If safe mode is enabled, no repository is published",
			Destination: &cfg.SafeMode,
		},
		&cli.StringFlag{
			Name:        "provider",
			Aliases:     []string{"pr"},
			Value:       "github",
			Usage:       "Provider where tweetable content comes from",
			Destination: &cfg.Provider,
		},
	}
}
