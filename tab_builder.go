package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jspc/gordon/types"
)

const charsWidthMax = 90

var (
	contentWidth float32
)

func CreateTab(url string) *container.TabItem {
	start := time.Now()

	p, err := ReadPage(url)
	if err != nil {
		return errorPage(err)
	}

	bs := binding.NewString()

	ti := container.NewTabItem(
		p.p.Title,
		container.NewVBox(
			p.StatusBar(bs),
			container.NewHBox(
				p.LinksBar(),
				p.ContentPane(),
			),
		),
	)

	bs.Set("(" + time.Now().Sub(start).String() + ")")

	return ti
}

func errorPage(err error) *container.TabItem {
	return container.NewTabItem(
		"Error",
		container.NewVBox(
			widget.NewLabelWithStyle("An Error Occurred!", fyne.TextAlignLeading, fyne.TextStyle{
				Bold: true,
			}),
			widget.NewLabel(err.Error()),
			widget.NewButtonWithIcon("Close", theme.CancelIcon(), func() {
				tabs.Remove(tabs.Selected())
			}),
		),
	)
}

func pageRefToURL(pr types.PageRef, cur string) string {
	if pr.Server == "" {
		pr.Server = cur
	}

	return fmt.Sprintf("//%s/%s", pr.Server, pr.Page)
}

func h1(s string) fyne.CanvasObject {
	t := canvas.NewText(s, theme.ForegroundColor())
	t.TextSize = 2 * t.TextSize
	t.TextStyle.Bold = true

	return t
}

func h2(s string) fyne.CanvasObject {
	t := canvas.NewText(s, theme.ForegroundColor())
	t.TextSize = 1.5 * t.TextSize
	t.TextStyle.Bold = true

	return t
}

func preamble(s string) fyne.CanvasObject {
	t := canvas.NewText(s, theme.ForegroundColor())
	t.TextStyle.Italic = true

	return t
}

func predicateString(p types.Predicate, reverse bool) string {
	switch p {
	case types.PredicateExtends:
		if reverse {
			return "extended by"
		}

		return "extends"

	case types.PredicateHasChild:
		if reverse {
			return "has parent"
		}

		return "has child"

	case types.PredicateSupercedes:
		if reverse {
			return "superceded by"
		}

		return "superceded"

	case types.PredicateSupplements:
		if reverse {
			return "supplemented by"
		}

		return "supplements"
	}

	return "related to"
}
