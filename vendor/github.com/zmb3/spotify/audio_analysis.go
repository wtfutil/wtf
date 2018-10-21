package spotify

import (
	"fmt"
)

// AudioAnalysis contains a detailed audio analysis for a single track
// identified by its unique Spotify ID. See:
// https://developer.spotify.com/web-api/get-audio-analysis/
type AudioAnalysis struct {
	Bars     []Marker      `json:"bars"`
	Beats    []Marker      `json:"beats"`
	Meta     AnalysisMeta  `json:"meta"`
	Sections []Section     `json:"sections"`
	Segments []Segment     `json:"segments"`
	Tatums   []Marker      `json:"tatums"`
	Track    AnalysisTrack `json:"track"`
}

// Marker represents beats, bars, tatums and are used in segment and section
// descriptions.
type Marker struct {
	Start      float64 `json:"start"`
	Duration   float64 `json:"duration"`
	Confidence float64 `json:"confidence"`
}

// AnalysisMeta describes details about Spotify's audio analysis of the track
type AnalysisMeta struct {
	AnalyzerVersion string  `json:"analyzer_version"`
	Platform        string  `json:"platform"`
	DetailedStatus  string  `json:"detailed_status"`
	StatusCode      int     `json:"status"`
	Timestamp       int64   `json:"timestamp"`
	AnalysisTime    float64 `json:"analysis_time"`
	InputProcess    string  `json:"input_process"`
}

// Section represents a large variation in rhythm or timbre, e.g. chorus, verse,
// bridge, guitar solo, etc. Each section contains its own descriptions of
// tempo, key, mode, time_signature, and loudness.
type Section struct {
	Marker
	Loudness                float64 `json:"loudness"`
	Tempo                   float64 `json:"tempo"`
	TempoConfidence         float64 `json:"tempo_confidence"`
	Key                     Key     `json:"key"`
	KeyConfidence           float64 `json:"key_confidence"`
	Mode                    Mode    `json:"mode"`
	ModeConfidence          float64 `json:"mode_confidence"`
	TimeSignature           int     `json:"time_signature"`
	TimeSignatureConfidence float64 `json:"time_signature_confidence"`
}

// Segment is characterized by it's perceptual onset and duration in seconds,
// loudness (dB), pitch and timbral content.
type Segment struct {
	Marker
	LoudnessStart   float64   `json:"loudness_start"`
	LoudnessMaxTime float64   `json:"loudness_max_time"`
	LoudnessMax     float64   `json:"loudness_max"`
	LoudnessEnd     float64   `json:"loudness_end"`
	Pitches         []float64 `json:"pitches"`
	Timbre          []float64 `json:"timbre"`
}

// AnalysisTrack contains audio analysis data about the track as a whole
type AnalysisTrack struct {
	NumSamples              int64   `json:"num_samples"`
	Duration                float64 `json:"duration"`
	SampleMD5               string  `json:"sample_md5"`
	OffsetSeconds           int     `json:"offset_seconds"`
	WindowSeconds           int     `json:"window_seconds"`
	AnalysisSampleRate      int64   `json:"analysis_sample_rate"`
	AnalysisChannels        int     `json:"analysis_channels"`
	EndOfFadeIn             float64 `json:"end_of_fade_in"`
	StartOfFadeOut          float64 `json:"start_of_fade_out"`
	Loudness                float64 `json:"loudness"`
	Tempo                   float64 `json:"tempo"`
	TempoConfidence         float64 `json:"tempo_confidence"`
	TimeSignature           int     `json:"time_signature"`
	TimeSignatureConfidence float64 `json:"time_signature_confidence"`
	Key                     Key     `json:"key"`
	KeyConfidence           float64 `json:"key_confidence"`
	Mode                    Mode    `json:"mode"`
	ModeConfidence          float64 `json:"mode_confidence"`
	CodeString              string  `json:"codestring"`
	CodeVersion             float64 `json:"code_version"`
	EchoprintString         string  `json:"echoprintstring"`
	EchoprintVersion        float64 `json:"echoprint_version"`
	SynchString             string  `json:"synchstring"`
	SynchVersion            float64 `json:"synch_version"`
	RhythmString            string  `json:"rhythmstring"`
	RhythmVersion           float64 `json:"rhythm_version"`
}

// GetAudioAnalysis queries the Spotify web API for an audio analysis of a
// single track.
func (c *Client) GetAudioAnalysis(id ID) (*AudioAnalysis, error) {
	url := fmt.Sprintf("%saudio-analysis/%s", c.baseURL, id)

	temp := AudioAnalysis{}

	err := c.get(url, &temp)
	if err != nil {
		return nil, err
	}

	return &temp, nil
}
