package config

type iconValues struct {
	Name  string
	Scale float32
}

// Exported variable
var Sizes = []iconValues{
	{"mdpi", 1},
	{"hdpi", 1.5},
	{"xhdpi", 2},
	{"xxhdpi", 3},
	{"xxxhdpi", 4},
}
