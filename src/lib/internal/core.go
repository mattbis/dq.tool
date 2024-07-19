package internal

const VERSION = "0.1.0.20"

type DQ struct {
    Version string
}

func NewDQ() *DQ {
    return &DQ{Version: VERSION}
}

func (dq *DQ) GetVersion() string {
    return dq.Version
}
