package config

type ModConfig struct {
	Video struct {
		BorderlessWindow   bool   `ini:"borderless_window"`
		Windowed           bool   `ini:"windowed"`
		BattleResolution   string `ini:"battle_resolution"`
		CampaignResolution string `ini:"campaign_resolution"`
		Bloom              bool   `ini:"bloom"`
	} `ini:"video"`
}

type AGOConfig struct {
	Debug struct {
		EnableLogging bool `ini:"enable_logging"`
		DevDebug      bool `ini:"dev_debug"`
		LogToConsole  bool `ini:"log_to_console"`
	} `ini:"debug"`

	Sorting struct {
		EnableSorting bool `ini:"enable_sorting"`
		SortMode1     int  `ini:"sortmode1"`
		SortMode2     int  `ini:"sortmode2"`
		SortMode3     int  `ini:"sortmode3"`
		SortPlayer    bool `ini:"sort_player"`
	} `ini:"sorting"`

	Limits struct {
		MaximumAncillaries int `ini:"maximum_ancillaries"`
		GuildCooldown      int `ini:"guild_cooldown"`
	} `ini:"limits"`

	Saving struct {
		PostBattleSaving bool `ini:"post_battle_saving"`
	} `ini:"saving"`

	Info struct {
		HideArmyInfo       bool `ini:"hide_army_info"`
		AIRaidNotification bool `ini:"ai_raid_notification"`
		WatchtowerRadius   int  `ini:"watchtower_radius"`
	} `ini:"info"`

	Scripts struct {
		NaturalDisasters   bool `ini:"natural_disasters"`
		RandomAaAiStart    bool `ini:"random_aa_ai_start"`
		MergeDolAmroth     bool `ini:"merge_dol_amroth"`
		RandomizedStart    bool `ini:"randomized_start"`
		ShatteredAlliances bool `ini:"shattered_alliances"`
		LastStandArmies    bool `ini:"last_stand_armies"`
	} `ini:"scripts"`

	Battle struct {
		NoDefaultSkirmish  bool `ini:"no_default_skirmish"`
		DefaultBattleSpeed int  `ini:"default_battle_speed"`
	} `ini:"battle"`

	Difficulty struct {
		AggressiveRebels bool `ini:"aggressive_rebels"`
		AIFreeGenerals   bool `ini:"ai_free_generals"`
	} `ini:"difficulty"`
}

type EOPConfig struct {
	UiCfg   struct{}
	GameCfg struct {
		IsOverrideBattleCamera       bool `json:"IsOverrideBattleCamera"`
		IsBlockLaunchWithoutEop      bool `json:"isBlockLaunchWithoutEop"`
		IsContextMenuNeeded          bool `json:"isContextMenuNeeded"`
		IsDXVKEnabled                bool `json:"isDXVKEnabled"`
		IsDeveloperModeNeeded        bool `json:"isDeveloperModeNeeded"`
		IsDiscordRichPresenceEnabled bool `json:"isDiscordRichPresenceEnabled"`
		IsFreecamIntegrationEnabled  bool `json:"isFreecamIntegrationEnabled"`
		IsSaveBackupEnabled          bool `json:"isSaveBackupEnabled"`
		IsTacticalMapViewerNeeded    bool `json:"isTacticalMapViewerNeeded"`
	}
	BattlesCfg struct {
		EnableAutoGeneration  bool `json:"enableAutoGeneration"`
		EnableResultsTransfer bool `json:"enableResultsTransfer"`
		IsPlannedRetreatRoute bool `json:"isPlannedRetreatRoute"`
	}
}
