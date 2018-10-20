package client

import (
	"github.com/pkg/errors"
)

type TallyIndication string

const (
	TallyOff     TallyIndication = "Off"
	TallyProgram TallyIndication = "Program"
	TallyPreview TallyIndication = "Preview"
)

type setStudioTallyParams struct {
	Indication TallyIndication `json:"Indication"`
}

func (c HTTPClient) SetStudioTally(indication TallyIndication) error {
	resp, err := c.makeRequest("SetStudioTally", setStudioTallyParams{
		Indication: indication,
	})
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}

	if resp.Result != apiResultSuccess {
		return errors.New("Operation failed")
	}

	return nil
}
