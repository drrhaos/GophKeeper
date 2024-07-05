// Package tuiclient реализует TUI инфтерфейс.
package tuiclient

import (
	"context"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"

	"gophkeeper/internal/client/configure"
	"gophkeeper/internal/client/grpcclient"
	"gophkeeper/pkg/proto"
	pb "gophkeeper/pkg/proto"

	"github.com/marcusolsson/tui-go"
)

// Form описывает структуру графического интерфейса.
type Form struct {
	ui         tui.UI
	listRows   *tui.List
	gridFields *tui.Grid
	cfg        configure.Config

	statusField    *tui.Label
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

	boxFile        *tui.Box
	fileNameLabel  *tui.Label
	fileNameEdit   *tui.TextEdit
	addFileButton  *tui.Button
	delFileButton  *tui.Button
	loadFileButton *tui.Button

	saveFileButton     *tui.Button
	cancelFileButton   *tui.Button
	saveLoadFileButton *tui.Button

	statusBar *tui.StatusBar

	cli        *grpcclient.GRPCClient
	listFields *proto.ListFielsdKeepResponse

	chainFocus tui.SimpleFocusChain

	change bool
}

// NewForm создает базовую форму.
func (fm *Form) NewForm(cfg configure.Config, buildVersion string, buildDate string) {
	fm.cfg = cfg
	var err error
	formLogin := FormLogin{}
	formLogin.NewFormLogin(cfg, buildVersion, buildDate)
	if formLogin.Cli == nil {
		return
	}
	fm.cli = formLogin.Cli

	fm.listRows = tui.NewList()

	fm.statusField = tui.NewLabel("")

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

	fm.boxFile = tui.NewHBox()
	fm.fileNameLabel = tui.NewLabel("Файл: ")
	fm.boxFile.Append(fm.fileNameLabel)
	fm.fileNameEdit = tui.NewTextEdit()
	fm.fileNameEdit.SetSizePolicy(tui.Expanding, tui.Minimum)
	fm.boxFile.Append(fm.fileNameEdit)
	fm.addFileButton = tui.NewButton("  Добавить  ")
	fm.boxFile.Append(fm.addFileButton)
	fm.saveFileButton = tui.NewButton("  Сохранить  ")
	fm.cancelFileButton = tui.NewButton("  Отменить  ")

	fm.saveLoadFileButton = tui.NewButton("  Сохранить  ")
	fm.delFileButton = tui.NewButton("  Удалить  ")
	fm.boxFile.Append(fm.delFileButton)
	fm.loadFileButton = tui.NewButton("  Скачать  ")
	fm.boxFile.Append(fm.loadFileButton)

	fieldSidebar := tui.NewVBox(gridFields, fm.boxFile, tui.NewSpacer())
	fieldSidebar.SetBorder(true)
	fieldSidebar.SetSizePolicy(tui.Expanding, tui.Expanding)

	fm.statusBar = tui.NewStatusBar("Соединен с сервером")

	fieldBox := tui.NewVBox(fm.statusField,
		fieldSidebar,
		boxButton,
	)
	fieldBox.SetBorder(true)
	boxGeneral := tui.NewVBox(
		tui.NewHBox(
			sidebarList,
			fieldBox,
		),
		fm.statusBar,
	)

	fm.ui, err = tui.New(boxGeneral)
	if err != nil {
		panic(err)
	}

	fm.ui.SetKeybinding("Esc", func() { fm.ui.Quit() })
	fm.ui.SetKeybinding("Up", fm.upList)
	fm.ui.SetKeybinding("Down", fm.downList)

	fm.listRows.OnSelectionChanged(fm.setFeld)

	fm.saveButton.OnActivated(fm.saveField)
	fm.addButton.OnActivated(fm.addField)
	fm.deleteButton.OnActivated(fm.deleteField)
	fm.cancelButton.OnActivated(fm.cancelField)

	fm.addFileButton.OnActivated(fm.addFile)
	fm.saveFileButton.OnActivated(fm.saveFile)
	fm.loadFileButton.OnActivated(fm.loadFile)
	fm.saveLoadFileButton.OnActivated(fm.saveLoadFile)

	fm.nameEdit.OnTextChanged(fm.setChanged)
	fm.loginEdit.OnTextChanged(fm.setChanged)
	fm.passwordEdit.OnTextChanged(fm.setChanged)
	fm.dataEdit.OnTextChanged(fm.setChanged)
	fm.cardNumberEdit.OnTextChanged(fm.setChanged)
	fm.cardCVCEdit.OnTextChanged(fm.setChanged)
	fm.cardDateEdit.OnTextChanged(fm.setChanged)
	fm.cardOwnerEdit.OnTextChanged(fm.setChanged)

	fm.loadItems()

	if len(fm.listFields.GetData()) > 0 {
		fm.listRows.SetSelected(0)
		fm.setFeld(fm.listRows)
	}
	fm.setChainFocus(0)
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
		FileName:   fm.fileNameEdit.Text(),
	}

	tmpFieldExt := &pb.EditFieldKeepRequest{
		Uuid: fm.listRows.SelectedItem(),
		Data: &tmpFieldKeep,
	}

	res, err := fm.cli.SaveField(tmpFieldExt)
	if err != nil {
		fm.statusBar.SetText(err.Error())
		return
	}

	filePath := filepath.Join(fm.cfg.StaticPath, fm.listRows.SelectedItem())
	_, err = os.Stat(filePath)
	if fm.fileNameEdit.Text() != "" && err == nil {
		fm.cli.Upload(context.Background(), filePath)
	} else {
		fm.statusBar.SetText("Не удалось загрузить файл")
	}

	fm.listFields.Data[fm.listRows.SelectedItem()] = res.Data

	fm.change = false
	fm.statusField.SetText("")
	fm.statusBar.SetText("Запись успешно сохранена")
}

func (fm *Form) addField(_ *tui.Button) {
	if fm.change {
		fm.statusBar.SetText("Сохраните изменения")
		return
	}
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
	fm.change = false
	fm.statusField.SetText("")
	fm.statusBar.SetText(fmt.Sprintf("Запись успешно добавлена %s", res.GetUuid()))
}

func (fm *Form) deleteField(_ *tui.Button) {
	if fm.change {
		fm.statusBar.SetText("Сохраните изменения")
		return
	}
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
	if !fm.change {
		return
	}
	if fm.listRows.Selected() < 0 {
		fm.statusBar.SetText("Запись не выбрана")
		return
	}
	id := fm.listRows.Selected()
	uuid := fm.listRows.SelectedItem()

	fm.loadItems()
	fm.listRows.Select(id)
	fm.statusBar.SetText(fmt.Sprintf("Изменения отменены %s", uuid))

	fm.change = false
	fm.statusField.SetText("")
}

func (fm *Form) setFeld(l *tui.List) {
	if fm.change {
		return
	}

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
	fm.fileNameEdit.SetText(mapList.FileName)
}

func (fm *Form) addFile(_ *tui.Button) {
	if fm.listRows.Selected() < 0 {
		fm.statusBar.SetText("Запись не выбрана")
		return
	}
	fm.boxFile.Remove(4)
	fm.boxFile.Remove(3)
	fm.boxFile.Remove(2)
	fm.boxFile.Append(fm.saveFileButton)
	fm.boxFile.Append(fm.cancelFileButton)
	fm.fileNameLabel.SetText("Укажите путь до файла: ")
	fm.setChainFocus(1)
}

func (fm *Form) saveFile(_ *tui.Button) {
	if fm.listRows.Selected() < 0 {
		fm.statusBar.SetText("Запись не выбрана")
		return
	}
	_, err := os.Stat(fm.fileNameEdit.Text())
	if err != nil {
		fm.statusBar.SetText("Файл не существует")
		return
	}

	srcFile, err := os.Open(fm.fileNameEdit.Text())
	if err != nil {
		fm.statusBar.SetText("Не удалось прочитать файл")
		return
	}
	defer srcFile.Close()

	destPath := filepath.Join(fm.cfg.StaticPath, fm.listRows.SelectedItem())

	destFile, err := os.Create(destPath)
	if err != nil {
		fm.statusBar.SetText("Не удалось записать файл")
		return
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)

	_, nameFile := filepath.Split(srcFile.Name())
	fm.fileNameEdit.SetText(nameFile)

	fm.fileNameLabel.SetText("Файл: ")
	fm.boxFile.Append(fm.addFileButton)
	fm.boxFile.Append(fm.delFileButton)
	fm.boxFile.Append(fm.loadFileButton)
	fm.boxFile.Remove(3)
	fm.boxFile.Remove(2)
	fm.setChainFocus(0)
	fm.statusBar.SetText("Файл добавлен")
	fm.setChainFocus(4)
	fm.change = true
}

func (fm *Form) loadFile(_ *tui.Button) {
	if fm.listRows.Selected() < 0 {
		fm.statusBar.SetText("Запись не выбрана")
		return
	}
	fm.boxFile.Remove(4)
	fm.boxFile.Remove(3)
	fm.boxFile.Remove(2)
	fm.boxFile.Append(fm.saveLoadFileButton)
	fm.boxFile.Append(fm.cancelFileButton)
	fm.fileNameLabel.SetText("Укажите путь для сохранения файла: ")
	fm.setChainFocus(2)
}

func (fm *Form) saveLoadFile(_ *tui.Button) {
	if fm.listRows.Selected() < 0 {
		fm.statusBar.SetText("Запись не выбрана")
		return
	}
	uuid := fm.listRows.SelectedItem()
	err := fm.cli.Download(context.Background(), uuid, fm.listFields.GetData()[uuid].FileName)
	if err != nil {
		fm.statusBar.SetText(err.Error())
		return
	}

	srcFile, err := os.Open(filepath.Join(fm.cfg.StaticPath, uuid))
	if err != nil {
		fm.statusBar.SetText("Не удалось прочитать файл")
		return
	}
	defer srcFile.Close()

	destPath := filepath.Join(fm.fileNameEdit.Text())

	destFile, err := os.Create(destPath)
	if err != nil {
		fm.statusBar.SetText("Не удалось создать файл")
		return
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fm.statusBar.SetText("Не удалось скопировать файл")
		return
	}

	fm.fileNameLabel.SetText("Файл: ")
	fm.boxFile.Remove(3)
	fm.boxFile.Remove(2)
	fm.boxFile.Append(fm.addFileButton)
	fm.boxFile.Append(fm.delFileButton)
	fm.boxFile.Append(fm.loadFileButton)
	fm.setChainFocus(0)
	fm.statusBar.SetText("Файл сохранен")
}

func (fm *Form) setChainFocus(name int) {
	switch name {
	case 0:
		fm.chainFocus.Set(
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
			fm.loadFileButton,
			fm.addButton,
			fm.saveButton,
			fm.deleteButton,
			fm.cancelButton,
		)
	case 1:
		fm.chainFocus.Set(
			fm.fileNameEdit,
			fm.saveFileButton,
			fm.cancelFileButton,
		)
	case 2:
		fm.chainFocus.Set(
			fm.fileNameEdit,
			fm.saveLoadFileButton,
			fm.cancelFileButton,
		)
	case 4:
		fm.chainFocus.Set(
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
			fm.loadFileButton,
			fm.saveButton,
			fm.cancelButton,
		)
	}
	fm.ui.SetFocusChain(&fm.chainFocus)
}

func (fm *Form) setChanged(_ *tui.TextEdit) {
	if fm.change == false {
		fm.change = true
		fm.statusField.SetText("* Запись изменена. Сохраните запись или отмените изменения.")
	}
}

func (fm *Form) upList() {
	if fm.change {
		return
	}
	if fm.listRows.Selected() > 0 {
		fm.listRows.Select(fm.listRows.Selected() - 1)
	} else {
		fm.listRows.Select(fm.listRows.Length() - 1)
	}
}

func (fm *Form) downList() {
	if fm.change {
		return
	}
	if fm.listRows.Selected() < fm.listRows.Length()-1 {
		fm.listRows.Select(fm.listRows.Selected() + 1)
	} else {
		fm.listRows.Select(0)
	}
}
