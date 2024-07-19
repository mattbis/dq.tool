package dirq

import (
    "github.com/mattbis/dq-go.tool/internal"
)

func GetVersion() string {
    dq := internal.NewDQ()
    return dq.GetVersion()
}
