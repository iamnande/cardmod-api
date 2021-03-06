package items

var (
	// item is the collection of item entities.
	items = [168]*item{
		{name: "Potion", purpose: "Restorative"},
		{name: "Potion+", purpose: "Restorative"},
		{name: "Hi-Potion", purpose: "Restorative"},
		{name: "Hi-Potion+", purpose: "Restorative"},
		{name: "X-Potion", purpose: "Restorative"},
		{name: "Mega-Potion", purpose: "Restorative"},
		{name: "Phoenix Down", purpose: "Revival"},
		{name: "Mega Phoenix", purpose: "Revival"},
		{name: "Elixir", purpose: "Forbidden Medicine"},
		{name: "Megalixir", purpose: "Forbidden Medicine"},
		{name: "Antidote", purpose: "Status Recovery"},
		{name: "Soft", purpose: "Status Recovery"},
		{name: "Eye Drops", purpose: "Status Recovery"},
		{name: "Echo Screen", purpose: "Status Recovery"},
		{name: "Holy Water", purpose: "Status Recovery"},
		{name: "Remedy", purpose: "Status Recovery"},
		{name: "Remedy+", purpose: "Status Recovery"},
		{name: "Hero-trial", purpose: "Invincibility"},
		{name: "Hero", purpose: "Invincibility"},
		{name: "Holy War-trial", purpose: "Invincibility"},
		{name: "Holy War", purpose: "Invincibility"},
		{name: "Shell Stone", purpose: "Spell Stones"},
		{name: "Protect Stone", purpose: "Spell Stones"},
		{name: "Aura Stone", purpose: "Spell Stones"},
		{name: "Death Stone", purpose: "Spell Stones"},
		{name: "Holy Stone", purpose: "Spell Stones"},
		{name: "Flare Stone", purpose: "Spell Stones"},
		{name: "Meteor Stone", purpose: "Spell Stones"},
		{name: "Ultima Stone", purpose: "Spell Stones"},
		{name: "Gysahl Greens", purpose: "GF Summon"},
		{name: "Phoenix Pinion", purpose: "GF Summon"},
		{name: "Friendship", purpose: "GF Summon"},
		{name: "Tent", purpose: "Shelters"},
		{name: "Pet House", purpose: "Shelters"},
		{name: "Cottage", purpose: "Shelters"},
		{name: "G-Potion", purpose: "GF Recovery"},
		{name: "G-Hi-Potion", purpose: "GF Recovery"},
		{name: "G-Mega-Potion", purpose: "GF Recovery"},
		{name: "G-Returner", purpose: "GF Recovery"},
		{name: "Rename Card", purpose: "Rename Card"},
		{name: "HP-J Scroll", purpose: "GF Ability"},
		{name: "Str-J Scroll", purpose: "GF Ability"},
		{name: "Vit-J Scroll", purpose: "GF Ability"},
		{name: "Mag-J Scroll", purpose: "GF Ability"},
		{name: "Spr-J Scroll", purpose: "GF Ability"},
		{name: "Luck-J Scroll", purpose: "GF Ability"},
		{name: "Aegis Amulet", purpose: "GF Ability"},
		{name: "Elem Atk", purpose: "GF Ability"},
		{name: "Elem Guard", purpose: "GF Ability"},
		{name: "Status Atk", purpose: "GF Ability"},
		{name: "Status Guard", purpose: "GF Ability"},
		{name: "Rosetta Stone", purpose: "GF Ability"},
		{name: "Magic Scroll", purpose: "Command Ability"},
		{name: "GF Scroll", purpose: "Command Ability"},
		{name: "Draw Scroll", purpose: "Command Ability"},
		{name: "Item Scroll", purpose: "Command Ability"},
		{name: "Gambler Spirit", purpose: "Command Ability"},
		{name: "Healing Ring", purpose: "Command Ability"},
		{name: "Phoenix Spirit", purpose: "Command Ability"},
		{name: "Med Kit", purpose: "Command Ability"},
		{name: "Bomb Spirit", purpose: "Command Ability"},
		{name: "Hungry Cookpot", purpose: "Command Ability"},
		{name: "Meg's Amulet", purpose: "Command Ability"},
		{name: "Stell Pipe", purpose: "GF Enhancement"},
		{name: "Star Fragment", purpose: "GF Enhancement"},
		{name: "Energy Crystal", purpose: "GF Enhancement"},
		{name: "Samantha Soul", purpose: "GF Enhancement"},
		{name: "Healing Mail", purpose: "GF Enhancement"},
		{name: "Silver Sail", purpose: "GF Enhancement"},
		{name: "Gold Armor", purpose: "GF Enhancement"},
		{name: "Diamond Armor", purpose: "GF Enhancement"},
		{name: "Regen Ring", purpose: "Character Ability"},
		{name: "Giant's Ring", purpose: "Character Ability"},
		{name: "Gaea's Ring", purpose: "Character Ability"},
		{name: "Strength Love", purpose: "Character Ability"},
		{name: "Power Wrist", purpose: "Character Ability"},
		{name: "Hyper Wrist", purpose: "Character Ability"},
		{name: "Turtle Shell", purpose: "Character Ability"},
		{name: "Orihalcon", purpose: "Character Ability"},
		{name: "Adamantine", purpose: "Character Ability"},
		{name: "Rune Armlet", purpose: "Character Ability"},
		{name: "Force Armlet", purpose: "Character Ability"},
		{name: "Magic Armlet", purpose: "Character Ability"},
		{name: "Circlet", purpose: "Character Ability"},
		{name: "Hypno Crown", purpose: "Character Ability"},
		{name: "Royal Crown", purpose: "Character Ability"},
		{name: "Jet Engine", purpose: "Character Ability"},
		{name: "Rocket Engine", purpose: "Character Ability"},
		{name: "Moon Curain", purpose: "Character Ability"},
		{name: "Steel Curtain", purpose: "Character Ability"},
		{name: "Glow Curtain", purpose: "Character Ability"},
		{name: "Accelerator", purpose: "Character Ability"},
		{name: "Monk's Code", purpose: "Character Ability"},
		{name: "Knight's Code", purpose: "Character Ability"},
		{name: "Doc's Code", purpose: "Character Ability"},
		{name: "Hundred Needles", purpose: "Character Ability"},
		{name: "Three Stars", purpose: "Character Ability"},
		{name: "Ribbon", purpose: "Character Ability"},
		{name: "Normal Ammo", purpose: "Ammo"},
		{name: "Shotgun Ammo", purpose: "Ammo"},
		{name: "Dark Ammo", purpose: "Ammo"},
		{name: "Fire Ammo", purpose: "Ammo"},
		{name: "Demolition Ammo", purpose: "Ammo"},
		{name: "Fast Ammo", purpose: "Ammo"},
		{name: "AP Ammo", purpose: "Ammo"},
		{name: "Pulse Ammo", purpose: "Ammo"},
		{name: "M-Stone Piece", purpose: "Refinement"},
		{name: "Magic Stone", purpose: "Refinement"},
		{name: "Wizard Stone", purpose: "Refinement"},
		{name: "Ochu Tentacle", purpose: "Refinement"},
		{name: "Healing Water", purpose: "Refinement"},
		{name: "Cockatrice Pinion", purpose: "Refinement"},
		{name: "Zombie Powder", purpose: "Refinement"},
		{name: "Lightweight", purpose: "Refinement"},
		{name: "Sharp Spike", purpose: "Refinement"},
		{name: "Screw", purpose: "Refinement"},
		{name: "Saw Blade", purpose: "Refinement"},
		{name: "Mesmerize Blade", purpose: "Refinement"},
		{name: "Vampire Fang", purpose: "Refinement"},
		{name: "Fury Fragment", purpose: "Refinement"},
		{name: "Betrayal Sword", purpose: "Refinement"},
		{name: "Sleep Powder", purpose: "Refinement"},
		{name: "Life Ring", purpose: "Refinement"},
		{name: "Dragon Fang", purpose: "Refinement"},
		{name: "Spider Web", purpose: "Blue Magic"},
		{name: "Coral Fragment", purpose: "Blue Magic"},
		{name: "Curse Spike", purpose: "Blue Magic"},
		{name: "Black Hole", purpose: "Blue Magic"},
		{name: "Water Crystal", purpose: "Blue Magic"},
		{name: "Missile", purpose: "Blue Magic"},
		{name: "Mystery Fluid", purpose: "Blue Magic"},
		{name: "Running Fire", purpose: "Blue Magic"},
		{name: "Inferno Fang", purpose: "Blue Magic"},
		{name: "Malboro Tentacle", purpose: "Blue Magic"},
		{name: "Whisper", purpose: "Blue Magic"},
		{name: "Laser Cannon", purpose: "Blue Magic"},
		{name: "Barrier", purpose: "Blue Magic"},
		{name: "Power Generator", purpose: "Blue Magic"},
		{name: "Dark Matter", purpose: "Blue Magic"},
		{name: "Bomb Fragment", purpose: "GF Compatibility"},
		{name: "Red Fang", purpose: "GF Compatibility"},
		{name: "Arctic Wind", purpose: "GF Compatibility"},
		{name: "North Wind", purpose: "GF Compatibility"},
		{name: "Dynamo Stone", purpose: "GF Compatibility"},
		{name: "Shear Feather", purpose: "GF Compatibility"},
		{name: "Venom Fang", purpose: "GF Compatibility"},
		{name: "Steel Orb", purpose: "GF Compatibility"},
		{name: "Moon Stone", purpose: "GF Compatibility"},
		{name: "Dino Bone", purpose: "GF Compatibility"},
		{name: "Windmill", purpose: "GF Compatibility"},
		{name: "Dragon Skin", purpose: "GF Compatibility"},
		{name: "Fish Fin", purpose: "GF Compatibility"},
		{name: "Dragon Fin", purpose: "GF Compatibility"},
		{name: "Silence Powder", purpose: "GF Compatibility"},
		{name: "Poison Powder", purpose: "GF Compatibility"},
		{name: "Dead Spirit", purpose: "GF Compatibility"},
		{name: "Chef's Knife", purpose: "GF Compatibility"},
		{name: "Cactus Thorn", purpose: "GF Compatibility"},
		{name: "Shaman Stone", purpose: "GF Compatibility"},
		{name: "Fuel", purpose: "Fuel"},
		{name: "HP Up", purpose: "Stat Boosting"},
		{name: "Str Up", purpose: "Stat Boosting"},
		{name: "Vit Up", purpose: "Stat Boosting"},
		{name: "Mag Up", purpose: "Stat Boosting"},
		{name: "Spr Up", purpose: "Stat Boosting"},
		{name: "Spd Up", purpose: "Stat Boosting"},
		{name: "Luck Up", purpose: "Stat Boosting"},
		{name: "LuvLuv G", purpose: "LuvLuv G"},
	}
)
