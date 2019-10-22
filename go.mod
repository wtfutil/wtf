module github.com/wtfutil/wtf

go 1.12

require (
	code.cloudfoundry.org/bytefmt v0.0.0-20190819182555-854d396b647c
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/PagerDuty/go-pagerduty v0.0.0-20191002190746-f60f4fc45222
	github.com/PuerkitoBio/goquery v1.5.0 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/VictorAvelar/devto-api-go v1.0.0
	github.com/adlio/trello v1.4.0
	github.com/alecthomas/chroma v0.6.7
	github.com/andygrunwald/go-gerrit v0.0.0-20190825170856-5959a9bf9ff8
	github.com/briandowns/openweathermap v0.0.0-20180804155945-5f41b7c9d92d
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/gdamore/tcell v1.3.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/godbus/dbus v4.1.0+incompatible // indirect
	github.com/google/go-github/v26 v26.1.3
	github.com/gophercloud/gophercloud v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hekmon/cunits v2.0.1+incompatible // indirect
	github.com/hekmon/transmissionrpc v0.0.0-20190525133028-1d589625bacd
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/logrusorgru/aurora v0.0.0-20190803045625-94edacc10f9b
	github.com/microsoft/azure-devops-go-api/azuredevops v0.0.0-20191014190507-26902c1d4325
	github.com/mmcdole/gofeed v1.0.0-beta2
	github.com/mmcdole/goxpp v0.0.0-20181012175147-0068e33feabf // indirect
	github.com/olebedev/config v0.0.0-20190528211619-364964f3a8e4
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pkg/profile v1.3.0
	github.com/radovskyb/watcher v1.0.7
	github.com/rivo/tview v0.0.0-20190829161255-f8bc69b90341
	github.com/shirou/gopsutil v2.19.9+incompatible
	github.com/sticreations/spotigopher v0.0.0-20181009182052-98632f6f94b0
	github.com/stretchr/testify v1.4.0
	github.com/wtfutil/todoist v0.0.1
	github.com/xanzy/go-gitlab v0.20.1
	github.com/yfronto/newrelic v0.0.0-00010101000000-000000000000
	github.com/zmb3/spotify v0.0.0-20191010212056-e12fb981aacb
	github.com/zorkian/go-datadog-api v2.24.0+incompatible
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.11.0
	gopkg.in/jarcoal/httpmock.v1 v1.0.0-20181110093347-3be5f16b70eb // indirect
	gopkg.in/yaml.v2 v2.2.4
	k8s.io/api v0.0.0-20191010143144-fbf594f18f80 // indirect
	k8s.io/apimachinery v0.0.0-20191016060620-86f2f1b9c076
	k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
	k8s.io/utils v0.0.0-20191010214722-8d271d903fe4 // indirect
)

replace github.com/yfronto/newrelic => ./vendor/github.com/yfronto/newrelic
