package peer
import "errors"

type Packet83_11 struct {
	SkipStats1 bool
	Stats_1_1 []byte
	Stats_1_2 float32
	Stats_1_3 float32
	Stats_1_4 float32
	Stats_1_5 bool

	SkipStats2 bool
	Stats_2_1 []byte
	Stats_2_2 float32
	Stats_2_3 uint32
	Stats_2_4 bool
	
	AvgPingMs float32
	AvgPhysicsSenderPktPS float32
	TotalDataKBPS float32
	TotalPhysicsKBPS float32
	DataThroughputRatio float32
}

func DecodePacket83_11(packet *UDPPacket, context *CommunicationContext) (interface{}, error) {
	var err error
	inner := &Packet83_11{}
	thisBitstream := packet.Stream
	inner.SkipStats1, err = thisBitstream.ReadBool()
	if err != nil {
		return inner, err
	}
	if !inner.SkipStats1 {
		stringLen, err := thisBitstream.ReadUint32BE()
		if err != nil {
			return inner, err
		}
		inner.Stats_1_1, err = thisBitstream.ReadString(int(stringLen))
		if err != nil {
			return inner, err
		}

		inner.Stats_1_2, err = thisBitstream.ReadFloat32BE()
		if err != nil {
			return inner, err
		}
		inner.Stats_1_3, err = thisBitstream.ReadFloat32BE()
		if err != nil {
			return inner, err
		}
		inner.Stats_1_4, err = thisBitstream.ReadFloat32BE()
		if err != nil {
			return inner, err
		}
		inner.Stats_1_5, err = thisBitstream.ReadBool()
		if err != nil {
			return inner, err
		}
		print("Receive stats1", inner.Stats_1_1, ",", inner.Stats_1_2, ",", inner.Stats_1_3, ",", inner.Stats_1_4, ",", inner.Stats_1_5)
	}

	inner.SkipStats2, err = thisBitstream.ReadBool()
	if err != nil {
		return inner, err
	}
	if !inner.SkipStats2 {
		stringLen, err := thisBitstream.ReadUint32BE()
		if err != nil {
			return inner, err
		}
		inner.Stats_2_1, err = thisBitstream.ReadString(int(stringLen))
		if err != nil {
			return inner, err
		}

		inner.Stats_2_2, err = thisBitstream.ReadFloat32BE()
		if err != nil {
			return inner, err
		}
		inner.Stats_2_3, err = thisBitstream.ReadUint32BE()
		if err != nil {
			return inner, err
		}
		inner.Stats_2_4, err = thisBitstream.ReadBool()
		if err != nil {
			return inner, err
		}
		print("Receive stats2", inner.Stats_2_1, ",", inner.Stats_2_2, ",", inner.Stats_2_3, ",", inner.Stats_2_4)
	}

	inner.AvgPingMs, err = thisBitstream.ReadFloat32BE()
	if err != nil {
		return inner, err
	}
	inner.AvgPhysicsSenderPktPS, err = thisBitstream.ReadFloat32BE()
	if err != nil {
		return inner, err
	}
	inner.TotalDataKBPS, err = thisBitstream.ReadFloat32BE()
	if err != nil {
		return inner, err
	}
	inner.TotalPhysicsKBPS, err = thisBitstream.ReadFloat32BE()
	if err != nil {
		return inner, err
	}
	inner.DataThroughputRatio, err = thisBitstream.ReadFloat32BE()
	if err != nil {
		return inner, err
	}

	return inner, nil
}

func (layer *Packet83_11) Serialize(isClient bool, context *CommunicationContext, stream *ExtendedWriter) error {
    return errors.New("packet 83_11 not implemented!")
}
