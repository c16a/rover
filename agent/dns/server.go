package dns

import (
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

func StartDnsServerWithHandler(handler *Handler, logger *zap.Logger) {
	server := &dns.Server{
		Addr:      ":1053",
		Net:       "udp",
		Handler:   handler,
		UDPSize:   65535,
		ReusePort: true,
	}

	logger.Info("starting DNS server on port 53")
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("failed to start DNS server", zap.Error(err))
	}
}

func StartDnsServer(logger *zap.Logger) {
	handler := NewHandler(logger)
	StartDnsServerWithHandler(handler, logger)
}
