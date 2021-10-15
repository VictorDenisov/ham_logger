package main

import (
	"strings"
)

func parseReadingTemplate(line string) []FieldSetter {

	verbs := strings.Split(line, "\t")
	setters := make([]FieldSetter, len(verbs))
	for i, v := range verbs {
		switch v {
		case FREQUENCY:
			setters[i] = FrequencySetter
		case CALL:
			setters[i] = CallSetter
		case DATE:
			setters[i] = DateSetter
		case TIME:
			setters[i] = TimeSetter
		case MODE:
			setters[i] = ModeSetter
		}
	}
	return setters
}

func parseWritingTemplate(line string) []FieldGetter {

	verbs := strings.Split(line, "\t")
	getters := make([]FieldGetter, len(verbs))
	for i, v := range verbs {
		switch v {
		case FREQUENCY:
			getters[i] = FrequencyGetter
		case CALL:
			getters[i] = CallGetter
		case DATE:
			getters[i] = DateGetter
		case TIME:
			getters[i] = TimeGetter
		case MODE:
			getters[i] = ModeGetter
		}
	}
	return getters
}
