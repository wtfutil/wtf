package helix

import "time"

// UserBitTotal ...
type UserBitTotal struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Rank     int    `json:"rank"`
	Score    int    `json:"score"`
}

// ManyUserBitTotals ...
type ManyUserBitTotals struct {
	Total         int            `json:"total"`
	DateRange     DateRange      `json:"date_range"`
	UserBitTotals []UserBitTotal `json:"data"`
}

// BitsLeaderboardResponse ...
type BitsLeaderboardResponse struct {
	ResponseCommon
	Data ManyUserBitTotals
}

// BitsLeaderboardParams ...
type BitsLeaderboardParams struct {
	Count     int       `query:"count,10"`   // Maximum 100
	Period    string    `query:"period,all"` // "all" (default), "day", "week", "month" and "year"
	StartedAt time.Time `query:"started_at"`
	UserID    string    `query:"user_id"`
}

// GetBitsLeaderboard gets a ranked list of Bits leaderboard
// information for an authorized broadcaster.
//
// Required Scope: bits:read
func (c *Client) GetBitsLeaderboard(params *BitsLeaderboardParams) (*BitsLeaderboardResponse, error) {
	resp, err := c.get("/bits/leaderboard", &ManyUserBitTotals{}, params)
	if err != nil {
		return nil, err
	}

	bits := &BitsLeaderboardResponse{}
	bits.StatusCode = resp.StatusCode
	bits.Header = resp.Header
	bits.Error = resp.Error
	bits.ErrorStatus = resp.ErrorStatus
	bits.ErrorMessage = resp.ErrorMessage
	bits.Data.Total = resp.Data.(*ManyUserBitTotals).Total
	bits.Data.DateRange = resp.Data.(*ManyUserBitTotals).DateRange
	bits.Data.UserBitTotals = resp.Data.(*ManyUserBitTotals).UserBitTotals

	return bits, nil
}
