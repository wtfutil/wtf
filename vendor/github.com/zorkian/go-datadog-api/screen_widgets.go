package datadog

import "encoding/json"

type PrecisionT string

// UnmarshalJSON is a Custom Unmarshal for PrecisionT. The Datadog API can
// return 1 (int), "1" (number, but a string type) or something like "100%" or
// "*" (string).
func (p *PrecisionT) UnmarshalJSON(data []byte) error {
	var err error
	var precisionNum json.Number
	if err = json.Unmarshal(data, &precisionNum); err == nil {
		*p = PrecisionT(precisionNum)
		return nil
	}

	var precisionStr string
	if err = json.Unmarshal(data, &precisionStr); err == nil {
		*p = PrecisionT(precisionStr)
		return nil
	}

	var p0 PrecisionT
	*p = p0

	return err
}

type TileDef struct {
	Events     []TileDefEvent   `json:"events,omitempty"`
	Markers    []TileDefMarker  `json:"markers,omitempty"`
	Requests   []TileDefRequest `json:"requests,omitempty"`
	Viz        *string          `json:"viz,omitempty"`
	CustomUnit *string          `json:"custom_unit,omitempty"`
	Autoscale  *bool            `json:"autoscale,omitempty"`
	Precision  *PrecisionT      `json:"precision,omitempty"`
	TextAlign  *string          `json:"text_align,omitempty"`

	// For hostmap
	NodeType      *string       `json:"nodeType,omitempty"`
	Scope         []*string     `json:"scope,omitempty"`
	Group         []*string     `json:"group,omitempty"`
	NoGroupHosts  *bool         `json:"noGroupHosts,omitempty"`
	NoMetricHosts *bool         `json:"noMetricHosts,omitempty"`
	Style         *TileDefStyle `json:"style,omitempty"`
}

type TileDefEvent struct {
	Query *string `json:"q"`
}

type TileDefMarker struct {
	Label *string `json:"label,omitempty"`
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type TileDefRequest struct {
	Query *string `json:"q,omitempty"`

	// For Hostmap
	Type *string `json:"type,omitempty"`

	// For Process
	QueryType  *string   `json:"query_type,omitempty"`
	Metric     *string   `json:"metric,omitempty"`
	TextFilter *string   `json:"text_filter,omitempty"`
	TagFilters []*string `json:"tag_filters"`
	Limit      *int      `json:"limit,omitempty"`

	ConditionalFormats []ConditionalFormat        `json:"conditional_formats,omitempty"`
	Style              *TileDefRequestStyle       `json:"style,omitempty"`
	Aggregator         *string                    `json:"aggregator,omitempty"`
	CompareTo          *string                    `json:"compare_to,omitempty"`
	ChangeType         *string                    `json:"change_type,omitempty"`
	OrderBy            *string                    `json:"order_by,omitempty"`
	OrderDir           *string                    `json:"order_dir,omitempty"`
	ExtraCol           *string                    `json:"extra_col,omitempty"`
	IncreaseGood       *bool                      `json:"increase_good,omitempty"`
	Metadata           map[string]TileDefMetadata `json:"metadata,omitempty"`
}

type TileDefMetadata struct {
	Alias *string `json:"alias,omitempty"`
}

type ConditionalFormat struct {
	Color         *string `json:"color,omitempty"`
	Palette       *string `json:"palette,omitempty"`
	Comparator    *string `json:"comparator,omitempty"`
	Invert        *bool   `json:"invert,omitempty"`
	CustomBgColor *string `json:"custom_bg_color,omitempty"`
	Value         *string `json:"value,omitempty"`
	ImageURL      *string `json:"image_url,omitempty"`
}

type TileDefRequestStyle struct {
	Palette *string `json:"palette,omitempty"`
	Type    *string `json:"type,omitempty"`
	Width   *string `json:"width,omitempty"`
}

type TileDefStyle struct {
	Palette     *string      `json:"palette,omitempty"`
	PaletteFlip *string      `json:"paletteFlip,omitempty"`
	FillMin     *json.Number `json:"fillMin,omitempty"`
	FillMax     *json.Number `json:"fillMax,omitempty"`
}

type Time struct {
	LiveSpan *string `json:"live_span,omitempty"`
}

type Widget struct {
	// Common attributes
	Type       *string `json:"type,omitempty"`
	Title      *bool   `json:"title,omitempty"`
	TitleText  *string `json:"title_text,omitempty"`
	TitleAlign *string `json:"title_align,omitempty"`
	TitleSize  *int    `json:"title_size,omitempty"`
	Height     *int    `json:"height,omitempty"`
	Width      *int    `json:"width,omitempty"`
	X          *int    `json:"x,omitempty"`
	Y          *int    `json:"y,omitempty"`

	// For Timeseries, TopList, EventTimeline, EvenStream, AlertGraph, CheckStatus, ServiceSummary, LogStream widgets
	Time *Time `json:"time,omitempty"`

	// For Timeseries, QueryValue, HostMap, Change, Toplist, Process widgets
	TileDef *TileDef `json:"tile_def,omitempty"`

	// For FreeText widget
	Text  *string `json:"text,omitempty"`
	Color *string `json:"color,omitempty"`

	// For AlertValue widget
	TextSize  *string     `json:"text_size,omitempty"`
	Unit      *string     `json:"unit,omitempty"`
	Precision *PrecisionT `json:"precision,omitempty"`

	// AlertGraph widget
	VizType *string `json:"viz_type,omitempty"`

	// For AlertValue, QueryValue, FreeText, Note widgets
	TextAlign *string `json:"text_align,omitempty"`

	// For FreeText, Note widgets
	FontSize *string `json:"font_size,omitempty"`

	// For AlertValue, AlertGraph widgets
	AlertID     *int  `json:"alert_id,omitempty"`
	AutoRefresh *bool `json:"auto_refresh,omitempty"`

	// For Timeseries, QueryValue, Toplist widgets
	Legend     *bool   `json:"legend,omitempty"`
	LegendSize *string `json:"legend_size,omitempty"`

	// For EventTimeline, EventStream, Hostmap, LogStream widgets
	Query *string `json:"query,omitempty"`

	// For Image, IFrame widgets
	URL *string `json:"url,omitempty"`

	// For CheckStatus widget
	Tags     []*string `json:"tags,omitempty"`
	Check    *string   `json:"check,omitempty"`
	Grouping *string   `json:"grouping,omitempty"`
	GroupBy  []*string `json:"group_by,omitempty"`
	Group    *string   `json:"group,omitempty"`

	// Note widget
	TickPos  *string `json:"tick_pos,omitempty"`
	TickEdge *string `json:"tick_edge,omitempty"`
	HTML     *string `json:"html,omitempty"`
	Tick     *bool   `json:"tick,omitempty"`
	Bgcolor  *string `json:"bgcolor,omitempty"`

	// EventStream widget
	EventSize *string `json:"event_size,omitempty"`

	// Image widget
	Sizing *string `json:"sizing,omitempty"`
	Margin *string `json:"margin,omitempty"`

	// For ServiceSummary (trace_service) widget
	Env                  *string `json:"env,omitempty"`
	ServiceService       *string `json:"serviceService,omitempty"`
	ServiceName          *string `json:"serviceName,omitempty"`
	SizeVersion          *string `json:"sizeVersion,omitempty"`
	LayoutVersion        *string `json:"layoutVersion,omitempty"`
	MustShowHits         *bool   `json:"mustShowHits,omitempty"`
	MustShowErrors       *bool   `json:"mustShowErrors,omitempty"`
	MustShowLatency      *bool   `json:"mustShowLatency,omitempty"`
	MustShowBreakdown    *bool   `json:"mustShowBreakdown,omitempty"`
	MustShowDistribution *bool   `json:"mustShowDistribution,omitempty"`
	MustShowResourceList *bool   `json:"mustShowResourceList,omitempty"`

	// For MonitorSummary (manage_status) widget
	DisplayFormat          *string `json:"displayFormat,omitempty"`
	ColorPreference        *string `json:"colorPreference,omitempty"`
	HideZeroCounts         *bool   `json:"hideZeroCounts,omitempty"`
	ManageStatusShowTitle  *bool   `json:"showTitle,omitempty"`
	ManageStatusTitleText  *string `json:"titleText,omitempty"`
	ManageStatusTitleSize  *string `json:"titleSize,omitempty"`
	ManageStatusTitleAlign *string `json:"titleAlign,omitempty"`
	Params                 *Params `json:"params,omitempty"`

	// For LogStream widget
	Columns *string `json:"columns,omitempty"`
	Logset  *string `json:"logset,omitempty"`

	// For Uptime
	// Widget is undocumented, subject to breaking API changes, and without customer support
	Timeframes []*string           `json:"timeframes,omitempty"`
	Rules      map[string]*Rule    `json:"rules,omitempty"`
	Monitor    *ScreenboardMonitor `json:"monitor,omitempty"`
}

type Params struct {
	Sort  *string `json:"sort,omitempty"`
	Text  *string `json:"text,omitempty"`
	Count *string `json:"count,omitempty"`
	Start *string `json:"start,omitempty"`
}

type Rule struct {
	Threshold *json.Number `json:"threshold,omitempty"`
	Timeframe *string      `json:"timeframe,omitempty"`
	Color     *string      `json:"color,omitempty"`
}

type ScreenboardMonitor struct {
	Id *int `json:"id,omitempty"`
}
