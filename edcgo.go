package main

import (
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/walk"
	"time"
	"fmt"
)

type EdcgoMainWindow struct {
	*walk.MainWindow
	flightLogPathLabel *walk.Label
	eventLabel         *walk.Label
	flightLog *FlightLog
}

func (mw *EdcgoMainWindow) init() *MainWindow {
	flightLogPath := getDefaultFlightLogPath()
	mw.flightLogPathLabel = new(walk.Label)
	mw.flightLogPathLabel.SetText(flightLogPath)
	mw.eventLabel = new(walk.Label)
	mw.eventLabel.SetText("None")
	mw.flightLog = new(FlightLog)
	mw.flightLog.path = flightLogPath
	mw.flightLog.updateLatestFile()

	return &MainWindow{
		AssignTo: &mw.MainWindow,
		Title: "Elite: Dangerous  Companion",
		//MinSize: Size{100, 400},
		Layout: VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text:               "Flight log path:",
					},
					Label{
						AssignTo: &mw.flightLogPathLabel,
						Text:     flightLogPath,
					},
					HSpacer{},
				},
			},
			Composite{
				Layout:HBox{},
				Children: []Widget{
					Label{Text:"Event:"},
					Label{AssignTo:&mw.eventLabel, Text:"None"},
					HSpacer{},
				},
			},
			VSpacer{},
		},
		ToolBar: ToolBar{
			Items: []MenuItem{
			},
		},
	}
}

func (mw *EdcgoMainWindow) getJournalPath() string {
	return mw.flightLogPathLabel.Text()
}

func (mw *EdcgoMainWindow) setEvent(event string) {
	mw.eventLabel.SetText(event)
}

func loop(fl *FlightLog) {
	for true {
		//if fl.isUpdated() {
		//	flight_log_file := fl.getLatestLog()
		//	fl.parseLogFile(flight_log_file)
		//}
		time.Sleep(time.Second)
	}
}

func main() {
	//main_window := new(EdcgoMainWindow)
	//mw := main_window.init()
	//go loop(main_window.flightLog)
	//mw.Run()
	fl := new(FlightLog)
	fl.path = getDefaultFlightLogPath()
	fmt.Println(fl.updateLatestFile())
}
