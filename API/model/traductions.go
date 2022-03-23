package model

// Traduction stores attributes needed for the request
type Traduction struct {
	TextToTrad string `json:"q"`
	Source     string `json:"source"`
	Target     string `json:"target"`
}

// TraductionResults stores the result of the traduction
type TraductionResults struct {
	TextTranslated string `json:"translatedText"`
}
