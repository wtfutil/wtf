package pocket

import (
	"fmt"
	"io/ioutil"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"gopkg.in/yaml.v2"
)

type Widget struct {
	view.ScrollableWidget
	view.KeyboardWidget

	settings     *Settings
	client       *Client
	items        []Item
	archivedView bool
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common),
		settings:         settings,
		client:           NewClient(settings.consumerKey, "http://localhost"),
		archivedView:     false,
	}

	widget.CommonSettings()
	widget.View.SetInputCapture(widget.InputCapture)
	widget.SetRenderFunction(widget.Render)
	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)
	widget.KeyboardWidget.SetView(widget.View)
	widget.initializeKeyboardControls()
	widget.Selected = -1
	widget.SetItemCount(0)
	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Render() {

	widget.Redraw(widget.content)
}

func (widget *Widget) Refresh() {
	if widget.client.accessToken == nil {
		metaData, err := readMetaDataFromDisk()
		if err != nil || metaData.AccessToken == nil {
			widget.Redraw(widget.authorizeWorkFlow)
			return
		}
		widget.client.accessToken = metaData.AccessToken
	}

	state := Unread
	if widget.archivedView == true {
		state = Read
	}
	response, err := widget.client.GetLinks(state)
	if err != nil {
		widget.SetItemCount(0)
	}

	widget.items = orderItemResponseByKey(response)
	widget.SetItemCount(len(widget.items))
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

type pocketMetaData struct {
	AccessToken *string
}

func writeMetaDataToDisk(metaData pocketMetaData) error {

	fileData, err := yaml.Marshal(metaData)
	if err != nil {
		return fmt.Errorf("Could not write token to disk %w", err)
	}

	wtfConfigDir, err := cfg.WtfConfigDir()

	if err != nil {
		return nil
	}

	filePath := fmt.Sprintf("%s/%s", wtfConfigDir, "pocket.data")
	err = ioutil.WriteFile(filePath, fileData, 0644)

	return err
}

func readMetaDataFromDisk() (pocketMetaData, error) {
	wtfConfigDir, err := cfg.WtfConfigDir()
	var metaData pocketMetaData
	if err != nil {
		return metaData, err
	}
	filePath := fmt.Sprintf("%s/%s", wtfConfigDir, "pocket.data")
	fileData, err := utils.ReadFileBytes(filePath)

	if err != nil {
		return metaData, err
	}

	err = yaml.Unmarshal(fileData, &metaData)

	return metaData, err

}

/*
	Authorization workflow is documented at https://getpocket.com/developer/docs/authentication
	broken to 4 steps :
		1- Obtain a platform consumer key from http://getpocket.com/developer/apps/new.
		2- Obtain a request token
		3- Redirect user to Pocket to continue authorization
		4- Receive the callback from Pocket, this wont be used
		5- Convert a request token into a Pocket access token
*/
func (widget *Widget) authorizeWorkFlow() (string, string, bool) {
	title := widget.CommonSettings().Title

	if widget.settings.requestKey == nil {
		requestToken, err := widget.client.ObtainRequestToken()

		if err != nil {
			logger.Log(err.Error())
			return title, err.Error(), true
		}
		widget.settings.requestKey = &requestToken
		redirectURL := widget.client.CreateAuthLink(requestToken)
		content := fmt.Sprintf("Please click on %s to Authorize the app", redirectURL)
		return title, content, true
	}

	if widget.settings.accessToken == nil {
		accessToken, err := widget.client.GetAccessToken(*widget.settings.requestKey)
		if err != nil {
			logger.Log(err.Error())
			redirectURL := widget.client.CreateAuthLink(*widget.settings.requestKey)
			content := fmt.Sprintf("Please click on %s to Authorize the app", redirectURL)
			return title, content, true
		}
		content := "Authorized"
		widget.settings.accessToken = &accessToken

		metaData := pocketMetaData{
			AccessToken: &accessToken,
		}

		err = writeMetaDataToDisk(metaData)
		if err != nil {
			content = err.Error()
		}

		return title, content, true
	}

	content := "Authorized"
	return title, content, true

}

func (widget *Widget) toggleView() {
	widget.archivedView = !widget.archivedView
	widget.Refresh()
}

func (widget *Widget) openLink() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.items != nil && sel < len(widget.items) {
		item := &widget.items[sel]
		utils.OpenFile(item.GivenURL)
	}
}

func (widget *Widget) toggleLink() {
	sel := widget.GetSelected()
	action := Archive
	if widget.archivedView == true {
		action = ReAdd
	}

	if sel >= 0 && widget.items != nil && sel < len(widget.items) {
		item := &widget.items[sel]
		_, err := widget.client.ModifyLink(action, item.ItemID)
		if err != nil {
			logger.Log(err.Error())
		}
	}

	widget.Refresh()
}

func (widget *Widget) formatItem(item Item, isSelected bool) string {
	foreColor, backColor := widget.settings.common.Colors.RowTheme.EvenForeground, widget.settings.common.Colors.RowTheme.EvenBackground
	text := item.ResolvedTitle
	if isSelected == true {
		foreColor = widget.settings.common.Colors.RowTheme.HighlightedForeground
		backColor = widget.settings.common.Colors.RowTheme.HighlightedBackground

	}

	return fmt.Sprintf("[%s:%s]%s[white]", foreColor, backColor, tview.Escape(text))
}

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	currentViewTitle := "Reading List"
	if widget.archivedView == true {
		currentViewTitle = "Archived list"
	}

	title = fmt.Sprintf("%s-%s", title, currentViewTitle)
	content := ""

	for i, v := range widget.items {
		content += widget.formatItem(v, i == widget.Selected) + "\n"
	}

	return title, content, false
}
