// Package tuiclient реализует TUI инфтерфейс.
package tuiclient

import (
	"image"

	"github.com/marcusolsson/tui-go"
)

// Run функция запускает TUI интерфейс.
func Run() {
	listRows := tui.NewList()
	listRows.AddItems("first")
	listRows.AddItems("two")

	gridFields := tui.NewGrid(2, 8)
	gridFields.SetBorder(true)

	gridFields.SetCell(image.Point{0, 0}, tui.NewLabel("Имя"))
	nameEdit := tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 0}, nameEdit)

	gridFields.SetCell(image.Point{0, 1}, tui.NewLabel("Логин"))
	loginEdit := tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 1}, loginEdit)

	gridFields.SetCell(image.Point{0, 2}, tui.NewLabel("Пароль"))
	passwordEdit := tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 2}, passwordEdit)

	gridFields.SetCell(image.Point{0, 3}, tui.NewLabel("Описание"))
	dataEdit := tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 3}, dataEdit)

	gridFields.SetCell(image.Point{0, 4}, tui.NewLabel("Номер карты"))
	cardNumberEdit := tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 4}, cardNumberEdit)

	gridFields.SetCell(image.Point{0, 5}, tui.NewLabel("CVC карты"))
	cardCVCEdit := tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 5}, cardCVCEdit)

	gridFields.SetCell(image.Point{0, 6}, tui.NewLabel("Дата окончания срока действия карты"))
	cardDateEdit := tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 6}, cardDateEdit)

	gridFields.SetCell(image.Point{0, 7}, tui.NewLabel("Владелец карты"))
	cardOwnerEdit := tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 7}, cardOwnerEdit)

	sidebarList := tui.NewVBox(listRows)
	sidebarList.SetBorder(true)
	fieldSidebar := tui.NewVBox(gridFields, tui.NewSpacer())
	fieldSidebar.SetBorder(true)
	fieldSidebar.SetSizePolicy(tui.Expanding, tui.Expanding)

	boxButton := tui.NewHBox()
	boxButton.SetBorder(true)
	boxButton.Append(tui.NewSpacer())
	cancelButton := tui.NewButton("Отменить")
	boxButton.Append(cancelButton)
	saveButton := tui.NewButton("Сохранить")
	boxButton.Append(saveButton)

	boxGeneral := tui.NewHBox(
		sidebarList,
		tui.NewVBox(
			fieldSidebar,
			boxButton,
		),
	)

	ui, err := tui.New(boxGeneral)
	if err != nil {
		panic(err)
	}
	var chainF tui.SimpleFocusChain
	chainF.Set(listRows,
		nameEdit,
		loginEdit,
		passwordEdit,
		dataEdit,
		cardNumberEdit,
		cardCVCEdit,
		cardDateEdit,
		cardOwnerEdit,
		cancelButton,
		saveButton,
	)

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetFocusChain(&chainF)

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
