package netmon

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/net"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget
	settings   *Settings
	interfaces map[string]*NIC
}

type NIC struct {
	Name string
	Sent uint64
	Recv uint64
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),
		settings:   settings,
	}

	widget.SetupNICs()
	widget.View.SetWrap(true)

	return &widget
}
func (widget *Widget) SetupNICs() {
	widget.interfaces = make(map[string]*NIC)
	interfaces, _ := net.IOCounters(true)

	for _, nic := range interfaces {
		if widget.settings.ignoreLoopback && strings.HasPrefix(nic.Name, "lo") ||
			widget.settings.ignoreBridges && strings.HasPrefix(nic.Name, "br") ||
			widget.settings.ignoreDocker && strings.HasPrefix(nic.Name, "docker") ||
			widget.settings.ignoreVETH && strings.HasPrefix(nic.Name, "veth") {
			continue
		}
		widget.interfaces[nic.Name] = &NIC{
			Name: nic.Name,
			Sent: nic.BytesSent,
			Recv: nic.BytesRecv}
	}
}
func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	var content strings.Builder

	interfaces, err := net.IOCounters(true)
	if err != nil {
		return widget.CommonSettings().Title, err.Error(), true
	}
	for _, nic := range interfaces {
		if _, ok := widget.interfaces[nic.Name]; ok {
			prevSent := widget.interfaces[nic.Name].Sent
			prevRecv := widget.interfaces[nic.Name].Recv
			fmt.Fprintf(&content, "%s:\t%5v▲ \t%5v▼\n", nic.Name, (nic.BytesSent-prevSent)/1024, (nic.BytesRecv-prevRecv)/1024)
			widget.interfaces[nic.Name].Sent = nic.BytesSent
			widget.interfaces[nic.Name].Recv = nic.BytesRecv
		}

	}

	return widget.CommonSettings().Title, content.String(), true
}
