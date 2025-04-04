package document

type Document struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Src         []byte `json:"src"`
}

type uploadResponse struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}
