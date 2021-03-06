package main
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/gui"
import "github.com/therecipe/qt/core"
import "github.com/gskartwii/rbxfile"
import "github.com/gskartwii/rbxfile/bin"
import "github.com/gskartwii/roblox-dissector/peer"
import "os"

func showChildren(rootNode *gui.QStandardItem, children []*rbxfile.Instance) {
	for _, instance := range(children) {
		row := showReplicationInstance(instance)
		if len(instance.Children) > 0 {
			childrenRootItem := NewQStandardItemF("%d children", len(instance.Children))
			showChildren(childrenRootItem, instance.Children)
			row[0].AppendRow([]*gui.QStandardItem{childrenRootItem,nil,nil,nil,nil,nil})
		}
		rootNode.AppendRow(row)
	}
}

func stripInvalidTypes(instances []*rbxfile.Instance, defaultValues DefaultValues) {
	for _, instance := range instances {
		color, ok := instance.Properties["Color3uint8"]
		if ok {
			if _, ok = color.(rbxfile.ValueDefault); !ok {
				col := color.(rbxfile.ValueColor3uint8)
				instance.Properties["Color"] = rbxfile.ValueColor3{
					R: float32(col.R) / 255,
					G: float32(col.R) / 255,
					B: float32(col.B) / 255,
				}
			}
		}

		for name, property := range instance.Properties {
			thisType := property.Type()
			if thisType == rbxfile.TypeDefault {
				class, ok := defaultValues[instance.ClassName]
				if !ok {
					println("stripping instance", instance.ClassName)
					delete(instance.Properties, name)
					continue
				}
				value, ok := class[name]
				if !ok {
					println("stripping because nodef", instance.ClassName, name)
					delete(instance.Properties, name)
					continue
				}
				instance.Properties[name] = value
			} else if thisType >= rbxfile.TypeNumberSequenceKeypoint ||
			   thisType == rbxfile.TypeVector2int16 {
				println("stripping property", instance.ClassName, name, thisType.String())
				delete(instance.Properties, name)
				continue
			}
		}
		stripInvalidTypes(instance.Children, defaultValues)
	}
}

func NewDataModelBrowser(context *peer.CommunicationContext, dataModel *rbxfile.Root, defaultValues DefaultValues) {
	subWindow := widgets.NewQWidget(window, core.Qt__Window)
	subWindowLayout := widgets.NewQVBoxLayout2(subWindow)

	subWindow.SetWindowTitle("Data Model")

	context.MSchema.Lock()
	children := dataModel.Copy()
	context.MSchema.Unlock()
	
	takeSnapshotButton := widgets.NewQPushButton2("Save as RBXL...", nil)
	takeSnapshotButton.ConnectPressed(func() {
		location := widgets.QFileDialog_GetSaveFileName(subWindow, "Save as RBXL...", "", "Roblox place files (*.rbxl)", "", 0)
		writer, err := os.OpenFile(location, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			println("while opening file:", err.Error())
			return
		}

		writableClone := children.Copy()
		stripInvalidTypes(writableClone.Instances, defaultValues)

		err = bin.SerializePlace(writer, nil, writableClone)
		if err != nil {
			println("while serializing place:", err.Error())
			return
		}
	})
	subWindowLayout.AddWidget(takeSnapshotButton, 0, 0)

	instanceList := widgets.NewQTreeView(nil)
	standardModel := NewProperSortModel(nil)
	standardModel.SetHorizontalHeaderLabels([]string{"Name", "Type", "Value", "Referent", "Parent"})

	rootNode := standardModel.InvisibleRootItem()
	showChildren(rootNode, children.Instances)
	instanceList.SetModel(standardModel)
	instanceList.SetSelectionMode(0)
	instanceList.SetSortingEnabled(true)

	subWindowLayout.AddWidget(instanceList, 0, 0)

	subWindow.Show()
}
