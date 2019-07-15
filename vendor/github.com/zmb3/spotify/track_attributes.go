package spotify

// TrackAttributes contains various tuneable parameters that can be used for recommendations.
// For each of the tuneable track attributes, target, min and max values may be provided.
// Target:
//   Tracks with the attribute values nearest to the target values will be preferred.
//   For example, you might request TargetEnergy=0.6 and TargetDanceability=0.8.
//   All target values will be weighed equally in ranking results.
// Max:
//   A hard ceiling on the selected track attribute’s value can be provided.
//   For example, MaxInstrumentalness=0.35 would filter out most tracks
//   that are likely to be instrumental.
// Min:
//   A hard floor on the selected track attribute’s value can be provided.
//   For example, min_tempo=140 would restrict results to only those tracks
//   with a tempo of greater than 140 beats per minute.
type TrackAttributes struct {
	intAttributes   map[string]int
	floatAttributes map[string]float64
}

// NewTrackAttributes returns a new TrackAttributes instance with no attributes set.
// Attributes can then be chained following a builder pattern:
//	ta := NewTrackAttributes().
//			MaxAcousticness(0.15).
//			TargetPopularity(90)
func NewTrackAttributes() *TrackAttributes {
	return &TrackAttributes{
		intAttributes:   map[string]int{},
		floatAttributes: map[string]float64{},
	}
}

// MaxAcousticness sets the maximum acousticness
// Acousticness is a confidence measure from 0.0 to 1.0 of whether
// the track is acoustic.  A value of 1.0 represents high confidence
// that the track is acoustic.
func (ta *TrackAttributes) MaxAcousticness(acousticness float64) *TrackAttributes {
	ta.floatAttributes["max_acousticness"] = acousticness
	return ta
}

// MinAcousticness sets the minimum acousticness
// Acousticness is a confidence measure from 0.0 to 1.0 of whether
// the track is acoustic.  A value of 1.0 represents high confidence
// that the track is acoustic.
func (ta *TrackAttributes) MinAcousticness(acousticness float64) *TrackAttributes {
	ta.floatAttributes["min_acousticness"] = acousticness
	return ta
}

// TargetAcousticness sets the target acousticness
// Acousticness is a confidence measure from 0.0 to 1.0 of whether
// the track is acoustic.  A value of 1.0 represents high confidence
// that the track is acoustic.
func (ta *TrackAttributes) TargetAcousticness(acousticness float64) *TrackAttributes {
	ta.floatAttributes["target_acousticness"] = acousticness
	return ta
}

// MaxDanceability sets the maximum danceability
// Danceability describes how suitable a track is for dancing based on
// a combination of musical elements including tempo, rhythm stability,
// beat strength, and overall regularity.
// A value of 0.0 is least danceable and 1.0 is most danceable.
func (ta *TrackAttributes) MaxDanceability(danceability float64) *TrackAttributes {
	ta.floatAttributes["max_danceability"] = danceability
	return ta
}

// MinDanceability sets the minimum danceability
// Danceability describes how suitable a track is for dancing based on
// a combination of musical elements including tempo, rhythm stability,
// beat strength, and overall regularity.
// A value of 0.0 is least danceable and 1.0 is most danceable.
func (ta *TrackAttributes) MinDanceability(danceability float64) *TrackAttributes {
	ta.floatAttributes["min_danceability"] = danceability
	return ta
}

// TargetDanceability sets the target danceability
// Danceability describes how suitable a track is for dancing based on
// a combination of musical elements including tempo, rhythm stability,
// beat strength, and overall regularity.
// A value of 0.0 is least danceable and 1.0 is most danceable.
func (ta *TrackAttributes) TargetDanceability(danceability float64) *TrackAttributes {
	ta.floatAttributes["target_danceability"] = danceability
	return ta
}

// MaxDuration sets the maximum length of the track in milliseconds
func (ta *TrackAttributes) MaxDuration(duration int) *TrackAttributes {
	ta.intAttributes["max_duration_ms"] = duration
	return ta
}

// MinDuration sets the minimum length of the track in milliseconds
func (ta *TrackAttributes) MinDuration(duration int) *TrackAttributes {
	ta.intAttributes["min_duration_ms"] = duration
	return ta
}

// TargetDuration sets the target length of the track in milliseconds
func (ta *TrackAttributes) TargetDuration(duration int) *TrackAttributes {
	ta.intAttributes["target_duration_ms"] = duration
	return ta
}

// MaxEnergy sets the maximum energy
// Energy is a measure from 0.0 to 1.0 and represents a perceptual mesaure
// of intensity and activity.  Typically, energetic tracks feel fast, loud,
// and noisy.
func (ta *TrackAttributes) MaxEnergy(energy float64) *TrackAttributes {
	ta.floatAttributes["max_energy"] = energy
	return ta
}

// MinEnergy sets the minimum energy
// Energy is a measure from 0.0 to 1.0 and represents a perceptual mesaure
// of intensity and activity.  Typically, energetic tracks feel fast, loud,
// and noisy.
func (ta *TrackAttributes) MinEnergy(energy float64) *TrackAttributes {
	ta.floatAttributes["min_energy"] = energy
	return ta
}

// TargetEnergy sets the target energy
// Energy is a measure from 0.0 to 1.0 and represents a perceptual mesaure
// of intensity and activity.  Typically, energetic tracks feel fast, loud,
// and noisy.
func (ta *TrackAttributes) TargetEnergy(energy float64) *TrackAttributes {
	ta.floatAttributes["target_energy"] = energy
	return ta
}

// MaxInstrumentalness sets the maximum instrumentalness
// Instrumentalness predicts whether a track contains no vocals.
// "Ooh" and "aah" sounds are treated as instrumental in this context.
// Rap or spoken word tracks are clearly "vocal".
// The closer the instrumentalness value is to 1.0,
// the greater likelihood the track contains no vocal content.
// Values above 0.5 are intended to represent instrumental tracks,
// but confidence is higher as the value approaches 1.0.
func (ta *TrackAttributes) MaxInstrumentalness(instrumentalness float64) *TrackAttributes {
	ta.floatAttributes["max_instrumentalness"] = instrumentalness
	return ta

}

// MinInstrumentalness sets the minimum instrumentalness
// Instrumentalness predicts whether a track contains no vocals.
// "Ooh" and "aah" sounds are treated as instrumental in this context.
// Rap or spoken word tracks are clearly "vocal".
// The closer the instrumentalness value is to 1.0,
// the greater likelihood the track contains no vocal content.
// Values above 0.5 are intended to represent instrumental tracks,
// but confidence is higher as the value approaches 1.0.
func (ta *TrackAttributes) MinInstrumentalness(instrumentalness float64) *TrackAttributes {
	ta.floatAttributes["min_instrumentalness"] = instrumentalness
	return ta

}

// TargetInstrumentalness sets the target instrumentalness
// Instrumentalness predicts whether a track contains no vocals.
// "Ooh" and "aah" sounds are treated as instrumental in this context.
// Rap or spoken word tracks are clearly "vocal".
// The closer the instrumentalness value is to 1.0,
// the greater likelihood the track contains no vocal content.
// Values above 0.5 are intended to represent instrumental tracks,
// but confidence is higher as the value approaches 1.0.
func (ta *TrackAttributes) TargetInstrumentalness(instrumentalness float64) *TrackAttributes {
	ta.floatAttributes["target_instrumentalness"] = instrumentalness
	return ta

}

// MaxKey sets the maximum key
// Integers map to pitches using standard Pitch Class notation
// (https://en.wikipedia.org/wiki/Pitch_class).
func (ta *TrackAttributes) MaxKey(key int) *TrackAttributes {
	ta.intAttributes["max_key"] = key
	return ta
}

// MinKey sets the minimum key
// Integers map to pitches using standard Pitch Class notation
// (https://en.wikipedia.org/wiki/Pitch_class).
func (ta *TrackAttributes) MinKey(key int) *TrackAttributes {
	ta.intAttributes["min_key"] = key
	return ta
}

// TargetKey sets the target key
// Integers map to pitches using standard Pitch Class notation
// (https://en.wikipedia.org/wiki/Pitch_class).
func (ta *TrackAttributes) TargetKey(key int) *TrackAttributes {
	ta.intAttributes["target_key"] = key
	return ta
}

// MaxLiveness sets the maximum liveness
// Detects the presence of an audience in the recording.  Higher liveness
// values represent an increased probability that the track was performed live.
// A value above 0.8 provides strong likelihook that the track is live.
func (ta *TrackAttributes) MaxLiveness(liveness float64) *TrackAttributes {
	ta.floatAttributes["max_liveness"] = liveness
	return ta
}

// MinLiveness sets the minimum liveness
// Detects the presence of an audience in the recording.  Higher liveness
// values represent an increased probability that the track was performed live.
// A value above 0.8 provides strong likelihook that the track is live.
func (ta *TrackAttributes) MinLiveness(liveness float64) *TrackAttributes {
	ta.floatAttributes["min_liveness"] = liveness
	return ta
}

// TargetLiveness sets the target liveness
// Detects the presence of an audience in the recording.  Higher liveness
// values represent an increased probability that the track was performed live.
// A value above 0.8 provides strong likelihook that the track is live.
func (ta *TrackAttributes) TargetLiveness(liveness float64) *TrackAttributes {
	ta.floatAttributes["target_liveness"] = liveness
	return ta
}

// MaxLoudness sets the maximum loudness in decibels (dB)
// Loudness values are averaged across the entire track and are
// useful for comparing the relative loudness of tracks.
// Typical values range between -60 and 0 dB.
func (ta *TrackAttributes) MaxLoudness(loudness float64) *TrackAttributes {
	ta.floatAttributes["max_loudness"] = loudness
	return ta
}

// MinLoudness sets the minimum loudness in decibels (dB)
// Loudness values are averaged across the entire track and are
// useful for comparing the relative loudness of tracks.
// Typical values range between -60 and 0 dB.
func (ta *TrackAttributes) MinLoudness(loudness float64) *TrackAttributes {
	ta.floatAttributes["min_loudness"] = loudness
	return ta
}

// TargetLoudness sets the target loudness in decibels (dB)
// Loudness values are averaged across the entire track and are
// useful for comparing the relative loudness of tracks.
// Typical values range between -60 and 0 dB.
func (ta *TrackAttributes) TargetLoudness(loudness float64) *TrackAttributes {
	ta.floatAttributes["target_loudness"] = loudness
	return ta
}

// MaxMode sets the maximum mode
// Mode indicates the modality (major or minor) of a track.
func (ta *TrackAttributes) MaxMode(mode int) *TrackAttributes {
	ta.intAttributes["max_mode"] = mode
	return ta
}

// MinMode sets the minimum mode
// Mode indicates the modality (major or minor) of a track.
func (ta *TrackAttributes) MinMode(mode int) *TrackAttributes {
	ta.intAttributes["min_mode"] = mode
	return ta
}

// TargetMode sets the target mode
// Mode indicates the modality (major or minor) of a track.
func (ta *TrackAttributes) TargetMode(mode int) *TrackAttributes {
	ta.intAttributes["target_mode"] = mode
	return ta
}

// MaxPopularity sets the maximum popularity.
// The value will be between 0 and 100, with 100 being the most popular.
// The popularity is calculated by algorithm and is based, in the most part,
// on the total number of plays the track has had and how recent those plays are.
// Note: When applying track relinking via the market parameter, it is expected to find
// relinked tracks with popularities that do not match min_*, max_* and target_* popularities.
// These relinked tracks are accurate replacements for unplayable tracks
// with the expected popularity scores. Original, non-relinked tracks are
// available via the linked_from attribute of the relinked track response.
func (ta *TrackAttributes) MaxPopularity(popularity int) *TrackAttributes {
	ta.intAttributes["max_popularity"] = popularity
	return ta
}

// MinPopularity sets the minimum popularity.
// The value will be between 0 and 100, with 100 being the most popular.
// The popularity is calculated by algorithm and is based, in the most part,
// on the total number of plays the track has had and how recent those plays are.
// Note: When applying track relinking via the market parameter, it is expected to find
// relinked tracks with popularities that do not match min_*, max_* and target_* popularities.
// These relinked tracks are accurate replacements for unplayable tracks
// with the expected popularity scores. Original, non-relinked tracks are
// available via the linked_from attribute of the relinked track response.
func (ta *TrackAttributes) MinPopularity(popularity int) *TrackAttributes {
	ta.intAttributes["min_popularity"] = popularity
	return ta
}

// TargetPopularity sets the target popularity.
// The value will be between 0 and 100, with 100 being the most popular.
// The popularity is calculated by algorithm and is based, in the most part,
// on the total number of plays the track has had and how recent those plays are.
// Note: When applying track relinking via the market parameter, it is expected to find
// relinked tracks with popularities that do not match min_*, max_* and target_* popularities.
// These relinked tracks are accurate replacements for unplayable tracks
// with the expected popularity scores. Original, non-relinked tracks are
// available via the linked_from attribute of the relinked track response.
func (ta *TrackAttributes) TargetPopularity(popularity int) *TrackAttributes {
	ta.intAttributes["target_popularity"] = popularity
	return ta
}

// MaxSpeechiness sets the maximum speechiness.
// Speechiness detects the presence of spoken words in a track.
// The more exclusively speech-like the recording, the closer to 1.0
// the speechiness will be.
// Values above 0.66 describe tracks that are probably made entirely of
// spoken words.  Values between 0.33 and 0.66 describe tracks that may
// contain both music and speech, including such cases as rap music.
// Values below 0.33 most likely represent music and other non-speech-like tracks.
func (ta *TrackAttributes) MaxSpeechiness(speechiness float64) *TrackAttributes {
	ta.floatAttributes["max_speechiness"] = speechiness
	return ta

}

// MinSpeechiness sets the minimum speechiness.
// Speechiness detects the presence of spoken words in a track.
// The more exclusively speech-like the recording, the closer to 1.0
// the speechiness will be.
// Values above 0.66 describe tracks that are probably made entirely of
// spoken words.  Values between 0.33 and 0.66 describe tracks that may
// contain both music and speech, including such cases as rap music.
// Values below 0.33 most likely represent music and other non-speech-like tracks.
func (ta *TrackAttributes) MinSpeechiness(speechiness float64) *TrackAttributes {
	ta.floatAttributes["min_speechiness"] = speechiness
	return ta

}

// TargetSpeechiness sets the target speechiness.
// Speechiness detects the presence of spoken words in a track.
// The more exclusively speech-like the recording, the closer to 1.0
// the speechiness will be.
// Values above 0.66 describe tracks that are probably made entirely of
// spoken words.  Values between 0.33 and 0.66 describe tracks that may
// contain both music and speech, including such cases as rap music.
// Values below 0.33 most likely represent music and other non-speech-like tracks.
func (ta *TrackAttributes) TargetSpeechiness(speechiness float64) *TrackAttributes {
	ta.floatAttributes["target_speechiness"] = speechiness
	return ta
}

// MaxTempo sets the maximum tempo in beats per minute (BPM).
func (ta *TrackAttributes) MaxTempo(tempo float64) *TrackAttributes {
	ta.floatAttributes["max_tempo"] = tempo
	return ta
}

// MinTempo sets the minimum tempo in beats per minute (BPM).
func (ta *TrackAttributes) MinTempo(tempo float64) *TrackAttributes {
	ta.floatAttributes["min_tempo"] = tempo
	return ta
}

// TargetTempo sets the target tempo in beats per minute (BPM).
func (ta *TrackAttributes) TargetTempo(tempo float64) *TrackAttributes {
	ta.floatAttributes["target_tempo"] = tempo
	return ta

}

// MaxTimeSignature sets the maximum time signature
// The time signature (meter) is a notational convention to
// specify how many beats are in each bar (or measure).
func (ta *TrackAttributes) MaxTimeSignature(timeSignature int) *TrackAttributes {
	ta.intAttributes["max_time_signature"] = timeSignature
	return ta
}

// MinTimeSignature sets the minimum time signature
// The time signature (meter) is a notational convention to
// specify how many beats are in each bar (or measure).
func (ta *TrackAttributes) MinTimeSignature(timeSignature int) *TrackAttributes {
	ta.intAttributes["min_time_signature"] = timeSignature
	return ta
}

// TargetTimeSignature sets the target time signature
// The time signature (meter) is a notational convention to
// specify how many beats are in each bar (or measure).
func (ta *TrackAttributes) TargetTimeSignature(timeSignature int) *TrackAttributes {
	ta.intAttributes["target_time_signature"] = timeSignature
	return ta
}

// MaxValence sets the maximum valence.
// Valence is a measure from 0.0 to 1.0 describing the musical positiveness
/// conveyed by a track.
// Tracks with high valence sound more positive (e.g. happy, cheerful, euphoric),
// while tracks with low valence sound more negative (e.g. sad, depressed, angry).
func (ta *TrackAttributes) MaxValence(valence float64) *TrackAttributes {
	ta.floatAttributes["max_valence"] = valence
	return ta
}

// MinValence sets the minimum valence.
// Valence is a measure from 0.0 to 1.0 describing the musical positiveness
/// conveyed by a track.
// Tracks with high valence sound more positive (e.g. happy, cheerful, euphoric),
// while tracks with low valence sound more negative (e.g. sad, depressed, angry).
func (ta *TrackAttributes) MinValence(valence float64) *TrackAttributes {
	ta.floatAttributes["min_valence"] = valence
	return ta
}

// TargetValence sets the target valence.
// Valence is a measure from 0.0 to 1.0 describing the musical positiveness
/// conveyed by a track.
// Tracks with high valence sound more positive (e.g. happy, cheerful, euphoric),
// while tracks with low valence sound more negative (e.g. sad, depressed, angry).
func (ta *TrackAttributes) TargetValence(valence float64) *TrackAttributes {
	ta.floatAttributes["target_valence"] = valence
	return ta
}
