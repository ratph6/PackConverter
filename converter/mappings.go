package converter

// BlockTextureRenames: assets/minecraft/textures/blocks/ → block/
var BlockTextureRenames = map[string]string{
	// Wool
	"wool_colored_white":      "white_wool",
	"wool_colored_orange":     "orange_wool",
	"wool_colored_magenta":    "magenta_wool",
	"wool_colored_light_blue": "light_blue_wool",
	"wool_colored_yellow":     "yellow_wool",
	"wool_colored_lime":       "lime_wool",
	"wool_colored_pink":       "pink_wool",
	"wool_colored_gray":       "gray_wool",
	"wool_colored_silver":     "light_gray_wool",
	"wool_colored_cyan":       "cyan_wool",
	"wool_colored_purple":     "purple_wool",
	"wool_colored_blue":       "blue_wool",
	"wool_colored_brown":      "brown_wool",
	"wool_colored_green":      "green_wool",
	"wool_colored_red":        "red_wool",
	"wool_colored_black":      "black_wool",

	// Stained Glass
	"glass_white":      "white_stained_glass",
	"glass_orange":     "orange_stained_glass",
	"glass_magenta":    "magenta_stained_glass",
	"glass_light_blue": "light_blue_stained_glass",
	"glass_yellow":     "yellow_stained_glass",
	"glass_lime":       "lime_stained_glass",
	"glass_pink":       "pink_stained_glass",
	"glass_gray":       "gray_stained_glass",
	"glass_silver":     "light_gray_stained_glass",
	"glass_cyan":       "cyan_stained_glass",
	"glass_purple":     "purple_stained_glass",
	"glass_blue":       "blue_stained_glass",
	"glass_brown":      "brown_stained_glass",
	"glass_green":      "green_stained_glass",
	"glass_red":        "red_stained_glass",
	"glass_black":      "black_stained_glass",

	// Stained Glass Pane Tops
	"glass_pane_top_white":      "white_stained_glass_pane_top",
	"glass_pane_top_orange":     "orange_stained_glass_pane_top",
	"glass_pane_top_magenta":    "magenta_stained_glass_pane_top",
	"glass_pane_top_light_blue": "light_blue_stained_glass_pane_top",
	"glass_pane_top_yellow":     "yellow_stained_glass_pane_top",
	"glass_pane_top_lime":       "lime_stained_glass_pane_top",
	"glass_pane_top_pink":       "pink_stained_glass_pane_top",
	"glass_pane_top_gray":       "gray_stained_glass_pane_top",
	"glass_pane_top_silver":     "light_gray_stained_glass_pane_top",
	"glass_pane_top_cyan":       "cyan_stained_glass_pane_top",
	"glass_pane_top_purple":     "purple_stained_glass_pane_top",
	"glass_pane_top_blue":       "blue_stained_glass_pane_top",
	"glass_pane_top_brown":      "brown_stained_glass_pane_top",
	"glass_pane_top_green":      "green_stained_glass_pane_top",
	"glass_pane_top_red":        "red_stained_glass_pane_top",
	"glass_pane_top_black":      "black_stained_glass_pane_top",

	// Terracotta (Hardened Clay)
	"hardened_clay":                    "terracotta",
	"hardened_clay_stained_white":      "white_terracotta",
	"hardened_clay_stained_orange":     "orange_terracotta",
	"hardened_clay_stained_magenta":    "magenta_terracotta",
	"hardened_clay_stained_light_blue": "light_blue_terracotta",
	"hardened_clay_stained_yellow":     "yellow_terracotta",
	"hardened_clay_stained_lime":       "lime_terracotta",
	"hardened_clay_stained_pink":       "pink_terracotta",
	"hardened_clay_stained_gray":       "gray_terracotta",
	"hardened_clay_stained_silver":     "light_gray_terracotta",
	"hardened_clay_stained_cyan":       "cyan_terracotta",
	"hardened_clay_stained_purple":     "purple_terracotta",
	"hardened_clay_stained_blue":       "blue_terracotta",
	"hardened_clay_stained_brown":      "brown_terracotta",
	"hardened_clay_stained_green":      "green_terracotta",
	"hardened_clay_stained_red":        "red_terracotta",
	"hardened_clay_stained_black":      "black_terracotta",

	// Planks
	"planks_oak":     "oak_planks",
	"planks_spruce":  "spruce_planks",
	"planks_birch":   "birch_planks",
	"planks_jungle":  "jungle_planks",
	"planks_acacia":  "acacia_planks",
	"planks_big_oak": "dark_oak_planks",

	// Logs
	"log_oak":         "oak_log",
	"log_oak_top":     "oak_log_top",
	"log_spruce":      "spruce_log",
	"log_spruce_top":  "spruce_log_top",
	"log_birch":       "birch_log",
	"log_birch_top":   "birch_log_top",
	"log_jungle":      "jungle_log",
	"log_jungle_top":  "jungle_log_top",
	"log_acacia":      "acacia_log",
	"log_acacia_top":  "acacia_log_top",
	"log_big_oak":     "dark_oak_log",
	"log_big_oak_top": "dark_oak_log_top",

	// Leaves
	"leaves_oak":     "oak_leaves",
	"leaves_spruce":  "spruce_leaves",
	"leaves_birch":   "birch_leaves",
	"leaves_jungle":  "jungle_leaves",
	"leaves_acacia":  "acacia_leaves",
	"leaves_big_oak": "dark_oak_leaves",

	// Saplings
	"sapling_oak":        "oak_sapling",
	"sapling_spruce":     "spruce_sapling",
	"sapling_birch":      "birch_sapling",
	"sapling_jungle":     "jungle_sapling",
	"sapling_acacia":     "acacia_sapling",
	"sapling_roofed_oak": "dark_oak_sapling",

	// Doors
	"door_wood_upper":     "oak_door_top",
	"door_wood_lower":     "oak_door_bottom",
	"door_spruce_upper":   "spruce_door_top",
	"door_spruce_lower":   "spruce_door_bottom",
	"door_birch_upper":    "birch_door_top",
	"door_birch_lower":    "birch_door_bottom",
	"door_jungle_upper":   "jungle_door_top",
	"door_jungle_lower":   "jungle_door_bottom",
	"door_acacia_upper":   "acacia_door_top",
	"door_acacia_lower":   "acacia_door_bottom",
	"door_dark_oak_upper": "dark_oak_door_top",
	"door_dark_oak_lower": "dark_oak_door_bottom",
	"door_iron_upper":     "iron_door_top",
	"door_iron_lower":     "iron_door_bottom",

	// Stone types
	"stone_granite":         "granite",
	"stone_granite_smooth":  "polished_granite",
	"stone_diorite":         "diorite",
	"stone_diorite_smooth":  "polished_diorite",
	"stone_andesite":        "andesite",
	"stone_andesite_smooth": "polished_andesite",

	// Sandstone
	"sandstone_normal":     "sandstone",
	"sandstone_carved":     "chiseled_sandstone",
	"sandstone_smooth":     "cut_sandstone",
	"red_sandstone_normal": "red_sandstone",
	"red_sandstone_carved": "chiseled_red_sandstone",
	"red_sandstone_smooth": "cut_red_sandstone",

	// Stone Bricks
	"stonebrick":         "stone_bricks",
	"stonebrick_mossy":   "mossy_stone_bricks",
	"stonebrick_cracked": "cracked_stone_bricks",
	"stonebrick_carved":  "chiseled_stone_bricks",

	// Quartz
	"quartz_block_chiseled":     "chiseled_quartz_block",
	"quartz_block_chiseled_top": "chiseled_quartz_block_top",
	"quartz_block_lines":        "quartz_pillar",
	"quartz_block_lines_top":    "quartz_pillar_top",

	// Grass / Dirt
	"grass_top":          "grass_block_top",
	"grass_side":         "grass_block_side",
	"grass_side_overlay": "grass_block_side_overlay",
	"grass_side_snowed":  "grass_block_snow",
	"dirt_podzol_side":   "podzol_side",
	"dirt_podzol_top":    "podzol_top",

	// Nether
	"nether_brick":        "nether_bricks",
	"quartz_ore":          "nether_quartz_ore",
	"nether_wart_stage_0": "nether_wart_stage0",
	"nether_wart_stage_1": "nether_wart_stage1",
	"nether_wart_stage_2": "nether_wart_stage2",

	// End
	"end_bricks":    "end_stone_bricks",
	"endframe_top":  "end_portal_frame_top",
	"endframe_side": "end_portal_frame_side",
	"endframe_eye":  "end_portal_frame_eye",

	// Prismarine
	"prismarine_rough": "prismarine",
	"prismarine_dark":  "dark_prismarine",

	// Crops
	"wheat_stage_0":    "wheat_stage0",
	"wheat_stage_1":    "wheat_stage1",
	"wheat_stage_2":    "wheat_stage2",
	"wheat_stage_3":    "wheat_stage3",
	"wheat_stage_4":    "wheat_stage4",
	"wheat_stage_5":    "wheat_stage5",
	"wheat_stage_6":    "wheat_stage6",
	"wheat_stage_7":    "wheat_stage7",
	"carrots_stage_0":  "carrots_stage0",
	"carrots_stage_1":  "carrots_stage1",
	"carrots_stage_2":  "carrots_stage2",
	"carrots_stage_3":  "carrots_stage3",
	"potatoes_stage_0": "potatoes_stage0",
	"potatoes_stage_1": "potatoes_stage1",
	"potatoes_stage_2": "potatoes_stage2",
	"potatoes_stage_3": "potatoes_stage3",
	"beetroots_stage_0": "beetroots_stage0",
	"beetroots_stage_1": "beetroots_stage1",
	"beetroots_stage_2": "beetroots_stage2",
	"beetroots_stage_3": "beetroots_stage3",
	"melon_stem_disconnected":  "melon_stem",
	"melon_stem_connected":     "attached_melon_stem",
	"pumpkin_stem_disconnected": "pumpkin_stem",
	"pumpkin_stem_connected":    "attached_pumpkin_stem",
	"cocoa_stage_0": "cocoa_stage0",
	"cocoa_stage_1": "cocoa_stage1",
	"cocoa_stage_2": "cocoa_stage2",

	// Flowers
	"flower_rose":               "poppy",
	"flower_dandelion":          "dandelion",
	"flower_blue_orchid":        "blue_orchid",
	"flower_allium":             "allium",
	"flower_houstonia":          "azure_bluet",
	"flower_tulip_red":          "red_tulip",
	"flower_tulip_orange":       "orange_tulip",
	"flower_tulip_white":        "white_tulip",
	"flower_tulip_pink":         "pink_tulip",
	"flower_oxeye_daisy":        "oxeye_daisy",
	"flower_paeonia_top":        "peony_top",
	"flower_paeonia_bottom":     "peony_bottom",
	"double_plant_rose_top":     "rose_bush_top",
	"double_plant_rose_bottom":  "rose_bush_bottom",
	"double_plant_sunflower_top":    "sunflower_top",
	"double_plant_sunflower_bottom": "sunflower_bottom",
	"double_plant_sunflower_front":  "sunflower_front",
	"double_plant_sunflower_back":   "sunflower_back",
	"double_plant_syringa_top":      "lilac_top",
	"double_plant_syringa_bottom":   "lilac_bottom",
	"double_plant_grass_top":        "tall_grass_top",
	"double_plant_grass_bottom":     "tall_grass_bottom",
	"double_plant_fern_top":         "large_fern_top",
	"double_plant_fern_bottom":      "large_fern_bottom",

	// Redstone
	"redstone_lamp_off":  "redstone_lamp",
	"redstone_torch_on":  "redstone_torch",
	"comparator_off":     "comparator",
	"repeater_off":       "repeater",
	"piston_top_normal":  "piston_top",

	// Torches / Misc lighting
	"torch_on": "torch",

	// Furnace / Dispensers / Droppers
	"furnace_front_off":          "furnace_front",
	"dispenser_front_horizontal": "dispenser_front",
	"dispenser_front_vertical":   "dispenser_front_vertical",
	"dropper_front_horizontal":   "dropper_front",
	"dropper_front_vertical":     "dropper_front_vertical",

	// Pumpkin
	"pumpkin_face_off": "carved_pumpkin",
	"pumpkin_face_on":  "jack_o_lantern",

	// Rails
	"rail_normal":            "rail",
	"rail_normal_turned":     "rail_corner",
	"rail_golden":            "powered_rail",
	"rail_golden_powered":    "powered_rail_on",
	"rail_detector":          "detector_rail",
	"rail_detector_powered":  "detector_rail_on",
	"rail_activator":         "activator_rail",
	"rail_activator_powered": "activator_rail_on",

	// Misc blocks
	"web":              "cobweb",
	"sponge_wet":       "wet_sponge",
	"noteblock":        "note_block",
	"mob_spawner":      "spawner",
	"waterlily":        "lily_pad",
	"tallgrass":        "short_grass",
	"deadbush":         "dead_bush",
	"trapdoor":         "oak_trapdoor",
	"ice_packed":       "packed_ice",
	"slime":            "slime_block",
	"trip_wire_source":  "tripwire_hook",
	"trip_wire":         "tripwire",
	"portal":            "nether_portal",
	"brick":             "bricks",
	"stone_slab_top":    "smooth_stone",
	"stone_slab_side":   "smooth_stone_slab_side",
	"reeds":             "sugar_cane",

	// Farmland
	"farmland_dry": "farmland",
	"farmland_wet": "farmland_moist",

	// Mossy cobblestone (alternate old name)
	"cobblestone_mossy": "mossy_cobblestone",

	// Fire layers
	"fire_layer_0": "fire_0",
	"fire_layer_1": "fire_1",

	// Fence gate (no wood prefix in 1.8.9)
	"fence_gate": "oak_fence_gate",

	// Double plant paeonia (alternate naming)
	"double_plant_paeonia_top":    "peony_top",
	"double_plant_paeonia_bottom": "peony_bottom",

	// Anvil
	"anvil_base":          "anvil",
	"anvil_top_damaged_0": "anvil_top",
	"anvil_top_damaged_1": "chipped_anvil_top",
	"anvil_top_damaged_2": "damaged_anvil_top",

	// Mushroom blocks
	"mushroom_block_skin_brown": "brown_mushroom_block",
	"mushroom_block_skin_red":   "red_mushroom_block",
	"mushroom_block_skin_stem":  "mushroom_stem",
	// mushroom_block_inside stays the same name in 1.21.10
	"mushroom_brown":            "brown_mushroom",
	"mushroom_red":              "red_mushroom",

	// Command blocks (names unchanged in 1.21.10, but kept for reference)
}

// ItemTextureRenames: assets/minecraft/textures/items/ → item/
var ItemTextureRenames = map[string]string{
	// Food
	"apple_golden":     "golden_apple",
	"beef_raw":         "beef",
	"beef_cooked":      "cooked_beef",
	"chicken_raw":      "chicken",
	"chicken_cooked":   "cooked_chicken",
	"porkchop_raw":     "porkchop",
	"porkchop_cooked":  "cooked_porkchop",
	"mutton_raw":       "mutton",
	"mutton_cooked":    "cooked_mutton",
	"rabbit_raw":       "rabbit",
	"rabbit_cooked":    "cooked_rabbit",
	"fish_cod_raw":     "cod",
	"fish_cod_cooked":  "cooked_cod",
	"fish_salmon_raw":     "salmon",
	"fish_salmon_cooked":  "cooked_salmon",
	"fish_clownfish_raw":  "tropical_fish",
	"fish_pufferfish_raw": "pufferfish",
	"melon":               "melon_slice",
	"melon_speckled":      "glistering_melon_slice",
	"potato_baked":        "baked_potato",
	"potato_poisonous":    "poisonous_potato",
	"carrot_golden":       "golden_carrot",

	// Tools (gold → golden, wood → wooden)
	"gold_axe":     "golden_axe",
	"gold_hoe":     "golden_hoe",
	"gold_pickaxe": "golden_pickaxe",
	"gold_shovel":  "golden_shovel",
	"gold_sword":   "golden_sword",
	"wood_axe":     "wooden_axe",
	"wood_hoe":     "wooden_hoe",
	"wood_pickaxe": "wooden_pickaxe",
	"wood_shovel":  "wooden_shovel",
	"wood_sword":   "wooden_sword",

	// Armor (gold → golden)
	"gold_boots":      "golden_boots",
	"gold_chestplate": "golden_chestplate",
	"gold_helmet":     "golden_helmet",
	"gold_leggings":   "golden_leggings",

	// Horse armor
	"gold_horse_armor": "golden_horse_armor",

	// Doors (item textures)
	"door_wood":     "oak_door",
	"door_spruce":   "spruce_door",
	"door_birch":    "birch_door",
	"door_jungle":   "jungle_door",
	"door_acacia":   "acacia_door",
	"door_dark_oak": "dark_oak_door",
	"door_iron":     "iron_door",

	// Boats
	"boat": "oak_boat",

	// Buckets
	"bucket_empty": "bucket",
	"bucket_water": "water_bucket",
	"bucket_lava":  "lava_bucket",
	"bucket_milk":  "milk_bucket",

	// Potions
	"potion_bottle_drinkable": "potion",
	"potion_bottle_splash":    "splash_potion",
	"potion_bottle_empty":     "glass_bottle",
	"potion_bottle_lingering": "lingering_potion",

	// Dyes (1.8.9 → 1.14+ names)
	"dye_powder_black":      "ink_sac",
	"dye_powder_red":        "red_dye",
	"dye_powder_green":      "green_dye",
	"dye_powder_brown":      "cocoa_beans",
	"dye_powder_blue":       "lapis_lazuli",
	"dye_powder_purple":     "purple_dye",
	"dye_powder_cyan":       "cyan_dye",
	"dye_powder_silver":     "light_gray_dye",
	"dye_powder_gray":       "gray_dye",
	"dye_powder_pink":       "pink_dye",
	"dye_powder_lime":       "lime_dye",
	"dye_powder_yellow":     "yellow_dye",
	"dye_powder_light_blue": "light_blue_dye",
	"dye_powder_magenta":    "magenta_dye",
	"dye_powder_orange":     "orange_dye",
	"dye_powder_white":      "bone_meal",

	// Minecarts
	"minecart_normal":        "minecart",
	"minecart_chest":         "chest_minecart",
	"minecart_furnace":       "furnace_minecart",
	"minecart_hopper":        "hopper_minecart",
	"minecart_tnt":           "tnt_minecart",
	"minecart_command_block": "command_block_minecart",

	// Music discs
	"record_11":      "music_disc_11",
	"record_13":      "music_disc_13",
	"record_blocks":  "music_disc_blocks",
	"record_cat":     "music_disc_cat",
	"record_chirp":   "music_disc_chirp",
	"record_far":     "music_disc_far",
	"record_mall":    "music_disc_mall",
	"record_mellohi": "music_disc_mellohi",
	"record_stal":    "music_disc_stal",
	"record_strad":   "music_disc_strad",
	"record_wait":    "music_disc_wait",
	"record_ward":    "music_disc_ward",

	// Books
	"book_normal":    "book",
	"book_enchanted": "enchanted_book",
	"book_writable":  "writable_book",
	"book_written":   "written_book",

	// Bows
	"bow_standby": "bow",

	// Fishing rod
	"fishing_rod_uncast": "fishing_rod",
	"fishing_rod_cast":   "fishing_rod_cast",

	// Seeds
	"seeds_wheat":   "wheat_seeds",
	"seeds_melon":   "melon_seeds",
	"seeds_pumpkin": "pumpkin_seeds",

	// Maps
	"map_filled": "filled_map",

	// Misc items
	"fireball":                 "fire_charge",
	"fireworks":                "firework_rocket",
	"fireworks_charge":         "firework_star",
	"fireworks_charge_overlay": "firework_star_overlay",
	"netherbrick":              "nether_brick",
	"redstone_dust":            "redstone",
	"reeds":                    "sugar_cane",
	"slimeball":                "slime_ball",
	"spider_eye_fermented":     "fermented_spider_eye",
	"sign":                     "oak_sign",
	"skull_char":               "player_head",
	"skull_creeper":            "creeper_head",
	"skull_skeleton":           "skeleton_skull",
	"skull_wither":             "wither_skeleton_skull",
	"skull_zombie":             "zombie_head",
	"bed":                      "red_bed",
	"quartz":                   "quartz",
	"map_empty":                "map",
	"speckled_melon":           "glistering_melon_slice",
	"wooden_armorstand":        "armor_stand",
}

// EntityTextureRenames: paths relative to assets/minecraft/textures/entity/
var EntityTextureRenames = map[string]string{
	// Player skins (moved to subdirectories)
	"steve": "player/wide/steve",
	"alex":  "player/slim/alex",

	// Chicken (single file → variant directory)
	"chicken": "chicken/temperate_chicken",

	// Cow
	"cow/cow":       "cow/temperate_cow",
	"cow/mooshroom": "cow/red_mooshroom",

	// Pig
	"pig/pig": "pig/temperate_pig",

	// Iron Golem (flat file → directory)
	"iron_golem": "iron_golem/iron_golem",

	// Squid (flat file → directory)
	"squid": "squid/squid",

	// Sheep
	"sheep/sheep_fur": "sheep/sheep_wool",

	// Zombie Pigman → Zombified Piglin
	"zombie_pigman": "piglin/zombified_piglin",

	// Sign (flat file → directory with wood types)
	"sign": "signs/oak",

	// Arrow (flat file → projectiles directory)
	"arrow": "projectiles/arrow",

	// Boat (flat file → directory)
	"boat": "boat/oak",

	// End Crystal (renamed directory)
	"endercrystal/endercrystal":      "end_crystal/end_crystal",
	"endercrystal/endercrystal_beam": "end_crystal/end_crystal_beam",

	// Zombie villager moved to its own directory
	"zombie/zombie_villager": "zombie_villager/zombie_villager",

	// Villager professions (flat → subdirectory)
	"villager/farmer":    "villager/profession/farmer",
	"villager/librarian": "villager/profession/librarian",
	"villager/butcher":   "villager/profession/butcher",
	"villager/priest":    "villager/profession/cleric",
	"villager/smith":     "villager/profession/weaponsmith",

	// Zombie villager professions
	"zombie/zombie_villager_farmer":    "zombie_villager/profession/farmer",
	"zombie/zombie_villager_librarian": "zombie_villager/profession/librarian",
	"zombie/zombie_villager_priest":    "zombie_villager/profession/cleric",
	"zombie/zombie_villager_smith":     "zombie_villager/profession/weaponsmith",
	"zombie/zombie_villager_butcher":   "zombie_villager/profession/butcher",

	// Bed colors (silver → light_gray)
	"bed/silver": "bed/light_gray",

	// Horse armor moved from entity/horse/armor/ to entity/equipment/horse_body/
	// (handled separately in ArmorTextureRenames)
}

// ArmorLayerRenames: models/armor/ → entity/equipment/ (relative to textures/)
var ArmorLayerRenames = map[string]string{
	// Diamond armor
	"models/armor/diamond_layer_1": "entity/equipment/humanoid/diamond",
	"models/armor/diamond_layer_2": "entity/equipment/humanoid_leggings/diamond",

	// Iron armor
	"models/armor/iron_layer_1": "entity/equipment/humanoid/iron",
	"models/armor/iron_layer_2": "entity/equipment/humanoid_leggings/iron",

	// Gold armor
	"models/armor/gold_layer_1": "entity/equipment/humanoid/gold",
	"models/armor/gold_layer_2": "entity/equipment/humanoid_leggings/gold",

	// Chainmail armor
	"models/armor/chainmail_layer_1": "entity/equipment/humanoid/chainmail",
	"models/armor/chainmail_layer_2": "entity/equipment/humanoid_leggings/chainmail",

	// Leather armor
	"models/armor/leather_layer_1":         "entity/equipment/humanoid/leather",
	"models/armor/leather_layer_1_overlay": "entity/equipment/humanoid/leather_overlay",
	"models/armor/leather_layer_2":         "entity/equipment/humanoid_leggings/leather",
	"models/armor/leather_layer_2_overlay": "entity/equipment/humanoid_leggings/leather_overlay",
}

// HorseArmorRenames: entity/horse/armor/ → entity/equipment/horse_body/
var HorseArmorRenames = map[string]string{
	"entity/horse/armor/horse_armor_diamond": "entity/equipment/horse_body/diamond",
	"entity/horse/armor/horse_armor_gold":    "entity/equipment/horse_body/gold",
	"entity/horse/armor/horse_armor_iron":    "entity/equipment/horse_body/iron",
}

// MiscTextureRenames: other texture paths (relative to textures/)
var MiscTextureRenames = map[string]string{
	"misc/enchanted_item_glint": "misc/enchanted_glint_item",
}

// PaintingAtlasName: old painting atlas filename (without extension).
const PaintingAtlasName = "paintings_kristoffer_zetterstrand"

// PaintingRegion defines a painting's position and size in the atlas grid.
type PaintingRegion struct {
	Name   string // output filename (without extension)
	X, Y   int    // position in atlas (in 16px units)
	W, H   int    // size (in 16px units)
}

// PaintingAtlasRegions: coordinates in 16px grid units within the 256x256 atlas.
var PaintingAtlasRegions = []PaintingRegion{
	// 1x1 paintings (row 0)
	{"kebab", 0, 0, 1, 1},
	{"aztec", 1, 0, 1, 1},
	{"alban", 2, 0, 1, 1},
	{"aztec2", 3, 0, 1, 1},
	{"bomb", 4, 0, 1, 1},
	{"plant", 5, 0, 1, 1},
	{"wasteland", 6, 0, 1, 1},

	// 2x1 paintings (row 2)
	{"pool", 0, 2, 2, 1},
	{"courbet", 2, 2, 2, 1},
	{"sea", 4, 2, 2, 1},
	{"sunset", 6, 2, 2, 1},
	{"creebet", 8, 2, 2, 1},

	// 1x2 paintings (row 4)
	{"wanderer", 0, 4, 1, 2},
	{"graham", 1, 4, 1, 2},

	// 4x2 painting (row 6)
	{"fighters", 0, 6, 4, 2},

	// 2x2 paintings (row 8)
	{"match", 0, 8, 2, 2},
	{"bust", 2, 8, 2, 2},
	{"stage", 4, 8, 2, 2},
	{"void", 6, 8, 2, 2},
	{"skull_and_roses", 8, 8, 2, 2},
	{"wither", 10, 8, 2, 2},

	// 4x4 paintings (row 12)
	{"pointer", 0, 12, 4, 4},
	{"pigscene", 4, 12, 4, 4},
	{"burning_skull", 8, 12, 4, 4},

	// 4x3 paintings
	{"skeleton", 12, 4, 4, 3},
	{"donkey_kong", 12, 7, 4, 3},
}

// WidgetAtlasRegion defines a UI element's position in widgets.png.
type WidgetAtlasRegion struct {
	OutputPath string // relative to textures/gui/sprites/
	X, Y       int    // pixel position in atlas
	W, H       int    // pixel size
}

// WidgetAtlasRegions: pixel coordinates within the 256x256 widgets.png atlas.
var WidgetAtlasRegions = []WidgetAtlasRegion{
	{"hud/hotbar", 0, 0, 182, 22},
	{"hud/hotbar_selection", 0, 22, 24, 24},
	{"widget/button_disabled", 0, 46, 200, 20},
	{"widget/button", 0, 66, 200, 20},
	{"widget/button_highlighted", 0, 86, 200, 20},
}

// IconsAtlasRegion defines a HUD icon's position in icons.png.
type IconsAtlasRegion struct {
	OutputPath string
	X, Y       int
	W, H       int
}

// IconsAtlasRegions: pixel coordinates within the 256x256 icons.png atlas.
var IconsAtlasRegions = []IconsAtlasRegion{
	{"hud/crosshair", 0, 0, 15, 15},
}

// DirectoryRenames: old directory substrings → new (substring matching).
var DirectoryRenames = map[string]string{
	"textures/blocks/": "textures/block/",
	"textures/items/":  "textures/item/",
}

// SoundEventRenames: 1.8.9 sound event names → 1.21 names.
var SoundEventRenames = map[string]string{
	"mob.bat.death":            "entity.bat.death",
	"mob.bat.hurt":             "entity.bat.hurt",
	"mob.bat.idle":             "entity.bat.ambient",
	"mob.bat.loop":             "entity.bat.loop",
	"mob.bat.takeoff":          "entity.bat.takeoff",
	"mob.blaze.breathe":        "entity.blaze.ambient",
	"mob.blaze.death":          "entity.blaze.death",
	"mob.blaze.hit":            "entity.blaze.hurt",
	"mob.cat.hiss":             "entity.cat.hiss",
	"mob.cat.hitt":             "entity.cat.hurt",
	"mob.cat.meow":             "entity.cat.ambient",
	"mob.cat.purr":             "entity.cat.purr",
	"mob.cat.purreow":          "entity.cat.purreow",
	"mob.chicken.hurt":         "entity.chicken.hurt",
	"mob.chicken.plop":         "entity.chicken.egg",
	"mob.chicken.say":          "entity.chicken.ambient",
	"mob.chicken.step":         "entity.chicken.step",
	"mob.cow.hurt":             "entity.cow.hurt",
	"mob.cow.say":              "entity.cow.ambient",
	"mob.cow.step":             "entity.cow.step",
	"mob.creeper.death":        "entity.creeper.death",
	"mob.creeper.say":          "entity.creeper.primed",
	"mob.enderdragon.end":      "entity.ender_dragon.death",
	"mob.enderdragon.growl":    "entity.ender_dragon.growl",
	"mob.enderdragon.hit":      "entity.ender_dragon.hurt",
	"mob.enderdragon.wings":    "entity.ender_dragon.flap",
	"mob.endermen.death":       "entity.enderman.death",
	"mob.endermen.hit":         "entity.enderman.hurt",
	"mob.endermen.idle":        "entity.enderman.ambient",
	"mob.endermen.portal":      "entity.enderman.teleport",
	"mob.endermen.scream":      "entity.enderman.scream",
	"mob.endermen.stare":       "entity.enderman.stare",
	"mob.ghast.affectionate_scream": "entity.ghast.hurt",
	"mob.ghast.charge":         "entity.ghast.warn",
	"mob.ghast.death":          "entity.ghast.death",
	"mob.ghast.fireball":       "entity.ghast.shoot",
	"mob.ghast.moan":           "entity.ghast.ambient",
	"mob.ghast.scream":         "entity.ghast.scream",
	"mob.guardian.attack":      "entity.guardian.attack",
	"mob.guardian.curse":       "entity.elder_guardian.curse",
	"mob.guardian.death":       "entity.guardian.death",
	"mob.guardian.elder.death":  "entity.elder_guardian.death",
	"mob.guardian.elder.hit":    "entity.elder_guardian.hurt",
	"mob.guardian.elder.idle":   "entity.elder_guardian.ambient",
	"mob.guardian.flop":        "entity.guardian.flop",
	"mob.guardian.hit":         "entity.guardian.hurt",
	"mob.guardian.idle":        "entity.guardian.ambient",
	"mob.guardian.land.death":  "entity.guardian.death_land",
	"mob.guardian.land.hit":    "entity.guardian.hurt_land",
	"mob.guardian.land.idle":   "entity.guardian.ambient_land",
	"mob.horse.angry":          "entity.horse.angry",
	"mob.horse.armor":          "entity.horse.armor",
	"mob.horse.breathe":        "entity.horse.breathe",
	"mob.horse.death":          "entity.horse.death",
	"mob.horse.donkey.angry":   "entity.donkey.angry",
	"mob.horse.donkey.death":   "entity.donkey.death",
	"mob.horse.donkey.hit":     "entity.donkey.hurt",
	"mob.horse.donkey.idle":    "entity.donkey.ambient",
	"mob.horse.gallop":         "entity.horse.gallop",
	"mob.horse.hit":            "entity.horse.hurt",
	"mob.horse.idle":           "entity.horse.ambient",
	"mob.horse.jump":           "entity.horse.jump",
	"mob.horse.land":           "entity.horse.land",
	"mob.horse.leather":        "entity.horse.saddle",
	"mob.horse.skeleton.death":  "entity.skeleton_horse.death",
	"mob.horse.skeleton.hit":    "entity.skeleton_horse.hurt",
	"mob.horse.skeleton.idle":   "entity.skeleton_horse.ambient",
	"mob.horse.soft":           "entity.horse.step",
	"mob.horse.wood":           "entity.horse.step_wood",
	"mob.horse.zombie.death":    "entity.zombie_horse.death",
	"mob.horse.zombie.hit":      "entity.zombie_horse.hurt",
	"mob.horse.zombie.idle":     "entity.zombie_horse.ambient",
	"mob.irongolem.death":      "entity.iron_golem.death",
	"mob.irongolem.hit":        "entity.iron_golem.hurt",
	"mob.irongolem.throw":      "entity.iron_golem.attack",
	"mob.irongolem.walk":       "entity.iron_golem.step",
	"mob.magmacube.big":        "entity.magma_cube.squish",
	"mob.magmacube.jump":       "entity.magma_cube.jump",
	"mob.magmacube.small":      "entity.magma_cube.squish_small",
	"mob.pig.death":            "entity.pig.death",
	"mob.pig.say":              "entity.pig.ambient",
	"mob.pig.step":             "entity.pig.step",
	"mob.rabbit.death":         "entity.rabbit.death",
	"mob.rabbit.hop":           "entity.rabbit.jump",
	"mob.rabbit.hurt":          "entity.rabbit.hurt",
	"mob.rabbit.idle":          "entity.rabbit.ambient",
	"mob.sheep.say":            "entity.sheep.ambient",
	"mob.sheep.shear":          "entity.sheep.shear",
	"mob.sheep.step":           "entity.sheep.step",
	"mob.silverfish.hit":       "entity.silverfish.hurt",
	"mob.silverfish.kill":      "entity.silverfish.death",
	"mob.silverfish.say":       "entity.silverfish.ambient",
	"mob.silverfish.step":      "entity.silverfish.step",
	"mob.skeleton.death":       "entity.skeleton.death",
	"mob.skeleton.hurt":        "entity.skeleton.hurt",
	"mob.skeleton.say":         "entity.skeleton.ambient",
	"mob.skeleton.step":        "entity.skeleton.step",
	"mob.slime.attack":         "entity.slime.attack",
	"mob.slime.big":            "entity.slime.squish",
	"mob.slime.small":          "entity.slime.squish_small",
	"mob.snowman.death":        "entity.snow_golem.death",
	"mob.snowman.hurt":         "entity.snow_golem.hurt",
	"mob.spider.death":         "entity.spider.death",
	"mob.spider.say":           "entity.spider.ambient",
	"mob.spider.step":          "entity.spider.step",
	"mob.villager.death":       "entity.villager.death",
	"mob.villager.haggle":      "entity.villager.trade",
	"mob.villager.hit":         "entity.villager.hurt",
	"mob.villager.idle":        "entity.villager.ambient",
	"mob.villager.no":          "entity.villager.no",
	"mob.villager.yes":         "entity.villager.yes",
	"mob.wither.death":         "entity.wither.death",
	"mob.wither.hurt":          "entity.wither.hurt",
	"mob.wither.idle":          "entity.wither.ambient",
	"mob.wither.shoot":         "entity.wither.shoot",
	"mob.wither.spawn":         "entity.wither.spawn",
	"mob.wolf.bark":            "entity.wolf.ambient",
	"mob.wolf.death":           "entity.wolf.death",
	"mob.wolf.growl":           "entity.wolf.growl",
	"mob.wolf.howl":            "entity.wolf.howl",
	"mob.wolf.hurt":            "entity.wolf.hurt",
	"mob.wolf.panting":         "entity.wolf.pant",
	"mob.wolf.shake":           "entity.wolf.shake",
	"mob.wolf.step":            "entity.wolf.step",
	"mob.wolf.whine":           "entity.wolf.whine",
	"mob.zombie.death":         "entity.zombie.death",
	"mob.zombie.hurt":          "entity.zombie.hurt",
	"mob.zombie.infect":        "entity.zombie.infect",
	"mob.zombie.metal":         "entity.zombie.attack_iron_door",
	"mob.zombie.remedy":        "entity.zombie_villager.cure",
	"mob.zombie.say":           "entity.zombie.ambient",
	"mob.zombie.step":          "entity.zombie.step",
	"mob.zombie.unfect":        "entity.zombie_villager.converted",
	"mob.zombie.wood":          "entity.zombie.attack_wooden_door",
	"mob.zombiepig.zpig":       "entity.zombified_piglin.ambient",
	"mob.zombiepig.zpigangry":  "entity.zombified_piglin.angry",
	"mob.zombiepig.zpigdeath":  "entity.zombified_piglin.death",
	"mob.zombiepig.zpighurt":   "entity.zombified_piglin.hurt",
}
