package main
import "github.com/therecipe/qt/widgets"
import "github.com/therecipe/qt/gui"
import "github.com/gskartwii/roblox-dissector/peer"

func ShowPacket81(packetType byte, packet *peer.UDPPacket, context *peer.CommunicationContext, layers *peer.PacketLayers) {
	MainLayer := layers.Main.(*peer.Packet81Layer)

	layerLayout := NewBasicPacketViewer(packetType, packet, context, layers)
	for i := 0; i < 5; i++ {
		thisLabel := NewQLabelF("Unknown boolean %d: %v", i, MainLayer.Bools[i])
		layerLayout.AddWidget(thisLabel, 0, 0)
	}
	referentStringLabel := NewQLabelF("Unknown string: %s", MainLayer.ReferentString)
	layerLayout.AddWidget(referentStringLabel, 0, 0)

	deletedList := widgets.NewQTreeView(nil)
	standardModel := NewProperSortModel(nil)
	standardModel.SetHorizontalHeaderLabels([]string{"Class name", "Referent"})

	deletedListRootNode := standardModel.InvisibleRootItem()
	for i := 0; i < len(MainLayer.Items); i++ {
		classNameItem := NewQStandardItemF("%s", context.StaticSchema.Instances[MainLayer.Items[i].ClassID].Name)
		referenceItem := NewQStandardItemF("%s", MainLayer.Items[i].Instance.Reference)
		deletedListRootNode.AppendRow([]*gui.QStandardItem{classNameItem, referenceItem})
	}

	deletedList.SetModel(standardModel)
	deletedList.SetSelectionMode(0)
	deletedList.SetSortingEnabled(true)

	layerLayout.AddWidget(deletedList, 0, 0)
}
