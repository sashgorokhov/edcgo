package main

import "edcgo/edjournal"
import (
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/walk"
)

type EdcgoMainWindow struct {
	*walk.MainWindow
	journalPathLabel *walk.Label
}

func (mw *EdcgoMainWindow) init() {
	journalPath, _ := edjournal.GetDefaultJournalPath()
	mw.journalPathLabel = new(walk.Label)
	mw.journalPathLabel.SetText(journalPath)
}

func loop() {
}

func main() {

	mainWindow := new(EdcgoMainWindow);
	mainWindow.init()

	go loop()

	mw := MainWindow{
		AssignTo: &mainWindow.MainWindow,
		Title: "Elite: Dangerous  Companion",
		//MinSize: Size{100, 400},
		Layout: VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text:               "Journal path:",
					},
					Label{
						AssignTo: &mainWindow.journalPathLabel,
						Text: mainWindow.journalPathLabel.Text(),
					},
					PushButton{
						Text: "change",
						OnClicked: func(){
							dlg := new(walk.FileDialog)

							dlg.Title = "Select Elite: Dangerous journal log forlder"
							dlg.InitialDirPath = mainWindow.journalPathLabel.Text()

							if ok, err := dlg.ShowBrowseFolder(mainWindow); err != nil {
								return
							} else if !ok {
								return
							}

							mainWindow.journalPathLabel.SetText(dlg.FilePath)
						},
					},
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

	mw.Run()

}
