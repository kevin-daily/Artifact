package main

type Data struct {
	Slots               int    `json:"slots"`
	Expansions          int    `json:"expansions"`
	Next_Expansion_Cost int    `json:"next_expansion_cost"`
	Gold                int    `json:"gold"`
	Type                string `json:"type"`
	Code                string `json:"code"`
	//Character           []Character
	Character   Character   `json:"character"`
	Destination Destination `json:"destination"`
	Content     Content     `json:"content"`
	Cooldown    Cooldown    `json:"cooldown"`
}

type Character struct {
	Name                   string
	Account                string
	Skin                   string
	Level                  int
	XP                     int
	Max_XP                 int
	Gold                   int
	Speed                  int
	Mining_Level           int
	Mining_XP              int
	Mining_Max_XP          int
	Woodcutting_Level      int
	Woodcutting_XP         int
	Woodcutting_Max_XP     int
	Fishing_Level          int
	Fishing_XP             int
	Fishing_Max_XP         int
	Weaponcrafting_Level   int
	Weaponcrafting_XP      int
	Weaponcrafting_Max_XP  int
	Gearcrafting_Level     int
	Gearcrafting_XP        int
	Gearcrafting_Max_XP    int
	Jewelrycrafting_Level  int
	Jewelrycrafting_XP     int
	Jewelrycrafting_Max_XP int
	Cooking_Level          int
	Cooking_XP             int
	Cooking_Max_XP         int
	Alchemy_Level          int
	Alchemy_XP             int
	Alchemy_Max_XP         int
	HP                     int
	Max_HP                 int
	Haste                  int
	Critical_Strike        int
	Wisdom                 int
	Prospecting            int
	Attack_Fire            int
	Attack_Earth           int
	Attack_Water           int
	Attack_Air             int
	DMG                    int
	DMG_Fire               int
	DMG_Earth              int
	DMG_Water              int
	DMG_Air                int
	Res_Fire               int
	Res_Earth              int
	Res_Water              int
	Res_Air                int
	X                      int
	Y                      int
	Cooldown               int
	Cooldown_Expiration    string
	Weapon_Slot            string
	Rune_Slot              string
	Shield_Slot            string
	Helmet_Slot            string
	Body_Armor_Slot        string
	Leg_Armor_Slot         string
	Boots_Slot             string
	Ring1_Slot             string
	Ring2_Slot             string
	Amulet_Slot            string
	Artifact1_Slot         string
	Artifact2_Slot         string
	Artifact3_Slot         string
	Utility1_Slot          string
	Utility1_Slot_Quantity int
	Utility2_Slot          string
	Utility2_Slot_Quantity int
	Bag_Slot               string
	Task                   string
	Task_Type              string
	Task_Progress          int
	Task_Total             int
	Inventory_Max_Items    int
	Inventory              []Inventory
}

type Content struct {
	Type string
	Code string
}

type Cooldown struct {
	Total_Seconds     int
	Remaining_Seconds int
	Started_At        string
	Expiration        string
	Reason            string
}

type Destination struct {
	Name    string
	Skin    string
	X       int
	Y       int
	Content Content
}

type Inventory struct {
	Slot     int
	Code     string
	Quantity int
}
