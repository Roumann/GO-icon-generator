package config

type icons struct {
	Name  string
	Scale float32
}

type appIcons struct {
	Name       string
	Scale      float32
	MidPadding float32
}

var NotifSizes = []icons{
	{"mdpi", 1},
	{"hdpi", 1.5},
	{"xhdpi", 2},
	{"xxhdpi", 3},
	{"xxxhdpi", 4},
}

var AppSizes = []appIcons{
	{"mipmap-mdpi", 1, 4},
	{"mipmap-hdpi", 1.5, 10},
	{"mipmap-xhdpi", 2, 16},
	{"mipmap-xxhdpi", 3, 24},
	{"mipmap-xxxhdpi", 4, 32},
}
