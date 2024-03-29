package constant

type Lang struct {
	Name string
}

var Langs = map[string]*Lang{
	"zh-CN": {
		Name: "Chinese Simplified",
	},
	"en-US": {
		Name: "English USA",
	},
	"ja-JP": {
		Name: "Japanese",
	},
	"ru-RU": {
		Name: "Russian",
	},
	"ko-KR": {
		Name: "Korean",
	},
	"fr-FR": {
		Name: "French",
	},
	"es-ES": {
		Name: "Spanish",
	},
}
