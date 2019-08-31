package googleanalytics

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/wtfutil/wtf/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	gaV3 "google.golang.org/api/analytics/v3"
	gaV4 "google.golang.org/api/analyticsreporting/v4"
)

type websiteReport struct {
	Name           string
	Report         *gaV4.GetReportsResponse
	RealtimeReport *gaV3.RealtimeData
}

func (widget *Widget) Fetch() []websiteReport {
	secretPath, err := utils.ExpandHomeDir(widget.settings.secretFile)
	if err != nil {
		log.Fatalf("Unable to parse secretFile path")
	}

	serviceV4, err := makeReportServiceV4(secretPath)
	if err != nil {
		log.Fatalf("Unable to create v3 Google Analytics Reporting Service")
	}

	var serviceV3 *gaV3.Service
	if widget.settings.enableRealtime {
		serviceV3, err = makeReportServiceV3(secretPath)
		if err != nil {
			log.Fatalf("Unable to create v3 Google Analytics Reporting Service")
		}
	}

	visitorsDataArray := getReports(
		serviceV4, widget.settings.viewIds, widget.settings.months, serviceV3,
	)
	return visitorsDataArray
}

func buildNetClient(secretPath string) *http.Client {
	clientSecret, err := ioutil.ReadFile(secretPath)
	if err != nil {
		log.Fatalf("Unable to read secretPath. %v", err)
	}

	jwtConfig, err := google.JWTConfigFromJSON(clientSecret, gaV4.AnalyticsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to get config from JSON. %v", err)
	}

	return jwtConfig.Client(oauth2.NoContext)
}

func makeReportServiceV3(secretPath string) (*gaV3.Service, error) {
	var netClient = buildNetClient(secretPath)
	svc, err := gaV3.New(netClient)
	if err != nil {
		log.Fatalf("Failed to create v3 Google Analytics Reporting Service")
	}

	return svc, err
}

func makeReportServiceV4(secretPath string) (*gaV4.Service, error) {
	var netClient = buildNetClient(secretPath)
	svc, err := gaV4.New(netClient)
	if err != nil {
		log.Fatalf("Failed to create v4 Google Analytics Reporting Service")
	}

	return svc, err
}

func getReports(
	serviceV4 *gaV4.Service, viewIds map[string]interface{}, displayedMonths int, serviceV3 *gaV3.Service,
) []websiteReport {
	startDate := fmt.Sprintf("%s-01", time.Now().AddDate(0, -displayedMonths+1, 0).Format("2006-01"))
	var websiteReports []websiteReport

	for website, viewID := range viewIds {
		// For custom queries: https://ga-dev-tools.appspot.com/dimensions-metrics-explorer/

		req := &gaV4.GetReportsRequest{
			ReportRequests: []*gaV4.ReportRequest{
				{
					ViewId: viewID.(string),
					DateRanges: []*gaV4.DateRange{
						{StartDate: startDate, EndDate: "today"},
					},
					Metrics: []*gaV4.Metric{
						{Expression: "ga:sessions"},
					},
					Dimensions: []*gaV4.Dimension{
						{Name: "ga:month"},
					},
				},
			},
		}
		response, err := serviceV4.Reports.BatchGet(req).Do()

		if err != nil {
			log.Fatalf("GET request to analyticsreporting/v4 returned error with viewID: %s", viewID)
		}
		if response.HTTPStatusCode != 200 {
			log.Fatalf("Did not get expected HTTP response code")
		}

		report := websiteReport{Name: website, Report: response}
		if serviceV3 != nil {
			report.RealtimeReport = getLiveCount(serviceV3, viewID.(string))
		}
		websiteReports = append(websiteReports, report)
	}
	return websiteReports
}

func getLiveCount(service *gaV3.Service, viewID string) *gaV3.RealtimeData {
	res, err := service.Data.Realtime.Get("ga:"+viewID, "rt:activeUsers").Do()
	if err != nil {
		log.Fatalf("Failed to fetch real time data for view ID %s: %v.  Have you enrolled in the real time beta?  If not, do so here: https://docs.google.com/forms/d/1qfRFysCikpgCMGqgF3yXdUyQW4xAlLyjKuOoOEFN2Uw/viewform", viewID, err)
	}

	return res
}
