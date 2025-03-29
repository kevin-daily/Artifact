package main

// Top level
type ActionCraft struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Details   Details   `json:"details"`
	Character Character `json:"character"`
}

type ActionFight struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Fight     Fight     `json:"fight"`
	Character Character `json:"character"`
}

type ActionGathering struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Details   Details   `json:"details"`
	Character Character `json:"character"`
}

type ActionUnequip struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Slot      string    `json:"slot"`
	Item      Item      `json:"item"`
	Character Character `json:"character"`
}

type DWBank struct {
	Cooldown  Cooldown  `json:"cooldown"`
	Item      Item      `json:"item"`
	Bank      []Drops   `json:"bank"`
	Character Character `json:"character"`
}

type BankDetails struct {
	Slots               int `json:"slots"`
	Expansions          int `json:"expansions"`
	Next_Expansion_Cost int `json:"next_expansion_cost"`
	Gold                int `json:"gold"`
}

type Move struct {
	Cooldown    Cooldown    `json:"cooldown"`
	Destination Destination `json:"destination"`
	Character   Character   `json:"character"`
}

type Rest struct {
	Cooldown    Cooldown  `json:"cooldown"`
	HP_Restored int       `json:"hp_restored"`
	Character   Character `json:"character"`
}

type StatusData struct {
	Status            string          `json:"status"`
	Version           string          `json:"version"`
	Max_Level         int             `json:"max_level"`
	Characters_Online int             `json:"characters_online"`
	Server_Time       string          `json:"server_time"`
	Announcements     []Announcements `json:"announcements"`
	Last_Wipe         string          `json:"last_wipe"`
	Next_Wipe         string          `json:"next_wipe"`
}

// Helper level
type Announcements struct {
	Message    string `json:"message"`
	Created_At string `json:"created_at"`
}

type BankPayload struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type Character struct {
	Name                   string      `json:"name"`
	Account                string      `json:"account"`
	Skin                   string      `json:"skin"`
	Level                  int         `json:"level"`
	XP                     int         `json:"xp"`
	Max_XP                 int         `json:"max_xp"`
	Gold                   int         `json:"gold"`
	Speed                  int         `json:"speed"`
	Mining_Level           int         `json:"mining_level"`
	Mining_XP              int         `json:"mining_xp"`
	Mining_Max_XP          int         `json:"mining_max_xp"`
	Woodcutting_Level      int         `json:"woodcutting_level"`
	Woodcutting_XP         int         `json:"woodcutting_xp"`
	Woodcutting_Max_XP     int         `json:"woodcutting_max_xp"`
	Fishing_Level          int         `json:"fishing_level"`
	Fishing_XP             int         `json:"fishing_xp"`
	Fishing_Max_XP         int         `json:"fishing_max_xp"`
	Weaponcrafting_Level   int         `json:"weaponcrafting_level"`
	Weaponcrafting_XP      int         `json:"weaponcrafting_xp"`
	Weaponcrafting_Max_XP  int         `json:"weaponcrafting_max_xp"`
	Gearcrafting_Level     int         `json:"gearcrafting_level"`
	Gearcrafting_XP        int         `json:"gearcrafting_xp"`
	Gearcrafting_Max_XP    int         `json:"gearcrafting_max_xp"`
	Jewelrycrafting_Level  int         `json:"jewelrycrafting_level"`
	Jewelrycrafting_XP     int         `json:"jewelrycrafting_xp"`
	Jewelrycrafting_Max_XP int         `json:"jewelrycrafting_max_xp"`
	Cooking_Level          int         `json:"cooking_level"`
	Cooking_XP             int         `json:"cooking_xp"`
	Cooking_Max_XP         int         `json:"cooking_max_xp"`
	Alchemy_Level          int         `json:"alchemy_level"`
	Alchemy_XP             int         `json:"alchemy_xp"`
	Alchemy_Max_XP         int         `json:"alchemy_max_xp"`
	HP                     int         `json:"hp"`
	Max_HP                 int         `json:"max_hp"`
	Haste                  int         `json:"haste"`
	Critical_Strike        int         `json:"critical_strike"`
	Wisdom                 int         `json:"wisdom"`
	Prospecting            int         `json:"prospecting"`
	Attack_Fire            int         `json:"attack_fire"`
	Attack_Earth           int         `json:"attack_earth"`
	Attack_Water           int         `json:"attack_water"`
	Attack_Air             int         `json:"attack_air"`
	DMG                    int         `json:"dmg"`
	DMG_Fire               int         `json:"dmg_fire"`
	DMG_Earth              int         `json:"dmg_earth"`
	DMG_Water              int         `json:"dmg_water"`
	DMG_Air                int         `json:"dmg_air"`
	Res_Fire               int         `json:"res_fire"`
	Res_Earth              int         `json:"res_earth"`
	Res_Water              int         `json:"res_water"`
	Res_Air                int         `json:"res_air"`
	X                      int         `json:"x"`
	Y                      int         `json:"y"`
	Cooldown               int         `json:"cooldown"`
	Cooldown_Expiration    string      `json:"cooldown_expiration"`
	Weapon_Slot            string      `json:"weapon_slot"`
	Rune_Slot              string      `json:"rune_slot"`
	Shield_Slot            string      `json:"shield_slot"`
	Helmet_Slot            string      `json:"helmet_slot"`
	Body_Armor_Slot        string      `json:"body_armor_slot"`
	Leg_Armor_Slot         string      `json:"leg_armor_slot"`
	Boots_Slot             string      `json:"boots_slot"`
	Ring1_Slot             string      `json:"ring1_slot"`
	Ring2_Slot             string      `json:"ring2_slot"`
	Amulet_Slot            string      `json:"amulet_slot"`
	Artifact1_Slot         string      `json:"artifact1_slot"`
	Artifact2_Slot         string      `json:"artifact2_slot"`
	Artifact3_Slot         string      `json:"artifact3_slot"`
	Utility1_Slot          string      `json:"utility1_slot"`
	Utility1_Slot_Quantity int         `json:"utility1_slot_quantity"`
	Utility2_Slot          string      `json:"utility2_slot"`
	Utility2_Slot_Quantity int         `json:"utility2_slot_quantity"`
	Bag_Slot               string      `json:"bag_slot"`
	Task                   string      `json:"task"`
	Task_Type              string      `json:"task_type"`
	Task_Progress          int         `json:"task_progress"`
	Task_Total             int         `json:"task_total"`
	Inventory_Max_Items    int         `json:"inventory_max_items"`
	Inventory              []Inventory `json:"inventory"`
}

type Content struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type Cooldown struct {
	Total_Seconds     int    `json:"total_seconds"`
	Remaining_Seconds int    `json:"remaining_seconds"`
	Started_At        string `json:"started_at"`
	Expiration        string `json:"expiration"`
	Reason            string `json:"reason"`
}

type Craft struct {
	Skill    string  `json:"skill"`
	Level    int     `json:"level"`
	Items    []Drops `json:"items"`
	Quantity int     `json:"quantity"`
}

type Destination struct {
	Name    string  `json:"name"`
	Skin    string  `json:"skin"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
	Content Content `json:"content"`
}

type Details struct {
	XP    int     `json:"xp"`
	Items []Drops `json:"items"`
}

type Drops struct {
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type Effects struct {
	Code  string `json:"code"`
	Value int    `json:"value"`
}

type Fight struct {
	XP                   int                  `json:"xp"`
	Gold                 int                  `json:"gold"`
	Drops                []Drops              `json:"drops"`
	Turns                int                  `json:"turns"`
	Monster_Blocked_Hits Monster_Blocked_Hits `json:"monster_blocked_hits"`
	Player_Blocked_Hits  Monster_Blocked_Hits `json:"player_blocked_hits"`
	Logs                 []string             `json:"logs"`
	Result               string               `json:"result"`
}

type Inventory struct {
	Slot     int    `json:"slot"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type Item struct {
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Level       int       `json:"level"`
	Type        string    `json:"type"`
	Subtype     string    `json:"subtype"`
	Description string    `json:"description"`
	Effects     []Effects `json:"effects"`
	Craft       Craft     `json:"craft"`
	Tradeable   bool      `json:"tradeable"`
}
type Monster_Blocked_Hits struct {
	Fire  int `json:"fire"`
	Earth int `json:"earth"`
	Water int `json:"water"`
	Air   int `json:"air"`
	Total int `json:"total"`
}
