package main
import "github.com/therecipe/qt/widgets"

func NewServerStartWidget(parent widgets.QWidget_ITF, settings *ServerSettings, callback func(*ServerSettings)()) {
	window := widgets.NewQWidget(parent, 1)
	window.SetWindowTitle("Start server...")
	layout := widgets.NewQVBoxLayout()

    dictionaryLabel := NewQLabelF("Dictionary location:")
    dictionaryTextBox := widgets.NewQLineEdit2(settings.DictionaryLocation, nil)
	browseButton := widgets.NewQPushButton2("Browse...", nil)
	browseButton.ConnectPressed(func() {
		dictionaryTextBox.SetText(widgets.QFileDialog_GetOpenFileName(window, "Find dictionaries...", "", "GOB files (*.gob)", "", 0))
	})
	layout.AddWidget(dictionaryLabel, 0, 0)
	layout.AddWidget(dictionaryTextBox, 0, 0)
	layout.AddWidget(browseButton, 0, 0)

    enumLabel := NewQLabelF("Enum schema location:")
	enumTextBox := widgets.NewQLineEdit2(settings.EnumSchemaLocation, nil)
	browseButton = widgets.NewQPushButton2("Browse...", nil)
	browseButton.ConnectPressed(func() {
		enumTextBox.SetText(widgets.QFileDialog_GetOpenFileName(window, "Find enum schema...", "", "GOB files (*.gob)", "", 0))
	})
	layout.AddWidget(enumLabel, 0, 0)
	layout.AddWidget(enumTextBox, 0, 0)
	layout.AddWidget(browseButton, 0, 0)

    instanceLabel := NewQLabelF("Instance schema location:")
	instanceTextBox := widgets.NewQLineEdit2(settings.InstanceSchemaLocation, nil)
	browseButton = widgets.NewQPushButton2("Browse...", nil)
	browseButton.ConnectPressed(func() {
		instanceTextBox.SetText(widgets.QFileDialog_GetOpenFileName(window, "Find instance schema...", "", "GOB files (*.gob)", "", 0))
	})
	layout.AddWidget(instanceLabel, 0, 0)
	layout.AddWidget(instanceTextBox, 0, 0)
	layout.AddWidget(browseButton, 0, 0)

	portLabel := NewQLabelF("Port number:")
	port := widgets.NewQLineEdit2(settings.Port, nil)
	layout.AddWidget(portLabel, 0, 0)
	layout.AddWidget(port, 0, 0)

	startButton := widgets.NewQPushButton2("Start", nil)
	startButton.ConnectPressed(func() {
		window.Destroy(true, true)
        settings.Port = port.Text()
        settings.EnumSchemaLocation = enumTextBox.Text()
        settings.InstanceSchemaLocation = instanceTextBox.Text()
        settings.DictionaryLocation = dictionaryTextBox.Text()
		callback(settings)
	})
	layout.AddWidget(startButton, 0, 0)

	window.SetLayout(layout)
	window.Show()
}
