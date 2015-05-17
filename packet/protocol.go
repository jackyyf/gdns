package packet

/*
 * This file is translated from dns-protocol.h in dnsmasq(http://www.thekelleys.org.uk/dnsmasq/doc.html)
 */

const (
	IETF_DNS_PORT   = 53
	IP6_SIZE        = 16
	IP4_SIZE        = 4
	PKT_MAX_SIZE    = 512
	LABEL_MAX_SIZE  = 63
	DOMAIN_MAX_SIZE = 253
)

type STATUS uint8

const (
	NOERROR STATUS = iota
	FORMERR
	SERVFAIL
	NXDOMAIN
	NOTIMP
	REFUSED
)

type RRTYPE uint16

const (
	T_A      RRTYPE = 1
	T_NS     RRTYPE = 2
	T_MD     RRTYPE = 3
	T_MF     RRTYPE = 4
	T_CNAME  RRTYPE = 5
	T_SOA    RRTYPE = 6
	T_MB     RRTYPE = 7
	T_MG     RRTYPE = 8
	T_MR     RRTYPE = 9
	T_PTR    RRTYPE = 12
	T_MINFO  RRTYPE = 14
	T_MX     RRTYPE = 15
	T_TXT    RRTYPE = 16
	T_RP     RRTYPE = 17
	T_AFSDB  RRTYPE = 18
	T_RT     RRTYPE = 21
	T_SIG    RRTYPE = 24
	T_PX     RRTYPE = 26
	T_AAAA   RRTYPE = 28
	T_NXT    RRTYPE = 30
	T_SRV    RRTYPE = 33
	T_NAPTR  RRTYPE = 35
	T_KX     RRTYPE = 36
	T_DNAME  RRTYPE = 39
	T_OPT    RRTYPE = 41
	T_DS     RRTYPE = 43
	T_RRSIG  RRTYPE = 46
	T_NSEC   RRTYPE = 47
	T_DNSKEY RRTYPE = 48
	T_NSEC3  RRTYPE = 50
	T_TKEY   RRTYPE = 249
	T_TSIG   RRTYPE = 250
	T_AXFR   RRTYPE = 252
	T_MAILB  RRTYPE = 253
	T_ANY    RRTYPE = 255
)

type Header struct {
	ID      uint16
	b3, b4  uint8
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

const (
	MASK_QR     uint8 = 0x80
	MASK_OPCODE uint8 = 0x78
	MASK_AA     uint8 = 0x04
	MASK_TC     uint8 = 0x02
	MASK_RD     uint8 = 0x01
	MASK_RA     uint8 = 0x80
	MASK_AD     uint8 = 0x20
	MASK_CD     uint8 = 0x10
	MASK_RCODE  uint8 = 0x0F
)

func (h *Header) QR() bool {
	return h.b3&MASK_QR == 0
}

func (h *Header) SetQR() {
	h.b3 = h.b3&^MASK_QR | MASK_QR
}

func (h *Header) ClearQR() {
	h.b3 &= ^MASK_QR
}

func (h *Header) Opcode() uint8 {
	return (h.b3 & MASK_OPCODE) >> 3
}

func (h *Header) SetOpcode(code uint8) {
	h.b3 = (h.b3 & ^MASK_OPCODE) | (code << 3)
}

func (h *Header) AA() bool {
	return h.b3&MASK_AA != 0
}

func (h *Header) SetAA() {
	h.b3 = h.b3&^MASK_AA | MASK_AA
}

func (h *Header) ClearAA() {
	h.b3 &= ^MASK_AA
}

func (h *Header) TC() bool {
	return h.b3&MASK_TC != 0
}

func (h *Header) SetTC() {
	h.b3 = h.b3&^MASK_TC | MASK_TC
}

func (h *Header) ClearTC() {
	h.b3 &= ^MASK_TC
}

func (h *Header) RD() bool {
	return h.b3&MASK_RD != 0
}

func (h *Header) SetRD() {
	h.b3 = h.b3&^MASK_RD | MASK_RD
}

func (h *Header) ClearRD() {
	h.b3 &= ^MASK_RD
}
