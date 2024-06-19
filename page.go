package main

import (
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jspc/gordon/client"
	"github.com/jspc/gordon/types"
	"github.com/mitchellh/go-wordwrap"
)

type Page struct {
	p    *types.Page
	addr client.Address
}

func ReadPage(url string) (p Page, err error) {
	p.addr, err = client.ParseAddress(url)
	if err != nil {
		return
	}

	p.p, err = client.DoRequest(types.VerbRead, p.addr)

	return
}

func (p Page) StatusBar(bs binding.String) fyne.CanvasObject {
	return container.NewHBox(
		widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
			tabs.Remove(tabs.Selected())
		}),
		widget.NewButtonWithIcon("", theme.ContentRedoIcon(), func() {
			tabs.Items[tabs.SelectedIndex()] = CreateTab(p.addr.String())
		}),
		widget.NewSeparator(),
		widget.NewLabel(p.addr.String()),
		widget.NewLabelWithData(bs),
	)
}

func (p Page) LinksBar() fyne.CanvasObject {
	vb := container.NewVBox()
	for _, cos := range [][]fyne.CanvasObject{
		p.metadata(),
		p.links(),
		p.relationships(),
		p.tags(),
		p.labels(),
	} {
		for _, co := range cos {
			vb.Add(co)
		}
	}

	return vb
}

func (p Page) ContentPane() fyne.CanvasObject {
	cp := container.NewVBox(
		h1(p.p.Title),
		preamble(p.p.Preamble),
		widget.NewSeparator(),
	)

	for _, s := range p.p.Sections {
		cp.Add(h2(s.Title))

		body := wordwrap.WrapString(s.Body, charsWidthMax)

		st := widget.NewEntry()
		st.SetText(body)
		st.Disable()
		st.MultiLine = true

		st.SetMinRowsVisible(len(strings.Split(body, "\n")))

		cp.Add(st)
		cp.Add(widget.NewSeparator())
	}

	scroller := container.NewVScroll(cp)

	size := scroller.MinSize()
	size.Width = contentWidth * 2.5

	scroller.SetMinSize(size)

	return scroller
}

func (p Page) metadata() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		widget.NewLabelWithStyle("Metadata", fyne.TextAlignCenter, fyne.TextStyle{
			Bold: true,
		}),

		widget.NewLabel("Author: " + p.p.Meta.Author),
		widget.NewLabel("Published: " + p.p.Meta.Published.Format(time.RFC1123)),

		widget.NewSeparator(),
	}
}

func (p Page) links() []fyne.CanvasObject {
	out := []fyne.CanvasObject{
		widget.NewLabelWithStyle("Links", fyne.TextAlignCenter, fyne.TextStyle{
			Bold: true,
		}),
	}

	for idx, l := range p.p.Links {
		out = append(out, (container.NewHBox(
			widget.NewLabel(strconv.Itoa(idx)),
			widget.NewButton(l.Page.String(), func() {
				t := CreateTab(pageRefToURL(l, p.addr.Server()))

				tabs.Append(t)
				tabs.Select(t)
			}),
		)))
	}

	return append(out, (widget.NewSeparator()))
}

func (p Page) relationships() []fyne.CanvasObject {
	out := []fyne.CanvasObject{
		widget.NewLabelWithStyle("Relationships", fyne.TextAlignCenter, fyne.TextStyle{
			Bold: true,
		}),
	}

	for _, r := range p.p.Relationships {
		isSub := r.Subject.Page == p.p.Meta.ID

		var button *widget.Button
		switch isSub {
		case true:
			button = widget.NewButton(r.Object.Page.String(), func() {
				tabs.Append(CreateTab(pageRefToURL(r.Object, p.addr.Server())))
			})

		default:
			button = widget.NewButton(r.Subject.Page.String(), func() {
				tabs.Append(CreateTab(pageRefToURL(r.Subject, p.addr.Server())))
			})
		}

		out = append(out, (container.NewVBox(
			widget.NewLabelWithStyle(predicateString(r.Predicate, !isSub), fyne.TextAlignLeading, fyne.TextStyle{Italic: true}),
			button,
		)))
	}

	return append(out, (widget.NewSeparator()))
}

func (p Page) tags() []fyne.CanvasObject {
	out := []fyne.CanvasObject{
		widget.NewLabelWithStyle("Tags", fyne.TextAlignCenter, fyne.TextStyle{
			Bold: true,
		}),
	}

	tags := container.NewHBox()
	for _, t := range p.p.Tags {
		tagsBackground := canvas.NewRectangle(fyne.CurrentApp().Settings().Theme().Color(TagsColour, fyne.CurrentApp().Settings().ThemeVariant()))
		c := container.NewStack(tagsBackground)
		c.Add(widget.NewLabelWithStyle(t, fyne.TextAlignLeading, fyne.TextStyle{
			Monospace: true,
		}))

		tags.Add(widget.NewSeparator())
		tags.Add(c)
	}

	return append(out, tags, widget.NewSeparator())
}

func (p Page) labels() []fyne.CanvasObject {
	out := []fyne.CanvasObject{
		widget.NewLabelWithStyle("Labels", fyne.TextAlignCenter, fyne.TextStyle{
			Bold: true,
		}),
	}

	tags := container.NewHBox()
	for k, v := range p.p.Labels {
		tagsBackground := canvas.NewRectangle(fyne.CurrentApp().Settings().Theme().Color(LabelsColour, fyne.CurrentApp().Settings().ThemeVariant()))
		c := container.NewStack(tagsBackground)
		c.Add(widget.NewLabelWithStyle(k+":"+v, fyne.TextAlignLeading, fyne.TextStyle{
			Monospace: true,
		}))

		tags.Add(widget.NewSeparator())
		tags.Add(c)
	}

	return append(out, tags, widget.NewSeparator())
}
