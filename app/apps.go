package app

import "embed"

var (
	//go:embed home/*
	//go:embed doodlergame/*
	//go:embed static/*
	//go:embed admin/dist/*
	Static embed.FS
)
