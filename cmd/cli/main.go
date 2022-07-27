package main

import (
	expand "github.com/openvenues/gopostal/expand"
	parser "github.com/openvenues/gopostal/parser"
	"go.uber.org/zap"
)

func main() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	if err := run(l); err != nil {
		panic(err)
	}
}

func run(logger *zap.Logger) error {
	logger.Sugar().Info("hello world")

	expanded := expand.ExpandAddress("504b Foundation ct, Nashville, TN 37209")
	logger.Sugar().Info(expanded)

	parsed := parser.ParseAddress("504B Foundation ct, Nashville TN 37209, Unuted States")
	logger.Sugar().Info(parsed)

	return nil
}
