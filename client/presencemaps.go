package main

// Empty ones need to be kept!
var largeImageActivityModes = map[string][]string{
	"anniversary": {"Desafios da Eternidade"},
	"beyondlight": {"Caçada do Império"},
	"control":     {"Controle: Competitivo", "Controle: Jogo Rápido"},
	"crucible": {"Fissura", "Jogo Rápido PvP", "Salvage", "Supremacy", "Caos", "Caustificação em Equipe", "Scorched", "Sobrevivencia", "Co-Op Competitivo",
		"Breakthrough", "PvP Competitivo", "Momentum", "Controle de Zonas", "Clash: Competitive", "Countdown", "Showdown", "Eliminação",
		"Briga", "Clash", "Lockdown", "Controle de Impulso", "The Crucible", "Classic Mix"},
	"destinylogo":        {},
	"doubles":            {"All Doubles"},
	"dungeon":            {},
	"explore":            {},
	"forge":              {"Forge Ignition"},
	"gambit":             {},
	"hauntedforest":      {},
	"ironbanner":         {},
	"lostsector":         {},
	"menagerie":          {"The Menagerie"},
	"nightmarehunt":      {},
	"privatecrucible":    {"Partida Privada"},
	"raid":               {},
	"reckoning":          {"The Reckoning"},
	"seasonchosen":       {},
	"seasonhaunted":      {},
	"seasonlost":         {},
	"seasonrisen":        {},
	"seasonsplicer":      {},
	"shadowkeep":         {},
	"socialall":          {"Social", "All"},
	"storypvecoopheroic": {"Aventura Heroica", "Ofensiva", "História", "PvE"},
	"strikes":            {"Scored Prestige Nightfall", "Scored Nightfall Strikes", "Assaltos do Anoitecer", "Anoitecer Prestígio", "Assalto", "Operação da Vanguarda", "Anoitecer"},
	"thenine":            {"Trials of the Nine", "Trials of the Nine Countdown", "Trials of the Nine Survival"},
	"thewitchqueen":      {},
	"trialsofosiris":     {"The Sundial", "Simulação do Farol"},
	"vexoffensive":       {},
	"wellspring":         {},
}

// Maps classID to class name
var classImages = map[int32]string{
	0: "Titã",
	1: "Caçador",
	2: "Arcano",
}

var scoredLostSectors = []string{
	"K1",                            // The Moon
	"Vazio Oculto", "Bunker E15", "Perdição", // Europa
	"Jardim da Êxodo 2A", "Labirinto de Veles", "A Pedreira", // The Cosmodrome
	"Covil do Oportunista", "Local de Escavação XII", // EDZ
	"Baia Dos Desejos Afogados", "Câmara da Luz Estelar", "Recanto do Afélio", // Dreaming City
	"Metamorfose", "Sepulcro", "Extração", // Savathûn's Throne World
	"O Confluxo", // Nessus
}
var storyMissions = map[string][]string{
	"shadowkeep":    {"A Mysterious Disturbance", "In Search of Answers", "Ghosts of Our Past", "In the Deep", "Beyond"},
	"beyondlight":   {"Darkness's Doorstep", "The New Kell", "Rising Resistance", "The Warrior", "The Technocrat", "The Kell of Darkness", "Sabotaging Salvation", "The Aftermath"},
	"thewitchqueen": {"The Arrival", "The Investigation", "The Ghosts", "The Communion", "The Mirror", "The Cunning", "The Last Chance", "The Ritual", "Memories of"},
	"lightfall":     {"First Contact", "Under Siege", "Downfall", "Breakneck", "On the Verge", "No Time Left", "Headlong", "Desperate Measures"},
}

var raidProgressionMap = map[string][]string{
	"Garden of Salvation": {"Embrace", "Undergrowth", "The Consecrated Mind", "The Sanctified Mind"},                                                   // 4
	"Last Wish":           {"Kalli, The Corrupted", "Shuro Chi, The Corrupted", "Morgeth, The Spirekeeper", "The Vault", "Riven of a Thousand Voices"}, // 5
	"Deep Stone Crypt":    {"Desolation / Crypt Security", "Atraks-1, Fallen Exo", "The Descent", "Taniks, The Abomination"},                           // 4
	"Vault of Glass":      {"Raise the Spire / Confluxes", "The Oracles", "The Templar", "The Gatekeeper", "Atheon, Time's Conflux"},                   // 5
	"Vow of the Disciple": {"Acquisition", "Collection", "Exhibition", "Dominion"},                                                                     // 4
	"King's Fall":         {"Totems", "Warpriest", "Golgoroth", "Daughters of Oryx", "Oryx, The Taken King"},                                           // 5
}