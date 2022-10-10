package hashcash

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

var (
	ErrMaxIterations = fmt.Errorf("max iterations exceeded")
)

const (
	version    = 1
	dateLayout = "060102150405" // YYMMDDhhmmss
	zeroByte   = 48
)

// Hash is an implementation of HashCash algorithm (see https://ru.wikipedia.org/wiki/Hashcash)
type Hash struct {
	Version  uint
	Bits     uint
	Date     time.Time
	Resource string
	Rand     string
	Counter  uint64
}

func NewHash(resource string, bits uint) (*Hash, error) {
	h := Hash{
		Version:  version,
		Bits:     bits,
		Date:     time.Now(),
		Resource: resource,
		Counter:  0,
	}

	var err error
	h.Rand, err = GenRandString(8)
	if err != nil {
		return nil, err
	}

	return &h, nil
}

func (h *Hash) GetHeader() Header {
	header := fmt.Sprintf("%d:%d:%s::%s::%s:%s",
		h.Version,
		h.Bits,
		h.Date.Format(dateLayout),
		h.Resource,
		h.Rand,
		encodeBase64UInt(h.Counter),
	)

	return Header(header)
}

func (h *Hash) Compute(maxIterations uint64) (Header, error) {
	for ; h.Counter < maxIterations; h.Counter++ {
		header := h.GetHeader()

		if h.verify(header) {
			return header, nil
		}
	}

	return Header{}, ErrMaxIterations
}

func (h *Hash) Verify() bool {
	header := h.GetHeader()

	return h.verify(header)
}

func (h *Hash) verify(header Header) bool {
	hash := header.ComputeHash()
	zerosAmount := int(h.Bits)

	if zerosAmount > len(hash) {
		return false
	}

	for _, c := range hash[:zerosAmount] {
		if c != zeroByte {
			return false
		}
	}

	return true
}

// Header is the hash computed header
type Header []byte

func (h Header) ComputeHash() string {
	hasher := sha1.New()
	hasher.Write(h)
	bs := hasher.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func (h Header) String() string {
	return string(h)
}

func GenRandString(length int) (string, error) {
	bs := make([]byte, length)
	_, err := rand.Read(bs)

	return encodeBase64Bytes(bs), err
}

func encodeBase64Bytes(val []byte) string {
	return base64.StdEncoding.EncodeToString(val)
}

func encodeBase64UInt(val uint64) string {
	str := strconv.FormatUint(val, 10)
	bs := []byte(str)

	return encodeBase64Bytes(bs)
}
