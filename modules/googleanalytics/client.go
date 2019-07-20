package googleanalytics

import (
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
	"time"

	"github.com/wtfutil/wtf/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	ga "google.golang.org/api/analyticsreporting/v4"
)

type websiteReport struct {
    Name string
    Report *ga.GetReportsResponse
}

func (widget *Widget) Fetch() ([]websiteReport) {
	secretPath, err := utils.ExpandHomeDir(widget.settings.secretFile)
	if err != nil {
		log.Fatalf("Unable to parse secretFile path")
	}

	service, err := makeReportService(secretPath)
	if err != nil {
		log.Fatalf("Unable to create Google Analytics Reporting Service")
	}

	visitorsDataArray := getReports(service, widget.settings.viewIds, widget.settings.months)
	return visitorsDataArray
}

func makeReportService(secretPath string) (*ga.Service, error) {
	clientSecret, err := ioutil.ReadFile(secretPath)
	if err != nil {
		log.Fatalf("Unable to read secretPath. %v", err)
	}

	jwtConfig, err := google.JWTConfigFromJSON(clientSecret, ga.AnalyticsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to get config from JSON. %v", err)
	}

	var netClient *http.Client
	netClient = jwtConfig.Client(oauth2.NoContext)
	svc, err := ga.New(netClient)
	if err != nil {
		log.Fatalf("Failed to create Google Analytics Reporting Service")
	}

	return svc, err
}

func getReports(service *ga.Service, viewIds map[string]interface{}, displayedMonths int) ([]websiteReport) {
	startDate := fmt.Sprintf("%s-01", time.Now().AddDate(0, -displayedMonths+1, 0).Format("2006-01"))
	var websiteReports []websiteReport = nil

	for website, viewId := range viewIds {
		// For custom queries: https://ga-dev-tools.appspot.com/dimensions-metrics-explorer/

		req := &ga.GetReportsRequest{
			ReportRequests: []*ga.ReportRequest{
				{
					ViewId: viewId.(string),
					DateRanges: []*ga.DateRange{
						{StartDate: startDate, EndDate: "today"},
					},
					Metrics: []*ga.Metric{
						{Expression: "ga:sessions"},
					},
					Dimensions: []*ga.Dimension{
						{Name: "ga:month"},
					},
				},
			},
		}
		response, err := service.Reports.BatchGet(req).Do()

		if err != nil {
			log.Fatalf("GET request to analyticsreporting/v4 returned error with viewID: %s", viewId)
		}
		if response.HTTPStatusCode != 200 {
			log.Fatalf("Did not get expected HTTP response code")
		}

		report := websiteReport{Name: website, Report: response,}
		websiteReports = append(websiteReports, report)
	}
	return websiteReports
}
