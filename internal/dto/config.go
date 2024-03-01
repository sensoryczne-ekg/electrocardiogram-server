package dto

type Config struct {
	DSN                      string `json:"DSN"`
	SigningSecret            string `json:"signingSecret"`
	MachineLearningServerURL string `json:"machineLearningServerURL"`
}
