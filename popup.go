package popup

import (
	"github.com/mmogo/gxui"
	"github.com/mmogo/gxui/drivers/gl"
	"github.com/mmogo/gxui/gxfont"
	"github.com/mmogo/gxui/themes/basic"
	"github.com/mmogo/gxui/themes/dark"
)

// Popup with title and text (And "OK" to close)
type Popup struct {
	Title  string
	Text   string
	IsBool bool
	Return bool
}

// P ...
func P(p *Popup) {
	gl.StartDriver(p.TextPopup)
}

// Message creates a popup window message
func Message(title, msg string) {
	p := new(Popup)
	p.Title = title
	p.Text = msg
	gl.StartDriver(p.TextPopup)
}

// Error creates a popup window titled "error"
func Error(err error) {
	p := new(Popup)
	p.Title = "Error"
	p.Text = err.Error()
	gl.StartDriver(p.TextPopup)
}

func Bool(title, msg string) bool {
	p := new(Popup)
	p.Title = title
	p.Text = msg
	p.IsBool = true
	gl.StartDriver(p.TextPopup)
	return p.Return
}

func (p *Popup) TextPopup(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	splitter := theme.CreateSplitterLayout()

	// header
	if p.Title != "" {
		header := theme.CreateLabel()
		font, err := driver.CreateFont(gxfont.Default, 75)
		if err != nil {
			panic(err)
		}
		header.SetFont(font)
		header.SetText(p.Title)
		splitter.AddChild(header)
	}
	// error box
	errbox := basic.CreateTextBox(theme.(*basic.Theme))
	errbox.SetDesiredWidth(200)
	errbox.SetText(p.Text)
	splitter.AddChild(errbox)

	// window
	window := theme.CreateWindow(200, 300, p.Title)

	if p.IsBool {
		trueBtn := theme.CreateButton()
		trueBtn.SetText("Yes")
		trueBtn.OnClick(func(gxui.MouseEvent) {
			p.Return = true
			window.Close()
			return
		})
		splitter.AddChild(trueBtn)
		falseBtn := theme.CreateButton()
		falseBtn.SetText("No")
		falseBtn.OnClick(func(gxui.MouseEvent) {
			p.Return = false
			window.Close()
			return
		})
		splitter.AddChild(falseBtn)
	} else {
		// ok button
		closeBtn := theme.CreateButton()
		closeBtn.SetText("OK")
		closeBtn.OnClick(func(gxui.MouseEvent) {
			window.Close()
			return
		})
		splitter.AddChild(closeBtn)
	}

	window.AddChild(splitter)
	window.OnClose(driver.Terminate)
}
