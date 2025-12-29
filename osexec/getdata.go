package osexec

import (
	"bytes"
	"encoding/xml"
	"io"
	"os/exec"
	"strings"
)

type Payload struct {
	Message string `xml:"message"`
}

func GetData(r io.Reader) string {
	decoder := xml.NewDecoder(r)
	var payload Payload
	decoder.Decode(&payload)
	return strings.ToUpper(payload.Message)
}

func GetReader() io.Reader {
	cmd := exec.Command("cat", "msg.xml")
	out, _ := cmd.StdoutPipe()

	cmd.Start()
	data, _ := io.ReadAll(out)
	cmd.Wait()

	return bytes.NewReader(data)
}
