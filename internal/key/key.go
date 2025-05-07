package key

// KeyWrapper: API responses do wrap the keys.
type KeyWrapper struct {
	Key Key `json:"key"`
}

// Key: Representation used in list key responses.
type Key struct {
	Name        string `json:"name" yaml:"name"`
	Fingerprint string `json:"fingerprint" yaml:"fingerprint"`
	Type        string `json:"type" yaml:"type"`
	Size        int    `json:"size" yaml:"size"`
	Data        string `json:"data" yaml:"data"`
	CreatedAt   string `json:"created_at" yaml:"created_at"`
}
