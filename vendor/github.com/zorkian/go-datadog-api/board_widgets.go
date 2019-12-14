/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */

package datadog

import (
	"encoding/json"
	"fmt"
)

const (
	ALERT_GRAPH_WIDGET             = "alert_graph"
	ALERT_VALUE_WIDGET             = "alert_value"
	CHANGE_WIDGET                  = "change"
	CHECK_STATUS_WIDGET            = "check_status"
	DISTRIBUTION_WIDGET            = "distribution"
	EVENT_STREAM_WIDGET            = "event_stream"
	EVENT_TIMELINE_WIDGET          = "event_timeline"
	FREE_TEXT_WIDGET               = "free_text"
	GROUP_WIDGET                   = "group"
	HEATMAP_WIDGET                 = "heatmap"
	HOSTMAP_WIDGET                 = "hostmap"
	IFRAME_WIDGET                  = "iframe"
	IMAGE_WIDGET                   = "image"
	LOG_STREAM_WIDGET              = "log_stream"
	MANAGE_STATUS_WIDGET           = "manage_status"
	NOTE_WIDGET                    = "note"
	QUERY_VALUE_WIDGET             = "query_value"
	QUERY_TABLE_WIDGET             = "query_table"
	SCATTERPLOT_WIDGET             = "scatterplot"
	SERVICE_LEVEL_OBJECTIVE_WIDGET = "slo"
	TIMESERIES_WIDGET              = "timeseries"
	TOPLIST_WIDGET                 = "toplist"
	TRACE_SERVICE_WIDGET           = "trace_service"
)

// BoardWidget represents the structure of any widget. However, the widget Definition structure is
// different according to widget type.
type BoardWidget struct {
	Definition interface{}   `json:"definition"`
	Id         *int          `json:"id,omitempty"`
	Layout     *WidgetLayout `json:"layout,omitempty"`
}

// WidgetLayout represents the layout for a widget on a "free" dashboard
type WidgetLayout struct {
	X      *float64 `json:"x,omitempty"`
	Y      *float64 `json:"y,omitempty"`
	Height *float64 `json:"height,omitempty"`
	Width  *float64 `json:"width,omitempty"`
}

func (widget *BoardWidget) GetWidgetType() (string, error) {
	switch widget.Definition.(type) {
	case AlertGraphDefinition:
		return ALERT_GRAPH_WIDGET, nil
	case AlertValueDefinition:
		return ALERT_VALUE_WIDGET, nil
	case ChangeDefinition:
		return CHANGE_WIDGET, nil
	case CheckStatusDefinition:
		return CHECK_STATUS_WIDGET, nil
	case DistributionDefinition:
		return DISTRIBUTION_WIDGET, nil
	case EventStreamDefinition:
		return EVENT_STREAM_WIDGET, nil
	case EventTimelineDefinition:
		return EVENT_TIMELINE_WIDGET, nil
	case FreeTextDefinition:
		return FREE_TEXT_WIDGET, nil
	case GroupDefinition:
		return GROUP_WIDGET, nil
	case HeatmapDefinition:
		return HEATMAP_WIDGET, nil
	case HostmapDefinition:
		return HOSTMAP_WIDGET, nil
	case IframeDefinition:
		return IFRAME_WIDGET, nil
	case ImageDefinition:
		return IMAGE_WIDGET, nil
	case LogStreamDefinition:
		return LOG_STREAM_WIDGET, nil
	case ManageStatusDefinition:
		return MANAGE_STATUS_WIDGET, nil
	case NoteDefinition:
		return NOTE_WIDGET, nil
	case QueryValueDefinition:
		return QUERY_VALUE_WIDGET, nil
	case QueryTableDefinition:
		return QUERY_TABLE_WIDGET, nil
	case ScatterplotDefinition:
		return SCATTERPLOT_WIDGET, nil
	case ServiceLevelObjectiveDefinition:
		return SERVICE_LEVEL_OBJECTIVE_WIDGET, nil
	case TimeseriesDefinition:
		return TIMESERIES_WIDGET, nil
	case ToplistDefinition:
		return TOPLIST_WIDGET, nil
	case TraceServiceDefinition:
		return TRACE_SERVICE_WIDGET, nil
	default:
		return "", fmt.Errorf("Unsupported widget type")
	}
}

// AlertGraphDefinition represents the definition for an Alert Graph widget
type AlertGraphDefinition struct {
	Type       *string     `json:"type"`
	AlertId    *string     `json:"alert_id"`
	VizType    *string     `json:"viz_type"`
	Title      *string     `json:"title,omitempty"`
	TitleSize  *string     `json:"title_size,omitempty"`
	TitleAlign *string     `json:"title_align,omitempty"`
	Time       *WidgetTime `json:"time,omitempty"`
}

// AlertValueDefinition represents the definition for an Alert Value widget
type AlertValueDefinition struct {
	Type       *string `json:"type"`
	AlertId    *string `json:"alert_id"`
	Precision  *int    `json:"precision,omitempty"`
	Unit       *string `json:"unit,omitempty"`
	TextAlign  *string `json:"text_align,omitempty"`
	Title      *string `json:"title,omitempty"`
	TitleSize  *string `json:"title_size,omitempty"`
	TitleAlign *string `json:"title_align,omitempty"`
}

// ChangeDefinition represents the definition for a Change widget
type ChangeDefinition struct {
	Type       *string         `json:"type"`
	Requests   []ChangeRequest `json:"requests"`
	Title      *string         `json:"title,omitempty"`
	TitleSize  *string         `json:"title_size,omitempty"`
	TitleAlign *string         `json:"title_align,omitempty"`
	Time       *WidgetTime     `json:"time,omitempty"`
}
type ChangeRequest struct {
	ChangeType   *string `json:"change_type,omitempty"`
	CompareTo    *string `json:"compare_to,omitempty"`
	IncreaseGood *bool   `json:"increase_good,omitempty"`
	OrderBy      *string `json:"order_by,omitempty"`
	OrderDir     *string `json:"order_dir,omitempty"`
	ShowPresent  *bool   `json:"show_present,omitempty"`
	// A ChangeRequest should implement exactly one of the following query types
	MetricQuery  *string              `json:"q,omitempty"`
	ApmQuery     *WidgetApmOrLogQuery `json:"apm_query,omitempty"`
	LogQuery     *WidgetApmOrLogQuery `json:"log_query,omitempty"`
	ProcessQuery *WidgetProcessQuery  `json:"process_query,omitempty"`
}

// CheckStatusDefinition represents the definition for a Check Status widget
type CheckStatusDefinition struct {
	Type       *string     `json:"type"`
	Check      *string     `json:"check"`
	Grouping   *string     `json:"grouping"`
	Group      *string     `json:"group,omitempty"`
	GroupBy    []string    `json:"group_by,omitempty"`
	Tags       []string    `json:"tags,omitempty"`
	Title      *string     `json:"title,omitempty"`
	TitleSize  *string     `json:"title_size,omitempty"`
	TitleAlign *string     `json:"title_align,omitempty"`
	Time       *WidgetTime `json:"time,omitempty"`
}

// DistributionDefinition represents the definition for a Distribution widget
type DistributionDefinition struct {
	Type       *string               `json:"type"`
	Requests   []DistributionRequest `json:"requests"`
	Title      *string               `json:"title,omitempty"`
	TitleSize  *string               `json:"title_size,omitempty"`
	TitleAlign *string               `json:"title_align,omitempty"`
	Time       *WidgetTime           `json:"time,omitempty"`
}
type DistributionRequest struct {
	Style *WidgetRequestStyle `json:"style,omitempty"`
	// A DistributionRequest should implement exactly one of the following query types
	MetricQuery  *string              `json:"q,omitempty"`
	ApmQuery     *WidgetApmOrLogQuery `json:"apm_query,omitempty"`
	LogQuery     *WidgetApmOrLogQuery `json:"log_query,omitempty"`
	ProcessQuery *WidgetProcessQuery  `json:"process_query,omitempty"`
}

// EventStreamDefinition represents the definition for an Event Stream widget
type EventStreamDefinition struct {
	Type       *string     `json:"type"`
	Query      *string     `json:"query"`
	EventSize  *string     `json:"event_size,omitempty"`
	Title      *string     `json:"title,omitempty"`
	TitleSize  *string     `json:"title_size,omitempty"`
	TitleAlign *string     `json:"title_align,omitempty"`
	Time       *WidgetTime `json:"time,omitempty"`
}

// EventTimelineDefinition represents the definition for an Event Timeline widget
type EventTimelineDefinition struct {
	Type       *string     `json:"type"`
	Query      *string     `json:"query"`
	Title      *string     `json:"title,omitempty"`
	TitleSize  *string     `json:"title_size,omitempty"`
	TitleAlign *string     `json:"title_align,omitempty"`
	Time       *WidgetTime `json:"time,omitempty"`
}

// FreeTextDefinition represents the definition for a Free Text widget
type FreeTextDefinition struct {
	Type      *string `json:"type"`
	Text      *string `json:"text"`
	Color     *string `json:"color,omitempty"`
	FontSize  *string `json:"font_size,omitempty"`
	TextAlign *string `json:"text_align,omitempty"`
}

// GroupDefinition represents the definition for an Group widget
type GroupDefinition struct {
	Type       *string       `json:"type"`
	LayoutType *string       `json:"layout_type"`
	Widgets    []BoardWidget `json:"widgets"`
	Title      *string       `json:"title,omitempty"`
}

// HeatmapDefinition represents the definition for a Heatmap widget
type HeatmapDefinition struct {
	Type       *string          `json:"type"`
	Requests   []HeatmapRequest `json:"requests"`
	Yaxis      *WidgetAxis      `json:"yaxis,omitempty"`
	Events     []WidgetEvent    `json:"events,omitempty"`
	Title      *string          `json:"title,omitempty"`
	TitleSize  *string          `json:"title_size,omitempty"`
	TitleAlign *string          `json:"title_align,omitempty"`
	Time       *WidgetTime      `json:"time,omitempty"`
}
type HeatmapRequest struct {
	Style *WidgetRequestStyle `json:"style,omitempty"`
	// A HeatmapRequest should implement exactly one of the following query types
	MetricQuery  *string              `json:"q,omitempty"`
	ApmQuery     *WidgetApmOrLogQuery `json:"apm_query,omitempty"`
	LogQuery     *WidgetApmOrLogQuery `json:"log_query,omitempty"`
	ProcessQuery *WidgetProcessQuery  `json:"process_query,omitempty"`
}

// HostmapDefinition represents the definition for a Hostmap widget
type HostmapDefinition struct {
	Type          *string          `json:"type"`
	Requests      *HostmapRequests `json:"requests"`
	NodeType      *string          `json:"node_type,omitempty"`
	NoMetricHosts *bool            `json:"no_metric_hosts,omitempty"`
	NoGroupHosts  *bool            `json:"no_group_hosts,omitempty"`
	Group         []string         `json:"group,omitempty"`
	Scope         []string         `json:"scope,omitempty"`
	Style         *HostmapStyle    `json:"style,omitempty"`
	Title         *string          `json:"title,omitempty"`
	TitleSize     *string          `json:"title_size,omitempty"`
	TitleAlign    *string          `json:"title_align,omitempty"`
}
type HostmapRequests struct {
	Fill *HostmapRequest `json:"fill,omitempty"`
	Size *HostmapRequest `json:"size,omitempty"`
}
type HostmapRequest struct {
	// A HostmapRequest should implement exactly one of the following query types
	MetricQuery  *string              `json:"q,omitempty"`
	ApmQuery     *WidgetApmOrLogQuery `json:"apm_query,omitempty"`
	LogQuery     *WidgetApmOrLogQuery `json:"log_query,omitempty"`
	ProcessQuery *WidgetProcessQuery  `json:"process_query,omitempty"`
}
type HostmapStyle struct {
	Palette     *string `json:"palette,omitempty"`
	PaletteFlip *bool   `json:"palette_flip,omitempty"`
	FillMin     *string `json:"fill_min,omitempty"`
	FillMax     *string `json:"fill_max,omitempty"`
}

// IframeDefinition represents the definition for an Iframe widget
type IframeDefinition struct {
	Type *string `json:"type"`
	Url  *string `json:"url"`
}

// ImageDefinition represents the definition for an Image widget
type ImageDefinition struct {
	Type   *string `json:"type"`
	Url    *string `json:"url"`
	Sizing *string `json:"sizing,omitempty"`
	Margin *string `json:"margin,omitempty"`
}

// LogStreamDefinition represents the definition for a Log Stream widget
type LogStreamDefinition struct {
	Type       *string     `json:"type"`
	Logset     *string     `json:"logset"`
	Query      *string     `json:"query,omitempty"`
	Columns    []string    `json:"columns,omitempty"`
	Title      *string     `json:"title,omitempty"`
	TitleSize  *string     `json:"title_size,omitempty"`
	TitleAlign *string     `json:"title_align,omitempty"`
	Time       *WidgetTime `json:"time,omitempty"`
}

// ManageStatusDefinition represents the definition for a Manage Status widget
type ManageStatusDefinition struct {
	Type            *string `json:"type"`
	Query           *string `json:"query"`
	Sort            *string `json:"sort,omitempty"`
	Count           *int    `json:"count,omitempty"`
	Start           *int    `json:"start,omitempty"`
	DisplayFormat   *string `json:"display_format,omitempty"`
	ColorPreference *string `json:"color_preference,omitempty"`
	HideZeroCounts  *bool   `json:"hide_zero_counts,omitempty"`
	Title           *string `json:"title,omitempty"`
	TitleSize       *string `json:"title_size,omitempty"`
	TitleAlign      *string `json:"title_align,omitempty"`
}

// NoteDefinition represents the definition for a Note widget
type NoteDefinition struct {
	Type            *string `json:"type"`
	Content         *string `json:"content"`
	BackgroundColor *string `json:"background_color,omitempty"`
	FontSize        *string `json:"font_size,omitempty"`
	TextAlign       *string `json:"text_align,omitempty"`
	ShowTick        *bool   `json:"show_tick,omitempty"`
	TickPos         *string `json:"tick_pos,omitempty"`
	TickEdge        *string `json:"tick_edge,omitempty"`
}

// QueryValueDefinition represents the definition for a Query Value widget
type QueryValueDefinition struct {
	Type       *string             `json:"type"`
	Requests   []QueryValueRequest `json:"requests"`
	Autoscale  *bool               `json:"autoscale,omitempty"`
	CustomUnit *string             `json:"custom_unit,omitempty"`
	Precision  *int                `json:"precision,omitempty"`
	TextAlign  *string             `json:"text_align,omitempty"`
	Title      *string             `json:"title,omitempty"`
	TitleSize  *string             `json:"title_size,omitempty"`
	TitleAlign *string             `json:"title_align,omitempty"`
	Time       *WidgetTime         `json:"time,omitempty"`
}
type QueryValueRequest struct {
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty"`
	Aggregator         *string                   `json:"aggregator,omitempty"`
	// A QueryValueRequest should implement exactly one of the following query types
	MetricQuery  *string              `json:"q,omitempty"`
	ApmQuery     *WidgetApmOrLogQuery `json:"apm_query,omitempty"`
	LogQuery     *WidgetApmOrLogQuery `json:"log_query,omitempty"`
	ProcessQuery *WidgetProcessQuery  `json:"process_query,omitempty"`
}

// QueryTableDefinition represents the definition for a Table widget
type QueryTableDefinition struct {
	Type       *string             `json:"type"`
	Requests   []QueryTableRequest `json:"requests"`
	Title      *string             `json:"title,omitempty"`
	TitleSize  *string             `json:"title_size,omitempty"`
	TitleAlign *string             `json:"title_align,omitempty"`
	Time       *WidgetTime         `json:"time,omitempty"`
}
type QueryTableRequest struct {
	Alias              *string                   `json:"alias,omitempty"`
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty"`
	Aggregator         *string                   `json:"aggregator,omitempty"`
	Limit              *int                      `json:"limit,omitempty"`
	Order              *string                   `json:"order,omitempty"`
	// A QueryTableRequest should implement exactly one of the following query types
	MetricQuery  *string              `json:"q,omitempty"`
	ApmQuery     *WidgetApmOrLogQuery `json:"apm_query,omitempty"`
	LogQuery     *WidgetApmOrLogQuery `json:"log_query,omitempty"`
	ProcessQuery *WidgetProcessQuery  `json:"process_query,omitempty"`
}

// ScatterplotDefinition represents the definition for a Scatterplot widget
type ScatterplotDefinition struct {
	Type          *string              `json:"type"`
	Requests      *ScatterplotRequests `json:"requests"`
	Xaxis         *WidgetAxis          `json:"xaxis,omitempty"`
	Yaxis         *WidgetAxis          `json:"yaxis,omitempty"`
	ColorByGroups []string             `json:"color_by_groups,omitempty"`
	Title         *string              `json:"title,omitempty"`
	TitleSize     *string              `json:"title_size,omitempty"`
	TitleAlign    *string              `json:"title_align,omitempty"`
	Time          *WidgetTime          `json:"time,omitempty"`
}
type ScatterplotRequests struct {
	X *ScatterplotRequest `json:"x"`
	Y *ScatterplotRequest `json:"y"`
}
type ScatterplotRequest struct {
	Aggregator *string `json:"aggregator,omitempty"`
	// A ScatterplotRequest should implement exactly one of the following query types
	MetricQuery  *string              `json:"q,omitempty"`
	ApmQuery     *WidgetApmOrLogQuery `json:"apm_query,omitempty"`
	LogQuery     *WidgetApmOrLogQuery `json:"log_query,omitempty"`
	ProcessQuery *WidgetProcessQuery  `json:"process_query,omitempty"`
}

// ServiceLevelObjectiveDefinition represents the definition for a Service Level Objective widget
type ServiceLevelObjectiveDefinition struct {
	// Common

	Type       *string `json:"type"`
	Title      *string `json:"title,omitempty"`
	TitleSize  *string `json:"title_size,omitempty"`
	TitleAlign *string `json:"title_align,omitempty"`

	// SLO specific
	ViewType                *string  `json:"view_type,omitempty"` // currently only "detail" is supported
	ServiceLevelObjectiveID *string  `json:"slo_id,omitempty"`
	ShowErrorBudget         *bool    `json:"show_error_budget,omitempty"`
	ViewMode                *string  `json:"view_mode,omitempty"`    // overall,component,both
	TimeWindows             []string `json:"time_windows,omitempty"` // 7d,30d,90d,week_to_date,previous_week,month_to_date,previous_month
}

// TimeseriesDefinition represents the definition for a Timeseries widget
type TimeseriesDefinition struct {
	Type       *string             `json:"type"`
	Requests   []TimeseriesRequest `json:"requests"`
	Yaxis      *WidgetAxis         `json:"yaxis,omitempty"`
	Events     []WidgetEvent       `json:"events,omitempty"`
	Markers    []WidgetMarker      `json:"markers,omitempty"`
	Title      *string             `json:"title,omitempty"`
	TitleSize  *string             `json:"title_size,omitempty"`
	TitleAlign *string             `json:"title_align,omitempty"`
	ShowLegend *bool               `json:"show_legend,omitempty"`
	LegendSize *string             `json:"legend_size,omitempty"`
	Time       *WidgetTime         `json:"time,omitempty"`
}
type TimeseriesRequest struct {
	Style       *TimeseriesRequestStyle `json:"style,omitempty"`
	Metadata    []WidgetMetadata        `json:"metadata,omitempty"`
	DisplayType *string                 `json:"display_type,omitempty"`
	// A TimeseriesRequest should implement exactly one of the following query types
	MetricQuery  *string              `json:"q,omitempty"`
	ApmQuery     *WidgetApmOrLogQuery `json:"apm_query,omitempty"`
	LogQuery     *WidgetApmOrLogQuery `json:"log_query,omitempty"`
	ProcessQuery *WidgetProcessQuery  `json:"process_query,omitempty"`
}
type TimeseriesRequestStyle struct {
	Palette   *string `json:"palette,omitempty"`
	LineType  *string `json:"line_type,omitempty"`
	LineWidth *string `json:"line_width,omitempty"`
}

// ToplistDefinition represents the definition for a Top list widget
type ToplistDefinition struct {
	Type       *string          `json:"type"`
	Requests   []ToplistRequest `json:"requests"`
	Title      *string          `json:"title,omitempty"`
	TitleSize  *string          `json:"title_size,omitempty"`
	TitleAlign *string          `json:"title_align,omitempty"`
	Time       *WidgetTime      `json:"time,omitempty"`
}
type ToplistRequest struct {
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty"`
	Style              *WidgetRequestStyle       `json:"style,omitempty"`
	// A ToplistRequest should implement exactly one of the following query types
	MetricQuery  *string              `json:"q,omitempty"`
	ApmQuery     *WidgetApmOrLogQuery `json:"apm_query,omitempty"`
	LogQuery     *WidgetApmOrLogQuery `json:"log_query,omitempty"`
	ProcessQuery *WidgetProcessQuery  `json:"process_query,omitempty"`
}

// TraceServiceDefinition represents the definition for a Trace Service widget
type TraceServiceDefinition struct {
	Type             *string     `json:"type"`
	Env              *string     `json:"env"`
	Service          *string     `json:"service"`
	SpanName         *string     `json:"span_name"`
	ShowHits         *bool       `json:"show_hits,omitempty"`
	ShowErrors       *bool       `json:"show_errors,omitempty"`
	ShowLatency      *bool       `json:"show_latency,omitempty"`
	ShowBreakdown    *bool       `json:"show_breakdown,omitempty"`
	ShowDistribution *bool       `json:"show_distribution,omitempty"`
	ShowResourceList *bool       `json:"show_resource_list,omitempty"`
	SizeFormat       *string     `json:"size_format,omitempty"`
	DisplayFormat    *string     `json:"display_format,omitempty"`
	Title            *string     `json:"title,omitempty"`
	TitleSize        *string     `json:"title_size,omitempty"`
	TitleAlign       *string     `json:"title_align,omitempty"`
	Time             *WidgetTime `json:"time,omitempty"`
}

// UnmarshalJSON is a Custom Unmarshal for BoardWidget. If first tries to unmarshal the data in a light
// struct that allows to get the widget type. Then based on the widget type, it will try to unmarshal the
// data using the corresponding widget struct.
func (widget *BoardWidget) UnmarshalJSON(data []byte) error {
	var widgetHandler struct {
		Definition *struct {
			Type *string `json:"type"`
		} `json:"definition"`
		Id     *int          `json:"id,omitempty"`
		Layout *WidgetLayout `json:"layout,omitempty"`
	}
	if err := json.Unmarshal(data, &widgetHandler); err != nil {
		return err
	}

	// Get the widget id
	widget.Id = widgetHandler.Id

	// Get the widget layout
	widget.Layout = widgetHandler.Layout

	// Get the widget definition based on the widget type
	switch *widgetHandler.Definition.Type {
	case ALERT_GRAPH_WIDGET:
		var alertGraphWidget struct {
			Definition AlertGraphDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &alertGraphWidget); err != nil {
			return err
		}
		widget.Definition = alertGraphWidget.Definition
	case ALERT_VALUE_WIDGET:
		var alertValueWidget struct {
			Definition AlertValueDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &alertValueWidget); err != nil {
			return err
		}
		widget.Definition = alertValueWidget.Definition
	case CHANGE_WIDGET:
		var changeWidget struct {
			Definition ChangeDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &changeWidget); err != nil {
			return err
		}
		widget.Definition = changeWidget.Definition
	case CHECK_STATUS_WIDGET:
		var checkStatusWidget struct {
			Definition CheckStatusDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &checkStatusWidget); err != nil {
			return err
		}
		widget.Definition = checkStatusWidget.Definition
	case DISTRIBUTION_WIDGET:
		var distributionWidget struct {
			Definition DistributionDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &distributionWidget); err != nil {
			return err
		}
		widget.Definition = distributionWidget.Definition
	case EVENT_STREAM_WIDGET:
		var eventStreamWidget struct {
			Definition EventStreamDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &eventStreamWidget); err != nil {
			return err
		}
		widget.Definition = eventStreamWidget.Definition
	case EVENT_TIMELINE_WIDGET:
		var eventTimelineWidget struct {
			Definition EventTimelineDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &eventTimelineWidget); err != nil {
			return err
		}
		widget.Definition = eventTimelineWidget.Definition
	case FREE_TEXT_WIDGET:
		var freeTextWidget struct {
			Definition FreeTextDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &freeTextWidget); err != nil {
			return err
		}
		widget.Definition = freeTextWidget.Definition
	case GROUP_WIDGET:
		var groupWidget struct {
			Definition GroupDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &groupWidget); err != nil {
			return err
		}
		widget.Definition = groupWidget.Definition
	case HEATMAP_WIDGET:
		var heatmapWidget struct {
			Definition HeatmapDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &heatmapWidget); err != nil {
			return err
		}
		widget.Definition = heatmapWidget.Definition
	case HOSTMAP_WIDGET:
		var hostmapWidget struct {
			Definition HostmapDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &hostmapWidget); err != nil {
			return err
		}
		widget.Definition = hostmapWidget.Definition
	case IFRAME_WIDGET:
		var iframeWidget struct {
			Definition IframeDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &iframeWidget); err != nil {
			return err
		}
		widget.Definition = iframeWidget.Definition
	case IMAGE_WIDGET:
		var imageWidget struct {
			Definition ImageDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &imageWidget); err != nil {
			return err
		}
		widget.Definition = imageWidget.Definition
	case LOG_STREAM_WIDGET:
		var logStreamWidget struct {
			Definition LogStreamDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &logStreamWidget); err != nil {
			return err
		}
		widget.Definition = logStreamWidget.Definition
	case MANAGE_STATUS_WIDGET:
		var manageStatusWidget struct {
			Definition ManageStatusDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &manageStatusWidget); err != nil {
			return err
		}
		widget.Definition = manageStatusWidget.Definition
	case NOTE_WIDGET:
		var noteWidget struct {
			Definition NoteDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &noteWidget); err != nil {
			return err
		}
		widget.Definition = noteWidget.Definition
	case QUERY_VALUE_WIDGET:
		var queryValueWidget struct {
			Definition QueryValueDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &queryValueWidget); err != nil {
			return err
		}
		widget.Definition = queryValueWidget.Definition
	case QUERY_TABLE_WIDGET:
		var queryTableWidget struct {
			Definition QueryTableDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &queryTableWidget); err != nil {
			return err
		}
		widget.Definition = queryTableWidget.Definition
	case SCATTERPLOT_WIDGET:
		var scatterplotWidget struct {
			Definition ScatterplotDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &scatterplotWidget); err != nil {
			return err
		}
		widget.Definition = scatterplotWidget.Definition
	case SERVICE_LEVEL_OBJECTIVE_WIDGET:
		var serviceLevelObjectiveWidget struct {
			Definition ServiceLevelObjectiveDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &serviceLevelObjectiveWidget); err != nil {
			return err
		}
		widget.Definition = serviceLevelObjectiveWidget.Definition
	case TIMESERIES_WIDGET:
		var timeseriesWidget struct {
			Definition TimeseriesDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &timeseriesWidget); err != nil {
			return err
		}
		widget.Definition = timeseriesWidget.Definition
	case TOPLIST_WIDGET:
		var toplistWidget struct {
			Definition ToplistDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &toplistWidget); err != nil {
			return err
		}
		widget.Definition = toplistWidget.Definition
	case TRACE_SERVICE_WIDGET:
		var traceServiceWidget struct {
			Definition TraceServiceDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &traceServiceWidget); err != nil {
			return err
		}
		widget.Definition = traceServiceWidget.Definition
	default:
		return fmt.Errorf("Cannot unmarshal widget of type: %s", *widgetHandler.Definition.Type)
	}

	return nil
}

//
// List of structs common to multiple widget definitions
//

type WidgetTime struct {
	LiveSpan *string `json:"live_span,omitempty"`
}

type WidgetAxis struct {
	Label       *string `json:"label,omitempty"`
	Scale       *string `json:"scale,omitempty"`
	Min         *string `json:"min,omitempty"`
	Max         *string `json:"max,omitempty"`
	IncludeZero *bool   `json:"include_zero,omitempty"`
}

type WidgetEvent struct {
	Query *string `json:"q"`
}

type WidgetMarker struct {
	Value       *string `json:"value"`
	DisplayType *string `json:"display_type,omitempty"`
	Label       *string `json:"label,omitempty"`
}

type WidgetMetadata struct {
	Expression *string `json:"expression"`
	AliasName  *string `json:"alias_name,omitempty"`
}

type WidgetConditionalFormat struct {
	Comparator    *string  `json:"comparator"`
	Value         *float64 `json:"value"`
	Palette       *string  `json:"palette"`
	CustomBgColor *string  `json:"custom_bg_color,omitempty"`
	CustomFgColor *string  `json:"custom_fg_color,omitempty"`
	ImageUrl      *string  `json:"image_url,omitempty"`
	HideValue     *bool    `json:"hide_value,omitempty"`
	Timeframe     *string  `json:"timeframe,omitempty"`
	Metric        *string  `json:"metric,omitempty"`
}

// WidgetApmOrLogQuery represents an APM or a Log query
type WidgetApmOrLogQuery struct {
	Index        *string                `json:"index"`
	Compute      *ApmOrLogQueryCompute  `json:"compute,omitempty"`
	MultiCompute []ApmOrLogQueryCompute `json:"multi_compute,omitempty"`
	Search       *ApmOrLogQuerySearch   `json:"search,omitempty"`
	GroupBy      []ApmOrLogQueryGroupBy `json:"group_by,omitempty"`
}
type ApmOrLogQueryCompute struct {
	Aggregation *string `json:"aggregation"`
	Facet       *string `json:"facet,omitempty"`
	Interval    *int    `json:"interval,omitempty"`
}
type ApmOrLogQuerySearch struct {
	Query *string `json:"query"`
}
type ApmOrLogQueryGroupBy struct {
	Facet *string                   `json:"facet"`
	Limit *int                      `json:"limit,omitempty"`
	Sort  *ApmOrLogQueryGroupBySort `json:"sort,omitempty"`
}
type ApmOrLogQueryGroupBySort struct {
	Aggregation *string `json:"aggregation"`
	Order       *string `json:"order"`
	Facet       *string `json:"facet,omitempty"`
}

// WidgetProcessQuery represents a Process query
type WidgetProcessQuery struct {
	Metric   *string  `json:"metric"`
	SearchBy *string  `json:"search_by,omitempty"`
	FilterBy []string `json:"filter_by,omitempty"`
	Limit    *int     `json:"limit,omitempty"`
}

// WidgetRequestStyle represents the style that can be apply to a request
type WidgetRequestStyle struct {
	Palette *string `json:"palette,omitempty"`
}
