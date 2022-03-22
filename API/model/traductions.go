package model

type Traduction struct {
	TextToTrad string `json:"q"`
	Source     string `json:"source"`
	Target     string `json:"target"`
}

type TraductionResults struct {
	TextTranslated string `json:"translatedText"`
}
