package ssl

type Sslcertkey struct {
	Bundle              string `json:"bundle,omitempty"`
	Cert                string `json:"cert,omitempty"`
	Certkey             string `json:"certkey,omitempty"`
	Clientcertnotafter  string `json:"clientcertnotafter,omitempty"`
	Clientcertnotbefore string `json:"clientcertnotbefore,omitempty"`
	Data                int    `json:"data,omitempty"`
	Daystoexpiration    int    `json:"daystoexpiration,omitempty"`
	Expirymonitor       string `json:"expirymonitor,omitempty"`
	Fipskey             string `json:"fipskey,omitempty"`
	Inform              string `json:"inform,omitempty"`
	Issuer              string `json:"issuer,omitempty"`
	Key                 string `json:"key,omitempty"`
	Linkcertkeyname     string `json:"linkcertkeyname,omitempty"`
	Nodomaincheck       bool   `json:"nodomaincheck,omitempty"`
	Notificationperiod  int    `json:"notificationperiod,omitempty"`
	Passcrypt           string `json:"passcrypt,omitempty"`
	Passplain           string `json:"passplain,omitempty"`
	Password            bool   `json:"password,omitempty"`
	Priority            int    `json:"priority,omitempty"`
	Publickey           string `json:"publickey,omitempty"`
	Publickeysize       int    `json:"publickeysize,omitempty"`
	Serial              string `json:"serial,omitempty"`
	Servicename         string `json:"servicename,omitempty"`
	Signaturealg        string `json:"signaturealg,omitempty"`
	Status              string `json:"status,omitempty"`
	Subject             string `json:"subject,omitempty"`
	Version             int    `json:"version,omitempty"`
}
