package requests

type AddNewCat struct {
	CatBreed             string   `json:"cat_breed"`
	CatOriginDescription string   `json:"cat_origin_description"`
	CatType              string   `json:"cat_type"`
	CatTypeInfo          *string  `json:"cat_type_info"`
	BodyType             []string `json:"body_type"`
	CoatPattern          string   `json:"coat_pattern"`
}

type UpdateCatTypeInfo struct {
	CatTypeInfo *string `json:"cat_type_info"`
}
