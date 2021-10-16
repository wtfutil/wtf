package app

import (
	"net/http"
	"testing"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/modules/feedreader"
	"gotest.tools/assert"
)

var feedreaderConfig string = `
enabled: true
feeds:
- http://localhost:3000/feed/rss.xml
auth:
- ["user", "password"]
- ["john doe", "supersecret"]

feedLimit: 10
position:
  top: 1
  left: 1
  width: 2
  height: 1
  refreshInterval: 14400`

var globalConfig string = `
wtf:
  colors:
    border:
      focusable: darkslateblue
      focused: orange
      normal: gray`

var sampleRSS string = `<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0">

<channel>
  <title>The Stoics</title>
  <link>https://google.com</link>
  <description>Testing feed to test private feeds</description>
  <item>
  <title>Marcus Aurelius</title>
  <link>https://en.wikipedia.org/wiki/Marcus_Aurelius</link>
  <description>
    "The soul becomes dyed with the color of its thoughts."
  </description>
  </item>
  <item>
  <title>Seneca</title>
  <link>https://en.wikipedia.org/wiki/Seneca_the_Younger</link>
  <description>"We suffer more often in imagination than in reality"</description>
  </item>
</channel>

</rss>
`

func setupPrivateRSSFeed() {
	http.HandleFunc("/feed/rss.xml", func(w http.ResponseWriter, r *http.Request) {
		username, password, exists := r.BasicAuth()
		if !exists || username != "john doe" || password != "supersecret" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(sampleRSS))
	})
	http.ListenAndServe(":3000", nil)
}

func TestFeedreaderAuth(t *testing.T) {
	ymlCfg, _ := config.ParseYaml(feedreaderConfig)
	globalCfg, _ := config.ParseYaml(globalConfig)

	settings := feedreader.NewSettingsFromYAML("feedreader", ymlCfg, globalCfg)
	widget := feedreader.NewWidget(nil, nil, settings)

	go setupPrivateRSSFeed()

	f, err := widget.Fetch([]string{"http://localhost:3000/feed/rss.xml"})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(f), 2)
}
