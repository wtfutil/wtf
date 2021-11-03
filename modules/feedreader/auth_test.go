package feedreader

import (
	"net/http"
	"testing"

	"github.com/olebedev/config"
	"gotest.tools/assert"
)

var listformat string = `
enabled: true
feeds:
  - https://news.ycombinator.com/rss
  - https://rss.cbc.ca/lineup/topstories.xml

feedLimit: 10
position:
  top: 1
  left: 1
  width: 2
  height: 1
  refreshInterval: 14400`

var mapformat string = `
enabled: true
feeds:
  http://localhost:3000/feed/rss.xml:
    username: johndoe
    password: supersecret

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

func Test_MapFormatFetch(t *testing.T) {
	ymlCfg, _ := config.ParseYaml(mapformat)
	globalCfg, _ := config.ParseYaml(globalConfig)

	settings := NewSettingsFromYAML("feedreader", ymlCfg, globalCfg)
	widget := NewWidget(nil, nil, settings)

	go setupPrivateRSSFeed()

	f, err := widget.Fetch(widget.settings.feeds)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(f), 2)
}

func Test_ListFormatFetch(t *testing.T) {
	ymlCfg, _ := config.ParseYaml(listformat)
	globalCfg, _ := config.ParseYaml(globalConfig)

	settings := NewSettingsFromYAML("feedreader", ymlCfg, globalCfg)
	widget := NewWidget(nil, nil, settings)

	go setupPrivateRSSFeed()

	f, err := widget.Fetch(widget.settings.feeds)
	if err != nil {
		t.Error(err)
	}

	if len(f) == 0 {
		t.Error("No articles fetched")
	}
}

func setupPrivateRSSFeed() {
	http.HandleFunc("/feed/rss.xml", func(w http.ResponseWriter, r *http.Request) {
		username, password, exists := r.BasicAuth()
		if !exists || username != "johndoe" || password != "supersecret" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/xml")

		_, err := w.Write([]byte(sampleRSS))
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
