package datadog

import (
	"encoding/json"
	"fmt"
)

const (
	ArithmeticProcessorType    = "arithmetic-processor"
	AttributeRemapperType      = "attribute-remapper"
	CategoryProcessorType      = "category-processor"
	DateRemapperType           = "date-remapper"
	GeoIPParserType            = "geo-ip-parser"
	GrokParserType             = "grok-parser"
	MessageRemapperType        = "message-remapper"
	NestedPipelineType         = "pipeline"
	ServiceRemapperType        = "service-remapper"
	StatusRemapperType         = "status-remapper"
	StringBuilderProcessorType = "string-builder-processor"
	TraceIdRemapperType        = "trace-id-remapper"
	UrlParserType              = "url-parser"
	UserAgentParserType        = "user-agent-parser"
)

// LogsProcessor struct represents the processor object from Config API.
type LogsProcessor struct {
	Name       *string     `json:"name"`
	IsEnabled  *bool       `json:"is_enabled"`
	Type       *string     `json:"type"`
	Definition interface{} `json:"definition"`
}

// ArithmeticProcessor struct represents unique part of arithmetic processor
// object from config API.
type ArithmeticProcessor struct {
	Expression       *string `json:"expression"`
	Target           *string `json:"target"`
	IsReplaceMissing *bool   `json:"is_replace_missing"`
}

// AttributeRemapper struct represents unique part of attribute remapper object
// from config API.
type AttributeRemapper struct {
	Sources            []string `json:"sources"`
	SourceType         *string  `json:"source_type"`
	Target             *string  `json:"target"`
	TargetType         *string  `json:"target_type"`
	PreserveSource     *bool    `json:"preserve_source"`
	OverrideOnConflict *bool    `json:"override_on_conflict"`
}

// CategoryProcessor struct represents unique part of category processor object
// from config API.
type CategoryProcessor struct {
	Target     *string    `json:"target"`
	Categories []Category `json:"categories"`
}

// Category represents category object from config API.
type Category struct {
	Name   *string              `json:"name"`
	Filter *FilterConfiguration `json:"filter"`
}

// SourceRemapper represents the object from config API that contains
// only a list of sources.
type SourceRemapper struct {
	Sources []string `json:"sources"`
}

// GeoIPParser represents geoIpParser object from config API.
type GeoIPParser struct {
	Sources []string `json:"sources"`
	Target  *string  `json:"target"`
}

type StringBuilderProcessor struct {
	Template         *string `json:"template"`
	Target           *string `json:"target"`
	IsReplaceMissing *bool   `json:"is_replace_missing"`
}

// GrokParser represents the grok parser processor object from config API.
type GrokParser struct {
	Source   *string   `json:"source"`
	Samples  []string  `json:"samples"`
	GrokRule *GrokRule `json:"grok"`
}

// GrokRule represents the rules for grok parser from config API.
type GrokRule struct {
	SupportRules *string `json:"support_rules"`
	MatchRules   *string `json:"match_rules"`
}

// NestedPipeline represents the pipeline as processor from config API.
type NestedPipeline struct {
	Filter     *FilterConfiguration `json:"filter"`
	Processors []LogsProcessor      `json:"processors,omitempty"`
}

// UrlParser represents the url parser from config API.
type UrlParser struct {
	Sources                []string `json:"sources"`
	Target                 *string  `json:"target"`
	NormalizeEndingSlashes *bool    `json:"normalize_ending_slashes"`
}

// UserAgentParser represents the user agent parser from config API.
type UserAgentParser struct {
	Sources   []string `json:"sources"`
	Target    *string  `json:"target"`
	IsEncoded *bool    `json:"is_encoded"`
}

// buildProcessor converts processor Definition of type interface{} to a map of string and interface{}.
// Simple cast from interface{} to map[string]interface{} will not work for our case here,
// since the underlying types of Definition are the processor structs.
func buildProcessor(definition interface{}) (map[string]interface{}, error) {
	inrec, err := json.Marshal(definition)
	if err != nil {
		return nil, err
	}
	var processor map[string]interface{}
	if err = json.Unmarshal(inrec, &processor); err != nil {
		return nil, err
	}
	return processor, err
}

// MarshalJSON serializes logsprocessor struct to config API compatible json object.
func (processor *LogsProcessor) MarshalJSON() ([]byte, error) {
	mapProcessor, err := buildProcessor(processor.Definition)
	if err != nil {
		return nil, err
	}
	mapProcessor["name"] = processor.Name
	mapProcessor["is_enabled"] = processor.IsEnabled
	mapProcessor["type"] = processor.Type
	jsn, err := json.Marshal(mapProcessor)
	if err != nil {
		return nil, err
	}
	return jsn, err
}

// UnmarshalJSON deserializes the config API json object to LogsProcessor struct.
func (processor *LogsProcessor) UnmarshalJSON(data []byte) error {
	var processorHandler struct {
		Type      *string `json:"type"`
		Name      *string `json:"name"`
		IsEnabled *bool   `json:"is_enabled"`
	}
	if err := json.Unmarshal(data, &processorHandler); err != nil {
		return err
	}

	processor.Name = processorHandler.Name
	processor.IsEnabled = processorHandler.IsEnabled
	processor.Type = processorHandler.Type

	switch *processorHandler.Type {
	case ArithmeticProcessorType:
		var arithmeticProcessor ArithmeticProcessor
		if err := json.Unmarshal(data, &arithmeticProcessor); err != nil {
			return err
		}
		processor.Definition = arithmeticProcessor
	case AttributeRemapperType:
		var attributeRemapper AttributeRemapper
		if err := json.Unmarshal(data, &attributeRemapper); err != nil {
			return err
		}
		processor.Definition = attributeRemapper
	case CategoryProcessorType:
		var categoryProcessor CategoryProcessor
		if err := json.Unmarshal(data, &categoryProcessor); err != nil {
			return err
		}
		processor.Definition = categoryProcessor
	case DateRemapperType,
		MessageRemapperType,
		ServiceRemapperType,
		StatusRemapperType,
		TraceIdRemapperType:
		var sourceRemapper SourceRemapper
		if err := json.Unmarshal(data, &sourceRemapper); err != nil {
			return err
		}
		processor.Definition = sourceRemapper
	case GeoIPParserType:
		var geoIPParser GeoIPParser
		if err := json.Unmarshal(data, &geoIPParser); err != nil {
			return err
		}
		processor.Definition = geoIPParser
	case GrokParserType:
		var grokParser GrokParser
		if err := json.Unmarshal(data, &grokParser); err != nil {
			return err
		}
		processor.Definition = grokParser
	case NestedPipelineType:
		var nestedPipeline NestedPipeline
		if err := json.Unmarshal(data, &nestedPipeline); err != nil {
			return err
		}
		processor.Definition = nestedPipeline
	case StringBuilderProcessorType:
		var stringBuilder StringBuilderProcessor
		if err := json.Unmarshal(data, &stringBuilder); err != nil {
			return err
		}
		processor.Definition = stringBuilder
	case UrlParserType:
		var urlParser UrlParser
		if err := json.Unmarshal(data, &urlParser); err != nil {
			return err
		}
		processor.Definition = urlParser
	case UserAgentParserType:
		var userAgentParser UserAgentParser
		if err := json.Unmarshal(data, &userAgentParser); err != nil {
			return err
		}
		processor.Definition = userAgentParser
	default:
		return fmt.Errorf("cannot unmarshal processor of type: %s", *processorHandler.Type)
	}
	return nil
}
