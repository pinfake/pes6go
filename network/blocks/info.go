package blocks;

import (
    "encoding/binary"
    "time"
    "bytes"
)

const dtLayout = "2006-01-02 15:04:05"

type Info struct {
    time time.Time;
    title string;
    text string;
}

type internal struct {
    header [6]byte;
    time   [20]byte;
    title  [64]byte;
    text   [128]byte;
}

func buildInternal(info Info) internal {
    // TODO: Hohoho, no luck on the conversion; will have to work on this...
    return internal{
        header: [6]byte{0x00, 0x00, 0x03, 0x10, 0x01, 0x00},
        time:   [20]byte(info.time.Format(dtLayout)),
        title:  [64]byte(info.title),
        text:   [128]byte(info.text),
    }

}

func (info Info) getData() []byte {
    buf := bytes.Buffer{};
    binary.Write(buf, binary.LittleEndian, buildInternal(info));
    return buf.Bytes();
}