package dash

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func LoadUi() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "DashBoard!"
	p.SetRect(0, 0, 25, 5)

	ui.Render(p)

	// bc := widgets.NewBarChart()
	// bc.Data = []float64{3, 2, 5, 3, 9, 3}
	// bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	// bc.Title = "Bar Chart"
	// bc.SetRect(50, 5, 100, 25)
	// bc.BarWidth = 5
	// bc.BarColors = []ui.Color{ui.ColorRed, ui.ColorGreen}
	// bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	// bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}
	// ui.Render(bc)

	// g0 := widgets.NewGauge()
	// g0.Title = "Slim Gauge"
	// g0.SetRect(20, 20, 30, 30)
	// g0.Percent = 75
	// g0.BarColor = ui.ColorRed
	// g0.BorderStyle.Fg = ui.ColorWhite
	// g0.TitleStyle.Fg = ui.ColorCyan

	// g2 := widgets.NewGauge()
	// g2.Title = "Slim Gauge"
	// g2.SetRect(0, 3, 50, 6)
	// g2.Percent = 60
	// g2.BarColor = ui.ColorYellow
	// g2.LabelStyle = ui.NewStyle(ui.ColorBlue)
	// g2.BorderStyle.Fg = ui.ColorWhite

	// g1 := widgets.NewGauge()
	// g1.Title = "Big Gauge"
	// g1.SetRect(0, 6, 50, 11)
	// g1.Percent = 30
	// g1.BarColor = ui.ColorGreen
	// g1.LabelStyle = ui.NewStyle(ui.ColorYellow)
	// g1.TitleStyle.Fg = ui.ColorMagenta
	// g1.BorderStyle.Fg = ui.ColorWhite

	// g3 := widgets.NewGauge()
	// g3.Title = "Gauge with custom label"
	// g3.SetRect(0, 11, 50, 14)
	// g3.Percent = 50
	// g3.Label = fmt.Sprintf("%v%% (100MBs free)", g3.Percent)

	// g4 := widgets.NewGauge()
	// g4.Title = "Gauge"
	// g4.SetRect(0, 14, 50, 17)
	// g4.Percent = 50
	// g4.Label = "Gauge with custom highlighted label"
	// g4.BarColor = ui.ColorGreen
	// g4.LabelStyle = ui.NewStyle(ui.ColorYellow)

	// ui.Render(g0, g1, g2, g3, g4)
	
	table1 := widgets.NewTable()
	table1.Rows = [][]string{
		{"header1", "header2", "header3"},
		{"你好吗", "Go-lang is so cool", "Im working on Ruby"},
		{"2016", "10", "11"},
	}
	table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	table1.SetRect(150, 0, 60, 10)

	ui.Render(table1)

	table2 := widgets.NewTable()
	table2.Rows = [][]string{
		{"header1", "header2", "header3"},
		{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		{"2016", "11", "11"},
	}
	table2.TextStyle = ui.NewStyle(ui.ColorWhite)
	table2.TextAlignment = ui.AlignCenter
	table2.RowSeparator = false
	table2.SetRect(150, 10, 20, 20)

	ui.Render(table2)

	table3 := widgets.NewTable()
	table3.Rows = [][]string{
		{"header1", "header2", "header3"},
		{"AAA", "BBB", "CCC"},
		{"DDD", "EEE", "FFF"},
		{"GGG", "HHH", "III"},
	}
	table3.TextStyle = ui.NewStyle(ui.ColorWhite)
	table3.RowSeparator = true
	table3.BorderStyle = ui.NewStyle(ui.ColorGreen)
	table3.SetRect(150, 50, 70, 20)
	table3.FillRow = true
	table3.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table3.RowStyles[2] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
	table3.RowStyles[3] = ui.NewStyle(ui.ColorYellow)

	ui.Render(table3)

	header := widgets.NewParagraph()
	header.Text = "Press q to quit, Press h or l to switch tabs"
	header.SetRect(0, 0, 50, 1)
	header.Border = false
	header.TextStyle.Bg = ui.ColorBlue

	p2 := widgets.NewParagraph()
	p2.Text = "Press q to quit\nPress h or l to switch tabs\n"
	p2.Title = "Keys"
	p2.SetRect(5, 5, 40, 15)
	p2.BorderStyle.Fg = ui.ColorYellow

	bc := widgets.NewBarChart()
	bc.Title = "Bar Chart"
	bc.Data = []float64{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bc.SetRect(5, 5, 35, 10)
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}

	tabpane := widgets.NewTabPane("pierwszy", "drugi", "trzeci", "żółw", "four", "five")
	tabpane.SetRect(0, 1, 50, 4)
	tabpane.Border = true

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(p2)
		case 1:
			ui.Render(bc)
		}
	}

	ui.Render(header, tabpane, p2)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "l":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		}
	}

 
	// uiEvents := ui.PollEvents()
	// for {
	// 	e := <-uiEvents
	// 	switch e.ID {
	// 	case "q", "<C-c>":
	// 		return
	// 	}
	// }
}
