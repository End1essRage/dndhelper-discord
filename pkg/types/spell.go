package types

//вынести в отдельный модуль
type Spell struct {
	Index         string     `json:"index"`
	Name          string     `json:"name"`
	Desc          []string   `json:"desc"`
	HigherLevel   []string   `json:"higher_level"`
	Range         string     `json:"range"`
	Components    []string   `json:"components"`
	Material      string     `json:"material"`
	Ritual        bool       `json:"ritual"`
	Duration      string     `json:"duration"`
	Concentration bool       `json:"concentration"`
	CastingTime   string     `json:"casting_time"`
	Level         int        `json:"level"`
	Damage        Damage     `json:"damage"`
	Dc            Dc         `json:"dc"`
	AreaOfEffect  Aoef       `json:"area_of_effect"`
	School        School     `json:"school"`
	Classes       []Class    `json:"classes"`
	Subclasses    []Subclass `json:"subclasses"`
	Url           string     `json:"url"`
}

// Damage represents a damage information
type Damage struct {
	DamageType        DamageType        `json:"damage_type"`
	DamageAtSlotLevel map[string]string `json:"damage_at_slot_level"`
}

// Dc represents a difficulty class information
type Dc struct {
	DcType    DcType `json:"dc_type"`
	DcSuccess string `json:"dc_success"`
}

// Aoef represents an area of effect information
type Aoef struct {
	Type string `json:"type"`
	Size int    `json:"size"`
}

// School represents a magic school information
type School struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

// Class represents a class information
type Class struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

// Subclass represents a subclass information
type Subclass struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

// DamageType represents a damage type information
type DamageType struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

// DcType represents a dc type information
type DcType struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}
