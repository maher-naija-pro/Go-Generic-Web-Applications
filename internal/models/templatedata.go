package models
import (
	"web_server/internal/forms"
)
// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap        map[string]string
	IntMap           map[string]int
	FloatMap         map[string]float32
	Data             map[string]interface{}
	IsAuthentifacted int
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Forms      *forms.Forms
}
