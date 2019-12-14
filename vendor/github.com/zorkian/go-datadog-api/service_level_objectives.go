/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2017 by authors and contributors.
 */

package datadog

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// Define the available machine-readable SLO types
const (
	ServiceLevelObjectiveTypeMonitorID int = 0
	ServiceLevelObjectiveTypeMetricID  int = 1
)

// Define the available human-readable SLO types
var (
	ServiceLevelObjectiveTypeMonitor = "monitor"
	ServiceLevelObjectiveTypeMetric  = "metric"
)

// ServiceLevelObjectiveTypeFromID maps machine-readable type to human-readable type
var ServiceLevelObjectiveTypeFromID = map[int]string{
	ServiceLevelObjectiveTypeMonitorID: ServiceLevelObjectiveTypeMonitor,
	ServiceLevelObjectiveTypeMetricID:  ServiceLevelObjectiveTypeMetric,
}

// ServiceLevelObjectiveTypeToID maps human-readable type to machine-readable type
var ServiceLevelObjectiveTypeToID = map[string]int{
	ServiceLevelObjectiveTypeMonitor: ServiceLevelObjectiveTypeMonitorID,
	ServiceLevelObjectiveTypeMetric:  ServiceLevelObjectiveTypeMetricID,
}

// ServiceLevelObjectiveThreshold defines an SLO threshold and timeframe
// For example it's the `<SLO: ex 99.999%> of <SLI> within <TimeFrame: ex 7d>
type ServiceLevelObjectiveThreshold struct {
	TimeFrame      *string  `json:"timeframe,omitempty"`
	Target         *float64 `json:"target,omitempty"`
	TargetDisplay  *string  `json:"target_display,omitempty"` // Read-Only for monitor type
	Warning        *float64 `json:"warning,omitempty"`
	WarningDisplay *string  `json:"warning_display,omitempty"` // Read-Only for monitor type
}

const thresholdTolerance float64 = 1e-8

// Equal check if one threshold is equal to another.
func (s *ServiceLevelObjectiveThreshold) Equal(o interface{}) bool {
	other, ok := o.(*ServiceLevelObjectiveThreshold)
	if !ok {
		return false
	}

	return s.GetTimeFrame() == other.GetTimeFrame() &&
		Float64AlmostEqual(s.GetTarget(), other.GetTarget(), thresholdTolerance) &&
		Float64AlmostEqual(s.GetWarning(), other.GetWarning(), thresholdTolerance)
}

// String implements Stringer
func (s ServiceLevelObjectiveThreshold) String() string {
	return fmt.Sprintf("Threshold{timeframe=%s target=%f target_display=%s warning=%f warning_display=%s",
		s.GetTimeFrame(), s.GetTarget(), s.GetTargetDisplay(), s.GetWarning(), s.GetWarningDisplay())
}

// ServiceLevelObjectiveMetricQuery represents a metric-based SLO definition query
// Numerator is the sum of the `good` events
// Denominator is the sum of the `total` events
type ServiceLevelObjectiveMetricQuery struct {
	Numerator   *string `json:"numerator,omitempty"`
	Denominator *string `json:"denominator,omitempty"`
}

// ServiceLevelObjectiveThresholds is a sortable array of ServiceLevelObjectiveThreshold(s)
type ServiceLevelObjectiveThresholds []*ServiceLevelObjectiveThreshold

// Len implements sort.Interface length
func (s ServiceLevelObjectiveThresholds) Len() int {
	return len(s)
}

// Less implements sort.Interface less comparator
func (s ServiceLevelObjectiveThresholds) Less(i, j int) bool {
	iDur, _ := ServiceLevelObjectiveTimeFrameToDuration(s[i].GetTimeFrame())
	jDur, _ := ServiceLevelObjectiveTimeFrameToDuration(s[j].GetTimeFrame())
	return iDur < jDur
}

// Swap implements sort.Interface swap method
func (s ServiceLevelObjectiveThresholds) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Equal check if one set of thresholds is equal to another.
func (s ServiceLevelObjectiveThresholds) Equal(o interface{}) bool {
	other, ok := o.(ServiceLevelObjectiveThresholds)
	if !ok {
		return false
	}

	if len(s) != len(other) {
		// easy case
		return false
	}

	// compare one set from another
	sSet := make(map[string]*ServiceLevelObjectiveThreshold, 0)
	for _, t := range s {
		sSet[t.GetTimeFrame()] = t
	}
	oSet := make(map[string]*ServiceLevelObjectiveThreshold, 0)
	for _, t := range other {
		oSet[t.GetTimeFrame()] = t
	}

	for timeframe, t := range oSet {
		threshold, ok := sSet[timeframe]
		if !ok {
			// other contains more
			return false
		}

		if !threshold.Equal(t) {
			// they differ
			return false
		}
		// drop from sSet for efficiency
		delete(sSet, timeframe)
	}
	// if there are any remaining then they differ
	if len(sSet) > 0 {
		return false
	}
	return true
}

// ServiceLevelObjective defines the Service Level Objective entity
type ServiceLevelObjective struct {
	// Common
	ID          *string                         `json:"id,omitempty"`
	Name        *string                         `json:"name,omitempty"`
	Description *string                         `json:"description,omitempty"`
	Tags        []string                        `json:"tags,omitempty"`
	Thresholds  ServiceLevelObjectiveThresholds `json:"thresholds,omitempty"`
	Type        *string                         `json:"type,omitempty"`
	TypeID      *int                            `json:"type_id,omitempty"` // Read-Only
	// SLI definition
	Query         *ServiceLevelObjectiveMetricQuery `json:"query,omitempty"`
	MonitorIDs    []int                             `json:"monitor_ids,omitempty"`
	MonitorSearch *string                           `json:"monitor_search,omitempty"`
	Groups        []string                          `json:"groups,omitempty"`

	// Informational
	MonitorTags []string `json:"monitor_tags,omitempty"` // Read-Only
	Creator     *Creator `json:"creator,omitempty"`      // Read-Only
	CreatedAt   *int     `json:"created_at,omitempty"`   // Read-Only
	ModifiedAt  *int     `json:"modified_at,omitempty"`  // Read-Only
}

// MarshalJSON implements custom marshaler to ignore some fields
func (s *ServiceLevelObjective) MarshalJSON() ([]byte, error) {
	var output struct {
		ID          *string                         `json:"id,omitempty"`
		Name        *string                         `json:"name,omitempty"`
		Description *string                         `json:"description,omitempty"`
		Tags        []string                        `json:"tags,omitempty"`
		Thresholds  ServiceLevelObjectiveThresholds `json:"thresholds,omitempty"`
		Type        *string                         `json:"type,omitempty"`
		// SLI definition
		Query         *ServiceLevelObjectiveMetricQuery `json:"query,omitempty"`
		MonitorIDs    []int                             `json:"monitor_ids,omitempty"`
		MonitorSearch *string                           `json:"monitor_search,omitempty"`
		Groups        []string                          `json:"groups,omitempty"`
	}

	output.ID = s.ID
	output.Name = s.Name
	output.Description = s.Description
	output.Tags = s.Tags
	output.Thresholds = s.Thresholds
	output.Type = s.Type
	output.Query = s.Query
	output.MonitorIDs = s.MonitorIDs
	output.MonitorSearch = s.MonitorSearch
	output.Groups = s.Groups
	return json.Marshal(&output)
}

var sloTimeFrameToDurationRegex = regexp.MustCompile(`(?P<quantity>\d+)(?P<unit>(d))`)

// ServiceLevelObjectiveTimeFrameToDuration will convert a timeframe into a duration
func ServiceLevelObjectiveTimeFrameToDuration(timeframe string) (time.Duration, error) {
	match := sloTimeFrameToDurationRegex.FindStringSubmatch(timeframe)
	result := make(map[string]string)
	for i, name := range sloTimeFrameToDurationRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	if len(result) != 2 {
		return 0, fmt.Errorf("invalid timeframe specified: '%s'", timeframe)
	}
	qty, err := json.Number(result["quantity"]).Int64()
	if err != nil {
		return 0, fmt.Errorf("invalid timeframe specified, could not convert quantity to number")
	}
	if qty <= 0 {
		return 0, fmt.Errorf("invalid timeframe specified, quantity must be a positive number")
	}

	switch result["unit"] {
	// FUTURE: will support more time frames, hence the switch here.
	default:
		// only matches on `d` currently, so this is simple
		return time.Hour * 24 * time.Duration(qty), nil
	}
}

// CreateServiceLevelObjective adds a new service level objective to the system. This returns a pointer
// to the service level objective so you can pass that to UpdateServiceLevelObjective or DeleteServiceLevelObjective
// later if needed.
func (client *Client) CreateServiceLevelObjective(slo *ServiceLevelObjective) (*ServiceLevelObjective, error) {
	var out reqServiceLevelObjectives

	if slo == nil {
		return nil, fmt.Errorf("no SLO specified")
	}

	if err := client.doJsonRequest("POST", "/v1/slo", slo, &out); err != nil {
		return nil, err
	}
	if out.Error != "" {
		return nil, fmt.Errorf(out.Error)
	}

	return out.Data[0], nil
}

// UpdateServiceLevelObjective takes a service level objective that was previously retrieved through some method
// and sends it back to the server.
func (client *Client) UpdateServiceLevelObjective(slo *ServiceLevelObjective) (*ServiceLevelObjective, error) {
	var out reqServiceLevelObjectives

	if slo == nil {
		return nil, fmt.Errorf("no SLO specified")
	}

	if _, ok := slo.GetIDOk(); !ok {
		return nil, fmt.Errorf("SLO must be created first")
	}

	if err := client.doJsonRequest("PUT", fmt.Sprintf("/v1/slo/%s", slo.GetID()), slo, &out); err != nil {
		return nil, err
	}

	if out.Error != "" {
		return nil, fmt.Errorf(out.Error)
	}

	return out.Data[0], nil

}

type reqServiceLevelObjectives struct {
	Data  []*ServiceLevelObjective `json:"data"`
	Error string                   `json:"error"`
}

// SearchServiceLevelObjectives searches for service level objectives by search criteria.
// limit will limit the amount of SLO's returned, the API will enforce a maximum and default to a minimum if not specified
func (client *Client) SearchServiceLevelObjectives(limit int, offset int, query string, ids []string) ([]*ServiceLevelObjective, error) {
	var out reqServiceLevelObjectives
	uriValues := make(url.Values, 0)
	if limit > 0 {
		uriValues.Set("limit", fmt.Sprintf("%d", limit))
	}
	if offset >= 0 {
		uriValues.Set("offset", fmt.Sprintf("%d", offset))
	}
	// Either use `query` or use `ids`
	hasQuery := strings.TrimSpace(query) != ""
	hasIDs := len(ids) > 0
	if hasQuery && hasIDs {
		return nil, fmt.Errorf("invalid search: must specify either ids OR query, not both")
	}

	// specify by query
	if hasQuery {
		uriValues.Set("query", query)
	}
	// specify by `ids`
	if hasIDs {
		uriValues.Set("ids", strings.Join(ids, ","))
	}

	uri := "/v1/slo?" + uriValues.Encode()

	if err := client.doJsonRequest("GET", uri, nil, &out); err != nil {
		return nil, err
	}

	if out.Error != "" {
		return nil, fmt.Errorf(out.Error)
	}

	return out.Data, nil
}

type reqSingleServiceLevelObjective struct {
	Data  *ServiceLevelObjective `json:"data"`
	Error string                 `json:"error"`
}

// GetServiceLevelObjective retrieves an service level objective by identifier.
func (client *Client) GetServiceLevelObjective(id string) (*ServiceLevelObjective, error) {
	var out reqSingleServiceLevelObjective

	if id == "" {
		return nil, fmt.Errorf("no SLO specified")
	}

	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/slo/%s", id), nil, &out); err != nil {
		return nil, err
	}
	if out.Error != "" {
		return nil, fmt.Errorf(out.Error)
	}

	return out.Data, nil
}

type reqDeleteResp struct {
	Data  []string `json:"data"`
	Error string   `json:"error"`
}

// DeleteServiceLevelObjective removes an service level objective from the system.
func (client *Client) DeleteServiceLevelObjective(id string) error {
	var out reqDeleteResp

	if id == "" {
		return fmt.Errorf("no SLO specified")
	}

	if err := client.doJsonRequest("DELETE", fmt.Sprintf("/v1/slo/%s", id), nil, &out); err != nil {
		return err
	}

	if out.Error != "" {
		return fmt.Errorf(out.Error)
	}

	return nil
}

// DeleteServiceLevelObjectives removes multiple service level objective from the system by id.
func (client *Client) DeleteServiceLevelObjectives(ids []string) error {
	var out reqDeleteResp

	if len(ids) == 0 {
		return fmt.Errorf("no SLOs specified")
	}

	if err := client.doJsonRequest("DELETE", "/v1/slo", ids, &out); err != nil {
		return err
	}

	if out.Error != "" {
		return fmt.Errorf(out.Error)
	}

	return nil
}

// ServiceLevelObjectiveDeleteTimeFramesResponse is the response unique to the delete individual time-frames request
// this is read-only
type ServiceLevelObjectiveDeleteTimeFramesResponse struct {
	DeletedIDs []string `json:"deleted"`
	UpdatedIDs []string `json:"updated"`
}

// ServiceLevelObjectiveDeleteTimeFramesError is the error specific to deleting individual time frames.
// It contains more detailed information than the standard error.
type ServiceLevelObjectiveDeleteTimeFramesError struct {
	ID        *string `json:"id"`
	TimeFrame *string `json:"timeframe"`
	Message   *string `json:"message"`
}

// Error computes the human readable error
func (e ServiceLevelObjectiveDeleteTimeFramesError) Error() string {
	return fmt.Sprintf("error=%s id=%s for timeframe=%s", e.GetMessage(), e.GetID(), e.GetTimeFrame())
}

type timeframesDeleteResp struct {
	Data   *ServiceLevelObjectiveDeleteTimeFramesResponse `json:"data"`
	Errors []*ServiceLevelObjectiveDeleteTimeFramesError  `json:"errors"`
}

// DeleteServiceLevelObjectiveTimeFrames will delete SLO timeframes individually.
// This is useful if you have a SLO with 3 time windows and only need to delete some of the time windows.
// It will do a full delete if all time windows are removed as a result.
//
// Example:
//    SLO `12345678901234567890123456789012` was defined with 2 time frames: "7d" and "30d"
// 	  SLO `abcdefabcdefabcdefabcdefabcdefab` was defined with 2 time frames: "30d" and "90d"
//
// 		When we delete `7d` from `12345678901234567890123456789012` we still have `30d` timeframe remaining, hence this is "updated"
// 		When we delete `30d` and `90d` from `abcdefabcdefabcdefabcdefabcdefab` we are left with 0 time frames, hence this is "deleted"
// 		     and the entire SLO config is deleted
func (client *Client) DeleteServiceLevelObjectiveTimeFrames(timeframeByID map[string][]string) (*ServiceLevelObjectiveDeleteTimeFramesResponse, error) {
	var out timeframesDeleteResp

	if len(timeframeByID) == 0 {
		return nil, fmt.Errorf("nothing specified")
	}

	if err := client.doJsonRequest("POST", "/v1/slo/bulk_delete", &timeframeByID, &out); err != nil {
		return nil, err
	}

	if out.Errors != nil && len(out.Errors) > 0 {
		errMsgs := make([]string, 0)
		for _, e := range out.Errors {
			errMsgs = append(errMsgs, e.Error())
		}
		return nil, fmt.Errorf("errors deleting timeframes: %s", strings.Join(errMsgs, ","))
	}

	return out.Data, nil
}

// ServiceLevelObjectivesCanDeleteResponse is the response for a check can delete SLO endpoint.
type ServiceLevelObjectivesCanDeleteResponse struct {
	Data struct {
		OK []string `json:"ok"`
	} `json:"data"`
	Errors map[string]string `json:"errors"`
}

// CheckCanDeleteServiceLevelObjectives checks if the SLO is referenced within Datadog.
// This is useful to prevent accidental deletion.
func (client *Client) CheckCanDeleteServiceLevelObjectives(ids []string) (*ServiceLevelObjectivesCanDeleteResponse, error) {
	var out ServiceLevelObjectivesCanDeleteResponse

	if len(ids) == 0 {
		return nil, fmt.Errorf("nothing specified")
	}

	uriValues := make(url.Values, 0)
	uriValues.Set("ids", strings.Join(ids, ","))
	uri := "/v1/slo/can_delete?" + uriValues.Encode()

	if err := client.doJsonRequest("GET", uri, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// ServiceLevelObjectiveHistorySeriesPoint is a convenient wrapper for (timestamp, value) history data response.
type ServiceLevelObjectiveHistorySeriesPoint [2]json.Number

// ServiceLevelObjectiveHistoryMetricSeriesData contains the `batch_query` like history data for `metric` based SLOs
type ServiceLevelObjectiveHistoryMetricSeriesData struct {
	Count    int64       `json:"count"`
	Sum      json.Number `json:"sum"`
	MetaData struct {
		QueryIndex int     `json:"query_index"`
		Aggregator string  `json:"aggr"`
		Scope      string  `json:"scope"`
		Metric     string  `json:"metric"`
		Expression string  `json:"expression"`
		Unit       *string `json:"unit"`
	} `json:"metadata"`
	Values []json.Number `json:"values"`
	Times  []int64       `json:"times"`
}

// ValuesAsFloats will transform all the values into a slice of float64
func (d *ServiceLevelObjectiveHistoryMetricSeriesData) ValuesAsFloats() ([]float64, error) {
	out := make([]float64, len(d.Values))
	for i := 0; i < len(d.Values); i++ {
		v, err := d.Values[i].Float64()
		if err != nil {
			return out, fmt.Errorf("could not deserialize value at index %d: %s", i, err)
		}
		out[i] = v
	}
	return out, nil
}

// ValuesAsInt64s will transform all the values into a slice of int64
func (d *ServiceLevelObjectiveHistoryMetricSeriesData) ValuesAsInt64s() ([]int64, error) {
	out := make([]int64, len(d.Values))
	for i := 0; i < len(d.Values); i++ {
		v, err := d.Values[i].Int64()
		if err != nil {
			return out, fmt.Errorf("could not deserialize value at index %d: %s", i, err)
		}
		out[i] = v
	}
	return out, nil
}

// ServiceLevelObjectiveHistoryMetricSeries defines the SLO history data response for `metric` type SLOs
type ServiceLevelObjectiveHistoryMetricSeries struct {
	ResultType      string      `json:"res_type"`
	Interval        int         `json:"interval"`
	ResponseVersion json.Number `json:"resp_version"`
	Query           string      `json:"query"`   // a CSV of <numerator>, <denominator> queries
	Message         string      `json:"message"` // optional message if there are specific query issues/warnings

	Numerator   *ServiceLevelObjectiveHistoryMetricSeriesData `json:"numerator"`
	Denominator *ServiceLevelObjectiveHistoryMetricSeriesData `json:"denominator"`
}

// ServiceLevelObjectiveHistoryMonitorSeries defines the SLO history data response for `monitor` type SLOs
type ServiceLevelObjectiveHistoryMonitorSeries struct {
	Uptime        float32                                   `json:"uptime"`
	SpanPrecision json.Number                               `json:"span_precision"`
	Name          string                                    `json:"name"`
	Precision     map[string]json.Number                    `json:"precision"`
	Preview       bool                                      `json:"preview"`
	History       []ServiceLevelObjectiveHistorySeriesPoint `json:"history"`
}

// ServiceLevelObjectiveHistoryOverall defines the overall SLO history data response
// for `monitor` type SLOs there is an additional `History` property that rolls up the overall state correctly.
type ServiceLevelObjectiveHistoryOverall struct {
	Uptime        float32                `json:"uptime"`
	SpanPrecision json.Number            `json:"span_precision"`
	Name          string                 `json:"name"`
	Precision     map[string]json.Number `json:"precision"`
	Preview       bool                   `json:"preview"`

	// Monitor extension
	History []ServiceLevelObjectiveHistorySeriesPoint `json:"history"`
}

// ServiceLevelObjectiveHistoryResponseData contains the SLO history data response.
// for `monitor` based SLOs use the `Groups` property for historical data along with the `Overall.History`
// for `metric` based SLOs use the `Metrics` property for historical data. This contains `batch_query` like response
//    data
type ServiceLevelObjectiveHistoryResponseData struct {
	Errors     []string                                  `json:"errors"`
	ToTs       int64                                     `json:"to_ts"`
	FromTs     int64                                     `json:"from_ts"`
	Thresholds map[string]ServiceLevelObjectiveThreshold `json:"thresholds"`
	Overall    *ServiceLevelObjectiveHistoryOverall      `json:"overall"`

	// metric based SLO
	Metrics *ServiceLevelObjectiveHistoryMetricSeries `json:"series"`

	// monitor based SLO
	Groups []*ServiceLevelObjectiveHistoryMonitorSeries `json:"groups"`
}

// ServiceLevelObjectiveHistoryResponse is the canonical response for SLO history data.
type ServiceLevelObjectiveHistoryResponse struct {
	Data  *ServiceLevelObjectiveHistoryResponseData `json:"data"`
	Error *string                                   `json:"error"`
}

// GetServiceLevelObjectiveHistory will retrieve the history data for a given SLO and provided from/to times
func (client *Client) GetServiceLevelObjectiveHistory(id string, fromTs time.Time, toTs time.Time) (*ServiceLevelObjectiveHistoryResponse, error) {
	var out ServiceLevelObjectiveHistoryResponse

	if id == "" {
		return nil, fmt.Errorf("nothing specified")
	}

	if !toTs.After(fromTs) {
		return nil, fmt.Errorf("toTs must be after fromTs")
	}

	uriValues := make(url.Values, 0)
	uriValues.Set("from_ts", fmt.Sprintf("%d", fromTs.Unix()))
	uriValues.Set("to_ts", fmt.Sprintf("%d", toTs.Unix()))

	uri := fmt.Sprintf("/v1/slo/%s/history?%s", id, uriValues.Encode())

	if err := client.doJsonRequest("GET", uri, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
