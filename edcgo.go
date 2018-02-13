package main

import (
	//. "github.com/lxn/walk/declarative"
	//"github.com/lxn/walk"
	"time"
	"github.com/sirupsen/logrus"
	"os"
)


var Logger = &logrus.Logger{
	Level:     logrus.DebugLevel,
	Formatter: new(logrus.TextFormatter),
	//Hooks:     make(logrus.LevelHooks),
	Out:       os.Stderr,
}

//type EdcgoMainWindow struct {
//	*walk.MainWindow
//	flightLogPathLabel *walk.Label
//	eventLabel         *walk.Label
//	flightLog *FlightLog
//}
//
//func (mw *EdcgoMainWindow) init() *MainWindow {
//	flightLogPath := getDefaultFlightLogPath()
//	mw.flightLogPathLabel = new(walk.Label)
//	mw.flightLogPathLabel.SetText(flightLogPath)
//	mw.eventLabel = new(walk.Label)
//	mw.eventLabel.SetText("None")
//	mw.flightLog = new(FlightLog)
//	mw.flightLog.path = flightLogPath
//	mw.flightLog.updateLatestFile()
//
//	return &MainWindow{
//		AssignTo: &mw.MainWindow,
//		Title: "Elite: Dangerous  Companion",
//		//MinSize: Size{100, 400},
//		Layout: VBox{},
//		Children: []Widget{
//			Composite{
//				Layout: HBox{},
//				Children: []Widget{
//					Label{
//						Text:               "Flight log path:",
//					},
//					Label{
//						AssignTo: &mw.flightLogPathLabel,
//						Text:     flightLogPath,
//					},
//					HSpacer{},
//				},
//			},
//			Composite{
//				Layout:HBox{},
//				Children: []Widget{
//					Label{Text:"Event:"},
//					Label{AssignTo:&mw.eventLabel, Text:"None"},
//					HSpacer{},
//				},
//			},
//			VSpacer{},
//		},
//		ToolBar: ToolBar{
//			Items: []MenuItem{
//			},
//		},
//	}
//}
//
//func (mw *EdcgoMainWindow) getJournalPath() string {
//	return mw.flightLogPathLabel.Text()
//}
//
//func (mw *EdcgoMainWindow) setEvent(event string) {
//	mw.eventLabel.SetText(event)
//}

func loop(fl *FlightLog) {
	config, _ := loadConfig(fl)
	for true {
		lines := process_flight_log(fl)
		if len(*lines) > 0 {
			Logger.Debugln("Lines to send", len(*lines))
			for _, line := range *lines {
				Logger.Debugln(line)
			}
			pushLines(lines, config.Token)
		}
		time.Sleep(time.Second)
	}
}

func process_flight_log(fl *FlightLog) *[]string {
	latest_log_file := fl.getLatestFile()
	if latest_log_file == "" {
		return &[]string{}
	}
	logger := Logger.WithField("filename", latest_log_file)
	fl.latest_log_file = latest_log_file
	logger.Infoln("Latest log file")
	lines := fl.getFlightLogLines(latest_log_file)
	lines_to_send := fl.getLinesToSend(lines)
	return lines_to_send
}


func main() {
	//main_window := new(EdcgoMainWindow)
	//mw := main_window.init()
	//go loop(main_window.flightLog)
	//mw.Run()
	fl := new(FlightLog)
	fl.path = getDefaultFlightLogPath()
	loop(fl)
}
