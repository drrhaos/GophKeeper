package tuiclient

import (
	"fmt"
	"image"

	"gophkeeper/internal/client/configure"
	"gophkeeper/internal/client/grpcclient"

	"github.com/marcusolsson/tui-go"
)

// FormLogin описывает структуру графического интерфейсаc входа пользователя.
type FormLogin struct {
	ui   tui.UI
	grid *tui.Grid

	loginEdit    *tui.TextEdit
	passwordEdit *tui.TextEdit

	signInButton *tui.Button
	signUpButton *tui.Button
	exitButton   *tui.Button

	labelInfo *tui.Label

	Cli *grpcclient.GRPCClient
	cfg configure.Config
}

// NewFormLogin создает форму авторизации.
func (fm *FormLogin) NewFormLogin(cfg configure.Config, buildVersion string, buildDate string) {
	fm.cfg = cfg
	gridFields := tui.NewGrid(1, 2)
	gridFields.SetBorder(true)

	fm.loginEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{0, 0}, fm.loginEdit)
	fm.passwordEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{0, 1}, fm.passwordEdit)

	boxButton := tui.NewHBox()
	boxButton.SetBorder(true)
	fm.signInButton = tui.NewButton("  Войти  ")
	boxButton.Append(fm.signInButton)
	fm.signUpButton = tui.NewButton("  Зарегистрироваться  ")
	boxButton.Append(fm.signUpButton)
	fm.exitButton = tui.NewButton("  Закрыть  ")
	boxButton.Append(fm.exitButton)

	fm.labelInfo = tui.NewLabel("")

	sidebarList := tui.NewVBox(tui.NewLabel(fmt.Sprintf("GophKeeper версия: %s дата сборки: %s", buildVersion, buildDate)), gridFields, fm.labelInfo, tui.NewSpacer(), boxButton)
	sidebarList.SetBorder(true)

	var err error
	fm.ui, err = tui.New(sidebarList)
	if err != nil {
		panic(err)
	}

	var chainF tui.SimpleFocusChain
	chainF.Set(
		fm.loginEdit,
		fm.passwordEdit,
		fm.signInButton,
		fm.signUpButton,
		fm.exitButton,
	)
	fm.ui.SetFocusChain(&chainF)

	fm.ui.SetKeybinding("Esc", func() { fm.ui.Quit() })
	fm.signInButton.OnActivated(fm.signIn)
	fm.signUpButton.OnActivated(fm.signUp)
	fm.exitButton.OnActivated(func(_ *tui.Button) {
		fm.ui.Quit()
	})

	if err := fm.ui.Run(); err != nil {
		panic(err)
	}
}

func (fm *FormLogin) signIn(_ *tui.Button) {
	var err error
	fm.Cli, err = grpcclient.Connect(fm.cfg, fm.loginEdit.Text(), fm.passwordEdit.Text())
	if err != nil {
		fm.labelInfo.SetText("Не вереное имя пользователя или пароль")
		return
	}
	fm.ui.Quit()
}

func (fm *FormLogin) signUp(_ *tui.Button) {
	var err error
	fm.Cli, err = grpcclient.Reg(fm.cfg, fm.loginEdit.Text(), fm.passwordEdit.Text())
	if err != nil {
		fm.labelInfo.SetText("Не удалось зарегистрировать пользователя")
		return
	}
	fm.ui.Quit()
}
