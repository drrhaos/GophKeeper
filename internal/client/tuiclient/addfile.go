package tuiclient

import (
	"image"

	"github.com/marcusolsson/tui-go"
)

// FormAddPath описывает структуру графического интерфейс добавления файла.
type FormAddPath struct {
	ui tui.UI

	pathEdit *tui.TextEdit

	saveButton   *tui.Button
	cancelButton *tui.Button

	labelPath *tui.Label
}

// NewFormAddPath создает форму добавления файла.
func (fm *FormAddPath) NewFormAddPath() {
	gridFields := tui.NewGrid(1, 2)
	gridFields.SetBorder(true)

	fm.pathEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{0, 0}, fm.pathEdit)
	fm.labelPath = tui.NewLabel("Путь до файла")
	pathBox := tui.NewHBox(fm.labelPath, fm.pathEdit)
	pathBox.SetBorder(true)

	boxButton := tui.NewHBox()
	boxButton.SetBorder(true)
	fm.saveButton = tui.NewButton("  Добавить  ")
	boxButton.Append(fm.saveButton)
	fm.cancelButton = tui.NewButton("  Отменить  ")
	boxButton.Append(fm.cancelButton)

	sidebarList := tui.NewVBox(pathBox, tui.NewSpacer(), boxButton)
	sidebarList.SetBorder(true)

	var err error
	fm.ui, err = tui.New(sidebarList)
	if err != nil {
		panic(err)
	}

	var chainF tui.SimpleFocusChain
	chainF.Set(
		fm.pathEdit,
		fm.saveButton,
		fm.cancelButton,
	)
	fm.ui.SetFocusChain(&chainF)

	fm.ui.SetKeybinding("Esc", func() { fm.ui.Quit() })
	// fm.signInButton.OnActivated(fm.signIn)
	// fm.signUpButton.OnActivated(fm.signUp)
	// fm.exitButton.OnActivated(func(_ *tui.Button) {
	// fm.ui.Quit()
	// })

	if err := fm.ui.Run(); err != nil {
		panic(err)
	}
}

// func (fm *FormLogin) signIn(_ *tui.Button) {
// 	var err error
// 	fm.Cli, err = grpcclient.Connect(fm.cfg, fm.loginEdit.Text(), fm.passwordEdit.Text())
// 	if err != nil {
// 		fm.labelInfo.SetText("Не вереное имя пользователя или пароль")
// 		return
// 	}
// 	fm.ui.Quit()
// }

// func (fm *FormLogin) signUp(_ *tui.Button) {
// 	var err error
// 	fm.Cli, err = grpcclient.Reg(fm.cfg, fm.loginEdit.Text(), fm.passwordEdit.Text())
// 	if err != nil {
// 		fm.labelInfo.SetText("Не удалось зарегистрировать пользователя")
// 		return
// 	}
// 	fm.ui.Quit()
// }
