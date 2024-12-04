package compiler

import "github.com/antlr4-go/antlr/v4"

// DiagnosticLevel represents the level of the diagnostic message.
type DiagnosticLevel int

const (
	Info DiagnosticLevel = iota
	Recommendation
	Warning
	Error
)

// Location represents a location in the code.
type Location struct {
	Filename   string      // The filename.
	StartToken antlr.Token // The start token.
	EndToken   antlr.Token // The end token.
}

// Diagnostic represents a diagnostic message.
type Diagnostic struct {
	Level    DiagnosticLevel
	Message  string
	Format   string
	Args     []interface{}
	Location Location
}

// DiagnosticBag represents a collection of diagnostics.
type DiagnosticBag struct {
	diagnostics          []Diagnostic
	InfosCount           int
	RecommendationsCount int
	WarningsCount        int
	ErrorsCount          int
}

// NewDiagnosticBag creates a new DiagnosticBag.
func NewDiagnosticBag() DiagnosticBag {
	return DiagnosticBag{}
}

// Report reports a diagnostic message.
func (d DiagnosticBag) Report(level DiagnosticLevel, location Location, message string, format string, args ...interface{}) {
	d.diagnostics = append(d.diagnostics, Diagnostic{
		Level:    level,
		Location: location,
		Message:  message,
		Format:   format,
		Args:     args,
	})

	switch level {
	case Info:
		d.InfosCount++
	case Recommendation:
		d.RecommendationsCount++
	case Warning:
		d.WarningsCount++
	case Error:
		d.ErrorsCount++
	}
}

// ReportInfo reports an info diagnostic message.
func (d DiagnosticBag) ReportInfo(location Location, message string, format string, args ...interface{}) {
	d.Report(Info, location, message, format, args...)
}

// ReportRecommendation reports a recommendation diagnostic message.
func (d DiagnosticBag) ReportRecommendation(location Location, message string, format string, args ...interface{}) {
	d.Report(Recommendation, location, message, format, args...)
}

// ReportWarning reports a warning diagnostic message.
func (d DiagnosticBag) ReportWarning(location Location, message string, format string, args ...interface{}) {
	d.Report(Warning, location, message, format, args...)
}

// ReportError reports an error diagnostic message.
func (d DiagnosticBag) ReportError(location Location, message string, format string, args ...interface{}) {
	d.Report(Error, location, message, format, args...)
}

// IsFailed returns true if there are any errors.
func (d DiagnosticBag) IsFailed() bool {
	return d.ErrorsCount > 0
}
