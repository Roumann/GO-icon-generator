package config

type icons struct {
	Name  string
	Scale float32
}

var NotifSizes = []icons{
	{"mdpi", 1},
	{"hdpi", 1.5},
	{"xhdpi", 2},
	{"xxhdpi", 3},
	{"xxxhdpi", 4},
}

var AppSizes = []icons{
	{"mipmap-mdpi", 1},
	{"mipmap-hdpi", 1.5},
	{"mipmap-xhdpi", 2},
	{"mipmap-xxhdpi", 3},
	{"mipmap-xxxhdpi", 4},
}
