package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/xackery/eqcp/server"
)

var (
	//Version of binary
	Version string
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("failed to create log.txt:", err)
		os.Exit(1)
	}
	defer f.Close()

	//prime file writer
	output := zerolog.ConsoleWriter{Out: f, TimeFormat: "2006-01-02 15:04:05", NoColor: true}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%3s", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	outputStd := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05", NoColor: false}
	if runtime.GOOS == "windows" {
		outputStd.NoColor = true
	}
	outputStd.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%3s", i))
	}
	outputStd.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	outputStd.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}
	outputStd.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	w := io.MultiWriter(output, outputStd)
	log.Logger = zerolog.New(w).With().Timestamp().Logger()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	err = run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Info().Msgf("starting eqcp %s", Version)

	s, err := server.New(ctx, cancel, "localhost:9090")
	if err != nil {
		return errors.Wrap(err, "server")
	}
	if s == nil {

	}

	select {
	case <-ctx.Done():
	}
	return nil
}
