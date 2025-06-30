package hardware

import (
	"context"
	"tugas-cli/internal/config"
	"tugas-cli/internal/utils"
	"time"
)

const TAG = "HARDWARE"

type Hardware struct {
	ctx context.Context
	log *utils.Logger
}

func NewHardware(config *config.Config) *Hardware {
	logger := utils.NewLogger(10)
	return &Hardware{log: logger}
}

func (h *Hardware) Init(ctx context.Context) {
	h.ctx = ctx
	go h.hardwareA()
	go h.hardwareB()
}

func (h *Hardware) hardwareA() {
	for {
		select {
		case <-h.ctx.Done():
			h.log.Add(TAG, "Stopping HW A")
			return
		default:
			h.log.Add(TAG, "HW A")
			time.Sleep(1 * time.Second)
		}
	}
}

func (h *Hardware) hardwareB() {
	for {
		h.log.Add(TAG, "HW A")
		time.Sleep(2 * time.Second)
	}
}

func (h *Hardware) Stop() {
	// h.ctx = ctx
}
