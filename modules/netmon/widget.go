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
			widget.settings.ignoreEthernet && strings.HasPrefix(nic.Name, "e") ||
			widget.settings.ignoreWireless && strings.HasPrefix(nic.Name, "wl") ||
			widget.settings.ignoreBridges && strings.HasPrefix(nic.Name, "br") ||
			widget.settings.ignoreDocker && strings.HasPrefix(nic.Name, "docker") ||
			widget.settings.ignoreVeth && strings.HasPrefix(nic.Name, "veth") {
			continue
		}
		if widget.settings.showOnly != "" {
			if widget.settings.showOnly == nic.Name {
				widget.addInterface(nic)
				break
			}
		} else {
			widget.addInterface(nic)
		}

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
		if NIC, ok := widget.interfaces[nic.Name]; ok {
			prevSent, prevRecv := NIC.Sent, NIC.Recv
			NIC.Sent, NIC.Recv = nic.BytesSent, nic.BytesRecv
			fmt.Fprintf(&content, "%s:\t▲%9s \t▼%9s\n",
				nic.Name,
				pretty(nic.BytesSent-prevSent),
				pretty(nic.BytesRecv-prevRecv))
		}
	}
	return widget.CommonSettings().Title, content.String(), true
}

func (widget *Widget) addInterface(nic net.IOCountersStat) {
	widget.interfaces[nic.Name] = &NIC{
		Name: nic.Name,
		Sent: nic.BytesSent,
		Recv: nic.BytesRecv,
	}
}

func pretty(bytes uint64) string {
	bits := bytes * 8
	bps := "bps"
	if bits > 1000000 {
		bits /= 1000000
		bps = "M" + bps
	} else if bits > 1000 {
		bits /= 1000
		bps = "K" + bps
	}
	return fmt.Sprintf("%v %s", bits, bps)
}
