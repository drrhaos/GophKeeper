// Package tuiclient реализует TUI инфтерфейс.
package tuiclient

import (
	"image"

	"gophkeeper/internal/client/configure"
	"gophkeeper/internal/client/grpcclient"
	"gophkeeper/internal/logger"
	pb "gophkeeper/pkg/proto"

	"github.com/marcusolsson/tui-go"
)

// Run функция запускает TUI интерфейс.
func Run(cfg configure.Config) {
	cli, err := grpcclient.NewGRPCClient(cfg, "test", "test")
	if err != nil {
		logger.Log.Panic("Не удалось установить соединение с сервером")
	}
	listFields := cli.GetListFields()

	listRows := tui.NewList()
	for fieldKeep := range listFields.GetData() {
		listRows.AddItems(fieldKeep)
	}

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
	addButton := tui.NewButton("  Добавить  ")
	boxButton.Append(addButton)
	boxButton.SetBorder(true)
	saveButton := tui.NewButton("  Сохранить  ")
	boxButton.Append(saveButton)
	deleteButton := tui.NewButton("  Удалить  ")
	boxButton.Append(deleteButton)
	boxButton.Append(tui.NewSpacer())
	cancelButton := tui.NewButton("  Отменить  ")
	boxButton.Append(cancelButton)

	statusBar := tui.NewStatusBar("Соединен с севревром")

	boxGeneral := tui.NewVBox(
		tui.NewHBox(
			sidebarList,
			tui.NewVBox(
				fieldSidebar,
				boxButton,
			),
		),
		statusBar,
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
		addButton,
		saveButton,
		deleteButton,
		cancelButton,
	)

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetFocusChain(&chainF)
	listRows.OnItemActivated(func(l *tui.List) {
		mapList := listFields.GetData()[l.SelectedItem()]
		nameEdit.SetText(mapList.Name)
		loginEdit.SetText(mapList.Login)
		passwordEdit.SetText(mapList.Password)
		dataEdit.SetText(mapList.Data)
		cardNumberEdit.SetText(mapList.CardNumber)
		cardCVCEdit.SetText(mapList.CardCVC)
		cardDateEdit.SetText(mapList.CardDate)
		cardOwnerEdit.SetText(mapList.CardOwner)
	})
	saveButton.OnActivated(func(_ *tui.Button) {
		if listRows.Selected() < 0 {
			statusBar.SetText("Запись не выбрана")
			return
		}
		tmpField := &pb.EditFieldKeepRequest{
			Uuid: listRows.SelectedItem(),
			Data: &pb.FieldKeep{
				Name:       nameEdit.Text(),
				Login:      loginEdit.Text(),
				Password:   passwordEdit.Text(),
				Data:       dataEdit.Text(),
				CardNumber: cardNumberEdit.Text(),
				CardCVC:    cardCVCEdit.Text(),
				CardDate:   dataEdit.Text(),
				CardOwner:  cardOwnerEdit.Text(),
			},
		}
		res, err := cli.SaveField(tmpField)
		if err != nil {
			statusBar.SetText(err.Error())
		}
		statusBar.SetText("Запись успешно сохранена")
		listFields.Data[listRows.SelectedItem()] = res.Data
	})

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
