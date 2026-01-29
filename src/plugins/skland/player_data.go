package skland

type PlayerDetail struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Data      struct {
		Detail struct {
			Base struct {
				ServerName    string `json:"serverName"`
				RoleID        string `json:"roleId"`
				Name          string `json:"name"`
				CreateTime    string `json:"createTime"`
				SaveTime      string `json:"saveTime"`
				LastLoginTime string `json:"lastLoginTime"`
				Exp           int    `json:"exp"`
				Level         int    `json:"level"`
				WorldLevel    int    `json:"worldLevel"`
				Gender        int    `json:"gender"`
				AvatarURL     string `json:"avatarUrl"`
				MainMission   struct {
					ID          string `json:"id"`
					Description string `json:"description"`
				} `json:"mainMission"`
				CharNum   int `json:"charNum"`
				WeaponNum int `json:"weaponNum"`
				DocNum    int `json:"docNum"`
			} `json:"base"`
			Chars []struct {
				CharData struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb             string `json:"atb"`
									AtkScale        string `json:"atk_scale"`
									DisplayAtkScale string `json:"display_atk_scale"`
									Poise           string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					LabelType       string   `json:"labelType"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				ID         string `json:"id"`
				Level      int    `json:"level"`
				UserSkills struct {
					One66D38Ccfe9F7486047E9721E7D6Adde struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"166d38ccfe9f7486047e9721e7d6adde"`
					Three299D991A0F8417C245Adc2997263Ae6 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"3299d991a0f8417c245adc2997263ae6"`
					B2C072B0B15F8157F0B11Da1C955De74 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"b2c072b0b15f8157f0b11da1c955de74"`
					Ff40Dad4B22B95C4A2A765Bd6B4F2D8E struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"ff40dad4b22b95c4a2a765bd6b4f2d8e"`
				} `json:"userSkills,omitempty"`
				BodyEquip struct {
					EquipID   string `json:"equipId"`
					EquipData struct {
						ID      string `json:"id"`
						Name    string `json:"name"`
						IconURL string `json:"iconUrl"`
						Rarity  struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"rarity"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Level struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"level"`
						Properties  []string `json:"properties"`
						IsAccessory bool     `json:"isAccessory"`
						Suit        struct {
							ID              string `json:"id"`
							Name            string `json:"name"`
							SkillID         string `json:"skillId"`
							SkillDesc       string `json:"skillDesc"`
							SkillDescParams struct {
								Duration   string `json:"duration"`
								SkillDmgUp string `json:"skill_dmg_up"`
								SpellDmgUp string `json:"spell_dmg_up"`
								StackCond  string `json:"stack_cond"`
							} `json:"skillDescParams"`
						} `json:"suit"`
						Function string `json:"function"`
						Pkg      string `json:"pkg"`
					} `json:"equipData"`
				} `json:"bodyEquip,omitempty"`
				ArmEquip struct {
					EquipID   string `json:"equipId"`
					EquipData struct {
						ID      string `json:"id"`
						Name    string `json:"name"`
						IconURL string `json:"iconUrl"`
						Rarity  struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"rarity"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Level struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"level"`
						Properties  []string `json:"properties"`
						IsAccessory bool     `json:"isAccessory"`
						Suit        struct {
							ID              string `json:"id"`
							Name            string `json:"name"`
							SkillID         string `json:"skillId"`
							SkillDesc       string `json:"skillDesc"`
							SkillDescParams struct {
								Duration    string `json:"duration"`
								Duration2   string `json:"duration2"`
								FireDmgUp   string `json:"fire_dmg_up"`
								NatureDmgUp string `json:"nature_dmg_up"`
								PhySpellUp  string `json:"phy_spell_up"`
							} `json:"skillDescParams"`
						} `json:"suit"`
						Function string `json:"function"`
						Pkg      string `json:"pkg"`
					} `json:"equipData"`
				} `json:"armEquip,omitempty"`
				FirstAccessory struct {
					EquipID   string `json:"equipId"`
					EquipData struct {
						ID      string `json:"id"`
						Name    string `json:"name"`
						IconURL string `json:"iconUrl"`
						Rarity  struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"rarity"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Level struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"level"`
						Properties  []string `json:"properties"`
						IsAccessory bool     `json:"isAccessory"`
						Suit        struct {
							ID              string `json:"id"`
							Name            string `json:"name"`
							SkillID         string `json:"skillId"`
							SkillDesc       string `json:"skillDesc"`
							SkillDescParams struct {
								Duration    string `json:"duration"`
								Duration2   string `json:"duration2"`
								FireDmgUp   string `json:"fire_dmg_up"`
								NatureDmgUp string `json:"nature_dmg_up"`
								PhySpellUp  string `json:"phy_spell_up"`
							} `json:"skillDescParams"`
						} `json:"suit"`
						Function string `json:"function"`
						Pkg      string `json:"pkg"`
					} `json:"equipData"`
				} `json:"firstAccessory,omitempty"`
				SecondAccessory struct {
					EquipID   string `json:"equipId"`
					EquipData struct {
						ID      string `json:"id"`
						Name    string `json:"name"`
						IconURL string `json:"iconUrl"`
						Rarity  struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"rarity"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Level struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"level"`
						Properties  []string `json:"properties"`
						IsAccessory bool     `json:"isAccessory"`
						Suit        struct {
							ID              string `json:"id"`
							Name            string `json:"name"`
							SkillID         string `json:"skillId"`
							SkillDesc       string `json:"skillDesc"`
							SkillDescParams struct {
								Duration    string `json:"duration"`
								Duration2   string `json:"duration2"`
								FireDmgUp   string `json:"fire_dmg_up"`
								NatureDmgUp string `json:"nature_dmg_up"`
								PhySpellUp  string `json:"phy_spell_up"`
							} `json:"skillDescParams"`
						} `json:"suit"`
						Function string `json:"function"`
						Pkg      string `json:"pkg"`
					} `json:"equipData"`
				} `json:"secondAccessory,omitempty"`
				TacticalItem struct {
					TacticalItemID   string `json:"tacticalItemId"`
					TacticalItemData struct {
						ID      string `json:"id"`
						Name    string `json:"name"`
						IconURL string `json:"iconUrl"`
						Rarity  struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"rarity"`
						ActiveEffectType struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"activeEffectType"`
						ActiveEffect       string `json:"activeEffect"`
						PassiveEffect      string `json:"passiveEffect"`
						ActiveEffectParams struct {
							Value  string `json:"value"`
							Value2 string `json:"value2"`
						} `json:"activeEffectParams"`
						PassiveEffectParams struct {
							Count  string `json:"count"`
							Param1 string `json:"param1"`
							Param2 string `json:"param2"`
						} `json:"passiveEffectParams"`
					} `json:"tacticalItemData"`
				} `json:"tacticalItem,omitempty"`
				EvolvePhase    int `json:"evolvePhase"`
				PotentialLevel int `json:"potentialLevel"`
				Weapon         struct {
					WeaponData struct {
						ID      string `json:"id"`
						Name    string `json:"name"`
						IconURL string `json:"iconUrl"`
						Rarity  struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"rarity"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Function    string `json:"function"`
						Description string `json:"description"`
						Skills      []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"skills"`
					} `json:"weaponData"`
					Level             int `json:"level"`
					RefineLevel       int `json:"refineLevel"`
					BreakthroughLevel int `json:"breakthroughLevel"`
					Gem               struct {
						ID   string `json:"id"`
						Icon string `json:"icon"`
					} `json:"gem"`
				} `json:"weapon"`
				Gender    string `json:"gender"`
				OwnTs     string `json:"ownTs"`
				CharData0 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills0 struct {
					ZeroC89F0A86C28E62038B8C7Fa6Ce47F2F struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"0c89f0a86c28e62038b8c7fa6ce47f2f"`
					Four9Ab9859462D2F66352F0B3C4D859Fb9 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"49ab9859462d2f66352f0b3c4d859fb9"`
					Six86Ef52Fdb1E963E0De4D87F2Ad9D4De struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"686ef52fdb1e963e0de4d87f2ad9d4de"`
					Seven89Ac67442F3F9Ecbf77Dec139C832D1 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"789ac67442f3f9ecbf77dec139c832d1"`
				} `json:"userSkills,omitempty"`
				CharData1 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills1 struct {
					Three802B082179103Ad7Fcbd95415A1Df38 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"3802b082179103ad7fcbd95415a1df38"`
					FourC72Af23F69262Eb45Eb5E138E974328 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"4c72af23f69262eb45eb5e138e974328"`
					De3865F07F8B6Fc870C13272Bcd96Aba struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"de3865f07f8b6fc870c13272bcd96aba"`
					F8Fb961B68A6E9D2A07Fa6Ac56Cca8A3 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"f8fb961b68a6e9d2a07fa6ac56cca8a3"`
				} `json:"userSkills,omitempty"`
				CharData2 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb             string `json:"atb"`
									AtkScale        string `json:"atk_scale"`
									DisplayAtkScale string `json:"display_atk_scale"`
									Poise           string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills2 struct {
					Zero8A5381Dcb8C2Ba1A99F7A2B3Bde4062 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"08a5381dcb8c2ba1a99f7a2b3bde4062"`
					Two407197665E3F04Bc97202F14A699A3E struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"2407197665e3f04bc97202f14a699a3e"`
					Seven05337Bf3Cd331A535F8D948D6B798E9 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"705337bf3cd331a535f8d948d6b798e9"`
					E9D1A6D458C93D78Ccd6C3D31Cd41Ca1 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"e9d1a6d458c93d78ccd6c3d31cd41ca1"`
				} `json:"userSkills,omitempty"`
				CharData3 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills3 struct {
					One98C3E43C56836698B70Bebbc3C69511 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"198c3e43c56836698b70bebbc3c69511"`
					Two548D91Cd37Ee2A6Dc90A57E30321E59 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"2548d91cd37ee2a6dc90a57e30321e59"`
					Six076E76A11350E18E746Bb8B04A5A2A4 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"6076e76a11350e18e746bb8b04a5a2a4"`
					Nine76Fe54C0Eb2937015Ee117140763Bfe struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"976fe54c0eb2937015ee117140763bfe"`
				} `json:"userSkills,omitempty"`
				CharData4 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills4 struct {
					One58D9D2E94684269595Bcd9Afcb9E11E struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"158d9d2e94684269595bcd9afcb9e11e"`
					Two2A08Aeb0A12F95Ca3618796Fffd9809 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"22a08aeb0a12f95ca3618796fffd9809"`
					FourD0A265144641A2305E13Dfab4E58804 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"4d0a265144641a2305e13dfab4e58804"`
					B5457621458Eaaebd41D7Dc1E927E1B7 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"b5457621458eaaebd41d7dc1e927e1b7"`
				} `json:"userSkills,omitempty"`
				CharData5 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb             string `json:"atb"`
									AtkScale        string `json:"atk_scale"`
									DisplayAtkScale string `json:"display_atk_scale"`
									Poise           string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills5 struct {
					Zero5Eae3C5F15E64Ea065Ca3B8A8B5670A struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"05eae3c5f15e64ea065ca3b8a8b5670a"`
					ThreeA3A6C9Dac031Eaea4D730Bb56B7A5Bd struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"3a3a6c9dac031eaea4d730bb56b7a5bd"`
					Five9C33146250667Cd6F515Cb7C269Dd2A struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"59c33146250667cd6f515cb7c269dd2a"`
					NineB989356364917Cab939C6Deba48F67C struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"9b989356364917cab939c6deba48f67c"`
				} `json:"userSkills,omitempty"`
				CharData6 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills6 struct {
					Two86E7001B2057E650F64F0E391223954 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"286e7001b2057e650f64f0e391223954"`
					FiveBcb587D98299Ae6Eac0F261273629E5 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"5bcb587d98299ae6eac0f261273629e5"`
					B616572F709440D13664Fcc2D28A70Eb struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"b616572f709440d13664fcc2d28a70eb"`
					F49D5D010Be4480E353521654E215A11 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"f49d5d010be4480e353521654e215a11"`
				} `json:"userSkills,omitempty"`
				CharData7 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills7 struct {
					Four1583C2833E1375A854B010471D7Ded3 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"41583c2833e1375a854b010471d7ded3"`
					Four18E4D8A0893F9B78F80F09A1830Afdf struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"418e4d8a0893f9b78f80f09a1830afdf"`
					D6A30D0C43C7499C35Ce1941B131780A struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"d6a30d0c43c7499c35ce1941b131780a"`
					Fb85E0419523563D546Cf91Fc62Dc45B struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"fb85e0419523563d546cf91fc62dc45b"`
				} `json:"userSkills,omitempty"`
				CharData8 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills8 struct {
					TwoE3B9Eb2C40E84D627C3B0295Ab9Cad9 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"2e3b9eb2c40e84d627c3b0295ab9cad9"`
					FourD7941B34Dbbfa8493168A95B3Ef9293 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"4d7941b34dbbfa8493168a95b3ef9293"`
					Nine7B929E5Cccfa3Fd8C7D51Fa2A383120 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"97b929e5cccfa3fd8c7d51fa2a383120"`
					E11F2F0E8Cd68D67Fce73F3A6Afa2712 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"e11f2f0e8cd68d67fce73f3a6afa2712"`
				} `json:"userSkills,omitempty"`
				CharData9 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills9 struct {
					A8464594Aeaaecc68C62842B5B73F180 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"a8464594aeaaecc68c62842b5b73f180"`
					B26F3Cc291540Cc7510E16D1A3C16D6B struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"b26f3cc291540cc7510e16d1a3c16d6b"`
					C170410E4D2F19B1Bd1Bd75C7Ebcf8C7 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"c170410e4d2f19b1bd1bd75c7ebcf8c7"`
					Ed9D75F5De77E5F5648Eca0184Cdb94C struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"ed9d75f5de77e5f5648eca0184cdb94c"`
				} `json:"userSkills,omitempty"`
				CharData10 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills10 struct {
					ZeroB911Fbc68D12Fed35866D8C2938171F struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"0b911fbc68d12fed35866d8c2938171f"`
					Nine89A9Ded833366184B85Fb34F9D647F4 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"989a9ded833366184b85fb34f9d647f4"`
					A141Dba473F57Dcbc049F6219226E862 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"a141dba473f57dcbc049f6219226e862"`
					F63Df3D48F7Be190Da5C2D4F0Ad97114 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"f63df3d48f7be190da5c2d4f0ad97114"`
				} `json:"userSkills,omitempty"`
				CharData11 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills11 struct {
					Eight52Bfa7E80F6E049Ff9349Be75F2929B struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"852bfa7e80f6e049ff9349be75f2929b"`
					Nine223851E5B616C8A1F4407C02Cdecd22 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"9223851e5b616c8a1f4407c02cdecd22"`
					Nine3Abca3046D7054C80F3686Ac9C3C640 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"93abca3046d7054c80f3686ac9c3c640"`
					A32C3F07Cc26De185049024438Ff565B struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"a32c3f07cc26de185049024438ff565b"`
				} `json:"userSkills,omitempty"`
				CharData12 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb             string `json:"atb"`
									AtkScale        string `json:"atk_scale"`
									AtkScaleDisplay string `json:"atk_scale_display"`
									Poise           string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills12 struct {
					TwoAc4128794Ba91Ef070785E25Be82688 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"2ac4128794ba91ef070785e25be82688"`
					Eight59F2Ce5C5527Eb8B6Aade671Ae2Ffeb struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"859f2ce5c5527eb8b6aade671ae2ffeb"`
					Nine68B8B372C87E70D54400804Fe7Ff718 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"968b8b372c87e70d54400804fe7ff718"`
					Be8Caed0421804679038Def16Cf01437 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"be8caed0421804679038def16cf01437"`
				} `json:"userSkills,omitempty"`
				CharData13 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills13 struct {
					Four633C05F585F98Ef73943D2Cf93659Ca struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"4633c05f585f98ef73943d2cf93659ca"`
					NineB13B486D8E2250Be15Aeb8E751C82B6 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"9b13b486d8e2250be15aeb8e751c82b6"`
					NineC8B5D7Df6Aac4Ecb495F892979Bb7B3 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"9c8b5d7df6aac4ecb495f892979bb7b3"`
					C99Cb3E84246F53A9F025Dcbe0E43159 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"c99cb3e84246f53a9f025dcbe0e43159"`
				} `json:"userSkills,omitempty"`
				CharData14 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb             string `json:"atb"`
									AtkScale        string `json:"atk_scale"`
									AttackPoise     string `json:"attack_poise"`
									DisplayAtkScale string `json:"display_atk_scale"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills14 struct {
					OneB0C4204Ac2Bf9C7Aeef1E631Fe3C280 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"1b0c4204ac2bf9c7aeef1e631fe3c280"`
					SixC9C7E48De94424967500De2F22Dbd6A struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"6c9c7e48de94424967500de2f22dbd6a"`
					Eight896E8838E36A58E3E26585199A9C9Da struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"8896e8838e36a58e3e26585199a9c9da"`
					Eaa2Ba83A129098Ff6A9C9A101684414 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"eaa2ba83a129098ff6a9c9a101684414"`
				} `json:"userSkills,omitempty"`
				CharData15 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills15 struct {
					Four68Ed7Cbb35C3D0185Ab5A8A1047486E struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"468ed7cbb35c3d0185ab5a8a1047486e"`
					FiveF9Ba45D5F953C3739Ba0930A360148B struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"5f9ba45d5f953c3739ba0930a360148b"`
					EightF575B059D36758D56753416Dbf7C756 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"8f575b059d36758d56753416dbf7c756"`
					C386C9E9F367D8B91E63982Ef73Bab8F struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"c386c9e9f367d8b91e63982ef73bab8f"`
				} `json:"userSkills,omitempty"`
				CharData16 struct {
					ID          string `json:"id"`
					Name        string `json:"name"`
					AvatarSqURL string `json:"avatarSqUrl"`
					AvatarRtURL string `json:"avatarRtUrl"`
					Rarity      struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"rarity"`
					Profession struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"profession"`
					Property struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"property"`
					WeaponType struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"weaponType"`
					Skills []struct {
						ID   string `json:"id"`
						Name string `json:"name"`
						Type struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"type"`
						Property struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"property"`
						IconURL    string `json:"iconUrl"`
						Desc       string `json:"desc"`
						DescParams struct {
						} `json:"descParams"`
						DescLevelParams struct {
							Num1 struct {
								Level  string `json:"level"`
								Params struct {
									Atb      string `json:"atb"`
									AtkScale string `json:"atk_scale"`
									Poise    string `json:"poise"`
								} `json:"params"`
							} `json:"1"`
						} `json:"descLevelParams"`
					} `json:"skills"`
					IllustrationURL string   `json:"illustrationUrl"`
					Tags            []string `json:"tags"`
				} `json:"charData,omitempty"`
				UserSkills16 struct {
					TwoF22B543Adb0F02A2Ae6Adaf94414B19 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"2f22b543adb0f02a2ae6adaf94414b19"`
					Four92623Fb505F551Feb131Ebc21E5Fe3B struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"492623fb505f551feb131ebc21e5fe3b"`
					Five36131A4B1969225A7200F45F18Cbef5 struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"536131a4b1969225a7200f45f18cbef5"`
					Ae6F2C3347E9B741F6Ae7371Aa70183A struct {
						SkillID  string `json:"skillId"`
						Level    int    `json:"level"`
						MaxLevel int    `json:"maxLevel"`
					} `json:"ae6f2c3347e9b741f6ae7371aa70183a"`
				} `json:"userSkills,omitempty"`
			} `json:"chars"`
			Achieve struct {
				AchieveMedals []interface{} `json:"achieveMedals"`
				Display       struct {
				} `json:"display"`
				Count int `json:"count"`
			} `json:"achieve"`
			SpaceShip struct {
				Rooms []struct {
					ID    string `json:"id"`
					Type  int    `json:"type"`
					Level int    `json:"level"`
					Chars []struct {
						CharID           string  `json:"charId"`
						PhysicalStrength float64 `json:"physicalStrength"`
						Favorability     int     `json:"favorability"`
					} `json:"chars"`
					Reports struct {
					} `json:"reports,omitempty"`
					Reports0 struct {
						Num1769471743 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemBlocWorker1              int `json:"item_bloc_worker_1"`
								ItemFactechToolsAsseblingMc1 int `json:"item_factech_tools_assebling_mc_1"`
								ItemFactechTree1             int `json:"item_factech_tree_1"`
								ItemPortStorager1            int `json:"item_port_storager_1"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769471743"`
						Num1769551732 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemPortStorager1 int `json:"item_port_storager_1"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769551732"`
						Num1769631723 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemPortStorager1 int `json:"item_port_storager_1"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769631723"`
					} `json:"reports,omitempty"`
					Reports1 struct {
						Num1769471743 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemExpcardStage2High int `json:"item_expcard_stage2_high"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769471743"`
						Num1769551732 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemExpcardStage2High int `json:"item_expcard_stage2_high"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769551732"`
						Num1769631723 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemExpcardStage2High int `json:"item_expcard_stage2_high"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769631723"`
					} `json:"reports,omitempty"`
					Reports2 struct {
						Num1769079365 struct {
							Char   []interface{} `json:"char"`
							Output struct {
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769079365"`
					} `json:"reports,omitempty"`
					Reports3 struct {
						Num1769471743 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemPlantCrylplant11 int `json:"item_plant_crylplant_1_1"`
								ItemPlantCrylplant12 int `json:"item_plant_crylplant_1_2"`
								ItemPlantMushroom11  int `json:"item_plant_mushroom_1_1"`
								ItemPlantMushroom12  int `json:"item_plant_mushroom_1_2"`
								ItemPlantSpcstone11  int `json:"item_plant_spcstone_1_1"`
								ItemPlantSpcstone12  int `json:"item_plant_spcstone_1_2"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769471743"`
						Num1769551732 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemPlantCrylplant12 int `json:"item_plant_crylplant_1_2"`
								ItemPlantMushroom11  int `json:"item_plant_mushroom_1_1"`
								ItemPlantMushroom12  int `json:"item_plant_mushroom_1_2"`
								ItemPlantSpcstone11  int `json:"item_plant_spcstone_1_1"`
								ItemPlantSpcstone12  int `json:"item_plant_spcstone_1_2"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769551732"`
						Num1769631723 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemPlantCrylplant12 int `json:"item_plant_crylplant_1_2"`
								ItemPlantMushroom12  int `json:"item_plant_mushroom_1_2"`
								ItemPlantSpcstone12  int `json:"item_plant_spcstone_1_2"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769631723"`
					} `json:"reports,omitempty"`
					Reports4 struct {
						Num1769471743 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemExpcardStage2High int `json:"item_expcard_stage2_high"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769471743"`
						Num1769551732 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemExpcardStage2High int `json:"item_expcard_stage2_high"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769551732"`
						Num1769631723 struct {
							Char   []string `json:"char"`
							Output struct {
								ItemExpcardStage2High int `json:"item_expcard_stage2_high"`
							} `json:"output"`
							CreatedTimeTs string `json:"createdTimeTs"`
						} `json:"1769631723"`
					} `json:"reports,omitempty"`
				} `json:"rooms"`
			} `json:"spaceShip"`
			Domain []struct {
				DomainID    string `json:"domainId"`
				Level       int    `json:"level"`
				Settlements []struct {
					ID             string `json:"id"`
					Level          int    `json:"level"`
					RemainMoney    string `json:"remainMoney"`
					OfficerCharIds string `json:"officerCharIds"`
					Name           string `json:"name"`
				} `json:"settlements"`
				MoneyMgr    string `json:"moneyMgr"`
				Collections []struct {
					LevelID       string `json:"levelId"`
					PuzzleCount   int    `json:"puzzleCount"`
					TrchestCount  int    `json:"trchestCount"`
					PieceCount    int    `json:"pieceCount"`
					BlackboxCount int    `json:"blackboxCount"`
				} `json:"collections"`
				Factory interface{} `json:"factory"`
				Name    string      `json:"name"`
			} `json:"domain"`
			Dungeon struct {
				CurStamina string `json:"curStamina"`
				MaxTs      string `json:"maxTs"`
				MaxStamina string `json:"maxStamina"`
			} `json:"dungeon"`
			BpSystem struct {
				CurLevel int `json:"curLevel"`
				MaxLevel int `json:"maxLevel"`
			} `json:"bpSystem"`
			DailyMission struct {
				DailyActivation    int `json:"dailyActivation"`
				MaxDailyActivation int `json:"maxDailyActivation"`
			} `json:"dailyMission"`
			Config struct {
				CharSwitch bool     `json:"charSwitch"`
				CharIds    []string `json:"charIds"`
			} `json:"config"`
			CurrentTs   string `json:"currentTs"`
			Quickaccess []struct {
				Name string `json:"name"`
				Icon string `json:"icon"`
				Link string `json:"link"`
			} `json:"quickaccess"`
		} `json:"detail"`
	} `json:"data"`
}
