package app

import "embed"

var (
	//go:embed home/*
	//go:embed doodlergame/*
	//go:embed static/*
	Static embed.FS
)
