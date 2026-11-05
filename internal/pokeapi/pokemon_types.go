package pokeapi

type Pokemon struct {
	Id             int     `json:"id,omitempty"`
	Name           string  `json:"name,omitempty"`
	BaseExperience int     `json:"base_experience,omitempty"`
	Height         int     `json:"height,omitempty"`
	Weight         int     `json:"weight,omitempty"`
	Stats          []Stats `json:"stats,omitempty"`
	Types          []Types `json:"types,omitempty"`
}

type PokemonResp struct {
	Id                     int           `json:"id,omitempty"`
	Name                   string        `json:"name,omitempty"`
	BaseExperience         int           `json:"base_experience,omitempty"`
	Height                 int           `json:"height,omitempty"`
	IsDefault              bool          `json:"is_default,omitempty"`
	Order                  int           `json:"order,omitempty"`
	Weight                 int           `json:"weight,omitempty"`
	Abilities              []Abilities   `json:"abilities,omitempty"`
	PastAbilities          []interface{} `json:"past_abilities,omitempty"`
	Forms                  []Forms       `json:"forms,omitempty"`
	GameIndices            []GameIndices `json:"game_indices,omitempty"`
	HeldItems              []HeldItems   `json:"held_items,omitempty"`
	LocationAreaEncounters string        `json:"location_area_encounters,omitempty"`
	Moves                  []Moves       `json:"moves,omitempty"`
	Species                Species       `json:"species,omitempty"`
	Sprites                Sprites       `json:"sprites,omitempty"`
	Cries                  Cries         `json:"cries,omitempty"`
	Stats                  []Stats       `json:"stats,omitempty"`
	PastStats              []interface{} `json:"past_stats,omitempty"`
	Types                  []Types       `json:"types,omitempty"`
	PastTypes              []interface{} `json:"past_types,omitempty"`
}

type Ability struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Abilities struct {
	IsHidden bool    `json:"is_hidden,omitempty"`
	Slot     int     `json:"slot,omitempty"`
	Ability  Ability `json:"ability,omitempty"`
}

type Forms struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type GameIndices struct {
	GameIndex int     `json:"game_index,omitempty"`
	Version   Version `json:"version,omitempty"`
}

type Item struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type HeldItems struct {
	Item           Item             `json:"item,omitempty"`
	VersionDetails []VersionDetails `json:"version_details,omitempty"`
}

type Move struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type VersionGroup struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type MoveLearnMethod struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type VersionGroupDetails struct {
	LevelLearnedAt  int             `json:"level_learned_at,omitempty"`
	VersionGroup    VersionGroup    `json:"version_group,omitempty"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method,omitempty"`
	Order           interface{}     `json:"order,omitempty"`
}

type Moves struct {
	Move                Move                  `json:"move,omitempty"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details,omitempty"`
}

type Species struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Home struct {
	FrontShiny       string      `json:"front_shiny,omitempty"`
	FrontFemale      interface{} `json:"front_female,omitempty"`
	FrontDefault     string      `json:"front_default,omitempty"`
	FrontShinyFemale interface{} `json:"front_shiny_female,omitempty"`
}

type Showdown struct {
	BackShiny        string      `json:"back_shiny,omitempty"`
	BackFemale       interface{} `json:"back_female,omitempty"`
	FrontShiny       string      `json:"front_shiny,omitempty"`
	BackDefault      string      `json:"back_default,omitempty"`
	FrontFemale      interface{} `json:"front_female,omitempty"`
	FrontDefault     string      `json:"front_default,omitempty"`
	BackShinyFemale  interface{} `json:"back_shiny_female,omitempty"`
	FrontShinyFemale interface{} `json:"front_shiny_female,omitempty"`
}

type DreamWorld struct {
	FrontFemale  interface{} `json:"front_female,omitempty"`
	FrontDefault string      `json:"front_default,omitempty"`
}

type RedAndBlue struct {
	FrontDefault interface{} `json:"front_default,omitempty"`
}

type RedAndGreen struct {
	FrontDefault interface{} `json:"front_default,omitempty"`
}

type GenerationI struct {
	RedAndBlue  RedAndBlue  `json:"red-and-blue,omitempty"`
	RedAndGreen RedAndGreen `json:"red-and-green,omitempty"`
}

type GoldAndSilver struct {
	FrontDefault interface{} `json:"front_default,omitempty"`
}

type GenerationIi struct {
	GoldAndSilver GoldAndSilver `json:"gold-and-silver,omitempty"`
}

type Versions struct {
	GenerationI  GenerationI  `json:"generation-i,omitempty"`
	GenerationIi GenerationIi `json:"generation-ii,omitempty"`
}

type OfficialArtwork struct {
	Versions     Versions `json:"versions,omitempty"`
	FrontShiny   string   `json:"front_shiny,omitempty"`
	FrontDefault string   `json:"front_default,omitempty"`
}

type Other struct {
	Home            Home            `json:"home,omitempty"`
	Showdown        Showdown        `json:"showdown,omitempty"`
	DreamWorld      DreamWorld      `json:"dream_world,omitempty"`
	OfficialArtwork OfficialArtwork `json:"official-artwork,omitempty"`
}

type Sprites struct {
	Other            Other       `json:"other,omitempty"`
	Versions         Versions    `json:"versions,omitempty"`
	BackShiny        string      `json:"back_shiny,omitempty"`
	BackFemale       interface{} `json:"back_female,omitempty"`
	FrontShiny       string      `json:"front_shiny,omitempty"`
	BackDefault      string      `json:"back_default,omitempty"`
	FrontFemale      interface{} `json:"front_female,omitempty"`
	FrontDefault     string      `json:"front_default,omitempty"`
	BackShinyFemale  interface{} `json:"back_shiny_female,omitempty"`
	FrontShinyFemale interface{} `json:"front_shiny_female,omitempty"`
}

type Cries struct {
	Latest string `json:"latest,omitempty"`
	Legacy string `json:"legacy,omitempty"`
}

type Stat struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Stats struct {
	BaseStat int  `json:"base_stat,omitempty"`
	Effort   int  `json:"effort,omitempty"`
	Stat     Stat `json:"stat,omitempty"`
}

type Type struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Types struct {
	Slot int  `json:"slot,omitempty"`
	Type Type `json:"type,omitempty"`
}
