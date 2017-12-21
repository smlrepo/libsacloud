package sacloud

import (
	"encoding/json"
	"strings"
)

// EWebAccelDomainType ウェブアクセラレータ ドメイン種別
type EWebAccelDomainType string

const (
	// EWebAccelDomainTypeOwn 独自ドメイン
	EWebAccelDomainTypeOwn = EWebAccelDomainType("own_domain")
	// EWebAccelDomainTypeSubDomain サブドメイン
	EWebAccelDomainTypeSubDomain = EWebAccelDomainType("subdomain")
)

// EWebAccelStatus ウェブアクセラレータ ステータス
type EWebAccelStatus string

const (
	// EWebAccelStatusEnabled 状態:有効
	EWebAccelStatusEnabled = EWebAccelStatus("enabled")
	// EWebAccelStatusDisabled 状態:無効
	EWebAccelStatusDisabled = EWebAccelStatus("disabled")
)

// WebAccelSite ウェブアクセラレータ サイト
type WebAccelSite struct {
	ID                 string              // ID
	Name               string              `json:",omitempty"`
	DomainType         EWebAccelDomainType `json:",omitempty"`
	Domain             string              `json:",omitempty"`
	Subdomain          string              `json:",omitempty"`
	ASCIIDomain        string              `json:",omitempty"`
	Origin             string              `json:",omitempty"`
	HostHeader         string
	Status             EWebAccelStatus `json:",omitempty"`
	HasCertificate     bool            `json:",omitempty"`
	HasOldCertificate  bool            `json:",omitempty"`
	GibSentInLastWeek  int64           `json:",omitempty"`
	CertValidNotBefore int64           `json:",omitempty"`
	CertValidNotAfter  int64           `json:",omitempty"`
	*propCreatedAt
}

// WebAccelCert ウェブアクセラレータ証明書
type WebAccelCert struct {
	ID               string `json:",omitempty"`
	SiteID           string `json:",omitempty"`
	CertificateChain string `json:",omitempty"`
	Key              string `json:",omitempty"`
	*propCreatedAt   `json:",omitempty"`
	*propUpdatedAt   `json:",omitempty"`

	SerialNumber string `json:",omitempty"`
	NotBefore    int64  `json:",omitempty"`
	NotAfter     int64  `json:",omitempty"`
	Issuer       *struct {
		Country            string `json:",omitempty"`
		Organization       string `json:",omitempty"`
		OrganizationalUnit string `json:",omitempty"`
		CommonName         string `json:",omitempty"`
	} `json:",omitempty"`
	Subject *struct {
		Country            string `json:",omitempty"`
		Organization       string `json:",omitempty"`
		OrganizationalUnit string `json:",omitempty"`
		Locality           string `json:",omitempty"`
		Province           string `json:",omitempty"`
		StreetAddress      string `json:",omitempty"`
		PostalCode         string `json:",omitempty"`
		SerialNumber       string `json:",omitempty"`
		CommonName         string `json:",omitempty"`
	} `json:",omitempty"`
	DNSNames          []string `json:",omitempty"`
	SHA256Fingerprint string   `json:",omitempty"`
}

// WebAccelCertRequest ウェブアクセラレータ証明書API リクエスト
type WebAccelCertRequest struct {
	CertificateChain string
	Key              string
}

// WebAccelCertResponse ウェブアクセラレータ証明書API レスポンス
type WebAccelCertResponse struct {
	Certificate *WebAccelCertResponseBody `json:",omitempty"`
	ResultFlagValue
}

// WebAccelCertResponseBody ウェブアクセラレータ証明書API レスポンスボディ
type WebAccelCertResponseBody struct {
	Current *WebAccelCert   `json:",omitempty"`
	Old     []*WebAccelCert `json:",omitempty"`
}

// UnmarshalJSON JSONアンマーシャル(配列、オブジェクトが混在するためここで対応)
func (s *WebAccelCertResponse) UnmarshalJSON(data []byte) error {
	tmp := &struct {
		Certificate *WebAccelCertResponseBody `json:",omitempty"`
		ResultFlagValue
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.Certificate.Current != nil || len(tmp.Certificate.Old) > 0 {
		s.Certificate = tmp.Certificate
	}
	s.ResultFlagValue = tmp.ResultFlagValue
	return nil
}

// UnmarshalJSON JSONアンマーシャル(配列、オブジェクトが混在するためここで対応)
func (s *WebAccelCertResponseBody) UnmarshalJSON(data []byte) error {
	targetData := strings.Replace(strings.Replace(string(data), " ", "", -1), "\n", "", -1)
	if targetData == `[]` {
		return nil
	}

	tmp := &struct {
		Current *WebAccelCert   `json:",omitempty"`
		Old     []*WebAccelCert `json:",omitempty"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	s.Current = tmp.Current
	s.Old = tmp.Old
	return nil
}
