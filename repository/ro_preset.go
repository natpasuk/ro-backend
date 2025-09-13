package repository

import (
	"fmt"
	"ro-backend/appError"
	"time"
)

type PresetModel struct {
	Class    int `bson:"class" json:"class"`
	Level    int `bson:"level" json:"level"`
	JobLevel int `bson:"jobLevel" json:"jobLevel"`
	Str      int `bson:"str" json:"str"`
	JobStr   int `bson:"jobStr" json:"jobStr"`
	Agi      int `bson:"agi" json:"agi"`
	JobAgi   int `bson:"jobAgi" json:"jobAgi"`
	Vit      int `bson:"vit" json:"vit"`
	JobVit   int `bson:"jobVit" json:"jobVit"`
	Int      int `bson:"int" json:"int"`
	JobInt   int `bson:"jobInt" json:"jobInt"`
	Dex      int `bson:"dex" json:"dex"`
	JobDex   int `bson:"jobDex" json:"jobDex"`
	Luk      int `bson:"luk" json:"luk"`
	JobLuk   int `bson:"jobLuk" json:"jobLuk"`

	Pow    int `bson:"pow" json:"pow"`
	JobPow int `bson:"jobPow" json:"jobPow"`
	Sta    int `bson:"sta" json:"sta"`
	JobSta int `bson:"jobSta" json:"jobSta"`
	Wis    int `bson:"wis" json:"wis"`
	JobWis int `bson:"jobWis" json:"jobWis"`
	Spl    int `bson:"spl" json:"spl"`
	JobSpl int `bson:"jobSpl" json:"jobSpl"`
	Con    int `bson:"con" json:"con"`
	JobCon int `bson:"jobCon" json:"jobCon"`
	Crt    int `bson:"crt" json:"crt"`
	JobCrt int `bson:"jobCrt" json:"jobCrt"`

	SelectedAtkSkill   string        `bson:"selectedAtkSkill" json:"selectedAtkSkill"`
	RawOptionTxts      []interface{} `bson:"rawOptionTxts" json:"rawOptionTxts"`
	PropertyAtk        string        `bson:"propertyAtk" json:"propertyAtk,omitempty"`
	Ammo               int           `bson:"ammo" json:"ammo,omitempty"`
	Weapon             int           `bson:"weapon" json:"weapon,omitempty"`
	WeaponRefine       int           `bson:"weaponRefine" json:"weaponRefine"`
	WeaponGrade        string        `bson:"weaponGrade" json:"weaponGrade"`
	WeaponCard1        int           `bson:"weaponCard1" json:"weaponCard1,omitempty"`
	WeaponCard2        int           `bson:"weaponCard2" json:"weaponCard2,omitempty"`
	WeaponCard3        int           `bson:"weaponCard3" json:"weaponCard3,omitempty"`
	WeaponCard4        int           `bson:"weaponCard4" json:"weaponCard4,omitempty"`
	WeaponEnchant0     int           `bson:"weaponEnchant0" json:"weaponEnchant0,omitempty"`
	WeaponEnchant1     int           `bson:"weaponEnchant1" json:"weaponEnchant1,omitempty"`
	WeaponEnchant2     int           `bson:"weaponEnchant2" json:"weaponEnchant2,omitempty"`
	WeaponEnchant3     int           `bson:"weaponEnchant3" json:"weaponEnchant3,omitempty"`
	LeftWeapon         int           `bson:"leftWeapon" json:"leftWeapon,omitempty"`
	LeftWeaponRefine   int           `bson:"leftWeaponRefine" json:"leftWeaponRefine"`
	LeftWeaponGrade    string        `bson:"leftWeaponGrade" json:"leftWeaponGrade"`
	LeftWeaponCard1    int           `bson:"leftWeaponCard1" json:"leftWeaponCard1,omitempty"`
	LeftWeaponCard2    int           `bson:"leftWeaponCard2" json:"leftWeaponCard2,omitempty"`
	LeftWeaponCard3    int           `bson:"leftWeaponCard3" json:"leftWeaponCard3,omitempty"`
	LeftWeaponCard4    int           `bson:"leftWeaponCard4" json:"leftWeaponCard4,omitempty"`
	LeftWeaponEnchant0 int           `bson:"leftWeaponEnchant0" json:"leftWeaponEnchant0,omitempty"`
	LeftWeaponEnchant1 int           `bson:"leftWeaponEnchant1" json:"leftWeaponEnchant1,omitempty"`
	LeftWeaponEnchant2 int           `bson:"leftWeaponEnchant2" json:"leftWeaponEnchant2,omitempty"`
	LeftWeaponEnchant3 int           `bson:"leftWeaponEnchant3" json:"leftWeaponEnchant3,omitempty"`
	Shield             int           `bson:"shield" json:"shield,omitempty"`
	ShieldRefine       int           `bson:"shieldRefine" json:"shieldRefine"`
	ShieldGrade        string        `bson:"shieldGrade" json:"shieldGrade"`
	ShieldCard         int           `bson:"shieldCard" json:"shieldCard,omitempty"`
	ShieldEnchant1     int           `bson:"shieldEnchant1" json:"shieldEnchant1,omitempty"`
	ShieldEnchant2     int           `bson:"shieldEnchant2" json:"shieldEnchant2,omitempty"`
	ShieldEnchant3     int           `bson:"shieldEnchant3" json:"shieldEnchant3,omitempty"`
	HeadUpper          int           `bson:"headUpper" json:"headUpper,omitempty"`
	HeadUpperRefine    int           `bson:"headUpperRefine" json:"headUpperRefine"`
	HeadUpperGrade     string        `bson:"headUpperGrade" json:"headUpperGrade"`
	HeadUpperCard      int           `bson:"headUpperCard" json:"headUpperCard,omitempty"`
	HeadUpperEnchant1  int           `bson:"headUpperEnchant1" json:"headUpperEnchant1,omitempty"`
	HeadUpperEnchant2  int           `bson:"headUpperEnchant2" json:"headUpperEnchant2,omitempty"`
	HeadUpperEnchant3  int           `bson:"headUpperEnchant3" json:"headUpperEnchant3,omitempty"`
	HeadMiddle         int           `bson:"headMiddle" json:"headMiddle,omitempty"`
	HeadMiddleGrade    string        `bson:"headMiddleGrade" json:"headMiddleGrade,omitempty"`
	HeadMiddleCard     int           `bson:"headMiddleCard" json:"headMiddleCard,omitempty"`
	HeadMiddleEnchant1 int           `bson:"headMiddleEnchant1" json:"headMiddleEnchant1,omitempty"`
	HeadMiddleEnchant2 int           `bson:"headMiddleEnchant2" json:"headMiddleEnchant2,omitempty"`
	HeadMiddleEnchant3 int           `bson:"headMiddleEnchant3" json:"headMiddleEnchant3,omitempty"`
	HeadLower          int           `bson:"headLower" json:"headLower,omitempty"`
	HeadLowerGrade     string        `bson:"headLowerGrade" json:"headLowerGrade,omitempty"`
	HeadLowerEnchant1  int           `bson:"headLowerEnchant1" json:"headLowerEnchant1,omitempty"`
	HeadLowerEnchant2  int           `bson:"headLowerEnchant2" json:"headLowerEnchant2,omitempty"`
	HeadLowerEnchant3  int           `bson:"headLowerEnchant3" json:"headLowerEnchant3,omitempty"`
	Armor              int           `bson:"armor" json:"armor,omitempty"`
	ArmorRefine        int           `bson:"armorRefine" json:"armorRefine"`
	ArmorGrade         string        `bson:"armorGrade" json:"armorGrade"`
	ArmorCard          int           `bson:"armorCard" json:"armorCard,omitempty"`
	ArmorEnchant1      int           `bson:"armorEnchant1" json:"armorEnchant1,omitempty"`
	ArmorEnchant2      int           `bson:"armorEnchant2" json:"armorEnchant2,omitempty"`
	ArmorEnchant3      int           `bson:"armorEnchant3" json:"armorEnchant3,omitempty"`
	Garment            int           `bson:"garment" json:"garment,omitempty"`
	GarmentRefine      int           `bson:"garmentRefine" json:"garmentRefine"`
	GarmentGrade       string        `bson:"garmentGrade" json:"garmentGrade"`
	GarmentCard        int           `bson:"garmentCard" json:"garmentCard,omitempty"`
	GarmentEnchant1    int           `bson:"garmentEnchant1" json:"garmentEnchant1,omitempty"`
	GarmentEnchant2    int           `bson:"garmentEnchant2" json:"garmentEnchant2,omitempty"`
	GarmentEnchant3    int           `bson:"garmentEnchant3" json:"garmentEnchant3,omitempty"`
	Boot               int           `bson:"boot" json:"boot,omitempty"`
	BootRefine         int           `bson:"bootRefine" json:"bootRefine"`
	BootGrade          string        `bson:"bootGrade" json:"bootGrade"`
	BootCard           int           `bson:"bootCard" json:"bootCard,omitempty"`
	BootEnchant1       int           `bson:"bootEnchant1" json:"bootEnchant1,omitempty"`
	BootEnchant2       int           `bson:"bootEnchant2" json:"bootEnchant2,omitempty"`
	BootEnchant3       int           `bson:"bootEnchant3" json:"bootEnchant3,omitempty"`
	AccLeft            int           `bson:"accLeft" json:"accLeft,omitempty"`
	AccLeftRefine      int           `bson:"accLeftRefine" json:"accLeftRefine,omitempty"`
	AccLeftGrade       string        `bson:"accLeftGrade" json:"accLeftGrade,omitempty"`
	AccLeftCard        int           `bson:"accLeftCard" json:"accLeftCard,omitempty"`
	AccLeftEnchant1    int           `bson:"accLeftEnchant1" json:"accLeftEnchant1,omitempty"`
	AccLeftEnchant2    int           `bson:"accLeftEnchant2" json:"accLeftEnchant2,omitempty"`
	AccLeftEnchant3    int           `bson:"accLeftEnchant3" json:"accLeftEnchant3,omitempty"`
	AccRight           int           `bson:"accRight" json:"accRight,omitempty"`
	AccRightRefine     int           `bson:"accRightRefine" json:"accRightRefine,omitempty"`
	AccRightGrade      string        `bson:"accRightGrade" json:"accRightGrade,omitempty"`
	AccRightCard       int           `bson:"accRightCard" json:"accRightCard,omitempty"`
	AccRightEnchant1   int           `bson:"accRightEnchant1" json:"accRightEnchant1,omitempty"`
	AccRightEnchant2   int           `bson:"accRightEnchant2" json:"accRightEnchant2,omitempty"`
	AccRightEnchant3   int           `bson:"accRightEnchant3" json:"accRightEnchant3,omitempty"`

	Pet int `bson:"pet" json:"pet,omitempty"`

	CostumeUpper   int `bson:"costumeUpper" json:"costumeUpper,omitempty"`
	CostumeMiddle  int `bson:"costumeMiddle" json:"costumeMiddle,omitempty"`
	CostumeLower   int `bson:"costumeLower" json:"costumeLower,omitempty"`
	CostumeGarment int `bson:"costumeGarment" json:"costumeGarment,omitempty"`

	CostumeEnchantUpper    int `bson:"costumeEnchantUpper" json:"costumeEnchantUpper,omitempty"`
	CostumeEnchantMiddle   int `bson:"costumeEnchantMiddle" json:"costumeEnchantMiddle,omitempty"`
	CostumeEnchantLower    int `bson:"costumeEnchantLower" json:"costumeEnchantLower,omitempty"`
	CostumeEnchantGarment  int `bson:"costumeEnchantGarment" json:"costumeEnchantGarment,omitempty"`
	CostumeEnchantGarment2 int `bson:"costumeEnchantGarment2" json:"costumeEnchantGarment2,omitempty"`
	CostumeEnchantGarment4 int `bson:"costumeEnchantGarment4" json:"costumeEnchantGarment4,omitempty"`

    ShadowWeapon          int `bson:"shadowWeapon" json:"shadowWeapon,omitempty"`
    ShadowWeaponRefine    int `bson:"shadowWeaponRefine" json:"shadowWeaponRefine"`
    ShadowWeaponEnchant2  int `bson:"shadowWeaponEnchant1" json:"shadowWeaponEnchant1,omitempty"`
    ShadowWeaponEnchant2  int `bson:"shadowWeaponEnchant2" json:"shadowWeaponEnchant2,omitempty"`
    ShadowWeaponEnchant3  int `bson:"shadowWeaponEnchant3" json:"shadowWeaponEnchant3,omitempty"`
    ShadowArmor           int `bson:"shadowArmor" json:"shadowArmor,omitempty"`
    ShadowArmorRefine     int `bson:"shadowArmorRefine" json:"shadowArmorRefine"`
    ShadowArmorEnchant2   int `bson:"shadowArmorEnchant1" json:"shadowArmorEnchant1,omitempty"`
    ShadowArmorEnchant2   int `bson:"shadowArmorEnchant2" json:"shadowArmorEnchant2,omitempty"`
    ShadowArmorEnchant3   int `bson:"shadowArmorEnchant3" json:"shadowArmorEnchant3,omitempty"`
    ShadowShield          int `bson:"shadowShield" json:"shadowShield,omitempty"`
    ShadowShieldRefine    int `bson:"shadowShieldRefine" json:"shadowShieldRefine"`
    ShadowShieldEnchant2  int `bson:"shadowShieldEnchant1" json:"shadowShieldEnchant1,omitempty"`
    ShadowShieldEnchant2  int `bson:"shadowShieldEnchant2" json:"shadowShieldEnchant2,omitempty"`
    ShadowShieldEnchant3  int `bson:"shadowShieldEnchant3" json:"shadowShieldEnchant3,omitempty"`
    ShadowBoot            int `bson:"shadowBoot" json:"shadowBoot,omitempty"`
    ShadowBootRefine      int `bson:"shadowBootRefine" json:"shadowBootRefine"`
    ShadowBootEnchant2    int `bson:"shadowBootEnchant1" json:"shadowBootEnchant1,omitempty"`
    ShadowBootEnchant2    int `bson:"shadowBootEnchant2" json:"shadowBootEnchant2,omitempty"`
    ShadowBootEnchant3    int `bson:"shadowBootEnchant3" json:"shadowBootEnchant3,omitempty"`
    ShadowEarring         int `bson:"shadowEarring" json:"shadowEarring,omitempty"`
    ShadowEarringRefine   int `bson:"shadowEarringRefine" json:"shadowEarringRefine"`
    ShadowEarringEnchant2 int `bson:"shadowEarringEnchant1" json:"shadowEarringEnchant1,omitempty"`
    ShadowEarringEnchant2 int `bson:"shadowEarringEnchant2" json:"shadowEarringEnchant2,omitempty"`
    ShadowEarringEnchant3 int `bson:"shadowEarringEnchant3" json:"shadowEarringEnchant3,omitempty"`
    ShadowPendant         int `bson:"shadowPendant" json:"shadowPendant,omitempty"`
    ShadowPendantRefine   int `bson:"shadowPendantRefine" json:"shadowPendantRefine"`
    ShadowPendantEnchant2 int `bson:"shadowPendantEnchant1" json:"shadowPendantEnchant1,omitempty"`
    ShadowPendantEnchant2 int `bson:"shadowPendantEnchant2" json:"shadowPendantEnchant2,omitempty"`
    ShadowPendantEnchant3 int `bson:"shadowPendantEnchant3" json:"shadowPendantEnchant3,omitempty"`

	SkillBuffMap    map[string]int `bson:"skillBuffMap" json:"skillBuffMap"`
	SkillBuffs      []int          `bson:"skillBuffs" json:"skillBuffs"`
	ActiveSkillMap  map[string]int `bson:"activeSkillMap" json:"activeSkillMap"`
	ActiveSkills    []int          `bson:"activeSkills" json:"activeSkills"`
	PassiveSkillMap map[string]int `bson:"passiveSkillMap" json:"passiveSkillMap"`
	PassiveSkills   []int          `bson:"passiveSkills" json:"passiveSkills"`
	Consumables     []int          `bson:"consumables" json:"consumables"`
	Consumables2    []int          `bson:"consumables2" json:"consumables2"`
	AspdPotion      int            `bson:"aspdPotion" json:"aspdPotion,omitempty"`
	AspdPotions     []int          `bson:"aspdPotions" json:"aspdPotions"`
}

type RoPreset struct {
	Id          string      `bson:"id" json:"id"`
	UserId      string      `bson:"user_id" json:"userId"`
	UserName    string      `bson:"user_name" json:"userName"`
	Label       string      `bson:"label" json:"label"`
	Model       PresetModel `bson:"model" json:"model"`
	ClassId     int         `bson:"class_id" json:"classId"`
	CreatedAt   time.Time   `bson:"created_at" json:"createdAt"`
	UpdatedAt   time.Time   `bson:"updated_at" json:"updatedAt"`
	PublishName string      `bson:"publish_name" json:"publishName"`
	IsPublished bool        `bson:"is_published" json:"isPublished"`
	PublishedAt time.Time   `bson:"published_at" json:"publishedAt"`
}

func (i *PresetModel) Validate() error {
	moreThenZero := []int{i.Class, i.Level, i.JobLevel, i.Str, i.Agi}
	for _, v := range moreThenZero {
		if !(v > 0) {
			return fmt.Errorf(appError.ErrInvalidPresetInput)
		}
	}

	return nil
}

type CreatePresetInput struct {
	UserId   string      `bson:"user_id" json:"userId"`
	UserName string      `bson:"user_name" json:"userName"`
	Label    string      `bson:"label" json:"label"`
	Model    PresetModel `bson:"model" json:"model"`
}

func (i *CreatePresetInput) Validate() error {
	if i.Label == "" {
		return fmt.Errorf(appError.ErrInvalidPresetInput)
	}

	if err := i.Model.Validate(); err != nil {
		return fmt.Errorf(appError.ErrInvalidPresetInput)
	}

	return nil
}

type UpdatePresetInput struct {
	ClassId     int          `bson:"class_id,omitempty" json:"classId"`
	UserId      string       `bson:"user_id,omitempty" json:"userId"`
	UserName    string       `bson:"user_name,omitempty" json:"userName"`
	Label       string       `bson:"label,omitempty" json:"label"`
	Model       *PresetModel `bson:"model,omitempty" json:"model"`
	UpdatedAt   time.Time    `bson:"updated_at,omitempty" json:"updatedAt"`
	PublishName string       `bson:"publish_name,omitempty" json:"publishName"`
	IsPublished bool         `bson:"is_published,omitempty" json:"isPublished"`
	PublishedAt time.Time    `bson:"published_at,omitempty" json:"publishedAt"`
}

type UnPublishPresetInput struct {
	IsPublished bool      `bson:"is_published"`
	PublishedAt time.Time `bson:"published_at"`
}

type UpdateTagsInput struct {
	Id        string    `bson:"id,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

type BulkCreatePresetInput struct {
	UserId   string `bson:"user_id" json:"userId"`
	UserName string `bson:"user_name" json:"userName"`
	BulkData []struct {
		Label string      `bson:"label" json:"label"`
		Model PresetModel `bson:"model" json:"model"`
	} `json:"bulkData"`
}

type PartialSearchRoPresetInput struct {
	Id           *string `bson:"id,omitempty"`
	UserId       *string `bson:"user_id,omitempty"`
	ClassId      *int    `bson:"class_id,omitempty"`
	Label        *string `bson:"label,omitempty"`
	Skip         *int
	Take         *int
	InCludeModel bool
}

type PartialSearchRoPresetForUpdateInput struct {
	Id      string `bson:"id,omitempty"`
	UserId  string `bson:"user_id,omitempty"`
	ClassId int    `bson:"class_id,omitempty"`
	Label   string `bson:"label,omitempty"`
}

type IdSearchInput struct {
	Id string `bson:"id"`
}

type PartialSearchRoPresetResult struct {
	Items []RoPreset
	Total int64
}

type FindPresetByIdInput struct {
	Id           string
	InCludeModel bool
}

type PresetListSorting struct {
	UpdatedAt int `bson:"updated_at"`
}

type RoPresetRepository interface {
	FindPresetById(FindPresetByIdInput) (*RoPreset, error)
	FindPresetByIds([]string) ([]RoPreset, error)
	PartialSearchPresets(PartialSearchRoPresetInput) (*PartialSearchRoPresetResult, error)
	CreatePreset(CreatePresetInput) (*RoPreset, error)
	CreatePresets(BulkCreatePresetInput) ([]RoPreset, error)
	UpdatePreset(id string, i UpdatePresetInput) error
	UpdateUserName(userId, userName string) error
	UnpublishedPreset(id string) error
	DeletePresetById(string) (*int, error)
}
