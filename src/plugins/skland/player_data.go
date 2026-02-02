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
				} `json:"charData"`
				ID         string `json:"id"`
				Level      int    `json:"level"`
				UserSkills map[string]struct {
					SkillID  string `json:"skillId"`
					Level    int    `json:"level"`
					MaxLevel int    `json:"maxLevel"`
				} `json:"userSkills"`
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
							ID              string      `json:"id"`
							Name            string      `json:"name"`
							SkillID         string      `json:"skillId"`
							SkillDesc       string      `json:"skillDesc"`
							SkillDescParams interface{} `json:"skillDescParams"`
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
							ID              string      `json:"id"`
							Name            string      `json:"name"`
							SkillID         string      `json:"skillId"`
							SkillDesc       string      `json:"skillDesc"`
							SkillDescParams interface{} `json:"skillDescParams"`
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
							ID              string      `json:"id"`
							Name            string      `json:"name"`
							SkillID         string      `json:"skillId"`
							SkillDesc       string      `json:"skillDesc"`
							SkillDescParams interface{} `json:"skillDescParams"`
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
							ID              string      `json:"id"`
							Name            string      `json:"name"`
							SkillID         string      `json:"skillId"`
							SkillDesc       string      `json:"skillDesc"`
							SkillDescParams interface{} `json:"skillDescParams"`
						} `json:"suit"`
						Function string `json:"function"`
						Pkg      string `json:"pkg"`
					} `json:"equipData"`
				} `json:"secondAccessory,omitempty"`
				TacticalItem struct {
					TacticalItemId   string `json:"tacticalItemId"`
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
						ActiveEffect        string      `json:"activeEffect"`
						PassiveEffect       string      `json:"passiveEffect"`
						ActiveEffectParams  interface{} `json:"activeEffectParams"`
						PassiveEffectParams interface{} `json:"passiveEffectParams"`
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
				} `json:"weapon,omitempty"`
				Gender string `json:"gender"`
				OwnTs  string `json:"ownTs"`
			} `json:"chars"`
			Achieve struct {
				AchieveMedals []interface{} `json:"achieveMedals"`
				Display       interface{}   `json:"display"`
				Count         int           `json:"count"`
			} `json:"achieve"`
			SpaceShip struct {
				Rooms []struct {
					ID    string `json:"id"`
					Type  int    `json:"type"`
					Level int    `json:"level"`
					Chars []struct {
						CharId           string  `json:"charId"`
						PhysicalStrength float64 `json:"physicalStrength"`
						Favorability     int     `json:"favorability"`
					} `json:"chars"`
					Reports interface{} `json:"reports"`
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
