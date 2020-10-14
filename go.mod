module github.com/wtfutil/wtf

go 1.15

require (
	code.cloudfoundry.org/bytefmt v0.0.0-20190819182555-854d396b647c
	github.com/Azure/go-autorest v11.1.2+incompatible // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/PagerDuty/go-pagerduty v0.0.0-20191002190746-f60f4fc45222
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/VictorAvelar/devto-api-go v1.0.0
	github.com/adlio/trello v1.8.0
	github.com/alecthomas/chroma v0.8.1
	github.com/andygrunwald/go-gerrit v0.0.0-20190825170856-5959a9bf9ff8
	github.com/briandowns/openweathermap v0.0.0-20180804155945-5f41b7c9d92d
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/creack/pty v1.1.11
	github.com/digitalocean/godo v1.46.0
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/docker-credential-helpers v0.6.3
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dustin/go-humanize v1.0.0
	github.com/gdamore/tcell v1.4.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/godbus/dbus v4.1.0+incompatible // indirect
	github.com/google/go-github/v32 v32.1.0
	github.com/gophercloud/gophercloud v0.5.0 // indirect
	github.com/hekmon/cunits v2.0.1+incompatible // indirect
	github.com/hekmon/transmissionrpc v0.0.0-20190525133028-1d589625bacd
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/lib/pq v1.2.0 // indirect
	github.com/logrusorgru/aurora v0.0.0-20190803045625-94edacc10f9b
	github.com/microsoft/azure-devops-go-api/azuredevops v0.0.0-20191014190507-26902c1d4325
	github.com/mmcdole/gofeed v1.1.0
	github.com/nicklaw5/helix v0.7.0
	github.com/olebedev/config v0.0.0-20190528211619-364964f3a8e4
	github.com/olekukonko/tablewriter v0.0.4
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/ovh/cds v0.0.0-20201014170613-39429542624d
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.5.0
	github.com/radovskyb/watcher v1.0.7
	github.com/rivo/tview v0.0.0-20200108161608-1316ea7a4b35
	github.com/shirou/gopsutil v2.20.9+incompatible
	github.com/shurcooL/githubv4 v0.0.0-20200802174311-f27d2ca7f6d5
	github.com/shurcooL/graphql v0.0.0-20181231061246-d48a9a75455f // indirect
	github.com/spf13/cobra v0.0.5 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.1 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	github.com/stretchr/testify v1.6.1
	github.com/wtfutil/spotigopher v0.0.0-20191127141047-7d8168fe103a
	github.com/wtfutil/todoist v0.0.2-0.20191216004217-0ec29ceda61a
	github.com/xanzy/go-gitlab v0.38.1
	github.com/zmb3/spotify v0.0.0-20191010212056-e12fb981aacb
	github.com/zorkian/go-datadog-api v2.29.0+incompatible
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/text v0.3.3
	google.golang.org/api v0.33.0
	gopkg.in/jarcoal/httpmock.v1 v1.0.0-20181110093347-3be5f16b70eb // indirect
	gopkg.in/yaml.v2 v2.3.0
	gotest.tools v2.2.0+incompatible
	jaytaylor.com/html2text v0.0.0-20200412013138-3577fbdbcff7
	k8s.io/apimachinery v0.0.0-20190223094358-dcb391cde5ca
	k8s.io/client-go v10.0.0+incompatible
)

// These hacks are in place to work around this bug in coreos/etcd/proxy/grpcproxy > v1.30.0 that fails
// with this error:
//
//    google.golang.org/grpc/naming: module google.golang.org/grpc@latest found (v1.31.1), but does not contain package google.golang.org/grpc/naming
//
// See here for more details: https://github.com/etcd-io/etcd/issues/12124
replace google.golang.org/grpc v1.30.0 => google.golang.org/grpc v1.29.1

replace google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.29.1
