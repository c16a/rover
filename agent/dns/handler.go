package dns

import (
	"github.com/miekg/dns"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Authoritative = true

	for _, question := range r.Question {
		h.logger.Info("received question", zap.String("question", question.String()), zap.String("type", question.Name))
		answers, err := resolve(question.Name, question.Qtype)
		if err == nil {
			msg.Answer = append(msg.Answer, answers...)
		}
	}

	w.WriteMsg(msg)
}

func resolve(domain string, qtype uint16) ([]dns.RR, error) {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), qtype)
	m.RecursionDesired = true

	c := new(dns.Client)
	in, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		return nil, err
	}

	return in.Answer, err
}
