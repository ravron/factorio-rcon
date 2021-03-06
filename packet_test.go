package rcon

import (
	"bytes"
	"testing"
)

func TestNewPacket(t *testing.T) {
	p1 := NewPacket(Auth, "password")

	if p1.Size != 18 {
		t.Error("Expected packet size 18, got", p1.Size)
	}
	if p1.Type != Auth {
		t.Error("Expected packet type Auth(3), got", p1.Type)
	}
	if p1.Body != "password" {
		t.Error("Expected packet body \"password\", got", p1.Body)
	}

	p2 := NewPacket(ExecCommand, "status")
	if p2.Size != 16 {
		t.Error("Expected packet size 16, got", p2.Size)
	}
	if p2.Type != ExecCommand {
		t.Error("Expected packet type ExecCommand(2), got", p2.Type)
	}
	if p2.Body != "status" {
		t.Error("Expected packet body \"status\", got", p2.Body)
	}
}

func TestRandomID(t *testing.T) {
	ids := make(map[int32]bool)

	// generating 1000 random id's in a row is probably sufficient
	for i := 0; i < 1000; i++ {
		p := NewPacket(Auth, "pw")

		if ids[p.ID] {
			t.Error("Expected unique IDs, saw ID multiple times: ", p.ID)
		}
		ids[p.ID] = true
	}
}

func TestPayload(t *testing.T) {
	p := NewPacket(Auth, "password")
	payload, _ := p.Payload()

	size := payload[0:4]
	typ := payload[8:12]
	body := payload[12 : len(payload)-2]
	padding := payload[len(payload)-2:]

	if !bytes.Equal(size, []byte{18, 0, 0, 0}) {
		t.Error("Expected payload [0:4] to be bytes [18 0 0 0], got", payload[0:4])
	}
	if !bytes.Equal(typ, []byte{3, 0, 0, 0}) {
		t.Error("Expected payload [8:12] to be bytes [3 0 0 0], got", typ)
	}
	if !bytes.Equal(body, []byte("password")) {
		t.Error("Expected payload body to be bytes \"password\", got", body)
	}
	if !bytes.Equal(padding, []byte("\x00\x00")) {
		t.Error("Expected two bytes of null padding at end of payload, got", padding)
	}
}
