package main

import (
	"log"
	"os"

	"github.com/pavel-kiselyov/syncing-notifier/pkg"
	"github.com/urfave/cli"
)

func main() {
	// Init CLI
	app := cli.NewApp()
	app.Name = "syncing-notifier"
	app.Usage = "Sends Slack incoming webhook about Geth node syncing status"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.UintFlag{
			Name:  "interval",
			Usage: "notifications interval (ms)",
			Value: 60000,
		},
		cli.StringFlag{
			Name:  "webhook-url",
			Usage: "Slack incoming webhook URL",
		},
		cli.StringSliceFlag{
			Name:  "nodes",
			Usage: "Ethereum node RPC entrypoints",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		// Parse flags
		interval := ctx.Uint("interval")
		webhookURL := ctx.String("webhook-url")
		nodes := ctx.StringSlice("nodes")

		notifier, err := pkg.NewNotifier(nodes, webhookURL, interval)
		if err != nil {
			return err
		}

		notifier.Run()
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error runnign notifier: %s", err.Error())
	}
}
