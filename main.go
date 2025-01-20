package main

import "log/slog"
import "os"

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    logger.Info("MAL api start")
}
