// Package tuiclient реализует TUI инфтерфейс.
package tuiclient

import (
	"bufio"
	"context"
	"fmt"
	"image"
	"os"
	"strings"
	"syscall"

	"gophkeeper/internal/client/configure"
	"gophkeeper/internal/client/grpcclient"
	"gophkeeper/pkg/proto"
	pb "gophkeeper/pkg/proto"

	"github.com/marcusolsson/tui-go"
	"golang.org/x/term"
)

// Form описывает структуру графического интерфейса.
type Form struct {
	ui         tui.UI
	listRows   *tui.List
	gridFields *tui.Grid

	nameEdit       *tui.TextEdit
	loginEdit      *tui.TextEdit
	passwordEdit   *tui.TextEdit
	dataEdit       *tui.TextEdit
	cardNumberEdit *tui.TextEdit
	cardCVCEdit    *tui.TextEdit
	cardDateEdit   *tui.TextEdit
	cardOwnerEdit  *tui.TextEdit

	addButton    *tui.Button
	saveButton   *tui.Button
	deleteButton *tui.Button
	cancelButton *tui.Button

	fileNameEdit  *tui.TextEdit
	addFileButton *tui.Button
	delFileButton *tui.Button

	statusBar *tui.StatusBar

	cli        *grpcclient.GRPCClient
	listFields *proto.ListFielsdKeepResponse
}

func credentials() (bool, string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите 'да' если хотите зарегистрироваться: ")
	regText, err := reader.ReadString('\n')
	if err != nil {
		return false, "", "", err
	}
	var reg bool
	if regText == "да" {
		reg = true
	}
	fmt.Print("Введите имя пользователя: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return reg, "", "", err
	}

	fmt.Print("Введите пароль: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return reg, "", "", err
	}

	password := string(bytePassword)
	return reg, strings.TrimSpace(username), strings.TrimSpace(password), nil
}

// NewForm создает базовую форму.
func (fm *Form) NewForm(cfg configure.Config) {
	var err error
	formLogin := FormLogin{}
	formLogin.NewFormLogin(cfg)
	if formLogin.Cli == nil {
		return
	}
	fm.cli = formLogin.Cli

	fm.listRows = tui.NewList()

	gridFields := tui.NewGrid(2, 8)
	gridFields.SetBorder(true)

	gridFields.SetCell(image.Point{0, 0}, tui.NewLabel("Имя"))
	fm.nameEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 0}, fm.nameEdit)

	gridFields.SetCell(image.Point{0, 1}, tui.NewLabel("Логин"))
	fm.loginEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 1}, fm.loginEdit)

	gridFields.SetCell(image.Point{0, 2}, tui.NewLabel("Пароль"))
	fm.passwordEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 2}, fm.passwordEdit)

	gridFields.SetCell(image.Point{0, 3}, tui.NewLabel("Описание"))
	fm.dataEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 3}, fm.dataEdit)

	gridFields.SetCell(image.Point{0, 4}, tui.NewLabel("Номер карты"))
	fm.cardNumberEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 4}, fm.cardNumberEdit)

	gridFields.SetCell(image.Point{0, 5}, tui.NewLabel("CVC карты"))
	fm.cardCVCEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 5}, fm.cardCVCEdit)

	gridFields.SetCell(image.Point{0, 6}, tui.NewLabel("Дата окончания срока действия карты"))
	fm.cardDateEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 6}, fm.cardDateEdit)

	gridFields.SetCell(image.Point{0, 7}, tui.NewLabel("Владелец карты"))
	fm.cardOwnerEdit = tui.NewTextEdit()
	gridFields.SetCell(image.Point{1, 7}, fm.cardOwnerEdit)

	sidebarList := tui.NewVBox(fm.listRows)
	sidebarList.SetBorder(true)

	boxButton := tui.NewHBox()
	boxButton.SetBorder(true)
	fm.addButton = tui.NewButton("  Добавить  ")
	boxButton.Append(fm.addButton)
	boxButton.SetBorder(true)
	fm.saveButton = tui.NewButton("  Сохранить  ")
	boxButton.Append(fm.saveButton)
	fm.deleteButton = tui.NewButton("  Удалить  ")
	boxButton.Append(fm.deleteButton)
	boxButton.Append(tui.NewSpacer())
	fm.cancelButton = tui.NewButton("  Отменить  ")
	boxButton.Append(fm.cancelButton)

	boxFile := tui.NewHBox()
	fm.fileNameEdit = tui.NewTextEdit()
	boxFile.Append(fm.fileNameEdit)
	fm.addFileButton = tui.NewButton("  Добавить файл  ")
	boxFile.Append(fm.addFileButton)
	fm.delFileButton = tui.NewButton("  Удалить файл  ")
	boxFile.Append(fm.delFileButton)

	fieldSidebar := tui.NewVBox(gridFields, boxFile, tui.NewSpacer())
	fieldSidebar.SetBorder(true)
	fieldSidebar.SetSizePolicy(tui.Expanding, tui.Expanding)

	fm.statusBar = tui.NewStatusBar("Соединен с севревром")

	boxGeneral := tui.NewVBox(
		tui.NewHBox(
			sidebarList,
			tui.NewVBox(
				fieldSidebar,
				boxButton,
			),
		),
		fm.statusBar,
	)

	fm.ui, err = tui.New(boxGeneral)
	if err != nil {
		panic(err)
	}

	var chainF tui.SimpleFocusChain
	chainF.Set(
		fm.listRows,
		fm.nameEdit,
		fm.loginEdit,
		fm.passwordEdit,
		fm.dataEdit,
		fm.cardNumberEdit,
		fm.cardCVCEdit,
		fm.cardDateEdit,
		fm.cardOwnerEdit,
		fm.addFileButton,
		fm.delFileButton,
		fm.addButton,
		fm.saveButton,
		fm.deleteButton,
		fm.cancelButton,
	)
	fm.ui.SetFocusChain(&chainF)

	fm.ui.SetKeybinding("Esc", func() { fm.ui.Quit() })
	fm.listRows.OnSelectionChanged(fm.setFeld)
	fm.saveButton.OnActivated(fm.saveField)
	fm.addButton.OnActivated(fm.addField)
	fm.deleteButton.OnActivated(fm.deleteField)
	fm.cancelButton.OnActivated(fm.cancelField)

	fm.addFileButton.OnActivated(fm.addFile)

	fm.loadItems()

	if len(fm.listFields.GetData()) > 0 {
		fm.listRows.SetSelected(0)
		fm.setFeld(fm.listRows)
	}

	if err := fm.ui.Run(); err != nil {
		panic(err)
	}
}

func (fm *Form) loadItems() {
	fm.listFields = fm.cli.GetListFields()
	fm.listRows.RemoveItems()
	for fieldKeep := range fm.listFields.GetData() {
		fm.listRows.AddItems(fieldKeep)
	}
}

func (fm *Form) saveField(_ *tui.Button) {
	if fm.listRows.Selected() < 0 {
		fm.statusBar.SetText("Запись не выбрана")
		return
	}

	tmpFieldKeep := pb.FieldKeep{
		Name:       fm.nameEdit.Text(),
		Login:      fm.loginEdit.Text(),
		Password:   fm.passwordEdit.Text(),
		Data:       fm.dataEdit.Text(),
		CardNumber: fm.cardNumberEdit.Text(),
		CardCVC:    fm.cardCVCEdit.Text(),
		CardDate:   fm.dataEdit.Text(),
		CardOwner:  fm.cardOwnerEdit.Text(),
	}

	tmpFieldExt := &pb.EditFieldKeepRequest{
		Uuid: fm.listRows.SelectedItem(),
		Data: &tmpFieldKeep,
	}

	res, err := fm.cli.SaveField(tmpFieldExt)

	fm.cli.Upload(context.Background(), "./client")

	if err != nil {
		fm.statusBar.SetText(err.Error())
	}
	fm.listFields.Data[fm.listRows.SelectedItem()] = res.Data
	fm.statusBar.SetText("Запись успешно сохранена")
}

func (fm *Form) addField(_ *tui.Button) {
	tmpFieldKeep := pb.FieldKeep{}

	tmpField := &pb.AddFieldKeepRequest{
		Data: &tmpFieldKeep,
	}

	res, err := fm.cli.AddField(tmpField)
	if err != nil {
		fm.statusBar.SetText(err.Error())
	}

	fm.loadItems()
	var i int
	for fieldKeep := range fm.listFields.GetData() {
		if fieldKeep == res.GetUuid() {
			fm.listRows.Select(i)
		}
		i++
	}
	fm.statusBar.SetText(fmt.Sprintf("Запись успешно добавлена %s", res.GetUuid()))
}

func (fm *Form) deleteField(_ *tui.Button) {
	if fm.listRows.Selected() < 0 {
		fm.statusBar.SetText("Запись не выбрана")
		return
	}

	tmpField := &pb.DeleteFieldKeepRequest{
		Uuid: fm.listRows.SelectedItem(),
	}

	res, err := fm.cli.DelField(tmpField)
	if err != nil {
		fm.statusBar.SetText(err.Error())
	}
	fm.loadItems()
	if fm.listRows.Length() > 0 {
		fm.listRows.Select(0)
	} else {
		fm.listRows.Select(-1)
	}
	fm.statusBar.SetText(fmt.Sprintf("Запись успешно удалена %s", res.GetUuid()))
}

func (fm *Form) cancelField(_ *tui.Button) {
	if fm.listRows.Selected() < 0 {
		fm.statusBar.SetText("Запись не выбрана")
		return
	}
	id := fm.listRows.Selected()
	uuid := fm.listRows.SelectedItem()

	fm.loadItems()
	fm.listRows.Select(id)
	fm.statusBar.SetText(fmt.Sprintf("Изменения отменены %s", uuid))
}

func (fm *Form) setFeld(l *tui.List) {
	if l.Length() == 0 {
		return
	}
	mapList := fm.listFields.GetData()[l.SelectedItem()]
	fm.nameEdit.SetText(mapList.Name)
	fm.loginEdit.SetText(mapList.Login)
	fm.passwordEdit.SetText(mapList.Password)
	fm.dataEdit.SetText(mapList.Data)
	fm.cardNumberEdit.SetText(mapList.CardNumber)
	fm.cardCVCEdit.SetText(mapList.CardCVC)
	fm.cardDateEdit.SetText(mapList.CardDate)
	fm.cardOwnerEdit.SetText(mapList.CardOwner)
}

func (fm *Form) addFile(_ *tui.Button) {
	if fm.listRows.Selected() < 0 {
		fm.statusBar.SetText("Запись не выбрана")
		return
	}
	fm.fileNameEdit.SetFocused(true)
	// fm.statusBar.SetText(fmt.Sprintf("Изменения отменены %s", uuid))
}
